<h1> 
    ❓ Quest Service 
</h1>

<h3>
    Quest Service - это микросервис опросник 360
</h3>



</br>



<h2>
    🛠️ Инструменты:
</h2>

- Go
- Gin
- Pgx
- RabbitMQ 
- PostgreSQL
- Docker



</br>



<h2>
    🚀 Зпуск приложения:
</h2>

- Зпускаем локально:
    - устанавливаем Go 1.21.0
    - устанавливаем Postgresql 15
    - создаем базу данных quest или любое другое название (тогда нужно будет указать новое название в `.env` файле)
    - `git clone https://github.com/ElishaFlacon/quest-service.git`
    - `cd quest-service`
    - `cp .env-example .env`
    - `cp .env-docker .docker.env`
    - `go mod download`
    - `go run .` или `air` для запуска live mode (используйте `air -c .air.toml` при первом запуске)

- Зпускаем через Docker:
    - устанавливаем Docker
    - добавляем в конфиг Docker зеркала, <a href="https://dockerhub.timeweb.cloud/">тут гайд</a>
    - `git clone https://github.com/ElishaFlacon/quest-service.git`
    - `cd quest-service`
    - `cp .env-example .env`
    - `cp .env-docker .docker.env`
    - `docker-compose build`
    - `docker-compose up`
    - заходим в pgadmin по адресу `localhost:5050/browser`
    - для входа используем: 
        - email = `user@root.com` 
        - password = `root`
    - пкм по servers -> register -> server
    - подключаемся к базе данных:
        - вкладка general 
            - name = `quest-postgres`
        - вкладка connection 
            - host = `quest-postgres`
            - port = `5432`
            - username = `user`
            - password = `root`
    - проверяем что контейнер миграций отработал корректно (появится сообщение goose: successfully migrated database to version: 1), если отработал с ошибкой запускаем еще раз
    - также для этого проекта есть <a href="https://hub.docker.com/r/elishaflacon/quest-service">docker image</a>
<h3>
    Запускаем, не работет, ура! 🗿🚬
</h3>



</br>



<h2>
    🛢️ Схема базы данных:
</h2>

![image](https://github.com/ElishaFlacon/quest-service/assets/83610362/409a4f6d-e4db-46ca-b7e9-2dba0ef03711)

> Схема актуальна для 06.06.2024, проверьте дату последних изменений в migrations



</br>



<h2>
    ⚡ Немного дополнительной информации:
</h2>

- `APP_MODE`: `PRODUCTION` или `DEBUG`
- Документация `/api/v1/quest-service/docs/index.html`
- Миграции (выполняются в bash терминале, для windows он идет вместе с git)
    - принять последнюю миграцию `./goose.sh up`
    - откатить последнюю миграцию `./goose.sh down`
    - остальные команды в документации goose
- При добавлении нового роута - не забываем добавить его в `routes/index.route.go`
- Hide запросы - это запросы на скрытие элементов, используйте их как удаление элементов
- P.S. Все баги и недочеты - это фичи





<br/>
<br/>
<br/>
<br/>
<br/>
<br/>





<p align="center">
    <img src="https://capsule-render.vercel.app/api?type=waving&color=d179b8&height=64&section=footer"/>
</p>
