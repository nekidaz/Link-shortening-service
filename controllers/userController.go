package controllers

import (
	"fmt"
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
	fmt.Printf("Link %s", url.Link)
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
	// Генерируем короткую ссылку
	shortURL := helpers.Shorten(body.URL)

	// Сохраняем данные в базу данных
	initializers.DB.Create(&models.URLS{ID: shortURL, Link: body.URL})

	// Отправляем ответ клиенту
	c.JSON(http.StatusOK, gin.H{
		"short_url": shortURL,
	})
}
