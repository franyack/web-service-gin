FROM golang:1.16-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY ./ ./
COPY ./entrypoint.sh /entrypoint.sh
RUN go mod download

# Build
# RUN go build -o /web-service-gin cmd/api/main.go

# Install Compile Daemon for go. We'll use it to watch changes in go files
RUN go get github.com/githubnemo/CompileDaemon


# wait-for-it requires bash, which alpine doesn't ship with by default. Use wait-for instead
ADD https://raw.githubusercontent.com/eficode/wait-for/v2.1.0/wait-for /usr/local/bin/wait-for
RUN chmod +rx /usr/local/bin/wait-for /entrypoint.sh

# This is for documentation purposes only.
# To actually open the port, runtime parameters
# must be supplied to the docker command.
EXPOSE 8080

# Run
# CMD [ "/web-service-gin" ]
ENTRYPOINT [ "sh", "/entrypoint.sh" ]
