package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simpleCrud/models"
	"simpleCrud/repository"
	"strconv"
)

type UserController struct {
	repo repository.UserRepository
}

func NewUserController(repo repository.UserRepository) *UserController {
	return &UserController{repo: repo}
}

// Create a user
func (uc *UserController) Create(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdUser, err := uc.repo.Create(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdUser)
}

// Get all users
func (uc *UserController) GetAll(c *gin.Context) {
	users, err := uc.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func StringToUint(s string) uint {
	i, _ := strconv.Atoi(s)
	return uint(i)
}

// Get user by ID
func (uc *UserController) GetByID(c *gin.Context) {
	id := c.Param("id")
	user, err := uc.repo.GetByID(StringToUint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Update a user
func (uc *UserController) Update(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = StringToUint(id)
	updatedUser, err := uc.repo.Update(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedUser)
}

// Delete a user
func (uc *UserController) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := uc.repo.Delete(StringToUint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
