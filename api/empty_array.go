package api

import (
	"encoding/json"
	"fmt"
)

type ArrayRawMessage []json.RawMessage

func (m *ArrayRawMessage) UnmarshalJSON(data []byte) error {
	fmt.Println("@@@@@", string(data))
	if len(data) == 2 && data[0] == '[' && data[1] == ']' {
		return nil
	}

	return json.Unmarshal(data, m)
}
