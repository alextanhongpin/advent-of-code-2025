today := $(shell date +"%d")

day:
	go run $(today)/main.go

run:
	go run $(day)/main.go

cp:
	mkdir -p $(today)
	cp template.go $(today)/main.go
