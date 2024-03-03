package models

type Product struct {
	ID             string   `json:"id" bson:"_id"`
	Name           string   `json:"name" bson:"name"`
	Description    string   `json:"description" bson:"description"`
	Sku            string   `json:"sku" bson:"sku"`
	Price          float64  `json:"price" bson:"price"`
	Stock          int      `json:"quantity" bson:"quantity"`
	Active         bool     `json:"active" bson:"active"`
	NameUrl        string   `json:"name_url" bson:"name_url"`
	Categorys      []string `json:"categorys" bson:"categorys"`
	Brands         []string `json:"brands" bson:"brands"`
	Images         []string `json:"images" bson:"images"`
	PrincipalImage string   `json:"principal_image" bson:"principal_image"`
}
