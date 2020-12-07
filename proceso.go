package main

import(
	"fmt"
)
type Proceso struct {
	Id int
	Contador int 
}
type ListaDeProcesos struct{
	Lista []Proceso
}

func (p *Proceso) MostrarProceso(b bool){
	if b {
		fmt.Println( p.Id,":",p.Contador)
	}
}
func (l *ListaDeProcesos) AñadirProceso(p Proceso){
	l.Lista= append(l.Lista,p)
}
func (l *ListaDeProcesos) MostrarProcesos(b bool){
	
	for i:= 0; i<len(l.Lista);i++{
		go l.Lista[i].MostrarProceso(b)
	}

}

func (l *ListaDeProcesos) remove(p int){
	l1:= ListaDeProcesos{}
	for i:= 0;i<len(l.Lista);i++{
		if l.Lista[i].Id!=p{
			l1.AñadirProceso(l.Lista[i])
		}
		
	}
	l= &l1
}