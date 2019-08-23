FROM golang:latest as build
RUN apt-get update -y
RUN apt-get install git -y
RUN go get code.google.com/p/whispering-gophers/hello
WORKDIR /opt
COPY ./main.go /opt/
RUN go build -o hello
RUN ls -la

FROM golang:latest
COPY --from=build /opt/hello /opt/
ENTRYPOINT [ "/opt/hello"]
