package handlerTrashCategory

import (
	"net/http"
	"sistem-pengelolaan-bank-sampah/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handlerTrashCategory) DeleteTrashCategory(c *gin.Context) {
	// get trash category id from url
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// get trash category data
	trashCategory, err := h.TrashCategoryRepository.FindTrashCategoryByID(uint(id))
	if err != nil {
		response := dto.Result{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		c.JSON(http.StatusNotFound, response)
		return
	}

	// delete trash category data
	trashCategory, err = h.TrashCategoryRepository.DeleteTrashCategory(trashCategory)
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
