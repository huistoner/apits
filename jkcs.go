package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)



type Info struct{
	Value string

}



func get(){
	Response,_:=http.Get("http://yfs11.ynssx.com/api/Mldatasave/getapzhvalue")
	defer Response.Body.Close()

	s,_:=ioutil.ReadAll(Response.Body)

	if Response.StatusCode==200{

		var myInfo []Info

		var js []byte = s    //用byte数组存放要解析的内容


		err :=json.Unmarshal(js,&myInfo)
		if err !=nil {
			fmt.Printf("error%v\n",err)
		}

		fmt.Printf("%+v",myInfo)

	}else {
		fmt.Println("error")
	}
}


func main(){
ticker := time.NewTicker(time.Minute * 1)

for _=range ticker.C{
	//fmt.Printf("ticked at %v",time.Now())
	get()
}



}