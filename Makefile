all: drink eat

drink: drink.go
	go build drink.go
	mkdir -p bin
	mv drink bin

eat: eat.go
	go build eat.go
	mkdir -p bin
	mv eat bin

clean:
	-rm bin/*