# test-go-plugin
https://github.com/hashicorp/go-plugin を使ってみる

四則演算をプラグイン化し、起動時に使用する演算プラグインを指定して演算を行う。
## Example

```sh
# build binary
$ make build

# run with plugin_add
$ make run_add

# run with plugin_sub
$ make run_sub

# run with plugin_div
$ make run_div

# run with plugin_mul
$ make run_mul
```

## Environments
- go 1.18
- libprotoc 3.19.4
- protoc-gen-go v1.26.0
- protoc-gen-go-grpc 1.1.0
