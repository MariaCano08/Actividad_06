package main

import (
	"fmt"
	"time"
)

var imprimir = false

type Proceso struct {
	Id       int
	Contador int
}
type ListaDeProcesos struct {
	Lista []Proceso
}

func (p *Proceso) MostrarProceso() {
	for {
		fmt.Println(p.Id, ":", p.Contador)
		time.Sleep(time.Millisecond * 500)
		if !imprimir {
			return
		}
	}
}
func (p *Proceso) hacerPreceso() {
	for {
		p.Contador += 1
		time.Sleep(time.Millisecond * 500)
	}
}
func (l *ListaDeProcesos) AñadirProceso(p Proceso) {
	l.Lista = append(l.Lista, p)

	for i := 0; i < len(l.Lista); i++ {
		go l.Lista[i].hacerPreceso()
	}
}
func (l *ListaDeProcesos) MostrarProcesos() {

	for i := 0; i < len(l.Lista); i++ {
		go l.Lista[i].MostrarProceso()
	}

}

func (l *ListaDeProcesos) remove(p int) {
	l1 := ListaDeProcesos{}
	for i := 0; i < len(l.Lista); i++ {
		if l.Lista[i].Id != p {
			l1.AñadirProceso(l.Lista[i])
		}
	}
	*l = l1
}
