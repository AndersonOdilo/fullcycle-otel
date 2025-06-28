package api

import (
	"testing"

	"github.com/AndersonOdilo/otel/service-b/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestBuscarLocationComCepValido_RetornaLocationComTodosOsParametros(t *testing.T) {

	locationRepository := NewLocationRepository();

	cep, _ := entity.NewCep("85603260");

	location, err := locationRepository.Get(cep);

	assert.Nil(t, err);
	assert.Equal(t, "Rua Barra Mansa", location.Logradouro);
	assert.Equal(t, "Pinheirinho", location.Bairro);
	assert.Equal(t, "Francisco Beltrão", location.Cidade);
	assert.Equal(t, "Paraná", location.Estado);
	assert.Equal(t, "85603-260", location.Cep);

}

func TestBuscarLocationComCepInvalido_RetornaErrorCanNotFindZipcode(t *testing.T) {

	locationRepository := NewLocationRepository();

	cep, _ := entity.NewCep("99999999");

	location, err := locationRepository.Get(cep);

	assert.Equal(t, entity.Location{}, location);
	assert.Error(t, err, "can not find zipcode");

}