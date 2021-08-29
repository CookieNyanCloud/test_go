gin:
	go run main.go -type=gin

mux:
	go run -type=mux

main:
	go run main.go

.PHONY: gin mux main