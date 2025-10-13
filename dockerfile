FROM golang:1.23.4

# Create app directory, this is in our container/in our image
WORKDIR /home/go/app


COPY go.mod go.sum ./

RUN go mod download


# Bundle app source
COPY . .

# Generate Prisma client for the container environment
RUN make mgs

RUN go build 


EXPOSE 4000
CMD ["go run", "main.go"]


