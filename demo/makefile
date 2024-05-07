.PHONY: all

all:
	go build
	# rm ./build/ -r
	# mkdir ./build
	./demo.exe

builtin: all
	scp ./build/global_conf.json root@10.6.6.6:/gw/opt/lora_pkt_fwd/
	scp ./build/chirpstack-gateway-bridge.toml root@10.6.6.6:/gw/etc/chirpstack-gateway-bridge
	scp ./build/chirpstack-network-server.toml root@10.6.6.6:/gw/etc/chirpstack-network-server

bs: all
	scp ./build/* root@10.6.6.6:/gw/opt/lora_pkt_fwd/

pf: 
	scp ./build/* root@10.6.6.6:/gw/opt/lora_pkt_fwd/