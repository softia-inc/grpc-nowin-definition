SHELL=bash

generate:
	protoc --proto_path=./ --go_out="plugins=grpc:./" ./*.proto
	cp -f github.com/softia-inc/**/* ./
	rm -rf github.com
gen-node:
	yarn -s run grpc_tools_node_protoc --proto_path=./ --js_out=import_style=commonjs,binary:./ --grpc_out=./  --plugin=protoc-gen-grpc=./node_modules/.bin/grpc_tools_node_protoc_plugin  ./*.proto
	yarn -s run grpc_tools_node_protoc --proto_path=./  --plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts  --ts_out=./  ./*.proto