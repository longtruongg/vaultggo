package models

import "encoding/json"

func UnmarshalKv(data []byte) (Kv, error) {
	var r Kv
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Kv) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Kv struct {
	Key string `json:"a"`
}
