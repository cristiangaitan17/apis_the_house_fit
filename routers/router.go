package routers

import (
    "apis_the_house_fit/controllers"
    "github.com/beego/beego/v2/server/web"
)

func init() {
    // Categorias
    web.Router("/api/v1/categorias", &controllers.CategoriasController{}, "get:Get;post:Post")
    web.Router("/api/v1/categorias/:id", &controllers.CategoriasController{}, "get:GetOne;put:Put;delete:Delete")
    
    // Noticias
    web.Router("/api/v1/noticias", &controllers.NoticiasController{}, "get:Get;post:Post")
    web.Router("/api/v1/noticias/:id", &controllers.NoticiasController{}, "get:GetOne;put:Put;delete:Delete")
    
    // Articulos Secciones
    web.Router("/api/v1/articulos_secciones", &controllers.ArticulosSeccionesController{}, "get:Get;post:Post")
    web.Router("/api/v1/articulos_secciones/:id", &controllers.ArticulosSeccionesController{}, "get:GetOne;put:Put;delete:Delete")
    
    // Nutricion
    web.Router("/api/v1/nutricion", &controllers.NutricionController{}, "get:Get;post:Post")
    web.Router("/api/v1/nutricion/:id", &controllers.NutricionController{}, "get:GetOne;put:Put;delete:Delete")
    
    // Dieta Comidas
    web.Router("/api/v1/dieta_comidas", &controllers.DietaComidasController{}, "get:Get;post:Post")
    web.Router("/api/v1/dieta_comidas/:id", &controllers.DietaComidasController{}, "get:GetOne;put:Put;delete:Delete")
    
    // Comentarios Comunidad
    web.Router("/api/v1/comentarios_comunidad", &controllers.ComentariosComunidadController{}, "get:Get;post:Post")
    web.Router("/api/v1/comentarios_comunidad/:id", &controllers.ComentariosComunidadController{}, "get:GetOne;put:Put;delete:Delete")
    
    // Respuestas Comentario
    web.Router("/api/v1/respuestas_comentario", &controllers.RespuestasComentarioController{}, "get:Get;post:Post")
    web.Router("/api/v1/respuestas_comentario/:id", &controllers.RespuestasComentarioController{}, "get:GetOne;put:Put;delete:Delete")
}