FROM golang:alpine

WORKDIR /api
COPY . .

RUN go install
RUN go build -o github_api .

VOLUME ["/download"]

EXPOSE 3000

CMD ["./github_api"]