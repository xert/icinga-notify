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

	funcMap := template.FuncMap{
		"env": Env,
	}

	tpl := template.Must(template.New("").Funcs(funcMap).ParseFiles(fmt.Sprintf("%s.tpl", ntype)))
	if tpl.Lookup("Subject") == nil {
		log.Fatal("Template has no Subject defined")
	}
	if tpl.Lookup("Content") == nil {
		log.Fatal("Template has no Content part defined")
	}

	var buffer bytes.Buffer

	err := tpl.ExecuteTemplate(&buffer, "Subject", tpl)
	if err != nil {
		panic(err)
	}
	subject := buffer.String()

	err = tpl.ExecuteTemplate(&buffer, "Content", tpl)
	if err != nil {
		panic(err)
	}
	body := buffer.String()

	m := gomail.NewMessage(gomail.SetCharset("UTF-8"))
	m.SetHeaders(map[string][]string{
		"From":    {m.FormatAddress(viper.GetString("mail.from.mail"), viper.GetString("mail.from.name"))},
		"To":      {viper.GetString("USEREMAIL")},
		"Subject": {subject},
	})

	m.SetBody("text/txt", body)

	checkError(mail(m))
}

func Env(name string) string {
	env := make(map[string]string)
	for _, i := range os.Environ() {
		sep := strings.Index(i, "=")
		env[i[0:sep]] = i[sep+1:]
	}

	if val, ok := env[name]; ok {
		return val
	}

	return fmt.Sprintf("$%s", name)
}

func mail(message *gomail.Message) error {
	mailer := gomail.NewMailer(
		viper.GetString("mail.server"),
		viper.GetString("mail.user"),
		viper.GetString("mail.password"),
		viper.GetInt("mail.port"),
		// gomail.SetTLSConfig(&tls.Config{InsecureSkipVerify: true}),
	)

	return mailer.Send(message)
}
