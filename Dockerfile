# Start from the latest golang base image
FROM golang:1.15.6 as builder

# Add Maintainer Info
LABEL maintainer="Wiard van Rij"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -a -installsuffix cgo -o  main .

######## Start a new stage from scratch #######
FROM alpine:3.13

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

WORKDIR /appuser/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8000
EXPOSE 2112

# Command to run the executable
CMD ["./main"] 