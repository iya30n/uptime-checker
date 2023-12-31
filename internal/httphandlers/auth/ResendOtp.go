package auth

import (
	"errors"
	"net/http"
	"uptime/internal/models"
	"uptime/internal/validations/auth"
	"uptime/pkg/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ResendOtp(c *gin.Context) {
	params := auth.ResendOtpValidation{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": params.Parse(err)})
		return
	}

	user := models.User{}
	if err := user.First("email = ?", params.Email); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "you didn't register yet!"})
			return
		}

		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong!"})
		return
	}

	if user.HasVerified() {
		c.JSON(http.StatusOK, gin.H{"message": "your account already verified."})
		return
	}

	// TODO: use a queue for sending mails with retry
	if err := sendVerificationEmail(user.Email); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong..."})
		return
	}

	c.JSON(200, gin.H{"message": "Verification email sent. please check your email."})
}
