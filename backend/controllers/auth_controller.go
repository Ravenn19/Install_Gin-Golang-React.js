package controllers

import (
	"backend/database"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// REGISTER
func Register(c *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format input salah"})
		return
	}

	var existing models.User
	if err := database.DB.Where("email = ?", input.Email).First(&existing).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email sudah digunakan"})
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hash),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal simpan user"})
		return
	}

	// Setelah register langsung login
	c.JSON(http.StatusOK, gin.H{
		"message": "Registrasi berhasil. Lanjut login.",
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}

// LOGIN
func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format input salah"})
		return
	}

	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email tidak ditemukan"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password salah"})
		return
	}

	// Simulasi login sukses
	c.Set("user", user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login berhasil",
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}

// DASHBOARD (akses setelah login)
func Dashboard(c *gin.Context) {
	email := c.Query("email")

	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Selamat datang di Dashboard",
		"user_id":  user.ID,
		"user":     user.Name,
		"userMail": user.Email,
	})
}
