package vfields

type Forms []Form

type Form interface {
	Name() string
	Valid() error
}

type Field interface {
	Valid() error

	UnmarshalJSON([]byte) error
	MarshalJSON() ([]byte, error)
}

type FieldContainer struct {
	Field
	Name string

	// Name() string
	// Label() string
	// Description() string
}

// type FieldValues map[string]interface{}

// // type Hydrator(
// // 	formName string,
// // 	vals FieldValues,
// // ) Form {

// // }
