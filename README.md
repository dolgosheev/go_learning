# Изучение Golang

## guess.go
``
Игра отгадай число
``

## test/pointers.go
``
Тест функциональности работы с указателями
``

## test/redis.go
``
Подключение NoSQL БД Redis
``

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

### Для корректной работы приложения потребуется установка плагина

```bash
go get github.com/redis/go-redis/v9
```