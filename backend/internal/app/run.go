package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/C-dexTeam/codex/internal/config"
	"github.com/C-dexTeam/codex/internal/http"
	"github.com/C-dexTeam/codex/internal/http/middlewares"
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/C-dexTeam/codex/internal/http/server"
	repo "github.com/C-dexTeam/codex/internal/repos/out"
	"github.com/C-dexTeam/codex/internal/services"
	validatorService "github.com/C-dexTeam/codex/pkg/validator_service"
	"github.com/pressly/goose"

	_ "github.com/lib/pq"
)

func Run(cfg *config.Config) {
	// Postgres Client
	connStr := fmt.Sprintf("user=%v password=%v dbname=%v port=%v sslmode=%v host=%v", cfg.DatabaseConfig.Managment.ManagmentUsername, cfg.DatabaseConfig.Managment.ManagmentPassword, cfg.DatabaseConfig.DBName, cfg.DatabaseConfig.Port, cfg.DatabaseConfig.SSLMode, cfg.DatabaseConfig.Host)
	conn, err := sql.Open(cfg.DatabaseConfig.Driver, connStr)
	if err != nil {
		return
	}
	if err := conn.Ping(); err != nil && err.Error() != "pq: database system is starting up" {
		panic(err)
	}
	if err := goose.Up(conn, cfg.Application.MigrationsPath); err != nil {
		panic(err)
	}

	// Repos
	queries := repo.New(conn)

	// Utilities Initialize
	validatorService := validatorService.NewValidatorService()

	// Service Initialize
	allServices := services.CreateNewServices(
		validatorService,
		queries,
		conn,
		&cfg.Defaults,
	)

	// First Run & Creating Default Admin
	firstRun(queries, allServices.RoleService(), allServices.UserService(), cfg.Defaults.Roles.RoleAdmin)

	// Handler Initialize
	handlers := http.NewHandler(allServices, &cfg.Defaults)

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

func firstRun(repo *repo.Queries, roleService *services.RoleService, userService *services.UserService, roleAdmin string) {
	// SQL sorgusunu sqlc ile çalıştırıyoruz

	count, err := repo.CountUserByName(context.Background(), sql.NullString{String: "admin", Valid: true})
	if err != nil {
		log.Fatalf("Error checking for admin user: %v", err)
	}
	if count == 0 {
		// Admin rolünü alıyoruz
		adminRole, _ := roleService.GetByName(context.Background(), roleAdmin)

		// Kullanıcıyı kaydediyoruz
		userService.Register(context.Background(), "admin", "admin@gmail.com", "adminadmin", "adminadmin", "admin", "admin", adminRole.ID)
	}
}
