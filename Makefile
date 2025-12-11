today := $(shell date +"%d")
today = 10

day:
	go run $(today)/main.go | tee $(today)/output.txt

run:
	go run $(day)/main.go

cp:
	mkdir $(today)
	cp template.go $(today)/main.go
	touch $(today)/test.txt
	touch $(today)/input.txt
