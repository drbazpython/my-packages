hello:
	echo "Hello"

build:
	@echo "Building .... "
	@go build -o ./bin drbaz.com/my-packages
	@echo "Installing ...."
	@go install drbaz.com/my-packages
	@echo "my-packages built to bin and installed"
	@echo "Running my-package ...."
	@my-packages
	
run:
	go run main.go

tidy:
	go mod tidy
	
	