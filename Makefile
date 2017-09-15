PREFIX=hvkvp
DESCRIBE=$(git describe --tags)

build: basher

basher: Dockerfile
	mkdir -p out
	docker build -t basher-build -f Dockerfile .
	docker create --name basher-extract basher-build sh
	docker cp basher-extract:/workspace/bin/basher* ./out/
	#docker rm basher-extract || true
	#docker rm basher-build || true

clean:
	rm -f ./out/*
