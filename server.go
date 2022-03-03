package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
)
type MSG struct {
	FILE 	[]byte `json:"FILE"`
	NAME	string `json:"NAME"`
	CHANNEL string `json:"CHANNEL"`
}


func startServer() {
	
	listener, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("The Server not start: %s", err.Error())
	}
	defer listener.Close()
	log.Printf("started server and lisening on PORT: 9000")

	for{
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("connection not accept: %s", err.Error())
			continue
		}

		go newClient(conn)
		
	
	}

}


func newClient(conn net.Conn){
	var data MSG
	log.Printf("new client: %s", conn.RemoteAddr().String())
	err:= json.NewDecoder(conn).Decode(&data)
	if err != nil {
		fmt.Println("Error al intentar leer el msg", err)
		return
	}


	err= ioutil.WriteFile("./docs/" +data.NAME, data.FILE, 0666)

	if err != nil {
		log.Println("Error al escribir el archivo ", err)
		fmt.Println("Error al escribir el archivo ", err)
		return
	}
	fmt.Printf("Cargado existosamente!!! %s", conn.RemoteAddr().String())

}