package controllers

import (
	"github.com/gin-gonic/gin"
	"jwt-auth/helpers"
	"jwt-auth/initializers"
	"jwt-auth/models"
	"net/http"
)

func Home(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"Email": user.(models.Users).Email,
	})
}

func RedirectUrl(c *gin.Context) {
	id := c.Param("shorten")
	var url models.URLS
	initializers.DB.First(&url, "id=?", id)

	if url.ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error : failed to read body",
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, url.Link)
}

func CutUrl(c *gin.Context) {
	var body struct {
		URL string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error : failed to read body",
		})
		return
	}
	// Generating a short link
	shortURL := helpers.Shorten(body.URL)

	// Saving data to the database
	err := initializers.DB.Create(&models.URLS{ID: shortURL, Link: body.URL}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to save link",
		})
	}

	// Sending a response to the client
	c.JSON(http.StatusOK, gin.H{
		"short_url": shortURL,
	})
}
