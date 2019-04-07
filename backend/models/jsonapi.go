package models

//
// JSONAPIErrorObject is returned when an API endpoint fails
type JSONAPIErrorObject struct {
	Status int    `json:"status,omitempty"`
	Code   string `json:"code,omitempty"`
}

//
// JSONAPISuccessObject is returned when the API endpoint is ok, and no other data should be returned
type JSONAPISuccessObject struct {
	Status int    `json:"status,omitempty"`
	ID     string `json:"id,omitempty"`
}
