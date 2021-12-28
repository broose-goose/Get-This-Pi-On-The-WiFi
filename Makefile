
.PHONY: compile
compile:
	protoc ./api/v1/pi_wifi.proto \
		--proto_path=./api/v1 \
		--go_out=./api/v1/go/ \
		--go_opt=paths=source_relative \

	protoc ./api/v1/pi_wifi.proto \
		--proto_path=./api/v1 \
		--js_out=import_style=commonjs:./api/v1/ts  \
		--grpc-web_out=import_style=typescript,mode=grpcwebtext:./api/v1/ts