package controllers

import (
	"encoding/json"
	"net/http"

	"feeder-backend/internal/models"
	"feeder-backend/internal/services"
)

type FeederController struct {
	service services.FeederService
}

func NewFeederController(service services.FeederService) *FeederController {
	return &FeederController{
		service: service,
	}
}

type createFeederRequest struct {
	HouseID int64 `json:house_id`
	MacAddress string `json:mac_address`
	Name string `json:name`
	PetType string `json:pet_type`
}

func (c *FeederController) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req createFeederRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json body", http.StatusBadRequest)
		return
	}

	feeder := &models.Feeder{
		HouseID: req.HouseID,
		MacAddress: req.MacAddress,
		Name: req.Name,
		PetType: req.PetType,
	}

	if err:= c.service.RegisterFeeder(r.Context(), feeder); err != nil {
		http.Error(w, "error atregister feeder: " + err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(feeder)

}
