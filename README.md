# Набор тестовых приложений и конфигураций для работы CloudNative приложений
## Сборка
```
make all                Сборка бинарников
cd build                Перейти в папку, в которую выгрузились все бинарники
docker-compose build    Собрать контейнеры для запуска
```

## Запуск сервисов
```
docker-compose up
```

### Запуск теста
Перейти в папку go
ADAPTER=127.0.0.1:5001 go run ./app/client



## Приложения

* /app
 * client     // клиент   запускается отдельно
 * adapter    // аналог csconn, который преобразует TCP запросы в gRPC.  запускается в compose / kubernetes или во вне, в зависимости от необходимости
 *  orch      // модуль, который принимает grpc запрос и производит проход по всем сервисам для сбора данных 
 * svc_dummy  // простой функциональный сервис          
 * svc_db     // логирует сервис в БД и возвращает ID последней записи

## Данные в БД
```
CREATE TABLE table1 (
  id serial NOT NULL PRIMARY KEY,
  val text NOT NULL
);

select * from table1

insert into table1 (val) values ('test text1') returning Id;
```