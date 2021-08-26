gin:
	go run serverGin/server.go

mux:
	go run serverMux/server.go

main:
	go run main.go

.PHONY: gin mux main