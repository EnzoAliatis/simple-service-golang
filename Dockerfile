FROM golang:1.16-alpine AS build
RUN apk add --update git
WORKDIR /go/src/simple-api-service
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/simple-api-service main.go

FROM scratch
ENV PORT=8080
COPY --from=build /go/bin/simple-api-service /go/bin/simple-api-service
ENTRYPOINT [ "/go/bin/simple-api-service" ]

