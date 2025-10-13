FROM golang:1.24.0

# Create app directory, this is in our container/in our image
WORKDIR /home/go/app


COPY go.mod go.sum ./

RUN go mod download


# Bundle app source
COPY . .

RUN make mgs

RUN go build 


EXPOSE 4000
CMD ["./main"]


