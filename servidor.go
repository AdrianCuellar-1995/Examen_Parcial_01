package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"strings"
)

var clientes []net.Conn
var Nomclientes []string
var activosClientes []int64
var HistorialMensajes []string

func servidor() {
	s, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleClient(c)
	}
}

func handleClient(c net.Conn) {
	var cadena string
	err := gob.NewDecoder(c).Decode(&cadena)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		if cadena != "" {
			cadenadividida := strings.Split(cadena, "|")
			opc := cadenadividida[0]
			msg := cadenadividida[1]
			user := cadenadividida[2]
			switch opc {
			case "0":
				clientes = append(clientes, c)
				Nomclientes = append(Nomclientes, user)
				activosClientes = append(activosClientes, 1)
				respuesta := "0|Bienvenido " + user
				err2 := gob.NewEncoder(c).Encode(respuesta)
				if err2 != nil {
					fmt.Println(err)
				}
				fmt.Println("Se conecto: ", user)
			case "1":
				enviarMensajes(user, msg)
			case "4":
				fmt.Println("Se desconecto: ", user)
			default:
				fmt.Println("Opcion no valida")
			}
		}
	}
}

func enviarMensajes(usuario string, mensaje string) {
	UserMsg := usuario + ": " + mensaje
	HistorialMensajes = append(HistorialMensajes, UserMsg)
	fmt.Println(UserMsg)
	Destino := "1|" + usuario + ": " + mensaje
	origen := "1|" + "Tú: " + mensaje
	for i, v := range Nomclientes {
		if v == usuario {
			err := gob.NewEncoder(clientes[i]).Encode(origen)
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			err := gob.NewEncoder(clientes[i]).Encode(Destino)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}

func serverVerMensajes() {
	fmt.Println("*********************")
	for _, v := range HistorialMensajes {
		fmt.Println(v)
	}
	fmt.Println("*********************")
}

func main() {
	go servidor()
	var opc int64
	for {
		fmt.Println("1) ver chat")
		fmt.Println("0) salir")
		fmt.Scanln(&opc)
		switch opc {
		case 1:
			serverVerMensajes()
		case 0:
			return
		}
	}

}
