FROM golang
RUN go install github.com/wingrime/MasterPickupBot
ENTRYPOINT /go/bin/MasterPickupBot
