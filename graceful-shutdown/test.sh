set -ex

if [ ! -d layers ]; then
	echo "FILE not found: layers/"
fi

if [ ! -f http-server.go ]; then
	echo "FILE not found: http-server.go"
fi

go build -o pebble ../cmd/pebble/
go build -o server http-server.go
PEBBLE=$PWD ./pebble run -v

if [ -f pebble ]; then
	rm pebble
fi

if [ -f server ]; then
	rm server
fi
