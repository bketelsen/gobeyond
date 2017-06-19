package mocks

import "usefulinterfaces/exercises/impl"

func NewMockCustomerService() *MockCustomerService {
	mdb := make(map[int]*impl.Customer)

	return &MockCustomerService{
		mockdb: mdb,
	}
}

type MockCustomerService struct {
	mockdb map[int]*impl.Customer
}

func (c *MockCustomerService) Get(id int) (*impl.Customer, error) {
	panic("not implemented")
}
