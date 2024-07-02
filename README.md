# Task Tracker

REST API application for users task tracking.

## Using

This App uses the [Go Buffalo framework](http://gobuffalo.io). Make sure that Buffalo is [installed](https://gobuffalo.io/documentation/getting_started/installation/).

Set up the environment variables by creating a `.env` file and adding the following content:

```
GO_ENV=development
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_PORT=5432
CONTAINER_NAME=postgres-tasktracker
INFO_URL=https://2eeba4a1-f18e-4fe5-9309-5b01aab3e99d.mock.pstmn.io
```

Note: If you are running the application in a different environment, make sure to update the `GO_ENV` variable accordingly.

Start the PostgreSQL database container and run the database migration:

```
make db
```

Start the application:

```
make dev
```

The application will be accessible at [http://127.0.0.1:3000](http://127.0.0.1:3000) by default.

## Documentation

For detailed documentation on the Task Tracker API, please visit [https://rustam2202.github.io/effective-mobile-tasktracker/](https://rustam2202.github.io/effective-mobile-tasktracker/).

<!-- ## What Next? -->
<!-- {"people":
    {
        "name": "Иван",
        "surname": "Иванов",
        "patronymic": "Иванович",
        "address": "г. Москва, ул. Ленина, д. 5, кв. 1"
    }
} -->
