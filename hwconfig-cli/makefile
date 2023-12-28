.PHONY: linux win clean copy all build run

all: 
	go build -o ./build/hwconfig.exe

build:
	go build -o ./build/hwconfig.exe
	
run: build
	./build/hwconfig.exe -r US915 -b basic_station -s 8
	scp ./build/chirpstack-gateway-bridge.toml root@10.6.6.6:/gw/etc/chirpstack-gateway-bridge/chirpstack-gateway-bridge.toml
	scp ./build/chirpstack-network-server.toml root@10.6.6.6:/gw/etc/chirpstack-network-server/chirpstack-network-server.toml

scp-bin: build
	scp ./build/hwconfig root@10.6.6.6:/

scp-conf:
	scp ./build/chirpstack-gateway-bridge.toml root@10.6.6.6:/gw/etc/chirpstack-gateway-bridge/chirpstack-gateway-bridge.toml
	scp ./build/chirpstack-network-server.toml root@10.6.6.6:/gw/etc/chirpstack-network-server/chirpstack-network-server.toml
	scp ./build/global_conf.json root@10.6.6.6:/gw/opt/lora_pkt_fwd/

clean:
	rm build/*

linux:
	go env -w GOOS=linux
	go env -w GOARCH=arm64

win:
	go env -w GOOS=windows
	go env -w GOARCH=amd64