package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["the_house_fitt/controllers:CarritoController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:CarritoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:CarritoController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:CarritoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:CarritoController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:CarritoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:CarritoController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:CarritoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:CarritoController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:CarritoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:CarritoItemsController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:CarritoItemsController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:CarritoItemsController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:CarritoItemsController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:CarritoItemsController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:CarritoItemsController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:CarritoItemsController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:CarritoItemsController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:CarritoItemsController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:CarritoItemsController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:CategoriasProductoController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:CategoriasProductoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:CategoriasProductoController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:CategoriasProductoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:CategoriasProductoController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:CategoriasProductoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:CategoriasProductoController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:CategoriasProductoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:CategoriasProductoController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:CategoriasProductoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:PedidoItemsController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:PedidoItemsController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:PedidoItemsController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:PedidoItemsController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:PedidoItemsController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:PedidoItemsController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:PedidoItemsController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:PedidoItemsController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:PedidoItemsController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:PedidoItemsController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:PedidosController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:PedidosController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:PedidosController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:PedidosController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:PedidosController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:PedidosController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:PedidosController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:PedidosController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:PedidosController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:PedidosController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:ProductosController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:ProductosController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:ProductosController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:ProductosController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:ProductosController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:ProductosController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:ProductosController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:ProductosController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:ProductosController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:ProductosController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:ResenasProductoController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:ResenasProductoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:ResenasProductoController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:ResenasProductoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:ResenasProductoController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:ResenasProductoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:ResenasProductoController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:ResenasProductoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["the_house_fitt/controllers:ResenasProductoController"] = append(beego.GlobalControllerRouter["the_house_fitt/controllers:ResenasProductoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})
}