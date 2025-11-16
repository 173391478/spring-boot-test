
package data

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// DataRecord represents a single data record
type DataRecord struct {
	ID        string
	Name      string
	Email     string
	Age       int
	Timestamp time.Time
}

// Validator defines the interface for data validation
type Validator interface {
	Validate(DataRecord) error
}

// Transformer defines the interface for data transformation
type Transformer interface {
	Transform(DataRecord) DataRecord
}

// EmailValidator validates email format
type EmailValidator struct{}

func (ev EmailValidator) Validate(record DataRecord) error {
	if !strings.Contains(record.Email, "@") {
		return errors.New("invalid email format")
	}
	return nil
}

// AgeValidator validates age range
type AgeValidator struct {
	MinAge int
	MaxAge int
}

func (av AgeValidator) Validate(record DataRecord) error {
	if record.Age < av.MinAge || record.Age > av.MaxAge {
		return fmt.Errorf("age %d out of valid range [%d-%d]", record.Age, av.MinAge, av.MaxAge)
	}
	return nil
}

// NameTransformer transforms name to title case
type NameTransformer struct{}

func (nt NameTransformer) Transform(record DataRecord) DataRecord {
	if record.Name != "" {
		record.Name = strings.Title(strings.ToLower(record.Name))
	}
	return record
}

// DataProcessor orchestrates the data processing pipeline
type DataProcessor struct {
	validators   []Validator
	transformers []Transformer
}

func NewDataProcessor() *DataProcessor {
	return &DataProcessor{
		validators:   make([]Validator, 0),
		transformers: make([]Transformer, 0),
	}
}

func (dp *DataProcessor) AddValidator(v Validator) {
	dp.validators = append(dp.validators, v)
}

func (dp *DataProcessor) AddTransformer(t Transformer) {
	dp.transformers = append(dp.transformers, t)
}

// ProcessRecord processes a single record through the pipeline
func (dp *DataProcessor) ProcessRecord(record DataRecord) (DataRecord, error) {
	// Validation stage
	for _, validator := range dp.validators {
		if err := validator.Validate(record); err != nil {
			return DataRecord{}, fmt.Errorf("validation failed: %w", err)
		}
	}

	// Transformation stage
	processedRecord := record
	for _, transformer := range dp.transformers {
		processedRecord = transformer.Transform(processedRecord)
	}

	return processedRecord, nil
}

// ProcessBatch processes multiple records
func (dp *DataProcessor) ProcessBatch(records []DataRecord) ([]DataRecord, []error) {
	var processed []DataRecord
	var errors []error

	for i, record := range records {
		processedRecord, err := dp.ProcessRecord(record)
		if err != nil {
			errors = append(errors, fmt.Errorf("record %d (ID: %s): %w", i, record.ID, err))
			continue
		}
		processed = append(processed, processedRecord)
	}

	return processed, errors
}
