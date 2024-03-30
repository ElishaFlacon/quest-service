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
    - `go run` или `air` для запуска live mode (используйте `air -c .air.toml` при первом запуске)

- Зпускаем через Docker (может не работать):
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
    - принять последнюю миграцию `./goose.sh up` (unix)
    - откатить последнюю миграцию `./goose.sh down` (unix)
    - остальные команды в документации goose

<h3>
    Запускаем, не работет, ура! 🗿🚬
</h3>



</br>



</br>



<h2>
    📋 Техническое задание:
</h2>

- Админ - создание опроса:
    - админ выбирает команды для запуска опроса на странице листинга команд (выбирает через чекбоксы);
    - админ нажимает кнопку запустить опрос;
    - админа перекидывает на модалку для создания опроса;
    - у админа есть возможность выбрать шаблон опроса и использовать его;
    - админ может скрыть шаблоны или удалить, скрытые шаблоны можно показать и активировать снова, для этого будет чекбокс\свитчер "Показать скрытые";
    - у админа появляется компактный список выбранных команд;
    - админ заполняет название, описание, даты начала и конца пороса, выбирает индикаторы опроса;
    - форма создания опроса:
        - в самом верху выбор шаблона и чекбокс\свитчер (не обязательные поля);
        - название опроса (обязательное поле);
        - описание опроса (не обязательное поле);
        - начало опроса и конец опроса (не обязательные поля, если не заполнены, то опрос начнется сразу после отправки и никогда не закончится)
        - список команды, только названия и кнопка для удаления (должна быть хотя-бы одна команда);
        - индикаторы опроса (должен быть хотя-бы один индикатор);
        - кнопка "Отменить" (всегда активна), рядом кнопка "Создать опрос" (активна, когда выполнены условия выше), рядом чекбокс "Создать шаблон" (по стандарту - активен)
        - при успешной\неуспешной отправке появляется системное уведомление;

- Админ - листинг опросов:
    - админ заходит на страницу просмотра опросов;
    - на странице есть таблица опросов, колонки таблицы:
        - айди опроса
        - название опроса;
        - процент участников, которые прошли опрос;
        - начало и конец опроса;
        - статус опроса (подготовлен\запущен\завершен);
    - собртировать таблицу можно по всем колонкам;
    - есть поиск, который работает на название;
    - есть фильтры по:
        - количеству команд;
        - количеству индикаторов;
        - датам начало\конец;
        - статусу (по дефолту только статус подготовлен и запущен);

- Админ - деталка опроса:
    - админ, со страницы опросов, нажав на название опроса, переходит в деталку опроса;
    - в деталке есть форма:
        - должен быть список всех участников и напротив кажого статус по прохождению опроса;
        - название (нельзя редактировать);
        - описание (нельзя редактировать);
        - список команд (нельзя редактировать);
        - список выбранных индикаторов (нельзя редактировать);
        - начало и конец опроса (скорее всего нельзя редактировать);
        - в левой части кнопка "Завершить принудительно", в правой части кнопка "Отменить", кнопка "Сохранить";

- Пользователь - опрос:
    - когда пользователю назначается опрос, ему приходит уведомление на сайте, почте и тг (тг опционально);
    - также опрос появляется в отдельной вкладке опросы;
    - в сайдбаре, напротив кнопки перехода на страницу опросы должен быть красный бейдж, с количеством не пройденных опросов (если все опросы пройдены - бейджа нет);
    - форма опроса:
        - название опроса;
        - описание опроса;
        - название команды;
        - дальше идет список участников команды и под каждым участником индикаторы, рядом с индикатором слайдер для выбора оценки, при передвижении слайдера, появляется оценка, цвет слайдера и фон оценки зависит от оценки (по стандарту свитчер стоит на 5);
        - кнопка "Отменить", кнопка "Отправить";



</br>



<h2>
    ⚡ Немного дополнительной информации:
</h2>

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
