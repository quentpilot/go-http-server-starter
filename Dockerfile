FROM golang:1.21

# Installation de air pour le live reload
RUN go install github.com/cosmtrek/air@v1.40.4

WORKDIR /app

COPY . .

EXPOSE 8080

CMD ["air"]