package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/C-dexTeam/codex/internal/config"
	"github.com/C-dexTeam/codex/internal/domains"
	"github.com/C-dexTeam/codex/internal/http"
	"github.com/C-dexTeam/codex/internal/http/middlewares"
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/C-dexTeam/codex/internal/http/server"
	"github.com/C-dexTeam/codex/internal/repositories"
	"github.com/C-dexTeam/codex/internal/services"
	dbadapter "github.com/C-dexTeam/codex/pkg/db_adapters/core"
	validatorService "github.com/C-dexTeam/codex/pkg/validator_service"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Run(cfg *config.Config) {
	// Postgres Client
	connStr := fmt.Sprintf("user=%v password=%v dbname=%v port=%v sslmode=%v host=%v", cfg.DatabaseConfig.Managment.ManagmentUsername, cfg.DatabaseConfig.Managment.ManagmentPassword, cfg.DatabaseConfig.DBName, cfg.DatabaseConfig.Port, cfg.DatabaseConfig.SSLMode, cfg.DatabaseConfig.Host)
	dbAdapter, err := dbadapter.Init(dbadapter.POSTGRESQL, cfg.DatabaseConfig.Driver, connStr, cfg.Application.MigrationsPath)
	if err != nil {
		panic(err)
	}
	conn, err := dbAdapter.ConnectAndMigrateGoose()
	if err != nil {
		panic(err)
	}

	// Repository Initialize
	userRepository := repositories.NewUserRepository(conn)
	userProfileRepository := repositories.NewUserProfileRepository(conn)
	transactionRepository := repositories.NewTransactionRepository(conn)
	roleRepository := repositories.NewRoleRepository(conn)
	languageRepository := repositories.NewLanguageRepository(conn)
	rewardRepository := repositories.NewRewardsRepository(conn)
	attributeRepository := repositories.NewAttributesRepository(conn)

	a, _, b := attributeRepository.Filter(context.Background(), domains.AttributeFilter{}, 10, 1)
	fmt.Println(a, b)

	// Utilities Initialize
	validatorService := validatorService.NewValidatorService()

	// Service Initialize
	allServices := services.CreateNewServices(validatorService, userRepository, userProfileRepository, transactionRepository, roleRepository, languageRepository, rewardRepository, attributeRepository)

	// First Run & Creating Default Admin
	firstRun(conn, allServices.RoleService(), allServices.UserService())

	// Handler Initialize
	handlers := http.NewHandler(allServices)

	// Fiber İnitialize
	fiberServer := server.NewServer(cfg, response.ResponseHandler)

	// Captcha Initialize
	go func() {
		err := fiberServer.Run(handlers.Init(cfg.Application.DevMode, &cfg.RedisConfig, middlewares.InitMiddlewares(cfg)...))
		if err != nil {
			log.Fatalf("Error while running fiber server: %v", err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	fmt.Println("Gracefully shutting down...")
	_ = fiberServer.Shutdown(context.Background())
	fmt.Println("Fiber was successful shutdown.")
}

func firstRun(db *sqlx.DB, roleService domains.IRoleService, userService domains.IUserService) {
	var count int
	err := db.Get(&count, "SELECT COUNT(*) FROM t_users WHERE username = $1", "admin")
	if err != nil {
		log.Fatalf("Error checking for admin user: %v", err)
	}
	if count == 0 {
		adminRole, _ := roleService.GetByName(context.Background(), domains.AdminRole)
		userService.Register(context.Background(), "admin", "admin@gmail.com", "adminadmin", "adminadmin", adminRole.GetID())
	}
}
