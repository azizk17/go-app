# Build stage
FROM golang:1.16-alpine3.13 AS builder
RUN apk --no-cache add curl

WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go version

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
# RUN go build -o main main.go
# RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

# FROM alpine:3.13

# RUN apk --no-cache add curl
# # Run stage
# WORKDIR /app

# COPY --from=builder /app .
# COPY --from=builder /app/migrate.linux-amd64 ./migrate


RUN ls
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration

# CMD [ "/app/main" ]
# ENTRYPOINT [ "/app/start.sh" ]
