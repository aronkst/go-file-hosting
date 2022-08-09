build:
	docker build -t go-file-hosting .
	docker container create --name go-file-hosting -e PORT=3000 -p 3000:3000 go-file-hosting

run:
	docker start -a go-file-hosting

start:
	docker start go-file-hosting

stop:
	docker container stop go-file-hosting

test-file:
	curl --request POST --url 'http://localhost:3000/file' --header 'Content-Type: multipart/form-data' --form file=@$(file)

test-url:
	curl --request POST --url 'http://localhost:3000/url' --header 'Content-Type: application/json' --data '{"url": "$(url)"}'
