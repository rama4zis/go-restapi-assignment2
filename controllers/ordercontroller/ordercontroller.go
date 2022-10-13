package ordercontroller

import (
	"net/http"

	"github.com/rama4zis/go-restapi-assignment2/models"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var orders []models.Order

	models.DB.Preload("Items").Find(&orders)
	c.JSON(http.StatusOK, gin.H{"data": orders})
}

func Show(c *gin.Context) {

}

func Create(c *gin.Context) {
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Create(&order)
	// models.DB.Model(&order).Update("CreatedAt", time.Now())
	c.JSON(http.StatusOK, gin.H{"data": order})
}

func Update(c *gin.Context) {
	var order models.Order
	id := c.Param("orderId")

	if err := c.ShouldBindJSON(&order); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid json provided!"})
		return
	}

	if models.DB.Model(&order).Where("id = ?", id).Updates(order).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error updating order"})
		return
	}
	// models.DB.Model(&order).Update("UpdatedAt", time.Now())
	c.JSON(http.StatusOK, gin.H{"Success": "Order updated successfully"})
	c.JSON(http.StatusOK, gin.H{"data": order})
}

func Delete(c *gin.Context) {
	var order models.Order
	id := c.Param("orderId")

	// delete with associated products
	models.DB.Preload("Items").Find(&order, id)

	models.DB.Model(&order).Association("Items").Clear()
	if models.DB.Delete(&order).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error deleting order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Successfully deleted order"})
}
