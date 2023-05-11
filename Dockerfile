# Start with the official Golang image
FROM golang:latest

# Copy the app files to the container
WORKDIR /app
COPY . .

# Build the app
RUN git config --global --add url."git@github.com:".insteadOf "https://github.com/"
RUN go build -o manufacturing-service

# Start the app
CMD ["./manufacturing-service"]
