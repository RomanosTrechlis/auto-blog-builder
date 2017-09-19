build:
	docker build -t generate-image .

run:
	docker run -it --rm -v $(pwd)/../artifacts:/artifacts --name generate-service generate-image -p 50053:50051
