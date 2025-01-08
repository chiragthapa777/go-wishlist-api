package utils

import (
	"database/sql"
	"encoding/json"
)

type NullableString sql.NullString

func (ns *NullableString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*ns = NullableString{String: "", Valid: false}
		return nil
	}
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*ns = NullableString{String: s, Valid: true}
	return nil
}

func (ns NullableString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(ns.String)
}
