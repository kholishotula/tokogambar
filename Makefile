run:
	docker build -t tokogambar -f ./Dockerfile .
	docker run -p 7124:7124 tokogambar