package auth

import (
	"errors"
	"net/http"
	"time"
	"uptime/internal/models"
	"uptime/internal/validations/auth"
	"uptime/pkg/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Verify(c *gin.Context) {
	// validate email (max:70, type:email)
	// validate code (5 digits)
	params := auth.VerifyValidation{}
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

	// check if user already verified
	if user.HasVerified() {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "your account already verified!"})
		return
	}

	// check if otp validation (using IsValid method on otp model)
	// if otp is not valid, return an error and redirect to resend otp code page (in front).
	otp := models.Otp{}
	if err := otp.First(params.Email, params.Code); err != nil || !otp.IsValid() {
		// TODO: fill the redirect_url
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Code!", "redirect_url": ""})
		return
	}

	// else update otp record used field to TRUE, then find the user and update the verified_at field
	if err := otp.Update(map[string]interface{}{"used": true}); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong!"})
		return
	}

	if err := user.Update(map[string]interface{}{"email_verified_at": time.Now()}); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong!"})
		return
	}

	// then redirect user to the login page.
	// TODO: fill the redirect_url
	c.JSON(http.StatusOK, gin.H{"Message": "Your account has been verified.", "redirect_url": ""})

	// NOTE: for redirection, my solution is to return the url to redirect to the front-end (["redirect_url" => "http://something..."])
	// NOTE: i think it's better to wrap whole queries into a transaction, think about it
}
