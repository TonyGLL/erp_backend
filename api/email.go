package api

import (
	"bytes"
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/TonyGLL/erp_backend/util"
	"github.com/gin-gonic/gin"
)

type Sender struct {
	auth smtp.Auth
}

type Message struct {
	To          []string
	CC          []string
	BCC         []string
	Subject     string
	Body        string
	Attachments map[string][]byte
}

func New(config util.Config) *Sender {
	auth := smtp.PlainAuth("", config.SMTP_FROM, config.SMTP_PASSWORD, config.SMTP_HOST)
	return &Sender{auth}
}

func (s *Sender) Send(m *Message, config util.Config) error {
	return smtp.SendMail(fmt.Sprintf("%s:%s", config.SMTP_HOST, config.SMTP_PORT), s.auth, config.SMTP_FROM, m.To, m.ToBytes())
}

func NewMessage(s, b string) *Message {
	return &Message{Subject: s, Body: b, Attachments: make(map[string][]byte)}
}

func (m *Message) AttachFile(src string) error {
	b, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	_, fileName := filepath.Split(src)
	m.Attachments[fileName] = b
	return nil
}

func (m *Message) ToBytes() []byte {
	buf := bytes.NewBuffer(nil)
	withAttachments := len(m.Attachments) > 0
	buf.WriteString(fmt.Sprintf("Subject: %s\n", m.Subject))
	buf.WriteString(fmt.Sprintf("To: %s\n", strings.Join(m.To, ",")))
	if len(m.CC) > 0 {
		buf.WriteString(fmt.Sprintf("Cc: %s\n", strings.Join(m.CC, ",")))
	}

	if len(m.BCC) > 0 {
		buf.WriteString(fmt.Sprintf("Bcc: %s\n", strings.Join(m.BCC, ",")))
	}

	buf.WriteString("MIME-Version: 1.0\n")
	writer := multipart.NewWriter(buf)
	boundary := writer.Boundary()
	if withAttachments {
		buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n", boundary))
		buf.WriteString(fmt.Sprintf("--%s\n", boundary))
	} else {
		buf.WriteString("Content-Type: text/plain; charset=utf-8\n")
	}

	buf.WriteString(m.Body)
	if withAttachments {
		for k, v := range m.Attachments {
			buf.WriteString(fmt.Sprintf("\n\n--%s\n", boundary))
			buf.WriteString(fmt.Sprintf("Content-Type: %s\n", http.DetectContentType(v)))
			buf.WriteString("Content-Transfer-Encoding: base64\n")
			buf.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=%s\n", k))

			b := make([]byte, base64.StdEncoding.EncodedLen(len(v)))
			base64.StdEncoding.Encode(b, v)
			buf.Write(b)
			buf.WriteString(fmt.Sprintf("\n--%s", boundary))
		}

		buf.WriteString("--")
	}

	return buf.Bytes()
}

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
