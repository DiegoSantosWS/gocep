FROM golang:1.10-alpine as build-stage
LABEL maintainer="diegosantosws1@gmail.com"
WORKDIR /go/src/github.com/DiegoSantosWS/restcep
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch
LABEL maintainer="diegosantosws1@gmail.com"
COPY --from=build-stage /go/src/github.com/DiegoSantosWS/restcep /
CMD ["/main"]