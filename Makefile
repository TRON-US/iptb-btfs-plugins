IPTB_ROOT ?=$(HOME)/testbed
NODES ?=5

all: iptb

iptb:
	(cd iptb; go build)
CLEAN += iptb/iptb

btfslocal:
	(cd localbtfs/plugin; go build -buildmode=plugin -o ../../build/localbtfs.so)
CLEAN += build/localbtfs.so

install:
	(cd iptb; go install)

clean:
	rm ${CLEAN}

start:
	iptb auto -type localbtfs -count $(NODES)
	iptb run -- btfs config profile apply storage-host
	iptb start
	sleep 10
	iptb logs > iptb_logs.txt
	go run standalone/rewrite_config.go

start_dev:
	iptb auto -type localbtfs -count $(NODES)
	iptb run -- btfs config profile apply storage-host-dev
	iptb run -- btfs config optin
	iptb start
	sleep 10
	iptb logs > iptb_logs.txt
	go run standalone/rewrite_config.go

start_testnet:
	iptb auto -type localbtfs -count $(NODES)
	iptb run -- btfs config profile apply storage-host-testnet
	iptb run -- btfs config optin
	iptb start
	sleep 10
	iptb logs > iptb_logs.txt
	go run standalone/rewrite_config.go

stop:
	iptb stop

.PHONY: all clean btfslocal iptb start start_dev start_testnet stop
