FROM golang:jessie

WORKDIR /go/src/igor
COPY . .

RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

EXPOSE 5000
CMD ["go-wrapper", "run"] # ["igor"]
