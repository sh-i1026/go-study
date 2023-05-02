//go:build wireinject
// +build wireinject

package registry

import (
	"github.com/google/wire"
	"github.com/study/go-study/clean-architecture/infra/client"
	"github.com/study/go-study/clean-architecture/interface/controller"
	"github.com/study/go-study/clean-architecture/interface/presenter"
	"github.com/study/go-study/clean-architecture/usecase"
	ucClie "github.com/study/go-study/clean-architecture/usecase/client"
)

func InitAddress() *controller.AddressController {
	wire.Build(
		controller.NewAddressController,
		usecase.NewAddressInteractor,
		presenter.NewAddressPresenter,
		client.NewZipCloudImplement,
		wire.Bind(new(usecase.AddressInput), new(*usecase.AddressInteractor)),
		wire.Bind(new(usecase.AddressOutput), new(*presenter.AddressPresenter)),
		wire.Bind(new(ucClie.ZipCloudClient), new(*client.ZipCloudImplement)),
	)
	return nil
}
