package response

type Result struct {
	Status bool   `json:"status"`
	Error  *error `json:"error"`
}
