build:
	go build ./cmd/setfan
	go build ./cmd/readtemp
	go build ./cmd/adjustfan

install:
	chmod 755 ./deploy/install.sh
	./deploy/install.sh

uninstall:
	chmod 755 ./deploy/uninstall.sh
	./deploy/uninstall.sh
