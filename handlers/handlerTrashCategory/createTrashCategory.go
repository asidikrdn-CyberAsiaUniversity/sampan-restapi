package handlerTrashCategory

import (
	"net/http"
	"sistem-pengelolaan-bank-sampah/dto"
	"sistem-pengelolaan-bank-sampah/models"

	"github.com/gin-gonic/gin"
)

func (h *handlerTrashCategory) CreateTrashCategory(c *gin.Context) {
	var request dto.CreateTrashCategoryRequest

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

	// create new trash category
	trashCategory := models.MstTrashCategory{
		Category: request.Category,
		Price:    request.Price,
	}

	// save new trash category to database
	addedTrashCategory, err := h.TrashCategoryRepository.CreateTrashCategory(&trashCategory)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// reload data
	newTrashCategory, err := h.TrashCategoryRepository.FindTrashCategoryByID(addedTrashCategory.ID)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// send response
	response := dto.Result{
		Status:  http.StatusCreated,
		Message: "OK",
		Data:    convertTrashCategoryResponse(newTrashCategory),
	}
	c.JSON(http.StatusCreated, response)
}
