FROM golang:buster as build

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN make compile

FROM alpine
WORKDIR /root/app
COPY --from=build /build/out .

CMD [ "/root/app/oxford-bot" ]
