package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// GetUsers godoc
// @Summary Get all users
// @Description Returns all registered users
// @Tags users
// @Produce json
// @Success 200 {array} user.User
// @Failure 500 {object} map[string]string
// @Router /users [get]
// @Security BearerAuth
func (h *Handler) GetUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch users",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

// Register godoc
// @Summary Register new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body user.RegisterRequest true "User data"
// @Success 201 {object} user.RegisterResponse
// @Failure 400 {object} map[string]string
// @Router /register [post]
func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Register(req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := RegisterResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	c.JSON(201, response)
}

// Login godoc
// @Summary User login
// @Tags users
// @Accept json
// @Produce json
// @Param credentials body user.LoginRequest true "Login credentials"
// @Success 200 {object} user.LoginResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /login [post]
func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	token, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
