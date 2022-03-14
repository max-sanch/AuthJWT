# Аутентификация на JWT-токенах


* [Инструкции](#guides)
    * [Запуск приложения](#launch-app)
        * [Используя Docker](#Docker)
        * [Запуск без Docker](#wDocker)
    * [Список эндпоинтов](#endpoints)
        * [POST /api/users](#createUser)
        * [DELETE /api/users](#deleteUser)
        * [POST /auth/sing-in](#singIn)
        * [POST /auth/refresh](#refresh)

## <a name="guides"></a> Инструкции

### <a name="launch-app"></a>Запуск приложения

При изменении порта в config файле, необходимо указать тот же порт в файле docker-compose.


#### <a name="Docker"></a>Используя Docker

Вам необходимы `Docker`, `docker-compose` и `git`.

1) Скачайте проект с GitHub:

       git clone https://github.com/max-sanch/AuthJWT.git

2) Перейдите в директорию проекта и введите:

       docker-compose up --build

#### <a name="wDocker"></a>Запуск без Docker

Вам необходимы `git`, `golang 1.16.5` и `MongoDB 4.4.6` с переменными окружениями которые указаны в файле `.env`.
Также нужно указать соответствующие параметры в файле `configs/config.yml`

1) Скачайте проект с GitHub:

       git clone https://github.com/max-sanch/AuthJWT.git

4) Перейдите в директорию проекта и введите:

       go run cmd/main.go

### <a name="endpoints"></a>Список эндпоинтов

#### <a name="createUser"></a>POST /api/users

Создаёт пользователя в базе данных. Принимает только одно значение `guid` типа `string`.
В случае успеха — возвращает `guid`. В случае ошибки вернёт текст ошибки `message`.

#### <a name="deleteUser"></a>DELETE /api/users

Удаляет пользователя из базы данных. Принимает только одно значение `guid` типа `string`.
В случае успеха — возвращает `guid`. В случае ошибки вернёт текст ошибки `message`.

#### <a name="singIn"></a>POST /auth/sing-in

Выдает пару Access и Refresh токенов. Принимает только одно значение `guid` типа `string`.
В случае успеха — возвращает `access_token` и `refresh_token` типа `string`. В случае ошибки вернёт текст ошибки `message`.

#### <a name="refresh"></a>POST /auth/refresh

Выполняет Refresh операцию. Принимает `access_token` и `refresh_token` типа `string`.
В случае успеха — возвращает `access_token` и `refresh_token` типа `string`. В случае ошибки вернёт текст ошибки `message`.
