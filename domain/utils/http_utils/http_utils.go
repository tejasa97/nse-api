package utils

type Response struct {
	StatusCode  int       `json:"statusCode"`
	Headers     map[string]string  `json:"headers"`
	Body        string    `json:"body"`
}