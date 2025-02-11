FROM golang:1.15.3

RUN go get github.com/gin-gonic/gin

WORKDIR /go/src/app
COPY . .

RUN go install -v
RUN go get -u github.com/gin-gonic/gin

CMD ["app"]