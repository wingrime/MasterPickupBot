FROM golang
RUN go install github.com/wingrime/masterbot
ENTRYPOINT /go/bin/masterbot
