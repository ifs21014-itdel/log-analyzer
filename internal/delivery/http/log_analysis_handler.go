package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ifs21014-itdel/log-analyzer/internal/domain"
	uc "github.com/ifs21014-itdel/log-analyzer/internal/usecase"
	"github.com/ifs21014-itdel/log-analyzer/pkg/jwt"
)

type LogAnalysisHandler struct {
	uc *uc.LogAnalysisUsecase
}

// ===================== HANDLER =====================
func NewLogAnalysisHandler(rg *gin.RouterGroup, uc *uc.LogAnalysisUsecase) {
	h := &LogAnalysisHandler{uc: uc}
	protected := rg.Group("/analyses")
	protected.Use(jwt.AuthMiddleware())
	protected.POST("/", h.Create)
	protected.GET("/", h.GetAll)
	protected.GET("/:id", h.GetByID)
	protected.PUT("/:id", h.Update)
	protected.DELETE("/:id", h.Delete)
}

// Create new log analysis
func (h *LogAnalysisHandler) Create(c *gin.Context) {
	var input domain.LogAnalysis
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ambil userID dari context JWT
	userID, _ := c.Get("userID")
	input.UserID = userID.(uint)

	if err := h.uc.Create(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, input)
}

// Get all log analyses
func (h *LogAnalysisHandler) GetAll(c *gin.Context) {
	logs, err := h.uc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, logs)
}

// Get by ID
func (h *LogAnalysisHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	log, err := h.uc.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, log)
}

// Update log analysis
func (h *LogAnalysisHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var input domain.LogAnalysis
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.ID = uint(id)

	if err := h.uc.Update(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, input)
}

// Delete log analysis
func (h *LogAnalysisHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	if err := h.uc.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
