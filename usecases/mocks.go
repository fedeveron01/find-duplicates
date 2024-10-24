package usecases 


import (
	"github.com/stretchr/testify/mock"
	"contact_duplicate_detector/entities/contacts"

)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetAllContacts() ([]contacts.Contact, error) {
	args := m.Mock.Called()
	return args.Get(0).([]contacts.Contact), args.Error(1)
}
