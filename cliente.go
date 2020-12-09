package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"
)

var nickname string

func cliente() {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := "0|null|" + nickname
	err = gob.NewEncoder(c).Encode(msg)
	for {
		var msge string
		err = gob.NewDecoder(c).Decode(&msge)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			mens := strings.Split(msge, "|")

			opcion := mens[0]
			mensaje := mens[1]
			if opcion == "0" {
				fmt.Println(mensaje)
			}
			if opcion == "1" {
				fmt.Println(mensaje)
			}
		}
	}
	c.Close()
}

func menuCliente() {
	var opc int64
	for {
		fmt.Println("1) Enviar Mensaje")
		fmt.Println("0) Salir")
		fmt.Scanln(&opc)

		switch opc {
		case 1:
			c, err := net.Dial("tcp", ":9999")
			if err != nil {
				fmt.Println(err)
				return
			}
			var mensaje string

			consoleReader := bufio.NewReader(os.Stdin)
			fmt.Println("chat: ")
			input, _ := consoleReader.ReadString('\n')

			mensF := strings.Split(input, "\n")
			mensaje = mensF[0]
			msg := "1|" + mensaje + "|" + nickname
			err = gob.NewEncoder(c).Encode(msg)
			c.Close()
		case 0:
			return
		}
	}
}

func clienteEND() {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := "4|null|" + nickname
	err = gob.NewEncoder(c).Encode(msg)
	c.Close()
}

func main() {
	fmt.Println("Nickname: ")
	fmt.Scanln(&nickname)
	go cliente()
	menuCliente()
	clienteEND()
	fmt.Println(nickname, "Se ha Desconectado")
}
