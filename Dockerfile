FROM golang:1.13

ENV APP_DIR=/go/src/personal-website
ENV APP_BIN=personal-website.bin

# Create app directory and copy over src
# Set current dir for further commands
RUN mkdir $APP_DIR
ADD . $APP_DIR
WORKDIR $APP_DIR

# Download and install any dependencies
RUN go get -v -d ./...
RUN go install -v ./...

# Build binary from current dir
RUN go build -o $APP_BIN .

# Our start command which kicks off our newly created binary executable
CMD $APP_DIR/$APP_BIN

# Expose the port
EXPOSE 8080