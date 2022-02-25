run:
	docker build -t hex-tokogambar -f ./build/package/server/Dockerfile .
	docker run -p 7124:7124 hex-tokogambar