package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	ErrorFailedToSendVerificationEmail = errors.New("unable to send verification email. Please check the email server configuration and try again")
	ErrorVerificationCodeRequired      = errors.New("a verification code must be provided in the request")
	ErrorEmailVerificationFailed       = errors.New("email verification was unsuccessful. Please ensure the code is correct and try again")
	ErrorVerificationExpiredOrInvalid  = errors.New("the verification code is either expired or invalid. Please request a new code")
)

var (
	SuccessfullySentVerificationEmail = "verification email sent successfully. Please check your inbox"
	SuccessfullyVerifiedEmail         = "your email has been verified successfully. You can now proceed with your account"
)

func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.services.UsersService.List()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if len(users) == 0 {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *Handler) SendVerificationEmail(c *gin.Context) {
	id, err := GetUserIdFromAccessToken(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	success, err := h.services.UsersService.SendEmailVerification(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if !success {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": ErrorFailedToSendVerificationEmail})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": SuccessfullySentVerificationEmail})
}

func (h *Handler) VerifyEmail(c *gin.Context) {
	id, err := GetUserIdFromAccessToken(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	code := c.Param("code")
	if len(code) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorVerificationCodeRequired)
		return
	}

	isVerified, err := h.services.UsersService.VerifyEmail(id, code)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorEmailVerificationFailed)
		return
	}

	if !isVerified {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorVerificationExpiredOrInvalid)
		return
	}

	c.JSON(http.StatusOK, SuccessfullyVerifiedEmail)
}
