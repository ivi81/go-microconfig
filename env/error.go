//error.go - содержит описание типов ошибок
package env

type envError string

const (
	ErrorEnvTagFormat envError = "error:env tag format"
)

func (e envError) Error() string {
	return string(e)
}
