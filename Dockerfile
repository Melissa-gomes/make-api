FROM golang:1.19.0 as builder

ENV WDIR /app

RUN mkdir -p ${WDIR}

WORKDIR ${WDIR}

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /api -mod=readonly ./src/cmd/rest/main.go

FROM alpine as api

RUN apk add --no-cache ca-certificates

COPY --from=builder /api /api

ENTRYPOINT [ "/api" ]
