package handlers

import (
	"net/http"
	"portfolio-backend/internal/models"
	"portfolio-backend/internal/repository"
	"portfolio-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type PortfolioHandler struct {
	repo       repository.PortfolioRepository
	contactSvc service.ContactService
}

func NewPortfolioHandler(repo repository.PortfolioRepository, contactSvc service.ContactService) *PortfolioHandler {
	return &PortfolioHandler{
		repo:       repo,
		contactSvc: contactSvc,
	}
}

// --- Technology Handlers ---

func (h *PortfolioHandler) CreateTechnology(c *gin.Context) {
	var tech models.Technology
	if err := c.ShouldBindJSON(&tech); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.CreateTechnology(c.Request.Context(), tech); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create technology"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Technology created successfully"})
}

func (h *PortfolioHandler) GetTechnologies(c *gin.Context) {
	techs, err := h.repo.GetTechnologies(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch technologies"})
		return
	}
	c.JSON(http.StatusOK, techs)
}

func (h *PortfolioHandler) UpdateTechnology(c *gin.Context) {
	id := c.Param("id")
	var tech models.Technology
	if err := c.ShouldBindJSON(&tech); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.UpdateTechnology(c.Request.Context(), id, tech); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Technology updated successfully"})
}

func (h *PortfolioHandler) DeleteTechnology(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.DeleteTechnology(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// --- Experience Handlers ---

func (h *PortfolioHandler) CreateExperience(c *gin.Context) {
	var exp models.Experience
	if err := c.ShouldBindJSON(&exp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.CreateExperience(c.Request.Context(), exp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create experience"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Experience created successfully"})
}

func (h *PortfolioHandler) GetExperience(c *gin.Context) {
	exps, err := h.repo.GetExperiences(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch experiences"})
		return
	}
	c.JSON(http.StatusOK, exps)
}

func (h *PortfolioHandler) UpdateExperience(c *gin.Context) {
	id := c.Param("id")
	var exp models.Experience
	if err := c.ShouldBindJSON(&exp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.UpdateExperience(c.Request.Context(), id, exp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Experience updated successfully"})
}

func (h *PortfolioHandler) DeleteExperience(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.DeleteExperience(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// --- Contact Handlers ---

func (h *PortfolioHandler) SendContact(c *gin.Context) {
	var contact models.ContactMessage
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Delegate business logic to Service Layer
	if err := h.contactSvc.ProcessContactMessage(c.Request.Context(), contact); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process message"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Message received successfully"})
}
