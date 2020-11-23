package domain

type Item struct {
	tableName   struct{} `pg:"todo"`
	ID          uint64   `json:"id" pg:"id,pk"`
	Title       string   `json:"title,omitempty" pg:"title"`
	Description string   `json:"description,omitempty" pg:"description"`
	Status      bool     `json:"status,omitempty" pg:"status"`
}
