package api

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/AndersonOdilo/otel/system-temp/internal/entity"
)


type ViaCepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Cidade      string `json:"localidade"`
	Estado      string `json:"estado"`
	Erro 	    string 	`json:"erro"`
}

type LocationRepository struct {
}

func NewLocationRepository() *LocationRepository {
	return &LocationRepository{}
}

func (r *LocationRepository) Get(cep *entity.Cep) (entity.Location, error) {
	
	res, err := http.Get("http://viacep.com.br/ws/" + cep.Get() + "/json/")

	if err != nil {
		log.Println(err);
		return entity.Location{}, err;
	}

	defer res.Body.Close()

	if (res.StatusCode == http.StatusBadRequest){
		return entity.Location{}, errors.New("can not find zipcode");
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Println(err);
		return entity.Location{}, err;
	}

	viaCepResponse := ViaCepResponse{}
	err = json.Unmarshal(body, &viaCepResponse)

	if err != nil {
		log.Println(err);
		return entity.Location{}, err;
	}

	if (viaCepResponse.Erro == "true"){
 		return entity.Location{}, errors.New("can not find zipcode");
	}

	return entity.Location{
		Cep: viaCepResponse.Cep,
		Logradouro: viaCepResponse.Logradouro,
		Bairro: viaCepResponse.Bairro,
		Cidade: viaCepResponse.Cidade,
		Estado: viaCepResponse.Estado,
		Complemento: viaCepResponse.Complemento,
	}, nil

}

