# GO Boilerplate

A MQTT and Rabbit MQ FIBER

## Features

- ⚙️ **Go** — The powerful and efficient programming language.
- ⚡ **Fiber** — An Express-inspired web framework for Go.
- 🧩 **sqlc** — Type-safe SQL query generator for Go.
- 🔁 **Air** — Live reload for Go apps during development.
- 🐇 **Rabbit MQ** — Broker.
- 🦟 **MQTT** — Lightweight Broker.

---

## Postman

ignore user route its used for project that use middleware, this project didn't use middleware yet

```
cd postman
vim GeoFetch API.postman_collection.json
```

---

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (>= 1.20)
- [Fiber](https://github.com/gofiber/fiber)
- [sqlc](https://github.com/kyleconroy/sqlc#installation)
- [Air](https://github.com/cosmtrek/air#installation)
- A SQL database (e.g., PostgreSQL)

---

### Installation

1. **Clone the repository:**

```bash
git clone https://github.com/Karula-k/go_geofetch
cd go_geofetch
```

2. **Copy ENV**

```bash
make new .env file using .env.example as reference
```

2. **Build**

```bash
docker-compose up --build -d
```

3. **Prerequisite**
   migrate the database

```bash
docker-compose exec api make migrateup
```

3. **Run the project**

```bash
docker-compose up
```

3. **open**

```bash
http://localhost:4000/swagger/
```

### Running Mock

```
cd scripts
go run .\mock.go
```

---

### Reference

[rabbitMQ EDA](https://github.com/Pungyeon/go-rabbitmq-example/blob/master/README.md)

[haversine Distance](https://medium.com/@abdurrehman-520/unlock-the-power-of-geofencing-in-flutter-with-haversine-formula-21b8203b1a5)
