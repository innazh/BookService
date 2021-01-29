#	STEP 1: BREAD.
FROM golang:alpine as builder

WORKDIR /app/server
# copy the contents of src folder to the workdir of the container
COPY ./src/ .
# get dependencies
RUN go get -d -v . 

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o bookserver

# STEP 2: LETTUCE
FROM scratch
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# COPY --from=builder /etc/passwd /etc/passwd
# WORKDIR /app/server
COPY --from=builder ./app/server/keys.txt keys.txt
COPY --from=builder ./app/server/config.yaml config.yaml
COPY --from=builder ./app/server/bookserver bookserver
ENTRYPOINT ["./bookserver"]
# CMD ["./bookserver"]

# FROM golang
# ENV GO111MODULE=on

# #assign a workdir inside the container
# WORKDIR /app/server
# #copy the contents of src folder to the workdir of the container
# COPY ./src/ .

# #get dependencies
# RUN go get -d -v . 
# #build a binary
# RUN go build -o bookserver

# CMD ["./bookserver"]