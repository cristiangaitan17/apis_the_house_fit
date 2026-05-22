package main
						import (
							"github.com/beego/beego/v2/client/orm/migration"
						)

						// DO NOT MODIFY
						type InsertCarrito_20260521_125245 struct {
							migration.Migration
						}

						// DO NOT MODIFY
						func init() {
							m := &InsertCarrito_20260521_125245{}
							m.Created = "20260521_125245"
							
							migration.Register("InsertCarrito_20260521_125245", m)
						}
					   
				// Run the migrations
				func (m *InsertCarrito_20260521_125245) Up() {
					m.SQL("INSERT INTO tienda.carrito (id, usuario_id, activo, \"Fecha_modificacion\", \"fecha_creacion\") VALUES (1, 1, true, '2026-05-21 12:52:45', '2026-05-21 12:52:45');")
					m.SQL("INSERT INTO tienda.carrito (id, usuario_id, activo, \"Fecha_modificacion\", \"fecha_creacion\") VALUES (2, 2, true, '2026-05-21 12:52:45', '2026-05-21 12:52:45');")
					m.SQL("INSERT INTO tienda.carrito (id, usuario_id, activo, \"Fecha_modificacion\", \"fecha_creacion\") VALUES (3, 3, true, '2026-05-21 12:52:45', '2026-05-21 12:52:45');")
					
					
				}
				// Reverse the migrations
				func (m *InsertCarrito_20260521_125245) Down() {
					m.SQL("DELETE FROM tienda.carrito WHERE id IN (1, 2, 3);")
					
					
				}
				