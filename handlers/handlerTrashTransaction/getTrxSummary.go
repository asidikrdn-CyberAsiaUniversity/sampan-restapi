package handlerTrashTransaction

import (
	"fmt"
	"net/http"
	"sistem-pengelolaan-bank-sampah/dto"

	"github.com/gin-gonic/gin"
)

func (h *handlerTrashTransaction) GetTrxSummary(c *gin.Context) {
	// get trash transaction data from database
	trxSummary, err := h.TrashTransactionRepository.GetTransactionSummary()
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
		Data:    trxSummary,
	}
	c.JSON(http.StatusOK, response)
}
