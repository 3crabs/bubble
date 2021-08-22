# bubble

bubble - простой задачник

## Генерация кода и установка библиотек
    protoc --proto_path=proto --go_out=plugins=grpc:. task.proto
    go get -u google.golang.org/grpc
