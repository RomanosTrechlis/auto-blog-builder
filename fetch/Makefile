build:
	docker build -t fetch-image .

run:
	docker run -it --rm -v $(pwd)/../artifacts:/artifacts --name fetch-service fetch-image -p 50051:50051
