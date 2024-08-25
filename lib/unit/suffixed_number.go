package unit

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const unmarshalError = "SuffixedNumber.UnmarshalJSON: %v"

type SuffixedNumber float64

func (t *SuffixedNumber) UnmarshalJSON(data []byte) error {
	input := ""

	if err := json.NewDecoder(bytes.NewBuffer(data)).Decode(&input); err != nil {
		return fmt.Errorf(unmarshalError, err)
	}

	if input == "null" || input == `""` {
		return nil
	}

	number, err := convert(input)

	if err != nil {
		return fmt.Errorf(unmarshalError, err)
	}

	*t = SuffixedNumber(number)

	return err
}
