gen:
    protoc -I=./messages ./messages/*.proto --go_out=plugins=grpc:.
	
.PHONY: gen gen1
