package payperday_test

import (
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/infrastructure"
	fakeDatastore "github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/infrastructure/datastore/mock"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/usecases/payperday"
)

type Mocks struct {
	DataStoreMock *fakeDatastore.DataStoreMock
}

func setupMocks() (payperday.PayPerDay, Mocks) {
	fakeDatastore := fakeDatastore.NewDataStoreMock()

	infra := infrastructure.NewInfrastructureInteractor(fakeDatastore)
	useCase := payperday.NewPayPerDay(infra)

	mocks := Mocks{
		DataStoreMock: fakeDatastore,
	}

	return *useCase, mocks
}
