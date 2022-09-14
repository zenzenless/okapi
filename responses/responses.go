package responses

import (
	"encoding/json"
	"fmt"
)

type (
	Basic struct {
		Code int    `json:"code,string"`
		Msg  string `json:"msg,omitempty"`
	}
)

type resMsg Basic

func (b *Basic) UnmarshalJSON(buf []byte) error {
	var r resMsg
	err := json.Unmarshal(buf, &r)
	if err != nil {
		return err
	}
	b = (*Basic)(&r)
	if b.Code != 0 {
		return fmt.Errorf("%+v", r)
	}

	return nil
}
