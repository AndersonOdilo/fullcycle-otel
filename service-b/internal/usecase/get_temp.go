package usecase

import (
	"context"

	"github.com/AndersonOdilo/otel/service-b/internal/entity"
	"go.opentelemetry.io/otel/trace"
)

type GetTempInputDTO struct {
	Cep    string  `json:"cep"`
}

type GetTempOutputDTO struct {
	Celsius         float64  	`json:"temp_C"`
	Fahrenheit      float64 	`json:"temp_F"`
	Kelvin        	float64 	`json:"temp_K"`
	City        	string 		`json:"city"`
}

type GetTempUseCase struct {
	LocationRepository 	entity.LocationRepositoryInterface
	TempRepository 		entity.TempRepositoryInterface
	Context 			context.Context
	Tracer            	trace.Tracer
}

func NewGetTempUseCase(locationRepository entity.LocationRepositoryInterface, tempRepository entity.TempRepositoryInterface, context context.Context, tracer trace.Tracer) *GetTempUseCase {
	return &GetTempUseCase{
		LocationRepository: locationRepository,
		TempRepository: tempRepository,
		Context: context,
		Tracer: tracer,
	}
}

func (g *GetTempUseCase) Execute(input GetTempInputDTO) (GetTempOutputDTO, error) {


	cep, err := entity.NewCep(input.Cep);

	if (err != nil) {
		return GetTempOutputDTO{}, err;
	}

	ctx, spanLocation := g.Tracer.Start(g.Context, "GetLocation")
	g.Context = ctx;

	location, err := g.LocationRepository.Get(cep);
	spanLocation.End();

	if (err != nil) {
		return GetTempOutputDTO{}, err;
	}

	ctx, spanTemp := g.Tracer.Start(g.Context, "GetTemp")
	g.Context = ctx;

	temp, err := g.TempRepository.Get(&location);
	spanTemp.End();

	if (err != nil) {
		return GetTempOutputDTO{}, err;
	}

	outputDTO := GetTempOutputDTO{
		Celsius:        temp.Celsius,
		Fahrenheit:     temp.Celsius * 1.8 + 32,
		Kelvin:        	temp.Celsius + 273,
		City:        	location.Cidade,
	}


	return outputDTO, nil
}


