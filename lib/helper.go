package lib

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	"html/template"
	"log"
	"net/http"
	"os"
)

var t *template.Template

// ParseTemplate parse html file from views folder with data
func ParseTemplate(w http.ResponseWriter, pathfile string, data interface{}) error {
	t, err := t.ParseFiles("views/base.html", "views/"+pathfile+".html")
	if err != nil {
		return err
	}

	err = t.ExecuteTemplate(w, "base", data)
	if err != nil {
		return err
	}
	return nil
}

// GetPort get value from PORT env
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("Err when open PORT env: ", port)
	}

	return ":" + port
}

// CustomMessageErrorValidate custom default validate message
func CustomMessageErrorValidate() {
	translator := en.New()
	uni := ut.New(translator, translator)

	trans, found := uni.GetTranslator("en")
	if !found {
		log.Fatal("translator not found")
	}

	v := validator.New()

	if err := entranslations.RegisterDefaultTranslations(v, trans); err != nil {
		log.Fatal(err)
	}

	_ = v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is required field", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})
}
