run:
	docker compose up

start:
	docker compose up -d

stop:
	docker compose stop

logs:
	docker compose logs -f

test-file:
	curl --request POST --url 'http://localhost:3000/file' --header 'Content-Type: multipart/form-data' --form file=@$(file)

test-url:
	curl --request POST --url 'http://localhost:3000/url' --header 'Content-Type: application/json' --data '{"url": "$(url)"}'
