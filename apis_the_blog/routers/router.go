// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"apis_the_blog/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/categorias",
			beego.NSInclude(
				&controllers.CategoriasController{},
			),
		),

		beego.NSNamespace("/Noticias",
			beego.NSInclude(
				&controllers.NoticiasController{},
			),
		),

		beego.NSNamespace("/articulos_secciones",
			beego.NSInclude(
				&controllers.ArticulosSeccionesController{},
			),
		),

		beego.NSNamespace("/Nutricion",
			beego.NSInclude(
				&controllers.NutricionController{},
			),
		),

		beego.NSNamespace("/dieta_comidas",
			beego.NSInclude(
				&controllers.DietaComidasController{},
			),
		),

		beego.NSNamespace("/comentarios_comunidad",
			beego.NSInclude(
				&controllers.ComentariosComunidadController{},
			),
		),

		beego.NSNamespace("/respuestas_comentario",
			beego.NSInclude(
				&controllers.RespuestasComentarioController{},
			),
		),

		beego.NSNamespace("/respuestas_comentario",
			beego.NSInclude(
				&controllers.RespuestasComentarioController{},
			),
		),

		beego.NSNamespace("/categorias",
			beego.NSInclude(
				&controllers.CategoriasController{},
			),
		),

		beego.NSNamespace("/noticias",
			beego.NSInclude(
				&controllers.NoticiasController{},
			),
		),

		beego.NSNamespace("/articulos_secciones",
			beego.NSInclude(
				&controllers.ArticulosSeccionesController{},
			),
		),

		beego.NSNamespace("/nutricion",
			beego.NSInclude(
				&controllers.NutricionController{},
			),
		),

		beego.NSNamespace("/dieta_comidas",
			beego.NSInclude(
				&controllers.DietaComidasController{},
			),
		),

		beego.NSNamespace("/comentarios_comunidad",
			beego.NSInclude(
				&controllers.ComentariosComunidadController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
