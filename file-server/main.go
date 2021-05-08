package main

import (
	"log"
	"net/http"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)

func getCommandOutput(command string, arguments ...string) string {
	out, _ := exec.Command(command, arguments...).Output()
	return string(out)
}

func main() {
	router := httprouter.New()
	response := getCommandOutput("/bin/pwd", "")
	response = response[:len(response)-1]

	router.ServeFiles("/*filepath",

		http.Dir(response))
	log.Fatal(http.ListenAndServe(":8000", router))
}
