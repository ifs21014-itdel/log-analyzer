package http

import (
	"github.com/gin-gonic/gin"
	usecaseAuth "github.com/ifs21014-itdel/log-analyzer/internal/usecase"
	usecaseLog "github.com/ifs21014-itdel/log-analyzer/internal/usecase"
)

func NewRouter(authUC *usecaseAuth.AuthUsecase, logUC *usecaseLog.LogAnalysisUsecase) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")

	// Auth endpoints
	NewAuthHandler(api, authUC)

	// Log analysis endpoints (protected)
	NewLogAnalysisHandler(api, logUC)
	NewUploadHandler(api, logUC)

	return r
}
