package main

import (
	SocketIO "github.com/yosuke-furukawa/socket.io-go-emitter"
	"log"
)

func main() {

	//conn := initialization()
	//
	//defer conn.Close()
	//log.Println("Connected to server...")
	//
	//conn.Write([]byte("-subscribe #tree \n"))
	//
	//for {
	//	message, err := bufio.NewReader(conn).ReadString('\n')
	//
	//	if err != nil {
	//		log.Println(err.Error())
	//	}
	//
	//	message = strings.Trim(message, "\r\n")

	emitter, err := SocketIO.NewEmitter(&SocketIO.EmitterOpts{
		Addr:     "localhost:3000",
		Protocol: "tcp",
	})

	if err != nil {
		log.Println(err)
	}

	e := emitter.To("main")
	e.Emit("message", "I love you!!")
	e.Close()

	//	log.Println(message)
	//}
}
