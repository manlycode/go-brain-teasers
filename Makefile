.PHONY: run
run:
	go run main.go

.PHONY: deps
deps:
	go install
	asdf reshim golang
