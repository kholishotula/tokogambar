run:
	docker build -t tokogambar -f ./build/package/server/Dockerfile .
	docker run -p 7124:7124 tokogambar