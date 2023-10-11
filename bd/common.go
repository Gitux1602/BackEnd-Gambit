package bd

import (
	"database/sql"
	"fmt"
	"os"

	"BackEnd-Gambit/models"
	"BackEnd-Gambit/secretm"

	_ "github.com/go-sql-driver/mysql"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}
func DbConnect() error {

	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {

		println(err.Error())
		return err
	}
	err = Db.Ping()
	if err != nil {
		println(err.Error())
	}

	println("Conexion exitosa a la DB")
	return nil
}

func ConnStr(claves models.SecretRDSJson) string {

	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = claves.Username
	authToken = claves.Password
	dbEndpoint = claves.Host
	dbName = "gambit"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true&parseTime=true", dbUser, authToken, dbEndpoint, dbName)
	fmt.Println(dsn)
	return dsn
}

func UserIsAdmin(userUUID string) (bool, string) {

	fmt.Println("Comienza user is Admin")
	err := DbConnect() //Hacemos la conexion a la DB

	if err != nil {
		return false, err.Error()
	}
	defer Db.Close()

	sentencia := "SELECT 1 FROM users WHERE User_UUID='" + userUUID + "' AND User_Status = 0"
	fmt.Println(sentencia)

	rows, err := Db.Query(sentencia)

	if err != nil {
		return false, err.Error()
	}

	var valor string
	rows.Next() //Se para en el primer registro
	rows.Scan(&valor)
	fmt.Println("UserIsAdmin - Ejecucion exitosa - valor devuelto" + valor)

	if valor == "1" {
		return true, ""
	}

	return false, "Este Usuario no es Admin"
}
