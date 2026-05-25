package main
						import (
							"github.com/beego/beego/v2/client/orm/migration"
						)

						// DO NOT MODIFY
						type InsertProductos_20260521_125259 struct {
							migration.Migration
						}

						// DO NOT MODIFY
						func init() {
							m := &InsertProductos_20260521_125259{}
							m.Created = "20260521_125259"
							
							migration.Register("InsertProductos_20260521_125259", m)
						}
					   
				// Run the migrations
				func (m *InsertProductos_20260521_125259) Up() {
					m.SQL("INSERT INTO tienda.productos (id_producto, nombre_producto, descripcion_producto, precio_producto, stock_producto, categoria_id) VALUES (1, 'Camiseta', 'Camiseta de algodón para hombres y mujeres.', 19.99, 100, 1);")
					m.SQL("INSERT INTO tienda.productos (id_producto, nombre_producto, descripcion_producto, precio_producto, stock_producto, categoria_id) VALUES (2, 'Teléfono Inteligente', 'Teléfono inteligente con pantalla táctil y cámara de alta resolución.', 499.99, 50, 2);")
					m.SQL("INSERT INTO tienda.productos (id_producto, nombre_producto, descripcion_producto, precio_producto, stock_producto, categoria_id) VALUES (3, 'Sofá', 'Sofá cómodo y elegante para tu sala de estar.', 899.99, 20, 3);")
					// use m.SQL("CREATE TABLE ...") to make schema update
					
					
				}
				// Reverse the migrations
				func (m *InsertProductos_20260521_125259) Down() {
					// use m.SQL("DROP TABLE ...") to reverse schema update
					m.SQL("DELETE FROM tienda.productos WHERE id_producto IN (1, 2, 3);")
				}
				