package main

import(
	"fmt"
	"log"
	"net/http"
	router "github.com/MahadevMaahi/jwt-api-authentication/router"
)

func main() {
	fmt.Println("Server Initiation Starting...")
	var mRouter router.RouterKit = &(router.MuxRouter{})
	r := mRouter.SetUpRouter()
	fmt.Println("Starting Server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}