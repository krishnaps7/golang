FROM golang
WORKDIR /apps
COPY ./image.go /apps/image.go
RUN go build image.go
CMD ["./image"]