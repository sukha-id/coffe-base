package validator

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	id_translations "github.com/go-playground/validator/v10/translations/id"
)

func IsRequestValid(m interface{}) (bool, error) {
	id := id.New()
	uni := ut.New(id, id)

	trans, _ := uni.GetTranslator("id")

	validate := validator.New()
	// register custom translation
	_ = validate.RegisterTranslation("profitloss", trans, RegTransProfitLoss, RegCustomTransContoh)

	// register custom validator
	_ = validate.RegisterValidation("profitloss", RegValidatorContoh)

	validate.RegisterTagNameFunc(customErrorMessage)

	id_translations.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(m)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			// return false, errors.New(err.Field() + " is " + err.Tag())

			return false, errors.New(err.Translate(trans))
		}
	}
	return true, nil
}

func customErrorMessage(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("comment"), ",", 2)[0]
	if name == "-" {
		return ""
	}
	return name
}

func RegTransProfitLoss(ut ut.Translator) error {
	return ut.Add("contoh", "{0} tidak boleh kosong", true) // see universal-translator for details
}

func RegCustomTransContoh(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T("contoh", fe.Field())
	return t
}

func RegValidatorContoh(fl validator.FieldLevel) bool {
	isValid := false
	if fl.Parent().FieldByName("contoh").IsValid() {
		isValid = true
	}

	return isValid
}
