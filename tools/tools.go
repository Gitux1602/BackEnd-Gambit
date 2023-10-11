package tools

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func FechaMySQL() string {

	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d", //Muestra en pantalla pero su salida es un string
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

}

func EscapeString(t string) string {
	//Funcion de ayuda para que en los datos a insertar si es que tienen comillas o \ los quite y no tener porblemas con las sentencias SQL
	desc := strings.ReplaceAll(t, "'", "") //Comilla simple("'") por nada("")
	desc = strings.ReplaceAll(desc, "\"", "")
	return desc
}

func ArmoSentencia(s string, fieldName string, typeField string, ValueN int, ValueF float64, ValueS string) string {
	if (typeField == "S" && len(ValueS) == 0) || //llega un string y si esta vacío entra ó
		(typeField == "F" && ValueF == 0) || //llega un float y el valor es 0 ó
		(typeField == "N" && ValueN == 0) { //llega un entero y el valor es 0
		return s //Así como llega la sentencia la regreso sin modificarle nada
	}

	if !strings.HasSuffix(s, "SET ") {
		s += ", "
	}

	switch typeField {
	case "S":
		s += fieldName + " = '" + EscapeString(ValueS) + "'"
	case "N":
		s += fieldName + " = " + strconv.Itoa(ValueN)
	case "F":
		s += fieldName + " = " + strconv.FormatFloat(ValueF, 'e', -1, 64)
	}

	return s
}
