# 编译
go build -o ./bin/  ./...

# 运行server
./bin/server -c cofnigs/config.yaml

# 运行client
./bin/client