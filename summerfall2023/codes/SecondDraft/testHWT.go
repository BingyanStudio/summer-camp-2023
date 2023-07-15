package main

import(
	"myproject/util"
	"fmt"
)

func main(){
	s:=util.GenJWT("13870879061","1234")
	fmt.Println(s)

	res :=util.ParseJWT(s)
	fmt.Println(res)
}