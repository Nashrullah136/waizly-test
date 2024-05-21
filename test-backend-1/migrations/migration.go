package migrations

import (
	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/gorm"
	"log"
	"os"
	"path"
	"strconv"
	"test-backend-1/entities"
	"test-backend-1/modules/actor"
	"test-backend-1/repositories"
	"time"
)

func Migrate(db *gorm.DB) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("Can't find current working directory")
		return
	}
	isMigrate, err := strconv.ParseBool(os.Getenv("MIGRATE"))
	if err != nil {
		log.Printf("Invalid MIGRATE valua of enviroment variable")
		return
	}
	if isMigrate {
		migrationSource := &migrate.FileMigrationSource{
			Dir: path.Join(dir, "migrations"),
		}
		sqlDb, _ := db.DB()
		n, err := migrate.Exec(sqlDb, "mysql", migrationSource, migrate.Up)
		if err != nil {
			log.Fatal("Can't do migration!\n", err.Error())
		}
		log.Printf("Success applied %d migrations", n)
	}
}

func Seed(db *gorm.DB, actorRepo repositories.ActorRepositoryInterface,
	roleRepo repositories.RoleRepositoryInterface,
) error {
	isMigrate, err := strconv.ParseBool(os.Getenv("MIGRATE"))
	if err != nil {
		log.Printf("Invalid MIGRATE value of enviroment variable")
		return err
	}
	isSeed, err := strconv.ParseBool(os.Getenv("SEED"))
	if err != nil {
		log.Printf("Invalid SEED valua of enviroment variable")
		return err
	}
	if isSeed || isMigrate {
		superRole := entities.Role{
			ID:       1,
			RoleName: "admin",
		}
		adminRole := entities.Role{
			ID:       2,
			RoleName: "user",
		}
		db.Save(adminRole)
		db.Save(superRole)
		actorUsecase := actor.NewUseCase(actorRepo, roleRepo)
		superAdminUsername := os.Getenv("SUPER_ADMIN_USERNAME")
		if superAdminUsername == "" {
			superAdminUsername = "super_admin"
		}
		superAdminPassword := os.Getenv("SUPER_ADMIN_PASSWORD")
		if superAdminPassword == "" {
			superAdminPassword = superAdminUsername
		}
		superAdmin := &entities.Actor{
			Username:  superAdminUsername,
			Password:  superAdminPassword,
			RoleID:    1,
			Active:    true,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}
		if err := actorUsecase.CreateActor(superAdmin); err != nil {
			panic("Can't create super admin")
		}
	}
	return nil
}
