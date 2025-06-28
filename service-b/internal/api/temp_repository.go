package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/AndersonOdilo/otel/service-b/configs"
	"github.com/AndersonOdilo/otel/service-b/internal/entity"
)

type WeatherApiResponse struct {
	Current struct {
		TempC            float64 `json:"temp_c"`
	} `json:"current"`
}

type TempRepository struct {
	Configs 	configs.Conf
}

func NewTempRepository(configs configs.Conf) *TempRepository {
	return &TempRepository{
		Configs: configs,
	}
}

func (r *TempRepository) Get(location *entity.Location) (entity.Temp, error) {
	
	params := url.Values{};
    params.Add("key", r.Configs.WeatherApiKey);
    params.Add("q", location.Cidade);
	params.Add("lang", "pt");

	res, err := http.Get("http://api.weatherapi.com/v1/current.json?"+params.Encode());

	if err != nil {
		log.Println(err);
		return entity.Temp{}, nil;
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Println(err);
		return entity.Temp{}, nil;
	}

	weatherApiResponse := WeatherApiResponse{}
	err = json.Unmarshal(body, &weatherApiResponse)

	if err != nil {
		log.Println(err);
		return entity.Temp{}, nil;
	}

	return entity.Temp{
		Celsius: weatherApiResponse.Current.TempC,
	}, nil;

}
