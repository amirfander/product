package models

type Product struct {
	Id          string `json:",omitempty" bson:"_id,omitempty"`
	Title       string
	Category    string
	Tags        []string
	Description string
}
