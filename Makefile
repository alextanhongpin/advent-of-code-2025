today := $(shell date +"%d")

day:
	go run $(today)/main.go

run:
	go run $(day)/main.go

cp:
	mkdir $(today)
	cp template.go $(today)/main.go
	touch $(today)/test.txt
	touch $(today)/input.txt
