package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

//read , write , close the connection;
func do(conn net.Conn){
  
   buf:=	make([]byte,1024) //creating a buffer of 1024 memory;
    _,error :=  conn.Read(buf); //blocking call; 
	if(error!=nil){
		log.Fatal(error);
	}
  fmt.Println("processing the request")
	time.Sleep(10*time.Second)


  data := []byte("HTTP/1.1 200 OK\r\n\r\n HELLO,WORLD\r\n")
	conn.Write(data); 
	conn.Close();
}

func main(){
	//net package/ listen is responsible for creating socket and binding to port
	listner , error := net.Listen("tcp", ":1729"); 
	if(error!=nil){
		log.Fatal(error);
	}

	for{
		// accepting the request on listener is an  blocking call; 

		fmt.Println("Waiting for client");
		conn, err := listner.Accept();  
	   if(err!=nil){
		log.Fatal(err)
	  }
		go  do(conn)
		 log.Print("client connected")
	}
}


