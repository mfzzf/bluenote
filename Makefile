.PHONY: docker
docker:
	@rm -rf ./bluenote
	@GOOS=linux GOARCH=arm go build -o bluenote .
	@docker image rm mfzzf/bluenote:v0.0.1
	@docker build -t mfzzf/bluenote:v0.0.1 .

