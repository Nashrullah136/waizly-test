package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/en"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"reflect"
	"strings"
	"test-backend-1/migrations"
	"test-backend-1/modules/actor"
	"test-backend-1/modules/authentication"
	"test-backend-1/repositories"
	"test-backend-1/utils/db"
	"test-backend-1/utils/translate"
	"time"
)

func registerTranslator() error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		if err := en.RegisterDefaultTranslations(v, translate.DefaultTranslator()); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}

	if err := registerTranslator(); err != nil {
		panic(err.Error())
	}

	engine := gin.Default()

	var dbConn *gorm.DB
	var err error
	for i := 0; i < 10; i++ {
		dbConn, err = db.DefaultConnection()
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		panic(err.Error())
	}

	migrations.Migrate(dbConn)

	actorRepo := repositories.NewActorRepository(dbConn)
	roleRepo := repositories.NewRoleRepository(dbConn)

	migrations.Seed(dbConn, actorRepo, roleRepo)

	actorRoute := actor.NewRoute(actorRepo, roleRepo)
	actorRoute.Handle(engine)

	actorUseCase := actor.NewUseCase(actorRepo, roleRepo)
	authRoute := authentication.NewRoute(actorUseCase)
	authRoute.Handle(engine)

	if err := engine.Run(); err != nil {
		panic(err.Error())
	}
}
