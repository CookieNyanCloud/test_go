package main

import (
	"flag"
	"fmt"
	"github.com/cookienyancloud/test_go/serverGin"
	"github.com/cookienyancloud/test_go/serverMux"
	"os"

	//"github.com/cookienyancloud/test_go/serverMux"
)

const (
	pickGin = "gin"
	pickMux = "mux"
)

const usageText = `
db:
  - init - creates version info table in the database.
  - up - runs all available migrations.
  - up [target] - runs available migrations up to the target one.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - version - prints current db version.
  - set_version [version] - sets db version without running migrations.
server:
  - gin - runs gin server, default.
  - mux - runs mux server.
Usage:
  go run *.go <command> [args]
`

func main()  {
	flag.Usage = usage
	var serverType string
	flag.StringVar(&serverType,"type","gin",usageText)
	flag.Parse()
	println(serverType)
	switch serverType {
	case pickGin:
		serverGin.RunGin()
	case pickMux:
		serverMux.RunMux()
	default:
		serverGin.RunGin()
	}
}

func usage() {
	fmt.Print(usageText)
	flag.PrintDefaults()
	os.Exit(2)
}