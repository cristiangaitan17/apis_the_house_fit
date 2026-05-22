package main

import (
	"fmt"
	"os"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"

	_ "apis_the_blog/routers"
)

func main() {
	// Cargar variables de entorno desde .env
	if err := godotenv.Load(); err != nil {
		fmt.Println("No se encontró archivo .env, usando variables del sistema")
	}

	// Leer variables de entorno
	dbUser := os.Getenv("APIS_THE_BLOG_PGUSER")
	dbPass := os.Getenv("APIS_THE_BLOG_PGPASS")
	dbHost := os.Getenv("APIS_THE_BLOG_PGHOST")
	dbPort := os.Getenv("APIS_THE_BLOG_PGPORT")
	dbName := os.Getenv("APIS_THE_BLOG_PGDB")
	dbSchema := os.Getenv("APIS_THE_BLOG_PGSCHEMA")

	// Valores por defecto
	if dbUser == "" {
		dbUser = "postgres"
	}
	if dbPass == "" {
		dbPass = "root"
	}
	if dbHost == "" {
		dbHost = "127.0.0.1"
	}
	if dbPort == "" {
		dbPort = "5432"
	}
	if dbName == "" {
		dbName = "blog_db"
	}
	if dbSchema == "" {
		dbSchema = "blog"
	}

	// Configuración de la base de datos
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=%s",
		dbHost, dbPort, dbUser, dbPass, dbName, dbSchema)

	fmt.Printf("Conectando a: postgres://%s:%s@%s:%s/%s?search_path=%s\n",
		dbUser, dbPass, dbHost, dbPort, dbName, dbSchema)

	orm.RegisterDataBase("default", "postgres", connStr)

	// Swagger en modo desarrollo
	if web.BConfig.RunMode == "dev" {
		web.BConfig.WebConfig.DirectoryIndex = true
		web.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// CORS
	web.InsertFilter("*", web.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{
			"Origin", "x-requested-with", "content-type",
			"accept", "origin", "authorization", "x-csrftoken",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	web.Run()
}