package api

import "fmt"

type CustomErrors struct {
	Text        string `json:"err"`
	HttpCode    int    `json:"-"`
	DescCode    string `json:"desc_code,omitempty"`
	InternalErr error  `json:"-"`
}

func NewError(text string, code int, descCode string) *CustomErrors {
	return &CustomErrors{
		Text:     text,
		HttpCode: code,
		DescCode: descCode,
	}
}

func (c *CustomErrors) Error() string {
	if c.InternalErr != nil {
		return fmt.Sprintf("text: %s, internalErr: %s", c.Text, c.InternalErr.Error())
	}

	return c.Text
}
