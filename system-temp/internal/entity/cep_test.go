package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyValor_WhenCreateANewCep_ThenShouldReceiveAnError(t *testing.T) {

	cep := Cep{}
	assert.Error(t, cep.IsValid(), "invalid zipcode")
}

func TestGivenAnSmallValor_WhenCreateANewCep_ThenShouldReceiveAnError(t *testing.T) {

	cep, err := NewCep("123456");
	assert.Nil(t, cep)
	assert.Error(t, err, "invalid zipcode")
}

func TestGivenAnLargeValor_WhenCreateANewCep_ThenShouldReceiveAnError(t *testing.T) {

	cep, err := NewCep("123456789");
	assert.Nil(t, cep)
	assert.Error(t, err, "invalid zipcode")
}

func TestGivenAValidParams_WhenICallNewCep_ThenIShouldReceiveCreateCepWithAllParams(t *testing.T) {

	cep, err := NewCep("80530908");
	assert.Equal(t, "80530908", cep.Get());
	assert.Nil(t, err)
}
