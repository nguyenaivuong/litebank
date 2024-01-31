package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nguyenaivuong/litebank/config"
	"github.com/nguyenaivuong/litebank/internal/handlers"
	"github.com/nguyenaivuong/litebank/internal/middleware"
	"github.com/nguyenaivuong/litebank/internal/models"
	"github.com/nguyenaivuong/litebank/internal/services"
	"github.com/nguyenaivuong/litebank/internal/token"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	AccountIDPath = "/accounts/:id"
)

func Start() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Panic("cannot load config")
	}

	log.Printf(config.Environment)

	// Initialize database and Gorm
	db := initializeDatabase(config)

	// Initialize token maker
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Panic("cannot make token", err)
	}

	// Initialize services
	loginService := services.NewLoginService(db)
	signupService := services.NewSignupService(db)
	accountService := services.NewAccountService(db)
	sessionService := services.NewSessionService(db)
	transferService := services.NewTransferService(db)

	// Initialize handlers
	loginHandler := handlers.NewLoginHandler(loginService)
	signupHandler := handlers.NewSignupHandler(signupService)
	accountHandler := handlers.NewAccountHandler(accountService)
	tokenHandler := handlers.NewTokenHandler(tokenMaker, config, sessionService)
	transferHandler := handlers.NewTransferHandlers(transferService)

	// Initialize Gin router
	r := gin.Default()

	// login
	r.POST("/login", loginHandler.LoginHandler)
	r.POST("/signup", signupHandler.SignupHandler)
	r.POST("/tokens/renew_access", tokenHandler.RenewAccessTokenHandler)

	// account
	authRoutes := r.Group("/").Use(middleware.AuthMiddleware)
	authRoutes.GET(AccountIDPath, accountHandler.GetAccountByID)
	authRoutes.PUT(AccountIDPath, accountHandler.UpdateAccount)
	authRoutes.DELETE(AccountIDPath, accountHandler.DeleteAccount)

	authRoutes.POST("/transfers", transferHandler.TransferHandler)

	// Start server
	log.Fatal(r.Run(":8080"))
}

func initializeDatabase(cfg config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.DBSource), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		panic("Failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.User{}, &models.Account{})
	if err != nil {
		return nil
	}

	return db
}
