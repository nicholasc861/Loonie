# syntax=docker/dockerfile:1.0

# Alpine is chosen for its small footprint
FROM golang as build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -a -o backend
COPY --from=build /app/backend /backend

CMD [ "/backend" ]