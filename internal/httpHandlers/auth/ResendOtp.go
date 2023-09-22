package auth

import (
	"errors"
	"fmt"
	"net/http"
	"uptime/internal/models/User"
	"uptime/internal/validations/auth"
	"uptime/pkg/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ResendOtp(c *gin.Context) {
	params := auth.ResendOtpValidation{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err.Error())})
		return
	}

	user := User.User{}
	if err := user.Find("email = ?", params.Email); err != nil {
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

// throttle resend request for every 3 minutes (check if there is any middleware)
// replace fmt.Sprintf with err.Error()
// move generate otp method to the otp model.

