package usecases

import (
	"testing"

	"contact_duplicate_detector/entities/contacts"

	"github.com/stretchr/testify/assert"
)

func TestFindDuplicates(t *testing.T) {
	mockRepo := new(MockRepository)

	mockRepo.On("GetAllContacts").Return([]contacts.Contact{
		{ID: "1001", FirstName: "C", LastName: "F", Email: "mollis.lectus.pede@outlook.net", ZipCode: "5555", Address: "123 Main St"},
		{ID: "1002", FirstName: "C", LastName: "French", Email: "mollis.lectus.pede@outlook.net", ZipCode: "5555", Address: "123 Elm St"},
		{ID: "1003", FirstName: "Ciara", LastName: "F", Email: "non.lacinia.at@zoho.ca", ZipCode: "6666", Address: "123 Oak St"},
	}, nil)

	duplicateFinder := NewDuplicateFinder(mockRepo)

	duplicates, err := duplicateFinder.FindDuplicates()

	assert.NoError(t, err)

	expected := []DuplicateResult{
		{SourceID: "1001", MatchID: "1002", Accuracy: "High"},
		{SourceID: "1001", MatchID: "1003", Accuracy: "Low"},
	}

	// Verifica que la longitud de los duplicados es la esperada
	assert.Equal(t, len(expected), len(duplicates), "Expected number of duplicates to be %d, got %d", len(expected), len(duplicates))

	// Itera a través de los duplicados y verifica cada uno
	for i := range expected {
		assert.Equal(t, expected[i].SourceID, duplicates[i].SourceID, "Expected SourceID to be %s, got %s", expected[i].SourceID, duplicates[i].SourceID)
		assert.Equal(t, expected[i].MatchID, duplicates[i].MatchID, "Expected MatchID to be %s, got %s", expected[i].MatchID, duplicates[i].MatchID)
		assert.Equal(t, expected[i].Accuracy, duplicates[i].Accuracy, "Expected Accuracy to be %s, got %s", expected[i].Accuracy, duplicates[i].Accuracy)
	}

}

func TestFindDuplicatesNoDuplicates(t *testing.T) {
	mockRepo := new(MockRepository)

	mockRepo.On("GetAllContacts").Return([]contacts.Contact{
		{ID: "1001", FirstName: "C", LastName: "F", Email: "unique@example.com", ZipCode: "4", Address: "123 Main St"},
		{ID: "1002", FirstName: "D", LastName: "G", Email: "another.unique@example.com", ZipCode: "5555", Address: "123 Elm St"},
	}, nil)

	duplicateFinder := NewDuplicateFinder(mockRepo)

	duplicates, err := duplicateFinder.FindDuplicates()

	assert.NoError(t, err)
	assert.Empty(t, duplicates, "Expected no duplicates, but found %d", len(duplicates))
}
