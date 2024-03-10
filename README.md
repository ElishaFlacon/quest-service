<h1> 
     ❓ Questionnaire Service 
</h1>

<h3>
 Questionnaire Service - это микросервис опросник
</h3>



</br>



<h2>
  🛠️ Инструменты, которые использовались при разработке приложения:
</h2>

- Go
- Gin
- Pgx
- PostgreSQL
- Docker



</br>



<h2>
  🚀 Зпуск приложения:
</h2>

- Зпускаем локально:
     - `git clone https://github.com/ElishaFlacon/questionnaire-service.git`
     - `cd questionnaire-service`
     - `go run`

- Зпускаем через Docker (может не работать):
     - устанавливаем Docker
     - `git clone https://github.com/ElishaFlacon/questionnaire-service.git`
     - `cd questionnaire-service`
     - `docker-compose build`
     - `docker-compose run`

- Инициализация базы данных:
    - устанавливаем Postgresql 15
    - создаем новую базу данных
    - устанавливаем свои значения для базы данных в `.env` файле
    - `./migrations/scripts/migrate.dev.ps1` (windows), `./migrations/scripts/migrate.dev.sh` (unix)

<h3>
    Запускаем, не работет, ура! 🗿🚬
</h3>



</br>



<h2>
⚡ Немного дополнительной информации:
</h2>

- https://habr.com/ru/companies/mvideo/articles/744434/
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
