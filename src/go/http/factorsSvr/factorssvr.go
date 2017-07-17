package main

import (
	"fmt"
	"net/http"
	"log"
	"strconv"
	"github.com/stvndall/languagetechstats/src/go/services/factors"
	"github.com/gorilla/mux"
)
func main(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/{numbers}", factorise)
	log.Fatal(http.ListenAndServe(":2500", router))
}

func factorise(w http.ResponseWriter, r *http.Request) {
	number, err := strconv.Atoi(mux.Vars(r)["numbers"])
	if(err != nil){
		fmt.Fprintf(w,"%v",err)
		return
	}
	fmt.Fprintf(w, "%v",factors.PrimeFactors(number) )
}