// Package controllers provides HTTP request handles for user authentication and authorization.
// It implements various endpoints for user management including registration, login, logout,
// and user profile retrieval.
package controllers

import (
	"backend/models"
	"backend/utils"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

// JWT Key user for signing JWT tokens.
// TODO: This should be loaded from environmental variables in production.
var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// `Login` handles user login logic and JWT token generation.
//
// It expects a JSON request body containing the user credentials (email and password).
// Upon successful authentication, it generates a JWT token and sets it to an HTTP-only cookie
// in the response.
//
// Request body:
// {
// 	"email": "johndoe@example.com"
//	"password": "johndoe123"
// }
//
// Responses:
// 	- 200: Successful login.
//	- 400: Invalid request body or credentials.
// 	- 500: Internal server error during JWT token generation.
func Login(c *gin.Context) {

	// Receive user credentials.
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User

	// Verify that user exists.
	models.DB.Where("email = ?", user.Email).First(&existingUser)
	if existingUser.ID == 0 {
		c.JSON(400, gin.H{"error": "User does not exist."})
		return
	}

	// Verify password.
	errHash := utils.CompareHashPassword(user.Password, existingUser.Password)

	if !errHash{
		c.JSON(400, gin.H{"error": "Invalid password."})
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &models.Claims{
		Role: existingUser.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: existingUser.Email,
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Generate JWT token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		c.JSON(500, gin.H{"error": "Could not generate JWT token."})
		return
	}

	// Set JWT token as the cookie in the user's browser.
	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)
	c.JSON(200, gin.H{"success": "User logged in."})
}

// `Register` handles user registration.
//
// The function creates a new user in the system. The password is hashed before
// storage. It prevents duplicate email registration.
//
// Request body:
// {
// 	"name": "John Doe"
//	"email": "johndoe@example.com"
//	"password": "johndoe123"
//	"role": "user"	// Optional, defaults to "user"
// }
//
// Responses:
// 	- 200: Successful registration.
//	- 400: Invalid request body or email already exists.
// 	- 500: Internal server error during password hashing.
func Register(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User

	models.DB.Where("email = ?", user.Email).First(&existingUser)

	if existingUser.ID != 0 {
		c.JSON(400, gin.H{"error": "User already exists."})
		return
	}

	var errHash error
	user.Password, errHash = utils.GenerateHashedPassword(user.Password)

	if errHash != nil {
		c.JSON(500, gin.H{"error": "Could not generate password hash."})
		return
	}

	models.DB.Create(&user)

	c.JSON(200, gin.H{"success": "User created."})
}

// `Home` handles requests to the protected Home page endpoint.
//
// This endpoint requires authentication via JWT token cookie.
// It verifies the user's role and returns the role information if
// the user is authorized.
//
// Required: Valid JWT token in the cookie named "token".
//
// Responses:
// 	- 200: Successfully accessed home page, returns user role.
//	- 401: Unauthorized access or invalid token.
func Home(c *gin.Context) {

	cookie, err := c.Cookie("token")

	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized."})
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized."})
	}

	if claims.Role != "user" && claims.Role != "admin" {
		c.JSON(401, gin.H{"error": "Unauthorized."})
		return
	}

	c.JSON(200, gin.H{"success": "Home page.", "role": claims.Role})

}

// `Logout` handles user session termination.
//
// It removes the JWT token cookie from the client's browser.
//
// Responses:
// 	- 200: Succesfully logged out.
func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(200, gin.H{"success": "User logged out."})
}

// `Me` retrieves the current authenticated user's profile information.
//
// This endpoint requires authentication via JWT token cookie.
// It returns the user's profile information excluding sensitive data (like
// passwords). The user is identified by the email stored in the JWT token
// claims.
//
// Required: Valid JWT token in cookie named "token".
//
// Responses:
// 	- 200: Success, returns user profile:
// 	{
//		"id": 1,
//		"name": "John Doe"
//		"email": "johndoe@example.com"
//		"role": "user"
// 	}
// 	- 400: Unauthorized access or invalid token.
//	- 404: User not found in database.
func Me(c *gin.Context) {

	cookie, err := c.Cookie("token")

	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized."})
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	// Verify role
	if claims.Role != "user" && claims.Role != "admin" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			return
	}

	var user models.User

	if err := models.DB.Where("email = ?", claims.Subject).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found."})
		return
	}

	// Return user information
	c.JSON(200, gin.H{
		"id": user.ID,
		"name": user.Name,
		"email": user.Email,
		"role": user.Role,
	})
}
