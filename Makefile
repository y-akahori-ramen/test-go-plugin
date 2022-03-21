modulename = github.com/y-akahori-ramen/test-go-plugin/pb

gen:
	rm ./pb/*go
	protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb --go_opt=module=$(modulename) --go-grpc_opt=module=$(modulename)

clean:
	rm -rf ./bin/

build:
	mkdir -p ./bin
	go build -o ./bin/plugin_add ./cmd/plugin_add/main.go
	go build -o ./bin/plugin_sub ./cmd/plugin_sub/main.go
	go build -o ./bin/plugin_div ./cmd/plugin_div/main.go
	go build -o ./bin/plugin_mul ./cmd/plugin_mul/main.go
	go build -o ./bin/client ./cmd/client/main.go

run:
	./bin/client

# addプラグインを指定して実行
run_add: export CALC_PLUGIN="./bin/plugin_add"
run_add:
	$(MAKE) run

# subプラグインを指定して実行
run_sub: export CALC_PLUGIN="./bin/plugin_sub"
run_sub:
	$(MAKE) run

# divプラグインを指定して実行
run_div: export CALC_PLUGIN="./bin/plugin_div"
run_div:
	$(MAKE) run

# mulプラグインを指定して実行
run_mul: export CALC_PLUGIN="./bin/plugin_mul"
run_mul:
	$(MAKE) run

.PHONY: gen clean build run_add run_sub run_div run_mul