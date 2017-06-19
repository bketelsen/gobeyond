// Adapted from http://www.alexedwards.net/blog/a-recap-of-request-handling
package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/bketelsen/trace"
	xtr "golang.org/x/net/trace"
	"gopkg.in/throttled/throttled.v2"
	"gopkg.in/throttled/throttled.v2/store/memstore"
)

type timeHandler struct {
	format string
	el     xtr.EventLog
}

// Inside the handler, use traces to capture request specific events and timings
// and the el member to log service specific events - like failures
func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t, _ := trace.NewContext(context.Background(), "webserver", "servehttp")
	defer t.Finish()

	tm := time.Now().Format(th.format)
	// log to the trace
	t.LazyPrintf("time %v", tm)
	w.Write([]byte("The time is: " + tm))
}

func main() {

	store, err := memstore.New(65536)
	if err != nil {
		log.Fatal(err)
	}

	quota := throttled.RateQuota{throttled.PerMin(20), 5}
	rateLimiter, err := throttled.NewGCRARateLimiter(store, quota)
	if err != nil {
		log.Fatal(err)
	}

	httpRateLimiter := throttled.HTTPRateLimiter{
		RateLimiter: rateLimiter,
		VaryBy:      &throttled.VaryBy{Path: true},
	}
	eventLog := trace.NewEventLog("webserver", "http")
	th := &timeHandler{
		format: time.RFC1123,
		el:     eventLog,
	}
	defer eventLog.Finish()

	http.Handle("/time", httpRateLimiter.RateLimit(th))
	eventLog.Printf("Listening on %s", ":3000")

	http.ListenAndServe(":3000", nil)
}
