<h1> 
    ❓ Quest Service 
</h1>

<h3>
    Quest Service - это микросервис опросник
</h3>



</br>



<h2>
    🛠️ Инструменты, которые использовались при разработке приложения:
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
    - `git clone https://github.com/ElishaFlacon/quest-service.git`
    - `cd questionnaire-service`
    - `cp .env-example .env`
    - `go mod download`
    - `go run .` или `air` для запуска live mode (используйте `air -c .air.toml` при первом запуске)

- Зпускаем через Docker (не работает):
    - устанавливаем Docker
    - `git clone https://github.com/ElishaFlacon/quest-service.git`
    - `cd questionnaire-service`
    - `docker-compose build`
    - `docker-compose run`

- Инициализация базы данных:
    - устанавливаем Postgresql 15
    - создаем новую базу данных
    - устанавливаем свои значения для базы данных в `.env` файле

- Миграции (выполняются в bash терминале, для windows он идет вместе с git)
    - принять последнюю миграцию `./goose.sh up`
    - откатить последнюю миграцию `./goose.sh down`
    - остальные команды в документации goose

<h3>
    Запускаем, не работет, ура! 🗿🚬
</h3>



</br>



<h2>
    ⚡ Немного дополнительной информации:
</h2>

- Документация `/docs/index.html`
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
