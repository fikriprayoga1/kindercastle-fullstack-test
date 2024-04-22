FROM golang

WORKDIR /go/src

COPY src/. ./

CMD ["go", "run", "."]