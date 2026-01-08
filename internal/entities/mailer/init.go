package mailer

import (
	"embed"
	"html/template"
	"study-stack/internal/shared/env"
	"sync"

	"github.com/resend/resend-go/v2"
)

const senderEmail = string("Study Stack <no-reply@phardev.online>")

//go:embed templates/*.html
var fs embed.FS

var (
	resendEmailClient *resend.Client
	templates         *template.Template
	once              sync.Once
)

func Init() {
	once.Do(func() {
		templates = template.Must(template.ParseFS(fs, "templates/*.html"))
		resendEmailClient = resend.NewClient(env.Config.RESEND_API_KEY)
	})
}
