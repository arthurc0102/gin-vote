FROM golang:1.10 AS build

WORKDIR $GOPATH/src/github.com/arthurc0102/gin-vote

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go get -v .
RUN go build -o app.out .
RUN cp ./app.out /app.out

# ==============================

FROM scratch

WORKDIR /app

COPY --from=build /app.out .
COPY static .
COPY public .
COPY log .

CMD ["/app/app.out"]
