package usecases

import (
	"contact_duplicate_detector/entities/contacts"
)

type ContactRepository interface {
	GetAllContacts() ([]contacts.Contact, error)
}

type DuplicateResult struct {
	SourceID string
	MatchID  string
	Accuracy string
}

type DuplicateFinder struct {
	repo ContactRepository
}

func NewDuplicateFinder(repo ContactRepository) *DuplicateFinder {
	return &DuplicateFinder{repo}
}

func (df *DuplicateFinder) FindDuplicates() ([]DuplicateResult, error) {
	contacts, err := df.repo.GetAllContacts()
	if err != nil {
		return nil, err
	}

	results := []DuplicateResult{}
	for i := 0; i < len(contacts); i++ {
		for j := i + 1; j < len(contacts); j++ {
			accuracy, _ := compareContacts(contacts[i], contacts[j])

			if accuracy != "No Match" {
				results = append(results, DuplicateResult{
					SourceID: contacts[i].ID,
					MatchID:  contacts[j].ID,
					Accuracy: accuracy,
				})
			}
		}
	}

	return results, nil
}

func compareContacts(contactA, contactB contacts.Contact) (string, error) {
	score := 0

	if contactA.FirstName == contactB.FirstName {
		score++
	}
	if contactA.LastName == contactB.LastName {
		score++
	}
	if contactA.Email == contactB.Email {
		score++
	}
	if contactA.ZipCode == contactB.ZipCode {
		score++
	}
	if contactA.Address == contactB.Address {
		score++
	}

	switch score {
	case 3, 4, 5:
		return "High", nil
	case 2:
		return "Medium", nil
	case 1:
		return "Low", nil
	default:
		return "No Match", nil
	}
}
