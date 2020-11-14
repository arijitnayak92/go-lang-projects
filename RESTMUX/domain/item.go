package domain

type Item struct {
	Id       int64  `json:"id,omitempty" bson:"id,omitempty" `
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	Price    uint64 `json:"price,omitempty" bson:"price,omitempty"`
	Quantity uint64 `json:"quantity,omitempty" bson:"quantity,omitempty"`
}

type FiboStruct struct {
	forID string
	value string
}
