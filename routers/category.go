package routers

import (
	"encoding/json"
	"strconv"

	//"github.com/aws/aws-lambda-go/events"

	"BackEnd-Gambit/bd"
	"BackEnd-Gambit/models"
)

func InsertCategory(body string, User string) (int, string) {
	var t models.Category

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if len(t.CategName) == 0 {
		return 400, "Debe especificar el Nombre (Title) de la Categoría"
	}
	if len(t.CategPath) == 0 {
		return 400, "Debe especificar el Path (Ruta) de la Categoría"
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	result, err2 := bd.InsertCategory(t)
	if err2 != nil {
		return 400, "Ocurrió un error al intentar realizar el registro de la categoría " + t.CategName + " > " + err2.Error()
	}
	return 200, "{ CategID: " + strconv.Itoa(int(result)) + "}"
}

func UpdateCategory(body string, User string, id int) (int, string) {

	var t models.Category
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}
	if len(t.CategName) == 0 && len(t.CategPath) == 0 {
		return 400, "Debe especificar el Nombre (Title) y Path (Ruta) de la Categoría"
	}
	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}
	t.CategID = id
	err2 := bd.UpdateCategory(t)
	if err2 != nil {
		return 400, "Ocurrió un error al intentar actualizar la categoría " + strconv.Itoa(t.CategID) + " > " + err2.Error()
	}

	return 200, "Update completado"
}
