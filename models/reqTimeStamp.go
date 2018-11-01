package models

type ReqestTimestamp struct {
	Limit int    `json:"limit"`
	Start string `json:"start"`
	End   string `json:"end"`
}
