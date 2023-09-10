package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type ViaCEPAddress struct {
	ZipCode      string `json:"cep"`
	Neighborhood string `json:"bairro"`
	City         string `json:"localidade"`
	State        string `json:"uf"`
	AreaCode     string `json:"ddd"`
	Complement   string `json:"complemento"`
	Street       string `json:"logradouro"`
}

type Address struct {
	ZipCode      string `json:"zipCode"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	AreaCode     string `json:"areaCode"`
	Complement   string `json:"complement"`
	Street       string `json:"street"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/addresses/", getAddressHandler)
	http.ListenAndServe(":8080", mux)
}

func getAddressHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	zipCode := r.URL.Path[len("/addresses/"):]

	if zipCode == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	address, err := getAddressByZipCode(zipCode)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if address == nil || address.ZipCode == "" {
		w.WriteHeader(http.StatusNotFound)

		errorResponse := map[string]string{
			"message": "Address not found",
		}

		json.NewEncoder(w).Encode(errorResponse)

		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(address)
}

func getAddressByZipCode(zipCode string) (*Address, error) {
	response, err := http.Get("https://viacep.com.br/ws/" + zipCode + "/json/")

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var viaCepAddress ViaCEPAddress
	json.Unmarshal(body, &viaCepAddress)

	address := Address(viaCepAddress)

	return &address, nil
}
