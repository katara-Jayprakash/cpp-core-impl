package main

import (
	"fmt"
	"log"
	"net"
)
func do(conn net.Conn){
  //creating a buffer of 1024 memory;
   buf:=	make([]byte,1024)
  _,error :=  conn.Read(buf); //blocking call; 
	if(error!=nil){
		log.Fatal(error);
	}

	//logic i wana do with my or api logic; 
	for i:=0; i<len(buf);i++{
		fmt.Println(string(buf[i]))
	}
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

	
	// accepting the request on listener; 
	conn, err := listner.Accept();  //blocking call; 
	if(err!=nil){
		log.Fatal(err)
	}
    //read , write , close the connection; 
 do(conn)
}


