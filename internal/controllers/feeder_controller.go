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

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&feeder); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Error: "invalid request body",
		})
		return
	}

	err := c.service.RegisterFeeder(r.Context(), &feeder) {
		 if err != nil {
			 if err == services.ErrMacAddressRequired {
				 w.WriteHeader(http.StatusBadRequest)
			 } else if err == services.ErrFeederAlreadyExists {
				 w.WriteHeader(http.StatusConflict)
			 } else {
				 w.WriteHeader(http.StatusInternalServerError)
			 }

			 json.NewEncoder(w).Encode(APIResponse{
				 Success: false,
				 Error: err.Error(),
			 })
			 return
		 }
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(APIResponse{
		Success: true,
		Data: feeder,
	})
}
