package main

import (
	"fmt"
	//"time"
	//"./Proceso"
)

func main() {
	var c, id, p int
	var imprimir bool
	lista := ListaDeProcesos{}
	id = 0
	imprimir = false
	for true {
		fmt.Println("1) Agregar Proceso\n2)Mostrar Procesos Corriendo\n3)Eliminar Procesos Especificos\n4)Salir")
		fmt.Scan(&c)
		switch c {
		case 1:
			lista.AñadirProceso(Proceso{id,0})
			id = id + 1
			break
		case 2:
			imprimir = !imprimir
			lista.MostrarProcesos(imprimir)
			break
		case 3:
			fmt.Println("Who?")
			fmt.Scan(&p)
			lista.remove(p)
			break
		case 4:
			return
		default:
			fmt.Println("VERIFIQUE")
			break
		}
	}

}
