package api

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/AndersonOdilo/otel/service-a/configs"
	"github.com/AndersonOdilo/otel/service-a/internal/entity"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)


type TempApiResponse struct {
	Celsius         float64  	`json:"temp_C"`
	Fahrenheit      float64 	`json:"temp_F"`
	Kelvin        	float64 	`json:"temp_K"`
	City        	string 		`json:"city"`
}


type TempRepository struct {
	Configs 	configs.Conf
}

func NewTempRepository(configs configs.Conf) *TempRepository {
	return &TempRepository{
		Configs: configs,
	}
}

func (r *TempRepository) Get(ctx context.Context, cep *entity.Cep) (entity.Temp, error) {
	
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://service-b:8081/temp/"+cep.Get(), nil)
	if err != nil {
		return entity.Temp{}, err;
	}

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return entity.Temp{}, err;
	}

	if (res.StatusCode == http.StatusNotFound){

		return entity.Temp{}, errors.New("can not find zipcode");

	} else if (res.StatusCode == http.StatusUnprocessableEntity) {
		
		return entity.Temp{}, errors.New("invalid zipcode");

	} else if (res.StatusCode != http.StatusOK) {
		
		return entity.Temp{}, errors.New("internal server error");

	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Println(err);
		return entity.Temp{}, nil;
	}

	tempApiResponse := TempApiResponse{}
	err = json.Unmarshal(body, &tempApiResponse)

	if err != nil {
		log.Println(err);
		return entity.Temp{}, nil;
	}

	return entity.Temp{
		Celsius:        tempApiResponse.Celsius,
		Fahrenheit:     tempApiResponse.Fahrenheit,
		Kelvin:        	tempApiResponse.Kelvin,
		City:        	tempApiResponse.City,
	}, nil;

}
