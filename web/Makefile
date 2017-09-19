build:
	docker build -t web-image .

run:
	docker run -it --rm -v $(pwd)/../artifacts:/artifacts --name web web-image -p 8080:50051
