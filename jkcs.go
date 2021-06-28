package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func get(){
	Response,_:=http.Get("http://yfs11.ynssx.com/api/Mldatasave/getapzhvalue")
	defer Response.Body.Close()

	s,_:=ioutil.ReadAll(Response.Body)
	fmt.Println(string(s))
	if Response.StatusCode==200{
		fmt.Println("ok")
	}else {
		fmt.Println("error")
	}
}
func (json.Marshaler)  {

}

func main(){

	get()

}