package main
						import (
							"github.com/beego/beego/v2/client/orm/migration"
						)

						// DO NOT MODIFY
						type InsertResenasProducto_20260521_125328 struct {
							migration.Migration
						}

						// DO NOT MODIFY
						func init() {
							m := &InsertResenasProducto_20260521_125328{}
							m.Created = "20260521_125328"
							
							migration.Register("InsertResenasProducto_20260521_125328", m)
						}
					   
				// Run the migrations
				func (m *InsertResenasProducto_20260521_125328) Up() {
					m.SQL("INSERT INTO tienda.resenas_producto (id_resena, producto_id, usuario_id, puntaje, comentario) VALUES (1, 1, 1, 5, 'Excelente producto.');")
					m.SQL("INSERT INTO tienda.resenas_producto (id_resena, producto_id, usuario_id, puntaje, comentario) VALUES (2, 1, 2, 4, 'Muy bueno.');")
					m.SQL("INSERT INTO tienda.resenas_producto (id_resena, producto_id, usuario_id, puntaje, comentario) VALUES (3, 2, 3, 5, 'Me encanta.');")
					// use m.SQL("CREATE TABLE ...") to make schema update
					
				}
				// Reverse the migrations
				func (m *InsertResenasProducto_20260521_125328) Down() {
					// use m.SQL("DROP TABLE ...") to reverse schema update
					m.SQL("DELETE FROM tienda.resenas_producto WHERE id_resena IN (1, 2, 3);")
				}
				