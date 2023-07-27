package handlerTrashTransaction

import (
	"fmt"
	"net/http"
	"sistem-pengelolaan-bank-sampah/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *handlerTrashTransaction) UpdateTrashTransaction(c *gin.Context) {
	var request dto.UpdateTrashTransactionRequest

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
	id, _ := uuid.Parse(c.Param("id"))

	// get trash transaction data from database
	trashTransaction, err := h.TrashTransactionRepository.FindTrashTransactionByID(id)
	if err != nil {
		fmt.Println(err.Error())
		response := dto.Result{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		c.JSON(http.StatusNotFound, response)
		return
	}

	// update user
	if request.UserID != "" {
		// check user
		userId, _ := uuid.Parse(request.UserID)
		customer, err := h.TrashTransactionRepository.FindUserByID(userId)
		if err != nil {
			response := dto.Result{
				Status:  http.StatusNotFound,
				Message: err.Error(),
			}
			c.JSON(http.StatusNotFound, response)
			return
		}
		// save customer data
		trashTransaction.UserID = customer.ID
	}

	// update Qty
	if request.Qty != 0 {
		trashTransaction.Qty = request.Qty
		trashTransaction.TotalPrice = request.Qty * trashTransaction.TrashCategory.Price
	}

	// update trash category
	if request.TrashID != 0 {
		// check trash category
		trash, err := h.TrashTransactionRepository.FindTrashCategoryByID(request.TrashID)
		if err != nil {
			response := dto.Result{
				Status:  http.StatusNotFound,
				Message: err.Error(),
			}
			c.JSON(http.StatusNotFound, response)
			return
		}
		// save trash data
		trashTransaction.TrashID = trash.ID
		// update qty
		if request.Qty != 0 {
			trashTransaction.TotalPrice = trash.Price * request.Qty
		} else {
			trashTransaction.TotalPrice = trash.Price * trashTransaction.Qty
		}
	}

	// save updated trash transaction data to database
	trashTransaction, err = h.TrashTransactionRepository.UpdateTrashTransaction(trashTransaction)
	if err != nil {
		fmt.Println(err.Error())
		response := dto.Result{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		c.JSON(http.StatusNotFound, response)
		return
	}

	// reload trash transaction data from database
	trashTransaction, err = h.TrashTransactionRepository.FindTrashTransactionByID(trashTransaction.ID)
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
		Data:    convertTrashTransactionResponse(trashTransaction),
	}
	c.JSON(http.StatusOK, response)
}
