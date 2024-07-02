# Task Tracker

REST API application for users task tracking.
[Powered by Buffalo](http://gobuffalo.io)

## Docs



### Enviroments

Set your enviroments in ```.env``` file. For example:
```
GO_ENV=development
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_PORT=5432
CONTAINER_NAME=postgres-tasktracker
INFO_URL=https://2eeba4a1-f18e-4fe5-9309-5b01aab3e99d.mock.pstmn.io
```

Change on ```GO_ENV=test``` for debuging tests.
```INFO_URL/info``` used for get and fill User info by request with passportNumber.

## Starting the Application

Run ``` make db``` for create PostgeSQL docker container and development database migrate.
Run ``` make dev``` for start application.
Run ``` make migrate-test``` for testing database migrate and using for debugging tests.

By default used url [http://127.0.0.1:3000](http://127.0.0.1:3000).



## What Next?




<!-- {"people":
    {
        "name": "Иван",
        "surname": "Иванов",
        "patronymic": "Иванович",
        "address": "г. Москва, ул. Ленина, д. 5, кв. 1"
    }
} -->