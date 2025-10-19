package main

import (
	"log"
	"os"

	"github.com/ifs21014-itdel/log-analyzer/config"
	"github.com/ifs21014-itdel/log-analyzer/internal/delivery/http"
	repo "github.com/ifs21014-itdel/log-analyzer/internal/repository"
	usecase "github.com/ifs21014-itdel/log-analyzer/internal/usecase"
	"github.com/joho/godotenv"
)

func main() {
	// load env
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Validasi JWT_SECRET
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET must be set in environment variables")
	}
	log.Println("JWT_SECRET loaded successfully (length:", len(jwtSecret), ")")

	db, err := config.NewDB()
	if err != nil {
		log.Fatal("db:", err)
	}

	// repo -> usecase -> handler
	userRepo := repo.NewUserRepository(db)
	authUC := usecase.NewAuthUsecase(userRepo)

	logRepo := repo.NewLogAnalysisRepo(db)
	logUC := usecase.NewLogAnalysisUsecase(logRepo)

	// router
	r := http.NewRouter(authUC, logUC)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("listen on :", port)
	r.Run(":" + port)
}
