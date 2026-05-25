package main
						import (
							"github.com/beego/beego/v2/client/orm/migration"
						)

						// DO NOT MODIFY
						type InsertPedidos_20260521_125306 struct {
							migration.Migration
						}

						// DO NOT MODIFY
						func init() {
							m := &InsertPedidos_20260521_125306{}
							m.Created = "20260521_125306"
							
							migration.Register("InsertPedidos_20260521_125306", m)
						}
					   
				// Run the migrations
				func (m *InsertPedidos_20260521_125306) Up() {
					m.SQL("INSERT INTO tienda.pedidos (id_pedido, usuario_id, total_pedido, estado_pedido, fecha_creacion) VALUES (1, 1, 59.97, 'Pendiente', '2026-05-21 12:53:06');")
					m.SQL("INSERT INTO tienda.pedidos (id_pedido, usuario_id, total_pedido, estado_pedido, fecha_creacion) VALUES (2, 2, 499.99, 'Enviado', '2026-05-21 12:53:06');")
					m.SQL("INSERT INTO tienda.pedidos (id_pedido, usuario_id, total_pedido, estado_pedido, fecha_creacion) VALUES (3, 3, 899.99, 'Entregado', '2026-05-21 12:53:06');")
					// use m.SQL("CREATE TABLE ...") to make schema update
					
				}
				// Reverse the migrations
				func (m *InsertPedidos_20260521_125306) Down() {
					// use m.SQL("DROP TABLE ...") to reverse schema update
					m.SQL("DELETE FROM tienda.pedidos WHERE id_pedido IN (1, 2, 3);")
				}
				