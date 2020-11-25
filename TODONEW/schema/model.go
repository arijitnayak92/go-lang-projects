package schema

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Note   string `json:"note"`
	Status bool   `json:"status"`
}
