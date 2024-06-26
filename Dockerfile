FROM golang:1.22-alpine as build-stage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /main

FROM gcr.io/distroless/base-debian11 AS release-stage

WORKDIR /

COPY .env .
COPY ./views /views
COPY ./static /static

USER nonroot:nonroot

COPY --from=build-stage /main /main

CMD ["/main"]
