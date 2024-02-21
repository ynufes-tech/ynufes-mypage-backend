setup-emulator:
	yarn install
	yarn emulator

test:
	env-cmd go test -p 1 ./...
