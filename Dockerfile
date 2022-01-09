FROM golang:1.16

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

RUN go install github.com/spf13/cobra/cobra@v1.3.0 && \
    go install github.com/golang/mock/mockgen@v1.5.0

RUN apt-get update && apt-get install sqlite3 -y

CMD ["tail", "-f", "/dev/null"]
