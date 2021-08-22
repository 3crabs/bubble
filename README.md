# bubble

bubble - простой задачник

## Генерация кода и установка библиотек
    protoc --proto_path=proto --go_out=plugins=grpc:. task.proto
    go get -u google.golang.org/grpc
    go get github.com/gofrs/uuid 

## Запуск программы
    go build -v -o bin/bubble
