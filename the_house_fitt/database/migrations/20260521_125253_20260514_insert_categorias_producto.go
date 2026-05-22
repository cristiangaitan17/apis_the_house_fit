package main
						import (
							"github.com/beego/beego/v2/client/orm/migration"
						)

						// DO NOT MODIFY
						type InsertCategoriasProducto_20260521_125253 struct {
							migration.Migration
						}

						// DO NOT MODIFY
						func init() {
							m := &InsertCategoriasProducto_20260521_125253{}
							m.Created = "20260521_125253"
							
							migration.Register("InsertCategoriasProducto_20260521_125253", m)
						}
					   
				// Run the migrations
				func (m *InsertCategoriasProducto_20260521_125253) Up() {
					m.SQL("INSERT INTO tienda.categorias_producto (id_categoria, nombre_categoria, descripcion_categoria) VALUES (1, 'Ropa', 'Categoría de ropa para hombres, mujeres y niños.');")
					m.SQL("INSERT INTO tienda.categorias_producto (id_categoria, nombre_categoria, descripcion_categoria) VALUES (2, 'Electrónica', 'Categoría de dispositivos electrónicos como teléfonos, computadoras y accesorios.');")
					m.SQL("INSERT INTO tienda.categorias_producto (id_categoria, nombre_categoria, descripcion_categoria) VALUES (3, 'Hogar', 'Categoría de productos para el hogar como muebles, decoración y electrodomésticos.');")
					// use m.SQL("CREATE TABLE ...") to make schema update
					
				}
				// Reverse the migrations
				func (m *InsertCategoriasProducto_20260521_125253) Down() {
					// use m.SQL("DROP TABLE ...") to reverse schema update
					m.SQL("DELETE FROM tienda.categorias_producto WHERE id_categoria IN (1, 2, 3);")
					
				}
				