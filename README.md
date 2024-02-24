Welcome

## Prerequisites

Before getting started, ensure you have the following prerequisites installed on your machine:

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Setup
0. Config
   copy the .env.example to .env and change the values

1. Clone the repository:
   ```bash
   git clone https://github.com/k-vanio/observabilidade-open-telemetry.git
   ```

2. Navigate to the project directory:
    ```bash
    cd observabilidade-open-telemetry
    ```

3. Run the application
    ```
   make build-up
    ```

7. URLS
   - [Service One](http://localhost:8080/?zipCode=07987110)
   - [JAEGER UI](http://localhost:16686)
   - [Zipkin](http://127.0.0.1:9411/zipkin/)
   - [Prometheus](http://localhost:9090)