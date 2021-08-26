package main

import (
	"flag"
	"github.com/cookienyancloud/test_go/serverGin"
	//"github.com/cookienyancloud/test_go/serverMux"
)

func main()  {
	serverType:= flag.String("type","Gin","server type")
	println(*serverType)

	serverGin.RunGin()
	//serverMux.RunMux()
}