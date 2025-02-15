all: configure build

#.PHONY: configure
configure:
	$(info Configuring the project)

#.PHONY: build
build:
	$(info Building the project...)
	CGO_ENABLED=0 go build -o ./out/kevin ./cmd/kevin/main.go

#.PHONY: build
run: build
	$(info Running the project...)
	./out/kevin

#.PHONY: clean
clean:
	$(info Cleaning up...)
	rm -f ./out/kevin
