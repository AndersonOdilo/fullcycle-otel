package usecase

import (
	"context"

	"github.com/AndersonOdilo/otel/service-a/internal/entity"
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
	TempRepository 		entity.TempRepositoryInterface
	Context 			context.Context
	Tracer            	trace.Tracer
}

func NewGetTempUseCase(tempRepository entity.TempRepositoryInterface, context context.Context, tracer trace.Tracer) *GetTempUseCase {
	return &GetTempUseCase{
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

	ctx, spanTemp := g.Tracer.Start(g.Context, "Chamando Serviço B")
	g.Context = ctx;

	temp, err := g.TempRepository.Get(ctx, cep);
	spanTemp.End();

	if (err != nil) {
		return GetTempOutputDTO{}, err;
	}

	outputDTO := GetTempOutputDTO{
		Celsius:        temp.Celsius,
		Fahrenheit:     temp.Fahrenheit,
		Kelvin:        	temp.Kelvin,
		City:        	temp.City,
	}


	return outputDTO, nil
}


