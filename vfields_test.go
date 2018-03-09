package vfields

import (
	"encoding/json"
	"strings"
	"testing"
)

type someForm struct {
	Fields []*FieldContainer
}

func (f someForm) Name() string {
	return "some1"
}

func (f someForm) Valid() error {
	// TODO: agg errors
	for _, field := range f.Fields {
		if err := field.Valid(); err != nil {
			return err
		}
	}
	return nil
}

func (f *someForm) UnmarshalJSON(dat []byte) error {
	fields := fieldsDecoder{}
	if err := json.Unmarshal(dat, &fields); err != nil {
		return err
	}
	return fields.Hydrate(f.Fields)
}

func TestSimpleDecode(t *testing.T) {
	jsonDat := `{"string1": "a", "integer1": 2, "location1": {"lng": 37.740070, "lat": 55.427397}}`
	f := someForm{
		Fields: []*FieldContainer{
			{
				&String{
					ValidatorRule: "required",
					Data:          "abc",
				},
				"string1",
			},
			{
				&Integer{
					ValidatorRule: "required,gte=0,lte=10",
					Data:          2,
				},
				"integer1",
			},
			{
				&Location{
					ValidatorRule: "inmsk",
					Data:          NullPoint{Y: 55.427397, X: 37.740070, Valid: true}, // Domodedovo
				},
				"location1",
			},
		},
	}
	if err := json.NewDecoder(strings.NewReader(jsonDat)).Decode(&f); err != nil {
		t.Errorf("unmarshal data %+v", err)
	}

	if f.Fields[0].Field.(*String).Data != "a" {
		t.Errorf("not expected strng, got %v want %v", f.Fields[0].Field.(*String).Data, "a")
	}

	if f.Fields[1].Field.(*Integer).Data != 2 {
		t.Errorf("not expected strng, got %v want %v", f.Fields[1].Field.(*Integer).Data, 2)
	}

	if f.Fields[2].Field.(*Location).Data.X != 37.740070 || f.Fields[2].Field.(*Location).Data.Y != 55.427397 {
		t.Errorf("not expected strng, got %v want %v", f.Fields[2].Field.(*Location).Data, NullPoint{Y: 55.427397, X: 37.740070, Valid: true})
	}
}

func TestSimpleValidation(t *testing.T) {
	f := someForm{
		Fields: []*FieldContainer{
			{
				&String{
					ValidatorRule: "required",
					Data:          "abc",
				},
				"string1",
			},
			{
				&Integer{
					ValidatorRule: "required,gte=0,lte=10",
					Data:          2,
				},
				"integer1",
			},
			{
				&Location{
					ValidatorRule: "inmsk",
					// Data:          NullPoint{Y: 1, X: 1, Valid: true},
					Data: NullPoint{Y: 55.427397, X: 37.740070, Valid: true}, // Domodedovo
				},
				"location1",
			},
		},
	}
	if err := f.Valid(); err != nil {
		t.Error("failed form validation", err)
	}
}
