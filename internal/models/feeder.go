package models

type Feeder struct {
	ID int64 `json:"id"`
	HouseID int64 `json:"house_id"`
	MacAddress string `json:"mac_address"`
	Name string `json:"name"`
	PetType string `json:"pet_type"`
	DeletedAt string `json:"deleted_at"`
}
