package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"os/user"
	"path"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v1"
)

func init() {
	usr, err := user.Current()
	checkError(err)

	cwd, err := os.Getwd()
	checkError(err)

	viper.AutomaticEnv()
	viper.SetConfigName("icinga-notify")
	viper.SetConfigType("toml")

	viper.AddConfigPath(usr.HomeDir)          // ~
	viper.AddConfigPath(cwd)                  // working directory
	viper.AddConfigPath(path.Dir(os.Args[0])) // program directory
	viper.AddConfigPath("/usr/local/etc/")    // fixed path
	viper.AddConfigPath("/etc/")              // fixed path

	checkError(viper.ReadInConfig())
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func Usage() {
	log.Fatalf("Usage %s host|service", path.Base(os.Args[0]))
}

func main() {
	if len(os.Args) != 2 {
		Usage()
	}

	ntype := os.Args[1]
	if ntype != "host" && ntype != "service" {
		Usage()
	}

	if viper.GetString("USEREMAIL") == "" {
		log.Fatal("Recipient is not set")
	}

	tpl := template.Must(template.ParseFiles("layout.tpl", fmt.Sprintf("%s.tpl", ntype)))
	if tpl.Lookup("Subject") == nil {
		log.Fatal("Template has no Subject defined")
	}
	if tpl.Lookup("Layout") == nil {
		log.Fatal("Template has no Layout part defined")
	}
	if tpl.Lookup("Content") == nil {
		log.Fatal("Template has no Content part defined")
	}
	if tpl.Lookup("Plain") == nil {
		log.Fatal("Template has no Plain part defined")
	}

	m := gomail.NewMessage(gomail.SetCharset("UTF-8"))
	m.SetHeaders(map[string][]string{
		"From":    {m.FormatAddress(viper.GetString("mail.from.mail"), viper.GetString("mail.from.name"))},
		"To":      {viper.GetString("USEREMAIL")},
		"Subject": {expand(tpl, "Subject")},
	})

	m.SetBody("text/html", expand(tpl, "Layout"))
	m.AddAlternative("text/plain", expand(tpl, "Plain"))

	checkError(mail(m))
}

func expand(tpl *template.Template, t string) string {
	var buffer bytes.Buffer

	env := make(map[string]string)
	for _, i := range os.Environ() {
		sep := strings.Index(i, "=")
		env[i[0:sep]] = i[sep+1:]
	}

	tpl.ExecuteTemplate(&buffer, t, env)

	return buffer.String()
}

func mail(message *gomail.Message) error {
	mailer := gomail.NewMailer(
		viper.GetString("mail.server"),
		viper.GetString("mail.user"),
		viper.GetString("mail.password"),
		viper.GetInt("mail.port"),
	)

	return mailer.Send(message)
}
