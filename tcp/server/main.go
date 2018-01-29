package main

import (
	"net"
	"fmt"
	"os"
	"io/ioutil"
	"log"
)

func main() {

	tcpAddr,err := net.ResolveTCPAddr("tcp",":1200")
	checkError(err)
	listener,err := net.ListenTCP("tcp",tcpAddr)
	checkError(err)
	for{
		conn,err := listener.Accept()
		if err != nil {
			continue
		}
		log.Printf("recieved local addr : %+v, \n remote addr : %+v: \n",conn.LocalAddr(),conn.RemoteAddr())
		body,_ := ioutil.ReadAll(conn)
		fmt.Println(string(body))
		conn.Close()
	}
}

func checkError(err error){
	if err != nil {
		fmt.Fprint(os.Stderr,"Fatal error : %s",err)
		os.Exit(1)
	}
}
