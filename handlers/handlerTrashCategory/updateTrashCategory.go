package handlerTrashCategory

import (
	"fmt"
	"net/http"
	"sistem-pengelolaan-bank-sampah/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handlerTrashCategory) UpdateTrashCategory(c *gin.Context) {
	var request dto.UpdateTrashCategoryRequest

	// get request data
	err := c.ShouldBind(&request)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

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

	// update category
	if request.Category != "" {
		trashCategory.Category = request.Category
	}

	// update price
	if request.Price != 0 {
		trashCategory.Price = request.Price
	}

	// send updated trash category to database
	trashCategory, err = h.TrashCategoryRepository.UpdateTrashCategory(trashCategory)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// reload data
	trashCategory, err = h.TrashCategoryRepository.FindTrashCategoryByID(trashCategory.ID)
	if err != nil {
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
