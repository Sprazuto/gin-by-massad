package models

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

type NullString sql.NullString

func (ns *NullString) Scan(value interface{}) error {
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return err
	}

	// If value is nil, set Valid to false
	if value == nil {
		ns.String, ns.Valid = "", false
		return nil
	}

	ns.String, ns.Valid = s.String, s.Valid
	return nil
}

func (ns NullString) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.String, nil
}

func (ns NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

func (ns *NullString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		ns.String, ns.Valid = "", false
		return nil
	}

	if err := json.Unmarshal(data, &ns.String); err != nil {
		return err
	}
	ns.Valid = true
	return nil
}

type NullInt64 struct {
	sql.NullInt64
}

func (ni NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
}

func (ni *NullInt64) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		ni.Int64, ni.Valid = 0, false
		return nil
	}

	if err := json.Unmarshal(data, &ni.Int64); err != nil {
		return err
	}
	ni.Valid = true
	return nil
}

type NullFloat64 struct {
	sql.NullFloat64
}

func (nf NullFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nf.Float64)
}

func (nf *NullFloat64) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		nf.Float64, nf.Valid = 0, false
		return nil
	}

	if err := json.Unmarshal(data, &nf.Float64); err != nil {
		return err
	}
	nf.Valid = true
	return nil
}

type NullStringSlice struct {
	Slice []string
	Valid bool
}

func (ns *NullStringSlice) Scan(value interface{}) error {
	if value == nil {
		ns.Slice, ns.Valid = nil, false
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		ns.Slice, ns.Valid = nil, false
		return nil
	}

	if err := json.Unmarshal(bytes, &ns.Slice); err != nil {
		ns.Slice, ns.Valid = nil, false
		return err
	}
	ns.Valid = true
	return nil
}

func (ns NullStringSlice) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.Slice)
}

func (ns *NullStringSlice) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		ns.Slice, ns.Valid = nil, false
		return nil
	}

	if err := json.Unmarshal(data, &ns.Slice); err != nil {
		return err
	}
	ns.Valid = true
	return nil
}

func (ns NullStringSlice) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return json.Marshal(ns.Slice)
}
