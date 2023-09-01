package jsonmap

import (
	"bytes"
	"encoding/json"
)

// MarshalJSON implements json.Marshaler interface.
// It marshals the map into JSON object.
//
//	data, err := json.Marshal(m)
func (m *Map) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('{')
	enc := json.NewEncoder(&buf)
	for i, key := range m.Keys() {
		if i > 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte('"')
		buf.WriteString(key)
		buf.WriteString(`":`)
		if err := enc.Encode(m.elements[key].value); err != nil {
			return nil, err
		}
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}
