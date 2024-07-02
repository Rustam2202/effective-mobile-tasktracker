# Task Tracker

REST API application for users task tracking.
[Powered by Buffalo](http://gobuffalo.io)
## Installation

To install and run the Task Tracker application, follow these steps:

1. Clone the repository:

    ```shell
    git clone https://github.com/rustamnurmiev/tasktracker.git
    ```

2. Navigate to the project directory:

    ```shell
    cd tasktracker
    ```

3. Set up the environment variables by creating a `.env` file and adding the following content:

    ```shell
    GO_ENV=development
    POSTGRES_USER=postgres
    POSTGRES_PASSWORD=postgres
    POSTGRES_PORT=5432
    CONTAINER_NAME=postgres-tasktracker
    INFO_URL=https://2eeba4a1-f18e-4fe5-9309-5b01aab3e99d.mock.pstmn.io
    ```

    Note: If you are running the application in a different environment, make sure to update the `GO_ENV` variable accordingly.

4. Start the PostgreSQL database container and run the database migration:

    ```shell
    make db
    ```

5. Start the application:

    ```shell
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