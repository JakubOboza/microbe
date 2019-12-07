package microscope

import "encoding/json"

// types
type messageResponse struct {
	Message string `form:"message" json:"message" xml:"message"`
}

// functions
func (c *client) Ping() (*string, error) {
	res, err := c.getApi("ping")
	if err != nil {
		return nil, err
	}

	msg := messageResponse{}
	json.Unmarshal([]byte(*res), &msg)

	return &msg.Message, nil

}
