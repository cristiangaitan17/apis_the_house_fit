// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"the_house_fitt/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/pedidos",
			beego.NSInclude(
				&controllers.PedidosController{},
			),
		),

		beego.NSNamespace("/productos",
			beego.NSInclude(
				&controllers.ProductosController{},
			),
		),

		beego.NSNamespace("/resenas_producto",
			beego.NSInclude(
				&controllers.ResenasProductoController{},
			),
		),

		beego.NSNamespace("/carrito",
			beego.NSInclude(
				&controllers.CarritoController{},
			),
		),

		beego.NSNamespace("/carrito_items",
			beego.NSInclude(
				&controllers.CarritoItemsController{},
			),
		),

		beego.NSNamespace("/categorias_producto",
			beego.NSInclude(
				&controllers.CategoriasProductoController{},
			),
		),

		beego.NSNamespace("/pedido_items",
			beego.NSInclude(
				&controllers.PedidoItemsController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
