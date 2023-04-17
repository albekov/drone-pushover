build:
	CGO_ENABLED=0 go build -o bin/drone-pushover -v -a -ldflags="-w -s" .

docker:
	docker build -t albekov/drone-pushover .
