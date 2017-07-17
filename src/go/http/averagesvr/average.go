package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"github.com/stvndall/languagetechstats/src/go/services/average"
)
// Request is the data type that will mashal the numbers to be averaged
type Request struct{
	Numbers []int `json:"numbers"`
}
func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r * http.Request){
		decoder := json.NewDecoder(r.Body)
		var t Request
		err := decoder.Decode(&t)
		defer r.Body.Close()
		if(err != nil){
			log.Fatal(err)
		}
		avg, err :=average.Numbers(t.Numbers)
		if(err != nil){
			fmt.Printf("error attempting to average the numbers: %v \n", err)
		}
		fmt.Fprintf(w, "%f",avg )
	})
	http.ListenAndServe(":2500",nil)
}