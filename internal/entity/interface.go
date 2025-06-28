package entity

type LocationRepositoryInterface interface {
	Get(cep *Cep) (Location, error)
}

type TempRepositoryInterface interface {
	Get(location *Location) (Temp, error)
}

