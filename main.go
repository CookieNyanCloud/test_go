package main

import (
	"flag"
	"github.com/cookienyancloud/test_go/serverGin"
	"github.com/cookienyancloud/test_go/serverMux"

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
	serverType:= flag.String("type","Gin","server type")
	//flag.Usage = usage
	//flag.Parse()
	println(*serverType)
	switch *serverType {
	case pickGin:
		serverGin.RunGin()
	case pickMux:
		serverMux.RunMux()
	default:
		serverGin.RunGin()
	}
}