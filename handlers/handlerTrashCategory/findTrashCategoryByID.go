package handlerTrashCategory

import (
	"fmt"
	"net/http"
	"sistem-pengelolaan-bank-sampah/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handlerTrashCategory) FindTrashCategoryByID(c *gin.Context) {
	// get id from url
	id, _ := strconv.Atoi(c.Param("id"))

	// get trash category data from database
	trashCategory, err := h.TrashCategoryRepository.FindTrashCategoryByID(uint(id))
	if err != nil {
		fmt.Println(err.Error())
		response := dto.Result{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		c.JSON(http.StatusNotFound, response)
		return
	}

	// send response
	response := dto.Result{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    convertTrashCategoryResponse(trashCategory),
	}
	c.JSON(http.StatusOK, response)
}
