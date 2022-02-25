package rest

import (
	"encoding/base64"
	"strings"
)

type searchReqBody struct {
	Data string `json:"data"`
}

func (rb searchReqBody) Validate() error {
	if len(rb.Data) == 0 {
		return NewErrBadRequest("missing `data`")
	}
	_, err := rb.GetByte()
	if err != nil {
		return NewErrBadRequest("invalid `data`")
	}
	return nil
}

func (rb searchReqBody) GetByte() ([]byte, error) {
	tokens := strings.Split(rb.Data, ",")
	if len(tokens) != 2 {
		return nil, NewErrBadRequest("unexpected number of tokens")
	}
	data, err := base64.StdEncoding.DecodeString(tokens[1])
	if err != nil {
		return nil, NewErrBadRequest(err.Error())
	}
	return data, nil
}
