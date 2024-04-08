# Start with a base image that includes Go runtime
FROM golang:latest

# Set necessary environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    DB_USER=root \
    DB_PASS=parth123 \
    DB_NAME=customerservice \
    JWT_KEY=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c



# Set the working directory inside the container
WORKDIR customerservice

# Copy the binary from the builder stage
COPY customerservice .

# # Expose port 8080 to the outside world
EXPOSE 5000

RUN chmod +x customerservice

# Command to run the executable
CMD ["./customerservice"]
