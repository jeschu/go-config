default: install

clean:
	@echo ">> clean <<"
	@rm -f version.go

generate:
	@echo ">> generate <<"
	@go generate -v ./...

build: generate
	@echo ">> build <<"
	@go build -v ./...

test: build
	@echo ">> test <<"
	@go test -v ./...

install: clean test
	@echo ">> install <<"
	@go install -v ./...