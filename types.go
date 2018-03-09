package vfields

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strings"

	validator "gopkg.in/go-playground/validator.v9"
)

// String

type String struct {
	Data          string
	ValidatorRule string
}

func (f String) Valid() error {
	return validator.New().Var(f.Data, f.ValidatorRule)
}

func (f String) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.Data)
}

func (f *String) UnmarshalJSON(dat []byte) error {
	return json.Unmarshal(dat, &f.Data)
}

// Integer

type Integer struct {
	Data          int64
	ValidatorRule string
}

func (f Integer) Valid() error {
	return validator.New().Var(f.Data, f.ValidatorRule)
}

func (f Integer) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.Data)
}

func (f *Integer) UnmarshalJSON(dat []byte) error {
	return json.Unmarshal(dat, &f.Data)
}

// Location

type Location struct {
	Data          NullPoint
	ValidatorRule string
}

func NullPointValuer(field reflect.Value) interface{} {
	point, ok := field.Interface().(NullPoint)
	if !ok {
		return false
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(point)

	return buf.String()
}

func (f Location) Valid() error {
	v := validator.New()

	v.RegisterCustomTypeFunc(NullPointValuer, NullPoint{})

	v.RegisterValidation("inmsk", func(fl validator.FieldLevel) bool {
		if fl.Field().Type().Kind() == reflect.Bool {
			return fl.Field().Bool()
		}
		if fl.Field().Type().Kind() != reflect.String {
			return false
		}
		point := &NullPoint{}
		if err := json.NewDecoder(strings.NewReader(fl.Field().String())).Decode(point); err != nil {
			return false
		}
		return int(point.Distance(MoscowCenterPoint)) <= MoscowRadius
	})

	return v.Var(&f.Data, f.ValidatorRule)
}

func (f Location) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.Data)
}

func (f *Location) UnmarshalJSON(dat []byte) error {
	return json.Unmarshal(dat, &f.Data)
}
