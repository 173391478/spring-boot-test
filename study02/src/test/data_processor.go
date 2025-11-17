
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type DataRecord struct {
	ID    int
	Name  string
	Value float64
}

func ProcessCSVFile(filename string) ([]DataRecord, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records := make([]DataRecord, 0)

	// Skip header
	_, err = reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read header: %w", err)
	}

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read row: %w", err)
		}

		if len(row) != 3 {
			return nil, fmt.Errorf("invalid row length: expected 3, got %d", len(row))
		}

		id, err := strconv.Atoi(row[0])
		if err != nil {
			return nil, fmt.Errorf("invalid ID format: %w", err)
		}

		name := row[1]

		value, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value format: %w", err)
		}

		records = append(records, DataRecord{
			ID:    id,
			Name:  name,
			Value: value,
		})
	}

	return records, nil
}

func ValidateRecords(records []DataRecord) error {
	for _, record := range records {
		if record.ID <= 0 {
			return fmt.Errorf("invalid ID: %d", record.ID)
		}
		if record.Name == "" {
			return fmt.Errorf("empty name for record ID: %d", record.ID)
		}
		if record.Value < 0 {
			return fmt.Errorf("negative value for record ID: %d", record.ID)
		}
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: data_processor <csv_file>")
		os.Exit(1)
	}

	records, err := ProcessCSVFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error processing file: %v\n", err)
		os.Exit(1)
	}

	err = ValidateRecords(records)
	if err != nil {
		fmt.Printf("Validation error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully processed %d records\n", len(records))
	for _, record := range records {
		fmt.Printf("ID: %d, Name: %s, Value: %.2f\n", record.ID, record.Name, record.Value)
	}
}
package main

import (
    "errors"
    "fmt"
    "strings"
)

type DataRecord struct {
    ID    int
    Name  string
    Email string
    Age   int
}

type ValidationRule func(DataRecord) error
type Transformation func(DataRecord) DataRecord

func ValidateEmail(record DataRecord) error {
    if !strings.Contains(record.Email, "@") {
        return errors.New("invalid email format")
    }
    return nil
}

func ValidateAge(record DataRecord) error {
    if record.Age < 18 || record.Age > 120 {
        return errors.New("age must be between 18 and 120")
    }
    return nil
}

func TransformEmailToLower(record DataRecord) DataRecord {
    record.Email = strings.ToLower(record.Email)
    return record
}

func ProcessData(records []DataRecord, validators []ValidationRule, transformers []Transformation) ([]DataRecord, []error) {
    var processed []DataRecord
    var errors []error

    for _, record := range records {
        valid := true
        for _, validator := range validators {
            if err := validator(record); err != nil {
                errors = append(errors, fmt.Errorf("record %d: %w", record.ID, err))
                valid = false
                break
            }
        }

        if !valid {
            continue
        }

        transformed := record
        for _, transformer := range transformers {
            transformed = transformer(transformed)
        }
        processed = append(processed, transformed)
    }

    return processed, errors
}

func main() {
    records := []DataRecord{
        {1, "John Doe", "JOHN@EXAMPLE.COM", 25},
        {2, "Jane Smith", "jane.smith@test.org", 30},
        {3, "Bob Wilson", "invalid-email", 17},
        {4, "Alice Brown", "ALICE@DOMAIN.COM", 150},
    }

    validators := []ValidationRule{ValidateEmail, ValidateAge}
    transformers := []Transformation{TransformEmailToLower}

    processed, errs := ProcessData(records, validators, transformers)

    fmt.Println("Processed records:")
    for _, record := range processed {
        fmt.Printf("ID: %d, Name: %s, Email: %s, Age: %d\n", 
            record.ID, record.Name, record.Email, record.Age)
    }

    if len(errs) > 0 {
        fmt.Println("\nValidation errors:")
        for _, err := range errs {
            fmt.Println(err)
        }
    }
}
package main

import (
    "fmt"
    "strings"
)

type DataRecord struct {
    ID    int
    Name  string
    Email string
    Active bool
}

type DataProcessor struct {
    records []DataRecord
}

func NewDataProcessor() *DataProcessor {
    return &DataProcessor{
        records: make([]DataRecord, 0),
    }
}

func (dp *DataProcessor) AddRecord(record DataRecord) {
    dp.records = append(dp.records, record)
}

func (dp *DataProcessor) FilterActive() []DataRecord {
    var active []DataRecord
    for _, record := range dp.records {
        if record.Active {
            active = append(active, record)
        }
    }
    return active
}

func (dp *DataProcessor) TransformEmails(domain string) []string {
    var emails []string
    for _, record := range dp.records {
        if record.Active && record.Email != "" {
            transformed := strings.ToLower(record.Email)
            if !strings.Contains(transformed, "@") {
                transformed = transformed + "@" + domain
            }
            emails = append(emails, transformed)
        }
    }
    return emails
}

func (dp *DataProcessor) GetRecordCount() int {
    return len(dp.records)
}

func main() {
    processor := NewDataProcessor()
    
    processor.AddRecord(DataRecord{1, "John Doe", "john@example.com", true})
    processor.AddRecord(DataRecord{2, "Jane Smith", "jane.smith", true})
    processor.AddRecord(DataRecord{3, "Bob Wilson", "bob@test.org", false})
    processor.AddRecord(DataRecord{4, "Alice Brown", "", true})
    
    fmt.Printf("Total records: %d\n", processor.GetRecordCount())
    
    active := processor.FilterActive()
    fmt.Printf("Active records: %d\n", len(active))
    
    emails := processor.TransformEmails("default.com")
    fmt.Println("Processed emails:")
    for _, email := range emails {
        fmt.Println(email)
    }
}