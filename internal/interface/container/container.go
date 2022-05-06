package container

import (
	"github.com/harlesbayu/kuncie/internal/infrastructure/psql/database"
	"github.com/harlesbayu/kuncie/internal/infrastructure/psql/repositories"
	"github.com/harlesbayu/kuncie/internal/interface/usecase/product"
	"github.com/harlesbayu/kuncie/internal/shared/config"
	"github.com/harlesbayu/kuncie/internal/shared/constants"
)

type Container struct {
	Config         *config.Config
	ProductService product.Service
}

func NewContainer(conf *config.Config) *Container {
	var dbRepo = repositories.NewDBRepository(database.NewDBConnection(&conf.Database))
	if conf.Env == constants.EnvLocal {
		err := database.MigrateAndSeed(dbRepo.DB())
		if err != nil {
			panic(err)
		}
	}

	return &Container{
		Config:         conf,
		ProductService: product.NewService(dbRepo, dbRepo),
	}
}
