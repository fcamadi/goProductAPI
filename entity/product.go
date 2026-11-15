package entity

/*
The json tags assigned to the struct field let’s the json Marshaller know that
this field is found with this particular name in a json file.
*/
type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	IsAvailable bool    `json:"isAvailable"`
}
