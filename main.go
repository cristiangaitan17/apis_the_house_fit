package main

import (
	"fmt"
	"log"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"

	_ "apis_the_house_fit/routers"
)

func init() {
    // Leer configuración de la base de datos
    dbDriver, _ := web.AppConfig.String("dbdriver")
    dbHost, _ := web.AppConfig.String("dbhost")
    dbPort, _ := web.AppConfig.String("dbport")
    dbUser, _ := web.AppConfig.String("dbuser")
    dbPassword, _ := web.AppConfig.String("dbpassword")
    dbName, _ := web.AppConfig.String("dbname")

    // Cadena de conexión para PostgreSQL
    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        dbHost, dbPort, dbUser, dbPassword, dbName)

    // Registrar base de datos con alias "default"
    err := orm.RegisterDataBase("default", dbDriver, connStr)
    if err != nil {
        log.Fatal("Error al conectar a la base de datos:", err)
    }

    // Sincronizar tablas (solo en desarrollo)
    orm.RunSyncdb("default", false, true)

    log.Println("✅ Conexión a PostgreSQL establecida con éxito")
}

func main() {
    web.Run()
}