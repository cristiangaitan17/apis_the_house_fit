package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["entrenadores/controllers:EntrenadorDocumentosController"] = append(beego.GlobalControllerRouter["entrenadores/controllers:EntrenadorDocumentosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["entrenadores/controllers:EntrenadorDocumentosController"] = append(beego.GlobalControllerRouter["entrenadores/controllers:EntrenadorDocumentosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["entrenadores/controllers:EntrenadorDocumentosController"] = append(beego.GlobalControllerRouter["entrenadores/controllers:EntrenadorDocumentosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["entrenadores/controllers:EntrenadorDocumentosController"] = append(beego.GlobalControllerRouter["entrenadores/controllers:EntrenadorDocumentosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["entrenadores/controllers:EntrenadorDocumentosController"] = append(beego.GlobalControllerRouter["entrenadores/controllers:EntrenadorDocumentosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["entrenadores/controllers:EntrenadoresController"] = append(beego.GlobalControllerRouter["entrenadores/controllers:EntrenadoresController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["entrenadores/controllers:EntrenadoresController"] = append(beego.GlobalControllerRouter["entrenadores/controllers:EntrenadoresController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["entrenadores/controllers:EntrenadoresController"] = append(beego.GlobalControllerRouter["entrenadores/controllers:EntrenadoresController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["entrenadores/controllers:EntrenadoresController"] = append(beego.GlobalControllerRouter["entrenadores/controllers:EntrenadoresController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["entrenadores/controllers:EntrenadoresController"] = append(beego.GlobalControllerRouter["entrenadores/controllers:EntrenadoresController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["entrenadores/controllers:SolicitudesEntrenadorController"] = append(beego.GlobalControllerRouter["entrenadores/controllers:SolicitudesEntrenadorController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["entrenadores/controllers:SolicitudesEntrenadorController"] = append(beego.GlobalControllerRouter["entrenadores/controllers:SolicitudesEntrenadorController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["entrenadores/controllers:SolicitudesEntrenadorController"] = append(beego.GlobalControllerRouter["entrenadores/controllers:SolicitudesEntrenadorController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["entrenadores/controllers:SolicitudesEntrenadorController"] = append(beego.GlobalControllerRouter["entrenadores/controllers:SolicitudesEntrenadorController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["entrenadores/controllers:SolicitudesEntrenadorController"] = append(beego.GlobalControllerRouter["entrenadores/controllers:SolicitudesEntrenadorController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
