package models

type Feeder struct {
	ID int `json:"id"`
	HouseID int `json:"house_id"`
	MacAddress string `json:"mac_address"`
	Name string `json:"name"`
	PetType string `json:"pet_type"`
	DeletedAt string `json:"deleted_at"`
}
