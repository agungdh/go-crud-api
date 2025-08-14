package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/agungdh/go-crud-api/entity"
	"github.com/agungdh/go-crud-api/service"
)

type ProjectHandler struct {
	svc *service.ProjectService
}

func NewProjectHandler(svc *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{svc: svc}
}

/*** DTOs ***/

type createProjectReq struct {
	Name        string     `json:"name" binding:"required,min=1,max=255"`
	Description string     `json:"description" binding:"required"`
	ReleaseDate *time.Time `json:"release_date,omitempty"` // format RFC3339 di JSON
}

type updateProjectReq struct {
	Name        *string     `json:"name,omitempty" binding:"omitempty,min=1,max=255"`
	Description *string     `json:"description,omitempty"`
	ReleaseDate **time.Time `json:"release_date,omitempty"` // pointer to pointer to allow explicit null
}

/*** Handlers ***/

func (h *ProjectHandler) Create(c *gin.Context) {
	var req createProjectReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	p := entity.NewProject(req.Name, req.Description, req.ReleaseDate)
	if err := h.svc.CreateProject(ctx, p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create project"})
		return
	}
	c.JSON(http.StatusCreated, p)
}

func (h *ProjectHandler) GetByUUID(c *gin.Context) {
	id := c.Param("uuid")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	p, err := h.svc.GetProjectByUUID(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "project not found"})
		return
	}
	c.JSON(http.StatusOK, p)
}

func (h *ProjectHandler) Update(c *gin.Context) {
	id := c.Param("uuid")
	u, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}

	var req updateProjectReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	// Ambil dulu entity yang ada (optional, kalau mau partial-validation)
	existing, err := h.svc.GetProjectByUUID(ctx, u.String())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "project not found"})
		return
	}

	// Patch fields
	if req.Name != nil {
		existing.Name = *req.Name
	}
	if req.Description != nil {
		existing.Description = *req.Description
	}
	if req.ReleaseDate != nil {
		// bisa set ke value atau null
		existing.ReleaseDate = *req.ReleaseDate
	}

	if err := h.svc.UpdateProject(ctx, existing); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update project"})
		return
	}
	c.JSON(http.StatusOK, existing)
}

func (h *ProjectHandler) Delete(c *gin.Context) {
	id := c.Param("uuid")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err := h.svc.DeleteProjectByUUID(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete project"})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *ProjectHandler) List(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	q := c.Query("q")

	if limit <= 0 || limit > 200 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	items, total, err := h.svc.ListProjects(ctx, limit, offset, q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list projects"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items":  items,
		"total":  total,
		"limit":  limit,
		"offset": offset,
		"q":      q,
	})
}
