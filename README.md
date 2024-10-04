# Song Library TZ


## Создание переменных окружения:
```bash
cp ./.envs/.env.example ./.envs/.env 
```

## Взаимодействие с проектом:

### 1. Запуск базы данных:

```bash
make up_db
```

### 2. Запуск приложения:

```bash
make up_song
```

### 3. Просмотр логов postgres:

```bash
make logs_db
```

### 4. Просмотр логов приложения:

```bash
make logs_song
```

### 5. Остановить базу данных:

```bash
make down_db
```

### 6. Остановить приложение:

```bash
make down_song
```

### 7. Остановить базу данных(и убить volumes):

```bash
make down_vol_db
```

### 8. Остановить приложение(и убить volumes):

```bash
make down_vol_song
```

### Если эндпоинты изменились, обновляем сваггер спецификацию:
```bash
swag init -g app/internal/controller/http/controller.go 
```


### После запуска базы данных и проекта, можно инициализировать готовые данные из файла init_data.json:
1. Устанавливаем пакеты из проекта:
```bash
go mod tidy
```
2. Запускаем файл init_data.go из корня проекта:
```bash
go run init_data.go
```


### TODO:
1. ~~Настроить инфраструктуру проекта~~
2. ~~Логирование~~
3. ~~База данных~~
4. ~~Rest API~~
5. Nginx + SSl (letsencrypt)

#### Ссылка на ТЗ - https://docs.google.com/document/d/1DeSr_ogIb802V9RVjm1rJ_t0iY7vDrvf1qYT6hggf5A/edit


#### Контакты:
* Telegram - https://t.me/go_rshok

