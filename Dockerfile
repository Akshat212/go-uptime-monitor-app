# Use an official Go image
FROM golang:1.24

# Set the working directory inside the container
WORKDIR /app

# Copy all files from your project directory into the container
COPY . .

# Build the Go application
RUN go build -o app .

# Expose the port the app runs on
EXPOSE 8080

# Run the app
CMD ["./app"]
