package email_repo

import (
	"bytes"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
	"html/template"
	"log"
	"os"
	"simi/internal/domain/model"
	"simi/internal/domain/model/entities"
)

type EmailRepository struct {
}

func NewEmailRepository() EmailRepository {
	return EmailRepository{}
}

func (r EmailRepository) SendEmail(user entities.User, order model.Order) error {
	// Cargar la plantilla desde el archivo HTML
	tmpl, err := template.ParseFiles("./internal/template/email_template.html")
	if err != nil {
		log.Println("Error al cargar la plantilla:", err)
		return err
	}

	// Buffer para almacenar el HTML generado
	var body bytes.Buffer

	// Ejecutar la plantilla con los datos del producto
	if err := tmpl.Execute(&body, order); err != nil {
		log.Println("Error al ejecutar la plantilla:", err)
		return err
	}

	// Crear un nuevo mensaje
	m := gomail.NewMessage()

	// Configurar los encabezados del correo
	m.SetHeader("From", "duartegaston07@gmail.com")
	m.SetHeader("To", user.Email) // destinatario
	m.SetHeader("Subject", "Confirmación de Compra")

	// Adjuntar un archivo, por ejemplo recibo
	//m.Attach("")

	// Cuerpo del correo en formato HTML
	m.SetBody("text/html", body.String())

	// Adjuntar la imagen y asociarla con el cid "logo_nike"
	m.Embed("./internal/image/nike/logo-nike.png", gomail.SetHeader(map[string][]string{
		"Content-ID": {"<logo_nike>"},
	}))

	_ = godotenv.Load(".env")
	// Configurar el servidor SMTP de Gmail
	d := gomail.NewDialer("smtp.gmail.com", 587, "duartegaston07@gmail.com", os.Getenv("EMAIL_PASSWORD"))

	// Enviar el correo
	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
		return err
	}

	log.Println("Correo enviado exitosamente")
	return nil
}
