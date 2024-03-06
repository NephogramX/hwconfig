.PHONY: buildin all

all:
	go build
	./demo

buildin:
	scp ./build/global_conf.json root@10.6.6.6:/gw/opt/lora_pkt_fwd/
	scp ./build/chirpstack-gateway-bridge.toml root@10.6.6.6:/gw/etc/chirpstack-gateway-bridge
	scp ./build/chirpstack-network-server.toml root@10.6.6.6:/gw/etc/chirpstack-network-server

bs:
	scp ./build/* root@10.6.6.6:/gw/opt/lora_pkt_fwd/