FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY ../go.mod ../go.sum ./

# Copy the local dependency
COPY /common ../common
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy everything from the current directory to the Working Directory inside the container
COPY ./ .

COPY . .
RUN go build -v -o /app ./...

CMD ["app"]



