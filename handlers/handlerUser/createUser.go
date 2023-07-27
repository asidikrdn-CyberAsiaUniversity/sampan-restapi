package handlerUser

import (
	"net/http"
	"sistem-pengelolaan-bank-sampah/dto"
	"sistem-pengelolaan-bank-sampah/models"
	"sistem-pengelolaan-bank-sampah/pkg/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *handlerUser) AddStaff(c *gin.Context) {
	var request dto.CreateStaffRequest

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

	// check email
	_, err = h.UserRepository.GetUserByEmailOrPhone(request.Email)
	if err == nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "Email already registered",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// check phone
	_, err = h.UserRepository.GetUserByEmailOrPhone(request.Phone)
	if err == nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "Phone number already registered",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// create new user
	user := models.MstUser{
		ID:              uuid.New(),
		FullName:        request.FullName,
		Email:           request.Email,
		IsEmailVerified: false,
		Phone:           request.Phone,
		IsPhoneVerified: false,
		Address:         request.Address,
		RoleID:          2,
		LoginAccess:     true,
	}

	// hashing password
	user.Password, err = bcrypt.HashingPassword(request.Password)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// get image from context (if exist)
	image, ok := c.Get("image")
	if ok {
		user.Image = image.(string)
	}

	// save new user to database
	addedUser, err := h.UserRepository.CreateUser(&user)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// reload data
	newUser, err := h.UserRepository.FindUserByID(addedUser.ID)
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
		Data:    convertUserResponse(newUser),
	}
	c.JSON(http.StatusCreated, response)
}

func (h *handlerUser) AddCustomer(c *gin.Context) {
	var request dto.CreateCustomerRequest

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

	// check email
	_, err = h.UserRepository.GetUserByEmailOrPhone(request.Email)
	if err == nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "Email already registered",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// check phone
	_, err = h.UserRepository.GetUserByEmailOrPhone(request.Phone)
	if err == nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "Phone number already registered",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// create new user
	user := models.MstUser{
		ID:              uuid.New(),
		FullName:        request.FullName,
		Email:           request.Email,
		IsEmailVerified: false,
		Phone:           request.Phone,
		IsPhoneVerified: false,
		Address:         request.Address,
		RoleID:          3,
		// LoginAccess:     true,
	}

	// get image from context (if exist)
	image, ok := c.Get("image")
	if ok {
		user.Image = image.(string)
	}

	// save new user to database
	addedUser, err := h.UserRepository.CreateUser(&user)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// reload data
	newUser, err := h.UserRepository.FindUserByID(addedUser.ID)
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
		Data:    convertUserResponse(newUser),
	}
	c.JSON(http.StatusCreated, response)
}

func (h *handlerUser) RegisterCustomer(c *gin.Context) {
	var request dto.CreateCustomerRequest

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

	// check email
	_, err = h.UserRepository.GetUserByEmailOrPhone(request.Email)
	if err == nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "Email already registered",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// check phone
	_, err = h.UserRepository.GetUserByEmailOrPhone(request.Phone)
	if err == nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "Phone number already registered",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// create new user
	user := models.MstUser{
		ID:              uuid.New(),
		FullName:        request.FullName,
		Email:           request.Email,
		IsEmailVerified: false,
		Phone:           request.Phone,
		IsPhoneVerified: false,
		Address:         request.Address,
		RoleID:          3,
		LoginAccess:     true,
	}

	// hashing password
	user.Password, err = bcrypt.HashingPassword(request.Password)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// get image from context (if exist)
	image, ok := c.Get("image")
	if ok {
		user.Image = image.(string)
	}

	// save new user to database
	addedUser, err := h.UserRepository.CreateUser(&user)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// reload data
	newUser, err := h.UserRepository.FindUserByID(addedUser.ID)
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
		Data:    convertUserResponse(newUser),
	}
	c.JSON(http.StatusCreated, response)
}
