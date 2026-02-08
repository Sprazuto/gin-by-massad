package controllers

import (
	"lke-app/forms"
	"lke-app/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController handles user authentication and management
// @title User Controller
// @version 1.0
// @description User authentication and management endpoints
type UserController struct{}

var userModel = new(models.UserModel)
var userForm = new(forms.UserForm)

// getUserID ...
func getUserID(c *gin.Context) (userID int64) {
	//MustGet returns the value for the given key if it exists, otherwise it panics.
	return c.MustGet("userID").(int64)
}

// Login User
// @Summary User login
// @Description Authenticate user and return access token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param loginForm body forms.LoginForm true "Login credentials"
// @Success 200 {object} map[string]interface{} "Successfully logged in"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 401 {object} map[string]interface{} "Invalid credentials"
// @Router /v1/user/login [post]
func (ctrl UserController) Login(c *gin.Context) {
	var loginForm forms.LoginForm

	if validationErr := c.ShouldBindJSON(&loginForm); validationErr != nil {
		message := userForm.Login(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	user, token, err := userModel.Login(loginForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Invalid login details"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged in", "user": user, "token": token})
}

// Register User
// @Summary User registration
// @Description Register a new user account
// @Tags Authentication
// @Accept json
// @Produce json
// @Param registerForm body forms.RegisterForm true "User registration data"
// @Success 200 {object} map[string]interface{} "Successfully registered"
// @Failure 400 {object} map[string]interface{} "Invalid request or validation failed"
// @Router /v1/user/register [post]
func (ctrl UserController) Register(c *gin.Context) {
	var registerForm forms.RegisterForm

	if validationErr := c.ShouldBindJSON(&registerForm); validationErr != nil {
		message := userForm.Register(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	user, err := userModel.Register(registerForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully registered", "user": user})
}

// Logout User
// @Summary User logout
// @Description Logout user and invalidate access token
// @Tags Authentication
// @Produce json
// @Security Bearer
// @Success 200 {object} map[string]interface{} "Successfully logged out"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Router /v1/user/logout [get]
func (ctrl UserController) Logout(c *gin.Context) {

	au, err := authModel.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "User not logged in"})
		return
	}

	deleted, delErr := authModel.DeleteAuth(au.AccessUUID)
	if delErr != nil || deleted == 0 { //if any goes wrong
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
