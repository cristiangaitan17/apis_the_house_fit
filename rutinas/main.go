package main

import (
	_ "rutinas/routers"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
	"fmt"
)

func main() {
	// Carga el archivo .env si está disponible, sin tumbar la app si falla
	_ = godotenv.Load()

	// Cambiamos '=' por ':=' para declarar 'err' por primera vez aquí
	err := beego.LoadAppConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}
    
    // El resto de tu código de inicialización (beego.Run(), etc.) sigue abajo igual...


	pgUser,_ := web.AppConfig.String("PGuser")
	pgPass,_ := web.AppConfig.String("PGpass")
	pgHost,_ := web.AppConfig.String("PGhost")
	pgPort,_ := web.AppConfig.String("PGport")
	pgDb, _ := web.AppConfig.String("PGdb")
	pgSchema,_ := web.AppConfig.String("PGschema")
	fmt.Printf("PostgreSQL connection string: postgres : //%s : %s@%s :%s/%s?sslmode=disable&search_path=%s\n", pgUser, pgPass, pgHost, pgPort, pgDb, pgSchema)
	
	orm.RegisterDataBase(
		"default",
		"postgres",
		"postgres://"+
			pgUser+":"+
			pgPass+"@"+
			pgHost+":"+
			pgPort+"/"+
			pgDb+
			"?sslmode=disable&search_path="+
			pgSchema,
	)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{
			"Origin",
			"X-Requested-With",
			"Content-Type",
			"Accept",
			"Authorization",
			"X-CsrfToken",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	beego.Run()
}