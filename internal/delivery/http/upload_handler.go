package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	uc "github.com/ifs21014-itdel/log-analyzer/internal/usecase"
	"github.com/ifs21014-itdel/log-analyzer/pkg/jwt"
)

type UploadHandler struct {
	uc *uc.LogAnalysisUsecase
}

func NewUploadHandler(rg *gin.RouterGroup, uc *uc.LogAnalysisUsecase) {
	h := &UploadHandler{uc: uc}
	protected := rg.Group("/upload")
	protected.Use(jwt.AuthMiddleware())
	protected.POST("/", h.Upload)
}

// POST /upload/
func (h *UploadHandler) Upload(c *gin.Context) {
	userID, _ := c.Get("userID")

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file required"})
		return
	}

	// simpan file sementara
	dst := fmt.Sprintf("./tmp/%s", file.Filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
		return
	}

	// panggil usecase untuk parse log concurrent
	err = h.uc.ParseAndSaveLog(dst, userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "file uploaded and analyzed"})
}
