package handlers

import (
	"authnz/internal/userService"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type UserHandlers struct {
	service userService.UserServiceIface
}

func NewUserHandler(s userService.UserServiceIface) *UserHandlers {
	return &UserHandlers{service: s}
}

type UserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *UserHandlers) Signup(c *gin.Context) {
	var input UserInput
	if c.Bind(&input) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read body"})
		return
	}
	if err := h.service.Register(input.Email, input.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user created"})
}

func (h *UserHandlers) Login(c *gin.Context) {
	var input UserInput
	if c.Bind(&input) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read body"})
		return
	}
	user, err := h.service.Authenticate(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create jwt token"})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*42*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "login successful"})
}

func (h *UserHandlers) Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"message": user})
}
