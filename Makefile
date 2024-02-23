setup-emulator:
	yarn install
	yarn emulator

test:
	yarn run env-cmd go test -p 1 ./...

make build:
	go build -o bin/ ./...
