package main

import (
    "contact_duplicate_detector/repository" 
    "contact_duplicate_detector/usecases" 
    "fmt"
)

func main() {
    // Create an instance of the repository that reads from a CSV file
    repo := &repositories.CSVContactRepository{FilePath: "data/input.csv"}

    // Create an instance of the use case
    duplicateFinder := usecases.NewDuplicateFinder(repo)

    // Call the use case to find duplicates
    duplicates, err := duplicateFinder.FindDuplicates()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Print the results of the duplicate search
    fmt.Println("Duplicates found:")
    for _, result := range duplicates {
        fmt.Printf("Source ID: %s, Match ID: %s, Accuracy: %s\n", 
            result.SourceID, result.MatchID, result.Accuracy)
    }
}
