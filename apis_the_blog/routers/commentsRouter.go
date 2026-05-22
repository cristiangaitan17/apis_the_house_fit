package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["apis_the_blog/controllers:ArticulosSeccionesController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:ArticulosSeccionesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:ArticulosSeccionesController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:ArticulosSeccionesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:ArticulosSeccionesController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:ArticulosSeccionesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:ArticulosSeccionesController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:ArticulosSeccionesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:ArticulosSeccionesController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:ArticulosSeccionesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:CategoriasController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:CategoriasController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:CategoriasController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:CategoriasController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:CategoriasController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:CategoriasController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:CategoriasController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:CategoriasController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:CategoriasController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:CategoriasController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:ComentariosComunidadController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:ComentariosComunidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:ComentariosComunidadController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:ComentariosComunidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:ComentariosComunidadController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:ComentariosComunidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:ComentariosComunidadController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:ComentariosComunidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:ComentariosComunidadController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:ComentariosComunidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:DietaComidasController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:DietaComidasController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:DietaComidasController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:DietaComidasController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:DietaComidasController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:DietaComidasController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:DietaComidasController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:DietaComidasController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:DietaComidasController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:DietaComidasController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:NoticiasController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:NoticiasController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:NoticiasController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:NoticiasController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:NoticiasController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:NoticiasController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:NoticiasController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:NoticiasController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:NoticiasController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:NoticiasController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:NutricionController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:NutricionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:NutricionController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:NutricionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:NutricionController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:NutricionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:NutricionController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:NutricionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:NutricionController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:NutricionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:RespuestasComentarioController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:RespuestasComentarioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:RespuestasComentarioController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:RespuestasComentarioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:RespuestasComentarioController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:RespuestasComentarioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:RespuestasComentarioController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:RespuestasComentarioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apis_the_blog/controllers:RespuestasComentarioController"] = append(beego.GlobalControllerRouter["apis_the_blog/controllers:RespuestasComentarioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
