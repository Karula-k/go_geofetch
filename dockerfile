FROM golang:1.24.9

# Create app directory, this is in our container/in our image
WORKDIR /home/go/app

# Install golang-migrate
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY go.mod go.sum ./

RUN go mod download

# Bundle app source
COPY . .

COPY .env .env

RUN go build -o main

EXPOSE 4000

CMD ["./main"]
# Run migrations: docker-compose exec api make migrateup

