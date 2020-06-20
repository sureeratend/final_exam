package todo

// Customer xxxx
type Customer struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

// Message delete
type Message struct {
	Message string `json:"message"`
}
