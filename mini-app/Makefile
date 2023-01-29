gen-wire:
	@echo "  >  Generate Wire..."
	wire wire/wire.go

# START APPLICATION
start-dev: gen-wire
	@echo "  >  Starting Program..."
	go run cmd/main.go -env dev

start-test: gen-wire
	@echo "  >  Starting Program..."
	go run cmd/main.go -env test

start-stag: gen-wire
	@echo "  >  Starting Program..."
	go run cmd/main.go -env stag

start-prod: gen-wire
	@echo "  >  Starting Program..."
	go run cmd/main.go -env prod