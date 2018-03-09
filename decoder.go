package vfields

import (
	"encoding/json"
	"fmt"
)

type fieldsDecoder map[string]json.RawMessage

func (f fieldsDecoder) Get(fieldName string) json.RawMessage {
	v, exists := f[fieldName]
	if !exists {
		return nil
	}
	return v
}

func (f fieldsDecoder) Hydrate(fields []*FieldContainer) error {
	errs := fieldsDecodeError{}
	for _, field := range fields {
		if dat := f.Get(field.Name); dat != nil {
			if err := json.Unmarshal(dat, field); err != nil {
				errs[field.Name] = fieldDecodeError{
					Name: field.Name,
					Err:  err,
					Data: dat,
				}
			}
		}
	}
	if len(errs) != 0 {
		return errs
	}
	return nil
}

type fieldDecodeError struct {
	Name string
	Err  error
	Data json.RawMessage
}

func (f fieldDecodeError) Error() string {
	return fmt.Sprintf("error decode, field_name=%v err=%v, dat=%v", f.Name, f.Err, string(f.Data))
}

type fieldsDecodeError map[string]fieldDecodeError

func (f fieldsDecodeError) Error() string {
	return fmt.Sprintf("error decode %d fields", len(f))
}
