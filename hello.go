package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func getEnv(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "Vaule is: ", os.Getenv(strings.ToUpper(ps.ByName("key"))), " My Ip is: ", os.Getenv("MY_POD_IP"))
}

func getKey(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//keyFile := strings.ToUpper(ps.ByName("key"))
	token, err := ioutil.ReadFile("/etc/keys/" + ps.ByName("key"))
	if err != nil {
		fmt.Fprint(w, err)
	}
	encodedKey := base64.StdEncoding.EncodeToString(token)

	fmt.Fprint(w, "Vaule is: ", encodedKey, " My Ip is: ", os.Getenv("MY_POD_IP"))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/env/:key", getEnv)
	router.GET("/key/:key", getKey)
	http.ListenAndServe(":9999", router)
}
