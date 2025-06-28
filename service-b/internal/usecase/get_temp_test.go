package usecase

import (
	"context"
	"testing"

	"github.com/AndersonOdilo/otel/service-b/configs"
	"github.com/AndersonOdilo/otel/service-b/internal/api"
	"github.com/AndersonOdilo/otel/service-b/internal/entity"
	"github.com/stretchr/testify/suite"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type GetTempUseCaseTestSuite struct {
	suite.Suite
	LocationRepository 	entity.LocationRepositoryInterface
	TempRepository 		entity.TempRepositoryInterface
	Tracer            	trace.Tracer
}

func (suite *GetTempUseCaseTestSuite) SetupSuite() {

	suite.LocationRepository = api.NewLocationRepository();
	suite.TempRepository = api.NewTempRepository(configs.Conf{
		WeatherApiKey: "376a4a90b4074c07a26165101252206",
	});
	suite.Tracer = otel.Tracer("temp-system-api-test")
}


func TestSuite(t *testing.T) {
	suite.Run(t, new(GetTempUseCaseTestSuite))
}

func (suite *GetTempUseCaseTestSuite) TestBuscarTempComCepMenor_RetornaErrorInvalidZipcode() {
	
	getTempUseCase := NewGetTempUseCase(suite.LocationRepository, suite.TempRepository, context.Background(), suite.Tracer);

	inputDTO := GetTempInputDTO{
		Cep: "1234567",
	}

	_, err := getTempUseCase.Execute(inputDTO);

	suite.Error(err, "invalid zipcode");
	
}


func (suite *GetTempUseCaseTestSuite) TestBuscarTempComCepMaior_RetornaErrorInvalidZipcode() {
	
	getTempUseCase := NewGetTempUseCase(suite.LocationRepository, suite.TempRepository, context.Background(), suite.Tracer);

	inputDTO := GetTempInputDTO{
		Cep: "123456789",
	}

	_, err := getTempUseCase.Execute(inputDTO);

	suite.Error(err, "invalid zipcode");
	
}


func (suite *GetTempUseCaseTestSuite) TestBuscarTempComCepInexistente_RetornaErrorInvalidZipcode() {
	

	getTempUseCase := NewGetTempUseCase(suite.LocationRepository, suite.TempRepository, context.Background(),  suite.Tracer);

	inputDTO := GetTempInputDTO{
		Cep: "99999999",
	}

	_, err := getTempUseCase.Execute(inputDTO);

	suite.Error(err, "can not find zipcode");
	
}

func (suite *GetTempUseCaseTestSuite) TestBuscarTempComCepValido_RetornaOutputComTodasTemperaturas() {

	getTempUseCase := NewGetTempUseCase(suite.LocationRepository, suite.TempRepository, context.Background(),  suite.Tracer);

	inputDTO := GetTempInputDTO{
		Cep: "90010170",
	}

	output, err := getTempUseCase.Execute(inputDTO);

	suite.NoError(err);
	suite.NotNil(output);
	suite.NotNil(output.Celsius);
	suite.NotNil(output.Fahrenheit);
	suite.NotNil(output.Kelvin);
	suite.Equal(output.City, "Porto Alegre");
	
}


