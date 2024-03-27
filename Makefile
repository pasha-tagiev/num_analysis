.SILENT:

.PHONY: lab1
lab1:
	go build -o build/out1 ./lab1
	./build/out1

.PHONY: lab2
lab2:
	go build -o build/out2 ./lab2
	./build/out2

.PHONY: lab3
lab3:
	go build -o build/out3 ./lab3
	./build/out3

clean:
	rm -rf ./build