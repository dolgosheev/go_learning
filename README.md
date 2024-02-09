# Изучение Golang

## guess.go
```
Игра отгадай число
```

## test/pointers.go
```
Тест функциональности работы с указателями
```

## test/redis.go
```
Подключение NoSQL БД Redis
```

## test/grpc_server.go
```
gRPC Сервер
Работает на порту 50051 c gRPC рефлексией
```

### Как подключить Redis

В папке с solution лежит файл docker-compose.yaml, его нужно запустить командой

Запуск в фоновом режиме
```bash
docker-compose -f docker-compose.yaml up -d
```

Удаление Redis (volume останется)
```bash
docker-compose -f docker-compose.yaml down
```

### Для корректной работы приложения потребуется 

#### Установка плагинов

```bash
go get github.com/redis/go-redis/v9
go get google.golang.org/grpc
go get google.golang.org/protobuf/protoc-gen-go
go get google.golang.org/grpc/reflection
```

#### Генерация proto файлов

```bash
chmod +x proto-compile.sh
./proto-compile.sh
```

В результате выполнения будет создана/переопределена дирректория ./gen/go