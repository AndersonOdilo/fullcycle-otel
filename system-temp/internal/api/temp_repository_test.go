package api

import (
	"testing"

	"github.com/AndersonOdilo/otel/system-temp/configs"
	"github.com/AndersonOdilo/otel/system-temp/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TempRepositoryTestSuite struct {
	suite.Suite
	Config configs.Conf
}

func (suite *TempRepositoryTestSuite) SetupSuite() {

	suite.Config = configs.Conf{
		WeatherApiKey: "376a4a90b4074c07a26165101252206",
	}
}


func TestSuite(t *testing.T) {
	suite.Run(t, new(TempRepositoryTestSuite))
}

func (suite *TempRepositoryTestSuite) TestBuscarTempComLocationValido_RetornaTempComTodosOsParametros() {
	
	tempRepository := NewTempRepository(suite.Config);

	location := entity.Location{
		Logradouro: "Rua Barra Mansa",
		Bairro: "Pinheirinho",
		Cidade: "Francisco Beltrao",
		Estado: "Paraná",
		Cep: "85603-260",
	}


	temp, err := tempRepository.Get(&location);

	suite.NoError(err)
	suite.NotNil(temp)
	suite.NotNil(temp.Celsius)
	
}
