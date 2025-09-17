package controllers

import (
	"authnz/internal/db"
	"authnz/internal/userService"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Signup godoc
// @Summary register new user
// @Tags auth
// @Accept json
// @Produce json
// @Param input body UserInput true "user data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /signup [post]
func Signup(c *gin.Context) {
	// get body
	var body struct {
		Email    string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read body"})
		return
	}
	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to hash password"})
		return
	}
	// create user in db
	user := userService.User{
		Model:    gorm.Model{},
		Email:    body.Email,
		Password: string(hash),
	}
	result := db.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to create user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// Login godoc
// @Summary login for registered user
// @Tags auth
// @Accept json
// @Produce json
// @Param input body UserInput true "user data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /login [post]
func Login(c *gin.Context) {
	// get body
	var body struct {
		Email    string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read body"})
		return
	}
	// get user by email
	var user userService.User
	result := db.DB.First(&user, "email = ?", body.Email)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email",
		})
		return
	}
	// compare password with stored hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid password",
		})
		return
	}
	// generate jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	// create jwt signed string
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create jwt token",
		})
		return
	}

	// send jwt back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*42*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}

// Validate godoc
// @Summary      validate jwt token
// @Tags         auth
// @Produce      json
// @Success      200 {object} map[string]interface{} "user data"
// @Failure      401 {object} map[string]interface{} "unauthorized"
// @Router       /validate [get]
func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
