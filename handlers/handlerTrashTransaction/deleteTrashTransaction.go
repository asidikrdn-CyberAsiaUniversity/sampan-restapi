package handlerTrashTransaction

import (
	"fmt"
	"net/http"
	"sistem-pengelolaan-bank-sampah/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *handlerTrashTransaction) DeleteTrashTransaction(c *gin.Context) {
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

	// delete trash transaction data from database
	trashTransaction, err = h.TrashTransactionRepository.DeleteTrashTransaction(trashTransaction)
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
