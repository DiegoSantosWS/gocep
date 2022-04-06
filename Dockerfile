FROM golang:alpine as build-stage
LABEL maintainer="diegosantosws1@gmail.com"
WORKDIR /go/src/github.com/DiegoSantosWS/restcep
COPY . .

#RUN apk --no-cache add git
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:3.15.4
LABEL maintainer="diegosantosws1@gmail.com"
COPY --from=build-stage /go/src/github.com/DiegoSantosWS/restcep /
RUN apk --no-cache add ca-certificates && update-ca-certificates
CMD ["/main"]