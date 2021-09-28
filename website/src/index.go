package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	//Loading enviroment variables
	godotenv.Load(".env")
	var portStr string = os.Getenv("WEB_SERVER_PORT")

	//TODO: COMO CASTEO EL STRING A ENTERO
	port, err := strconv.Atoi(portStr)

	fmt.Println("Puerto en string: " + portStr)

	if err != nil {
		fmt.Println("HUBO UN ERROR CON LA CONVERSION")
		fmt.Println(err)
	} else {
		fmt.Println("Puerto en int: " + strconv.Itoa(port))
	}

	/*
		fs := http.FileServer(http.Dir("./static"))
		http.Handle("/home/", http.StripPrefix("/home/", fs))

		http.HandleFunc("/", home)

		http.ListenAndServe(":55555", nil)
	*/

}

func home(writer http.ResponseWriter, request *http.Request) {
	html := "<html>"
	html += "<body>"
	html += "<h1>Hola Mundo</h1>"
	html += "</body>"
	html += "</html>"

	writer.Write([]byte(html))
}
