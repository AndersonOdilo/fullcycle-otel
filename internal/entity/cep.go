package entity

import "errors"

type Cep struct {
	valor string
}

func NewCep(valor string) (*Cep, error) {
	cep := &Cep{
		valor: valor,
	}

	err := cep.IsValid()
	if err != nil {
		return nil, err
	}
	return cep, nil
}

func (c *Cep) Get() string {

	return c.valor;
}

func (c *Cep) IsValid() error {

	if (len(c.valor) != 8){

		return errors.New("invalid zipcode");
	}

	return nil;
}