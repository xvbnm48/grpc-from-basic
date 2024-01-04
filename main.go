package main

import (
	"fmt"
	"log"
	"time"

	"github.com/xvbnm48/grpc-basic/basic"
)

type logWriter struct{}

func (writer logWriter) Write(byte []byte) (int, error) {
	return fmt.Print(time.Now().Format("15:04:05") + " " + string(byte))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
	basic.BasicHello()
	basic.BasicUser()
}
