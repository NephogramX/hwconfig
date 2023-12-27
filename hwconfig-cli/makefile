all: 
	go build -o ./build/hwconfig.exe

.PONEY = ./build/hwconfig.exe

build:
	go build -o ./build/hwconfig.exe
	
run: build
	./build/hwconfig.exe -r US915 -b basic_station -s 8
	scp ./build/chirpstack-gateway-bridge.toml root@10.6.6.6:/gw/etc/chirpstack-gateway-bridge/chirpstack-gateway-bridge.toml
	scp ./build/chirpstack-network-server.toml root@10.6.6.6:/gw/etc/chirpstack-network-server/chirpstack-network-server.toml

copy: build
	scp ./build/hwconfig root@10.6.6.6:/

clean:
	rm build/*