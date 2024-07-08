package api

import (
	"bytes"
	"database/sql"
	"encoding/csv"
	"fmt"
	"net/http"
	"time"

	db "github.com/TonyGLL/erp_backend/db/sql"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

/* GET USER */
type getUserRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

// Get User	godoc
// @Summary Get user
// @Description Get user by ID
// @Tags Users
// @Accept json
// @Produce application/json
// @Param			id	path		getUserRequest		true	"User ID"
// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization
// @Security JWT
// @Success 200 {object} db.GetUserRow
// @Failure		401			{string}	gin.H	"StatusUnauthorized"
// @Failure		400			{string}	gin.H	"StatusBadRequest"
// @Failure		404			{string}	gin.H	"StatusNotFound"
// @Failure		500			{string}	gin.H	"StatusInternalServerError"
// @Router /users/{id} [get]
func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

/* GET USERS */
type getUsersRequest struct {
	Page  int32 `form:"page" binding:"required,min=1"`
	Limit int32 `form:"limit" binding:"required,min=5,max=10"`
}

type getUsersResponse struct {
	Total int64           `json:"total"`
	Page  int32           `json:"page"`
	Users []db.GetUserRow `json:"users"`
}

// Get Users	godoc
// @Summary Get users
// @Description Get all users
// @Tags Users
// @Accept json
// @Produce application/json
// @Param	request	query	getUsersRequest true "Query Params"
// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization
// @Security JWT
// @Success 200 {object} getUsersResponse
// @Failure		401			{string}	gin.H	"StatusUnauthorized"
// @Failure		500			{string}	gin.H	"StatusInternalServerError"
// @Router /users [get]
func (server *Server) getUsers(ctx *gin.Context) {
	var req getUsersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetUsersParams{
		Limit:  req.Limit,
		Offset: (req.Page - 1) * req.Limit,
	}

	users, err := server.store.GetUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	total, err := server.store.CountUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := getUsersResponse{
		Total: total,
		Page:  req.Page,
		Users: users,
	}

	ctx.JSON(http.StatusOK, response)
}

/* CREATE USER */
type CreateUserRequest struct {
	RoleID         int    `json:"role_id" binding:"required"`
	Name           string `json:"name" binding:"required"`
	FirstLastName  string `json:"first_last_name" binding:"required"`
	SecondLastName string `json:"second_last_name" binding:"required"`
	Email          string `json:"email" binding:"required"`
	Age            int    `json:"age" binding:"required"`
	Phone          string `json:"phone" binding:"required"`
	Username       string `json:"username" binding:"required"`
	Avatar         string `json:"avatar" binding:"required"`
	Salary         int    `json:"salary" binding:"required"`
	Password       string `json:"password" binding:"required"`
}

// Create User	godoc
// @Summary Create user
// @Description Create user
// @Tags Users
// @Accept json
// @Produce application/json
// @Param   User       body CreateUserRequest true "User"
// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization
// @Security JWT
// @Success 201 {string} gin.H User created successfully.
// @Failure		401			{string}	gin.H	"StatusUnauthorized"
// @Failure		500			{string}	gin.H	"StatusInternalServerError"
// @Router /users [post]
func (server *Server) createUser(ctx *gin.Context) {
	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {

		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		RoleID:         req.RoleID,
		Name:           req.Name,
		FirstLastName:  req.FirstLastName,
		SecondLastName: req.SecondLastName,
		Email:          req.Email,
		Age:            req.Age,
		Phone:          req.Phone,
		Username:       req.Username,
		Avatar:         req.Avatar,
		Salary:         req.Salary,
	}

	user_id, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	passArg := db.CreatePasswordParams{
		UserID:   user_id,
		Password: string(hashedPass),
	}

	err = server.store.CreatePassword(ctx, passArg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully."})
}

/* UPDATE USER */

type UpdateUserRequestBody struct {
	Name           string `json:"name" binding:"required"`
	FirstLastName  string `json:"first_last_name" binding:"required"`
	SecondLastName string `json:"second_last_name" binding:"required"`
	Age            int    `json:"age" binding:"required"`
	Avatar         string `json:"avatar" binding:"required"`
	Salary         int    `json:"salary" binding:"required"`
}

type UpdateUserRequestUri struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

// Update User	godoc
// @Summary Update user
// @Description Update user
// @Tags Users
// @Accept json
// @Produce application/json
// @Param   User       body UpdateUserRequestBody true "User"
// @Param			id	path		UpdateUserRequestUri		true	"User ID"
// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization
// @Security JWT
// @Success 204 {string} gin.H User updated successfully.
// @Failure		401			{string}	gin.H	"StatusUnauthorized"
// @Failure		500			{string}	gin.H	"StatusInternalServerError"
// @Router /users/{id} [put]
func (server *Server) updateUser(ctx *gin.Context) {
	var reqBody UpdateUserRequestBody
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {

		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var reqUri UpdateUserRequestUri
	if err := ctx.ShouldBindUri(&reqUri); err != nil {

		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateUserParams{
		ID:             reqUri.ID,
		Name:           reqBody.Name,
		FirstLastName:  reqBody.FirstLastName,
		SecondLastName: reqBody.SecondLastName,
		Age:            reqBody.Age,
		Avatar:         reqBody.Avatar,
		Salary:         reqBody.Salary,
	}

	err := server.store.UpdateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"message": "User updated successfully."})
}

/* DELETE USER */
type DeleteUserRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

// Delete User	godoc
// @Summary Delete user
// @Description Delete user
// @Tags Users
// @Accept json
// @Produce application/json
// @Param			id	path		DeleteUserRequest		true	"User ID"
// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization
// @Security JWT
// @Success 204 {string} gin.H User deleted successfully.
// @Failure		401			{string}	gin.H	"StatusUnauthorized"
// @Failure		500			{string}	gin.H	"StatusInternalServerError"
// @Router /users/{id} [delete]
func (server *Server) deleteUser(ctx *gin.Context) {
	var req DeleteUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {

		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.DeleteUserParams{
		ID: req.ID,
	}

	err := server.store.DeleteUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"message": "User deleted successfully."})
}

/* DOWNLOAD CSV */
// Download CSV	godoc
// @Summary Download CSV
// @Description Download CSV
// @Tags Users
// @Accept json
// @Produce application/json
// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization
// @Security JWT
// @Success 200 {string} text/csv
// @Failure		401			{string}	gin.H	"StatusUnauthorized"
// @Failure		400			{string}	gin.H	"StatusBadRequest"
// @Failure		500			{string}	gin.H	"StatusInternalServerError"
// @Router /users/download/csv [get]
func (server *Server) downloadUsersCSV(ctx *gin.Context) {
	users, err := server.store.GetUsersForDownload(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Crear un buffer en memoria para el CSV
	buffer := &bytes.Buffer{}
	writer := csv.NewWriter(buffer)

	// Escribir el encabezado del CSV
	header := []string{"ID", "Role ID", "User Type ID", "Name", "First Last Name", "Second Last Name", "Email", "Age", "Phone", "Username", "Avatar", "Cellphone Verification", "Salary", "Deleted", "Created At", "Updated At"}
	if err := writer.Write(header); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo escribir el encabezado en el archivo CSV"})
		return
	}

	// Iterar sobre las filas y escribir los datos en el CSV
	for _, user := range users {
		record := []string{
			fmt.Sprintf("%d", user.ID),
			fmt.Sprintf("%d", user.RoleID),
			user.Name,
			user.FirstLastName,
			user.SecondLastName,
			user.Email,
			fmt.Sprintf("%d", user.Age),
			user.Phone,
			user.Username,
			user.Avatar,
			fmt.Sprintf("%t", user.CellphoneVerification),
			fmt.Sprintf("%f", user.Salary),
			fmt.Sprintf("%t", user.Deleted),
			user.CreatedAt.Format(time.RFC3339),
			user.UpdatedAt.Format(time.RFC3339),
		}
		if err := writer.Write(record); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo escribir la fila en el archivo CSV"})
			return
		}
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al escribir el CSV"})
		return
	}

	// Configurar los encabezados de la respuesta
	filename := fmt.Sprintf("users_%s.csv", time.Now().Format("20060102_150405"))
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	ctx.Data(http.StatusOK, "text/csv", buffer.Bytes())
}
