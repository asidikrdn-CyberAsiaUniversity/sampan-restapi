package handlerTrashTransaction

import (
	"fmt"
	"net/http"
	"sistem-pengelolaan-bank-sampah/dto"
	"sistem-pengelolaan-bank-sampah/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func (h *handlerTrashTransaction) CreateTrashTransaction(c *gin.Context) {
	var request dto.CreateTrashTransactionRequest
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

	// create new transaction
	newTrx := models.TrxTrashCustomer{
		ID:              uuid.New(),
		Qty:             request.Qty,
		TransactionTime: time.Now(),
	}

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
	newTrx.TrashID = trash.ID
	newTrx.TotalPrice = trash.Price * request.Qty

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
	newTrx.UserID = customer.ID

	// check creator
	// get jwt payload
	claims, ok := c.Get("userData")
	if !ok {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "User data not found",
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	// read user data from jwt payload
	creatorData := claims.(jwt.MapClaims)
	creatorId, err := uuid.Parse(creatorData["id"].(string))
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	// get user data from database
	creator, err := h.TrashTransactionRepository.FindUserByID(creatorId)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusUnauthorized,
			Message: "Creator is Unauthorized",
		}
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	// save creator data
	newTrx.CreatedBy = creator.ID

	// save trash transaction data to database
	trashTransaction, err := h.TrashTransactionRepository.CreateTrashTransaction(&newTrx)
	if err != nil {
		fmt.Println(err.Error())
		response := dto.Result{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		c.JSON(http.StatusNotFound, response)
		return
	}

	// reload transaction data
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
