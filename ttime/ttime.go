package ttime

import (
	"bytes"
	"encoding/json"
	"strconv"
	"time"
)

// PossibleTime is a time.Time object with a custom JSON unmarshaler to handle cases where the
// JSON value might be an empty or null string. If the JSON value is empty or null, the
// PossibleTime value will be a zero value of time.Time.
type PossibleTime struct {
	time.Time
}

// UnmarshalJSON decodes a possibly empty time string into a time.Time object.
func (p *PossibleTime) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	if bytes.Equal(data, []byte(`""`)) {
		return nil
	}
	if bytes.Equal(data, []byte(`null`)) {
		return nil
	}
	data = bytes.ReplaceAll(data, []byte(`"`), []byte(``))
	t, err := time.Parse(time.RFC3339, string(data))
	if err != nil {
		return err
	}
	*p = PossibleTime{t}
	return nil
}

// Timestamp is a time.Time object with a custom JSON unmarshaler to parse int64 timestamps.
type Timestamp struct {
	time.Time
}

// UnmarshalJSON decodes an int64 timestamp into a time.Time object.
func (p *Timestamp) UnmarshalJSON(bytes []byte) error {
	var raw int64
	err := json.Unmarshal(bytes, &raw)
	if err != nil {
		return err
	}
	switch len(strconv.FormatInt(raw, 10)) {
	case 13:
		// Milliseconds
		p.Time = time.UnixMilli(raw).UTC()
	case 19:
		// Nanoseconds
		p.Time = time.Unix(0, raw).UTC()
	default:
		p.Time = time.Unix(raw, 0).UTC()
	}
	return nil
}
