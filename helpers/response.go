package helpers

type Failed struct {
	Status bool   `json:"status"`
	Body   string `json:"body"`
}

func ResponseFailed(body string) *Failed {
	res := &Failed{
		Status: false,
		Body:   body,
	}
	return res
}
