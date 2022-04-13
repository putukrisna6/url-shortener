package models

type Redirect struct {
	Code      string `json:"code" msgpack:"code"`
	URL       string `json:"url" msgpack:"url" validate:"required,url"`
	CreatedAt int64  `json:"created_at" msgpack:"created_at"`
}
