#Multi-stage docker file
FROM golang:1.22.5 as base

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o webserver .

#Final stage using distroless image

FROM gcr.io/distroless/base

COPY --from=base /app/webserver .

COPY --from=base /app/static ./static

EXPOSE 8080

CMD [ "./webserver" ]