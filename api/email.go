package api

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/TonyGLL/erp_backend/util"
	"github.com/gin-gonic/gin"
)

// Send Email	godoc
// @Summary Send Email
// @Description Send Email
// @Tags SMTP
// @Accept json
// @Produce application/json
// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization
// @Security JWT
// @Success 200 {string} gin.H Email Sent Successfully!
// @Failure		401			{string}	gin.H	"StatusUnauthorized"
// @Failure		500			{string}	gin.H	"StatusInternalServerError"
// @Router /email [post]
func (server *Server) sendEmail(ctx *gin.Context) {
	sender := New(server.config)
	m := NewMessage("Test", "Body message.")
	m.To = []string{"tonygllambia@gmail.com"}
	m.CC = []string{"yaisch89@gmail.com"}

	users, err := server.store.GetUsersForDownload(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
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

	filename := fmt.Sprintf("users_%s.csv", time.Now().Format("20060102_150405"))

	buffers := buffer.Bytes()

	util.SaveToFile(filename, buffers)

	m.AttachFile(filename)

	if err := sender.Send(m, server.config); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = os.Remove(filename)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al borrar csv de local"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Email Sent Successfully!"})
}
