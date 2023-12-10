package main

import (
	user_service "entry-rpc/kitex_gen/user_service/userservice"
	"entry-rpc/modules"
	"log"
)

func main() {
	err := modules.Startup()
	if err != nil {
		log.Fatal(err)
	}
	svr := user_service.NewServer(new(UserServiceImpl))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
