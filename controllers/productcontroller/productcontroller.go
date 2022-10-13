package productcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/rama4zis/go-restapi-assignment2/models"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func Index(c *gin.Context) {

	var products []models.Product

	models.DB.Find(&products)
	// panic("test")
	c.JSON(http.StatusOK, gin.H{"data": products})

}

func Show(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong!"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func Create(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid json provided!"})
		return
	}

	models.DB.Create(&product)
	// models.DB.Model(&product).Update("CreatedAt", time.Now())
	c.JSON(http.StatusOK, gin.H{"data": product})
}

func Update(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid json provided!"})
		return
	}

	if models.DB.Model(&product).Where("id = ?", id).Updates(product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Update Error"})
		return
	}
	// models.DB.Model(&product).Update("UpdatedAt", time.Now())
	c.JSON(http.StatusOK, gin.H{"data": "Update Success"})

}

func Delete(c *gin.Context) {
	var product models.Product

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid json provided!"})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Delete Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Delete Success"})

}
