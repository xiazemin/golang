package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

func HomeHandler(http.ResponseWriter, *http.Request){

}

type add interface {
	Add(int,int)int
}

type addfunc func (int,int)int

func(f addfunc)Add(a int,b int)int{
	return f(a,b)
}
func myAdd(a,b int)int{
	return a+b
}
func main() {
	//fmt.Println(add(myAdd).Add(1,2))
	fmt.Println(addfunc(myAdd)(1,2))//3
	fmt.Println(addfunc(myAdd).Add(1,2))//3
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products", HomeHandler)
	r.HandleFunc("/articles", HomeHandler)
	http.Handle("/", r)

	http.ListenAndServe(":8080", r)
}