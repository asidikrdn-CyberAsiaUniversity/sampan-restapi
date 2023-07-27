package handlerUser

import (
	"fmt"
	"net/http"
	"sistem-pengelolaan-bank-sampah/dto"
	"sistem-pengelolaan-bank-sampah/pkg/helpers"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func (h *handlerUser) DeleteUser(c *gin.Context) {
	// get userid from param
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// get user data
	user, err := h.UserRepository.FindUserByID(id)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		c.JSON(http.StatusNotFound, response)
		return
	}

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

	// extract user data from jwt claims
	userData := claims.(jwt.MapClaims)

	// get roleId from user that access this function
	loginRoleId := uint(userData["roleId"].(float64))

	// if they are not superadmin
	if loginRoleId != 1 {
		// can't delete another role (admin only can delete customer)
		if user.RoleID != 3 {
			response := dto.Result{
				Status:  http.StatusBadRequest,
				Message: "Admin can delete customer only",
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	// delete user image if exist
	if user.Image != "" {
		if !helpers.DeleteFile(user.Image) {
			fmt.Println(err.Error())
		}
	}

	// delete user data
	user, err = h.UserRepository.DeleteUser(user)
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
		Data:    convertUserResponse(user),
	}
	c.JSON(http.StatusOK, response)
}
