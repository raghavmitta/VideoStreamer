FROM golang:1.19.5-alpine3.16
# Specify that we now need to execute any commands in this directory.
WORKDIR /app
COPY . .

# Copy everything from this project into the filesystem of the container.
# Obtain the package needed to run redis commands. Alternatively use GO Modules.
## Add this go mod download command to pull in any dependencies
RUN go mod download
## Our project will now successfully build with the necessary go libraries included.
RUN go build -o /VideoStreamer
## Our start command which kicks off
## our newly created binary executable
CMD ["/VideoStreamer"]