package repositories

import (
    "encoding/csv"
    "os"
    "log"

	"contact_duplicate_detector/entities/contacts"
)

type CSVContactRepository struct {
    FilePath string
}

func (repo *CSVContactRepository) GetAllContacts() ([]contacts.Contact, error) {
    file, err := os.Open(repo.FilePath)
    if err != nil {
        log.Println("Error opening CSV file:", err)
        return nil, err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        log.Println("Error reading CSV file:", err)
        return nil, err
    }

    var contactList []contacts.Contact
    for _, record := range records[1:] { // Skip header
        contact := contacts.Contact{
            ID:        record[0], // Asumiendo que el ID está en la primera columna
            FirstName: record[1],
            LastName:  record[2],
            Email:     record[3],
            ZipCode:   record[4],
            Address:   record[5],
        }
        contactList = append(contactList, contact)
    }

    return contactList, nil
}