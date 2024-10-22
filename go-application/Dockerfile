FROM golang:latest

RUN mkdir -p /app/go-application
WORKDIR /app/go-application
COPY . .

RUN go build -o goApplication
EXPOSE 8080

CMD [ "./goApplication" ]