package utils

import (
	"encoding/json"
)

// OptionalNullableProperty can hold a nullable property that may be missing or null.
type OptionalNullableProperty[T any] struct {
	Data  T
	Null  bool
	Set   bool // to ensure that key is not missing is json
	Valid bool
}

// UnmarshalJSON implements custom unmarshalling for OptionalNullableProperty.
func (ns *OptionalNullableProperty[T]) UnmarshalJSON(data []byte) error {
	ns.Set = true
	ns.Valid = true

	// If the data is null, mark as null.
	if string(data) == "null" {
		ns.Null = true
		ns.Valid = false
		return nil
	}

	// Unmarshal into the property data.
	if err := json.Unmarshal(data, &ns.Data); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements custom marshalling for OptionalNullableProperty.
func (ns *OptionalNullableProperty[T]) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.Data)
}
