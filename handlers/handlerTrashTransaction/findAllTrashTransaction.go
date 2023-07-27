package handlerTrashTransaction

import (
	"math"
	"net/http"
	"sistem-pengelolaan-bank-sampah/dto"
	"sistem-pengelolaan-bank-sampah/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handlerTrashTransaction) FindAllTrashTransaction(c *gin.Context) {
	var (
		trashTransaction      *[]models.TrxTrashCustomer
		err                   error
		totalTrashTransaction int64
		filterQuery           dto.TrashTransactionFilter
	)

	// get filter by UserId
	if c.Query("userId") != "" {
		filterQuery.UserID = c.Query("userId")
	}

	// get filter by TrashId
	if c.Query("trashId") != "" {
		trashId, _ := strconv.Atoi(c.Query("trashId"))
		filterQuery.TrashID = uint(trashId)
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

		// get trash transaction data from database
		trashTransaction, totalTrashTransaction, err = h.TrashTransactionRepository.FindAllTrashTransaction(limit, offset, filterQuery, searchQuery)
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
			TotalData:   totalTrashTransaction,
			TotalPages:  int(math.Ceil(float64(float64(totalTrashTransaction) / float64(limit)))),
			CurrentPage: page,
			Data:        convertMultipleTrashTransactionResponse(trashTransaction),
		}
		c.JSON(http.StatusOK, response)
	} else { // without pagination

		// get trashTransaction data from database
		trashTransaction, totalTrashTransaction, err = h.TrashTransactionRepository.FindAllTrashTransaction(-1, -1, filterQuery, searchQuery)
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
			TotalData:   totalTrashTransaction,
			TotalPages:  1,
			CurrentPage: 1,
			Data:        convertMultipleTrashTransactionResponse(trashTransaction),
		}
		c.JSON(http.StatusOK, response)
	}
}
