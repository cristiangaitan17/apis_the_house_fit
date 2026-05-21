package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"apis_gimnasios/controllers"
)

func init() {
	// SedesGimnasios
	beego.Router("/api/v1/sedes_gimnasios", &controllers.SedesGimnasiosController{}, "get:GetAll;post:Post")
	beego.Router("/api/v1/sedes_gimnasios/:id", &controllers.SedesGimnasiosController{}, "get:GetOne;put:Put;delete:Delete")

	// GimnasioClases
	beego.Router("/api/v1/gimnasio_clases", &controllers.GimnasioClasesController{}, "get:GetAll;post:Post")
	beego.Router("/api/v1/gimnasio_clases/:id", &controllers.GimnasioClasesController{}, "get:GetOne;put:Put;delete:Delete")

	// ResenasGimnasio
	beego.Router("/api/v1/resenas_gimnasio", &controllers.ResenasGimnasioController{}, "get:GetAll;post:Post")
	beego.Router("/api/v1/resenas_gimnasio/:id", &controllers.ResenasGimnasioController{}, "get:GetOne;put:Put;delete:Delete")
}