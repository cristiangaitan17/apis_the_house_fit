package main
						import (
							"github.com/beego/beego/v2/client/orm/migration"
						)

						// DO NOT MODIFY
						type InsertPedidoItems_20260521_125312 struct {
							migration.Migration
						}

						// DO NOT MODIFY
						func init() {
							m := &InsertPedidoItems_20260521_125312{}
							m.Created = "20260521_125312"
							
							migration.Register("InsertPedidoItems_20260521_125312", m)
						}
					   
				// Run the migrations
				func (m *InsertPedidoItems_20260521_125312) Up() {
					m.SQL("INSERT INTO tienda.pedido_items (id_item, pedido_id, producto_id, cantidad, precio_unitario) VALUES (1, 1, 1, 2, 19.99);")
					m.SQL("INSERT INTO tienda.pedido_items (id_item, pedido_id, producto_id, cantidad, precio_unitario) VALUES (2, 1, 2, 1, 499.99);")
					m.SQL("INSERT INTO tienda.pedido_items (id_item, pedido_id, producto_id, cantidad, precio_unitario) VALUES (3, 2, 2, 1, 499.99);")
					m.SQL("INSERT INTO tienda.pedido_items (id_item, pedido_id, producto_id, cantidad, precio_unitario) VALUES (4, 3, 3, 1, 899.99);")
					// use m.SQL("CREATE TABLE ...") to make schema update
					
				}
				// Reverse the migrations
				func (m *InsertPedidoItems_20260521_125312) Down() {
					// use m.SQL("DROP TABLE ...") to reverse schema update
					m.SQL("DELETE FROM tienda.pedido_items WHERE id_item IN (1, 2, 3, 4);")
				}
				