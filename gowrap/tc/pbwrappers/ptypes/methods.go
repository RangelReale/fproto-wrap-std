package pbwrappers_ptypes

import (
	"encoding/json"
	"fmt"
	"strconv"
)

//
// DoubleValue
//

// Scan implements the sql.Scanner interface.
func (u *DoubleValue) Scan(src interface{}) error {
	if src == nil {
		u.Valid = false
		return nil
	}

	u.Valid = true
	switch src := src.(type) {
	case []byte:
		return u.UnmarshalText(src)
	case string:
		return u.UnmarshalText([]byte(src))
	}

	return fmt.Errorf("DoubleValue: cannot convert %T to float64", src)
}

// MarshalJSON implements the json.Marshaler interface.
func (t DoubleValue) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return nil, nil
	}

	return json.Marshal(t.Value)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *DoubleValue) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		t.Valid = false
		return nil
	}

	return json.Unmarshal(data, t.Value)
}

// MarshalText implements the encoding.TextMarshaler interface.
func (t DoubleValue) MarshalText() ([]byte, error) {
	if !t.Valid {
		return nil, nil
	}
	return []byte(fmt.Sprintf("%f", t.Value)), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (t *DoubleValue) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		t.Valid = false
		return nil
	}

	tmp, err := strconv.ParseFloat(string(data), 64)
	if err == nil {
		t.Value = tmp
		t.Valid = true
	}
	return err
}

//
// FloatValue
//

// Scan implements the sql.Scanner interface.
func (u *FloatValue) Scan(src interface{}) error {
	if src == nil {
		u.Valid = false
		return nil
	}

	u.Valid = true
	switch src := src.(type) {
	case []byte:
		return u.UnmarshalText(src)
	case string:
		return u.UnmarshalText([]byte(src))
	}

	return fmt.Errorf("FloatValue: cannot convert %T to float32", src)
}

// MarshalJSON implements the json.Marshaler interface.
func (t FloatValue) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return nil, nil
	}

	return json.Marshal(t.Value)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *FloatValue) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		t.Valid = false
		return nil
	}

	return json.Unmarshal(data, t.Value)
}

// MarshalText implements the encoding.TextMarshaler interface.
func (t FloatValue) MarshalText() ([]byte, error) {
	if !t.Valid {
		return nil, nil
	}
	return []byte(fmt.Sprintf("%f", t.Value)), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (t *FloatValue) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		t.Valid = false
		return nil
	}

	tmp, err := strconv.ParseFloat(string(data), 32)
	if err == nil {
		t.Value = float32(tmp)
		t.Valid = true
	}
	return err
}

//
// Int64Value
//

// Scan implements the sql.Scanner interface.
func (u *Int64Value) Scan(src interface{}) error {
	if src == nil {
		u.Valid = false
		return nil
	}

	u.Valid = true
	switch src := src.(type) {
	case []byte:
		return u.UnmarshalText(src)
	case string:
		return u.UnmarshalText([]byte(src))
	}

	return fmt.Errorf("Int64Value: cannot convert %T to int64", src)
}

// MarshalJSON implements the json.Marshaler interface.
func (t Int64Value) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return nil, nil
	}

	return json.Marshal(t.Value)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *Int64Value) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		t.Valid = false
		return nil
	}

	return json.Unmarshal(data, t.Value)
}

// MarshalText implements the encoding.TextMarshaler interface.
func (t Int64Value) MarshalText() ([]byte, error) {
	if !t.Valid {
		return nil, nil
	}
	return []byte(fmt.Sprintf("%d", t.Value)), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (t *Int64Value) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		t.Valid = false
		return nil
	}

	tmp, err := strconv.ParseInt(string(data), 10, 64)
	if err == nil {
		t.Value = tmp
		t.Valid = true
	}
	return err
}

//
// UInt64Value
//

// Scan implements the sql.Scanner interface.
func (u *UInt64Value) Scan(src interface{}) error {
	if src == nil {
		u.Valid = false
		return nil
	}

	u.Valid = true
	switch src := src.(type) {
	case []byte:
		return u.UnmarshalText(src)
	case string:
		return u.UnmarshalText([]byte(src))
	}

	return fmt.Errorf("UInt64Value: cannot convert %T to uint64", src)
}

// MarshalJSON implements the json.Marshaler interface.
func (t UInt64Value) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return nil, nil
	}

	return json.Marshal(t.Value)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *UInt64Value) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		t.Valid = false
		return nil
	}

	return json.Unmarshal(data, t.Value)
}

// MarshalText implements the encoding.TextMarshaler interface.
func (t UInt64Value) MarshalText() ([]byte, error) {
	if !t.Valid {
		return nil, nil
	}
	return []byte(fmt.Sprintf("%d", t.Value)), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (t *UInt64Value) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		t.Valid = false
		return nil
	}

	tmp, err := strconv.ParseUint(string(data), 10, 64)
	if err == nil {
		t.Value = tmp
		t.Valid = true
	}
	return err
}

//
// Int32Value
//

// Scan implements the sql.Scanner interface.
func (u *Int32Value) Scan(src interface{}) error {
	if src == nil {
		u.Valid = false
		return nil
	}

	u.Valid = true
	switch src := src.(type) {
	case []byte:
		return u.UnmarshalText(src)
	case string:
		return u.UnmarshalText([]byte(src))
	}

	return fmt.Errorf("Int32Value: cannot convert %T to int32", src)
}

// MarshalJSON implements the json.Marshaler interface.
func (t Int32Value) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return nil, nil
	}

	return json.Marshal(t.Value)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *Int32Value) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		t.Valid = false
		return nil
	}

	return json.Unmarshal(data, t.Value)
}

// MarshalText implements the encoding.TextMarshaler interface.
func (t Int32Value) MarshalText() ([]byte, error) {
	if !t.Valid {
		return nil, nil
	}
	return []byte(fmt.Sprintf("%d", t.Value)), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (t *Int32Value) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		t.Valid = false
		return nil
	}

	tmp, err := strconv.ParseInt(string(data), 10, 32)
	if err == nil {
		t.Value = int32(tmp)
		t.Valid = true
	}
	return err
}

//
// UInt32Value
//

// Scan implements the sql.Scanner interface.
func (u *UInt32Value) Scan(src interface{}) error {
	if src == nil {
		u.Valid = false
		return nil
	}

	u.Valid = true
	switch src := src.(type) {
	case []byte:
		return u.UnmarshalText(src)
	case string:
		return u.UnmarshalText([]byte(src))
	}

	return fmt.Errorf("UInt32Value: cannot convert %T to uint32", src)
}

// MarshalJSON implements the json.Marshaler interface.
func (t UInt32Value) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return nil, nil
	}

	return json.Marshal(t.Value)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *UInt32Value) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		t.Valid = false
		return nil
	}

	return json.Unmarshal(data, t.Value)
}

// MarshalText implements the encoding.TextMarshaler interface.
func (t UInt32Value) MarshalText() ([]byte, error) {
	if !t.Valid {
		return nil, nil
	}
	return []byte(fmt.Sprintf("%d", t.Value)), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (t *UInt32Value) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		t.Valid = false
		return nil
	}

	tmp, err := strconv.ParseUint(string(data), 10, 32)
	if err == nil {
		t.Value = uint32(tmp)
		t.Valid = true
	}
	return err
}

//
// BoolValue
//

// Scan implements the sql.Scanner interface.
func (u *BoolValue) Scan(src interface{}) error {
	if src == nil {
		u.Valid = false
		return nil
	}

	u.Valid = true
	switch src := src.(type) {
	case []byte:
		return u.UnmarshalText(src)
	case string:
		return u.UnmarshalText([]byte(src))
	}

	return fmt.Errorf("BoolValue: cannot convert %T to bool", src)
}

// MarshalJSON implements the json.Marshaler interface.
func (t BoolValue) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return nil, nil
	}

	return json.Marshal(t.Value)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *BoolValue) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		t.Valid = false
		return nil
	}

	return json.Unmarshal(data, t.Value)
}

// MarshalText implements the encoding.TextMarshaler interface.
func (t BoolValue) MarshalText() ([]byte, error) {
	if !t.Valid {
		return nil, nil
	}
	return []byte(fmt.Sprintf("%t")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (t *BoolValue) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		t.Valid = false
		return nil
	}

	tmp, err := strconv.ParseBool(string(data))
	if err == nil {
		t.Value = tmp
		t.Valid = true
	}
	return err
}

//
// StringValue
//

// Scan implements the sql.Scanner interface.
func (u *StringValue) Scan(src interface{}) error {
	if src == nil {
		u.Valid = false
		return nil
	}

	u.Valid = true
	switch src := src.(type) {
	case []byte:
		return u.UnmarshalText(src)
	case string:
		return u.UnmarshalText([]byte(src))
	}

	return fmt.Errorf("StringValue: cannot convert %T to string", src)
}

// MarshalJSON implements the json.Marshaler interface.
func (t StringValue) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return nil, nil
	}

	return json.Marshal(t.Value)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *StringValue) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		t.Valid = false
		return nil
	}

	return json.Unmarshal(data, t.Value)
}

// MarshalText implements the encoding.TextMarshaler interface.
func (t StringValue) MarshalText() ([]byte, error) {
	if !t.Valid {
		return nil, nil
	}
	return []byte(t.Value), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (t *StringValue) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		t.Valid = false
		return nil
	}

	t.Value = string(data)
	t.Valid = true

	return nil
}

//
// BytesValue
//

// Scan implements the sql.Scanner interface.
func (u *BytesValue) Scan(src interface{}) error {
	if src == nil {
		u.Valid = false
		return nil
	}

	u.Valid = true
	switch src := src.(type) {
	case []byte:
		return u.UnmarshalText(src)
	case string:
		return u.UnmarshalText([]byte(src))
	}

	return fmt.Errorf("BytesValue: cannot convert %T to []byte", src)
}

// MarshalJSON implements the json.Marshaler interface.
func (t BytesValue) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return nil, nil
	}

	return json.Marshal(t.Value)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *BytesValue) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		t.Valid = false
		return nil
	}

	return json.Unmarshal(data, t.Value)
}

// MarshalText implements the encoding.TextMarshaler interface.
func (t BytesValue) MarshalText() ([]byte, error) {
	if !t.Valid {
		return nil, nil
	}
	return t.Value, nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (t *BytesValue) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		t.Valid = false
		return nil
	}

	t.Value = data
	t.Valid = true

	return nil
}
