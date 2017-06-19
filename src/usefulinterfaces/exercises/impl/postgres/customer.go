package postgres

import (
	"database/sql"
	"log"
	"usefulinterfaces/exercises/impl"
)

func NewCustomerService(db *sql.DB, l log.Logger) *CustomerService {
	return &CustomerService{
		db:     db,
		logger: l,
	}

}

type CustomerService struct {
	db     *sql.DB
	logger log.Logger
}

func (c *CustomerService) Get(id int) (*impl.Customer, error) {
	panic("not implemented")
}
