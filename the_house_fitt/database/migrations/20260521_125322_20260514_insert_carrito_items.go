package main
						import (
							"github.com/beego/beego/v2/client/orm/migration"
						)

						// DO NOT MODIFY
						type InsertCarritoItems_20260521_125322 struct {
							migration.Migration
						}

						// DO NOT MODIFY
						func init() {
							m := &InsertCarritoItems_20260521_125322{}
							m.Created = "20260521_125322"
							
							migration.Register("InsertCarritoItems_20260521_125322", m)
						}
					   
				// Run the migrations
				func (m *InsertCarritoItems_20260521_125322) Up() {
					m.SQL("INSERT INTO tienda.carrito_items (id_item, carrito_id, producto_id, cantidad, precio_unitario) VALUES (1, 1, 1, 2, 19.99);")
					m.SQL("INSERT INTO tienda.carrito_items (id_item, carrito_id, producto_id, cantidad, precio_unitario) VALUES (2, 1, 2, 1, 499.99);")
					m.SQL("INSERT INTO tienda.carrito_items (id_item, carrito_id, producto_id, cantidad, precio_unitario) VALUES (3, 2, 2, 1, 499.99);")
					m.SQL("INSERT INTO tienda.carrito_items (id_item, carrito_id, producto_id, cantidad, precio_unitario) VALUES (4, 3, 3, 1, 899.99);")
					// use m.SQL("CREATE TABLE ...") to make schema update
					
				}
				// Reverse the migrations
				func (m *InsertCarritoItems_20260521_125322) Down() {
					// use m.SQL("DROP TABLE ...") to reverse schema update
					m.SQL("DELETE FROM tienda.carrito_items WHERE id_item IN (1, 2, 3, 4);")
				}
				