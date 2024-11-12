package views

type ViewError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
