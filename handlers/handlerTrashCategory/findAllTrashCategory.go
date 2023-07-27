package handlerTrashCategory

import (
	"math"
	"net/http"
	"sistem-pengelolaan-bank-sampah/dto"
	"sistem-pengelolaan-bank-sampah/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handlerTrashCategory) FindAllTrashCategory(c *gin.Context) {
	var (
		trashCategory      *[]models.MstTrashCategory
		err                error
		totalTrashCategory int64
		filterQuery        dto.TrashCategoryFilter
	)

	// get filter by category
	if c.Query("category") != "" {
		filterQuery.Category = c.Query("category")
	}

	// get filter by price
	if c.Query("price") != "" {
		price, _ := strconv.Atoi(c.Query("price"))
		filterQuery.Price = uint(price)
	}

	// get search query
	searchQuery := c.Query("search")

	// with pagination
	if c.Query("page") != "" {
		var (
			limit  int
			offset int
		)

		// get page position
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			response := dto.Result{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		// set limit (if not exist, use default limit -> 5)
		if c.Query("limit") != "" {
			limit, err = strconv.Atoi(c.Query("limit"))
			if err != nil {
				response := dto.Result{
					Status:  http.StatusBadRequest,
					Message: err.Error(),
				}
				c.JSON(http.StatusBadRequest, response)
				return
			}
		} else {
			limit = 5

		}

		// set offset
		if page == 1 {
			offset = -1
		} else {
			offset = (page * limit) - limit
		}

		// get trash category data from database
		trashCategory, totalTrashCategory, err = h.TrashCategoryRepository.FindAllTrashCategory(limit, offset, filterQuery, searchQuery)
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
			Status:      http.StatusOK,
			Message:     "OK",
			TotalData:   totalTrashCategory,
			TotalPages:  int(math.Ceil(float64(float64(totalTrashCategory) / float64(limit)))),
			CurrentPage: page,
			Data:        convertMultipleTrashCategoryResponse(trashCategory),
		}
		c.JSON(http.StatusOK, response)
	} else { // without pagination

		// get trashCategory data from database
		trashCategory, totalTrashCategory, err = h.TrashCategoryRepository.FindAllTrashCategory(-1, -1, filterQuery, searchQuery)
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
			Status:      http.StatusOK,
			Message:     "OK",
			TotalData:   totalTrashCategory,
			TotalPages:  1,
			CurrentPage: 1,
			Data:        convertMultipleTrashCategoryResponse(trashCategory),
		}
		c.JSON(http.StatusOK, response)
	}
}
