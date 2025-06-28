package api

import (
	"context"
	"testing"

	"github.com/AndersonOdilo/otel/service-a/configs"
	"github.com/AndersonOdilo/otel/service-a/internal/entity"
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

	cep, _  := entity.NewCep("85603260")

	temp, err := tempRepository.Get(context.Background(), cep);

	suite.NoError(err)
	suite.NotNil(temp)
	suite.NotNil(temp.Celsius)
	
}
