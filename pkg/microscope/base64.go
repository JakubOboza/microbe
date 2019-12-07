package microscope

import (
	"encoding/json"
	"errors"
)

// types
type dataReqRes struct {
	Data  string `form:"data" json:"data" xml:"data"`
	Error string `form:"error" json:"error,omitempty" xml:"error" `
}

// functions
func (c *client) Base64Encode(data string) (*string, error) {

	reqData := dataReqRes{Data: data}

	marshaledData, err := json.Marshal(reqData)

	if err != nil {
		return nil, err
	}

	res, err := c.postApi("api/base64/encode", string(marshaledData))
	if err != nil {
		return nil, err
	}

	body := dataReqRes{}

	json.Unmarshal([]byte(*res), &body)

	if body.Error != "" {
		return nil, errors.New(body.Error)
	}

	return &body.Data, nil
}

func (c *client) Base64Decode(data string) (*string, error) {

	reqData := dataReqRes{Data: data}

	marshaledData, err := json.Marshal(reqData)

	if err != nil {
		return nil, err
	}

	res, err := c.postApi("api/base64/decode", string(marshaledData))
	if err != nil {
		return nil, err
	}

	body := dataReqRes{}

	json.Unmarshal([]byte(*res), &body)

	if body.Error != "" {
		return nil, errors.New(body.Error)
	}

	return &body.Data, nil

}
