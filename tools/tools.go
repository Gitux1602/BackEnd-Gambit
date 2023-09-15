package tools

import (
	"fmt"
	"strings"
	"time"
)

func FechaMySQL() string {

	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d", //Muestra en pantalla pero su salida es un string
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

}

func EscapeString(t string) string {
	desc := strings.ReplaceAll(t, "'", "") //Comilla simple("'") por nada("")
	desc = strings.ReplaceAll(desc, "\"", "")
	return desc
}
