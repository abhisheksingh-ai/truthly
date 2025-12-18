# ------------ Step-1 Build Stage ------- 
# go compiler image 
FROM golang:1.25-alpine AS builder

#create working directory inside container 
WORKDIR /app

#copy dependency files first (for caching)
COPY go.mod go.sum ./

#download dependencies 
RUN go mod download

#copy entire code source code 
COPY . .

#build the go binary 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/main



# --------------- Step - 2 Run stage   -----------
# lighweight image for running app 
FROM alpine:latest

# create workig directory 
WORKDIR /app

# create the log directory
RUN mkdir -p /app/logs

#document volume location
VOLUME [ "/app/logs" ]

#copy only the compiled binary 
COPY --from=builder /app/app .

#expose port to run the app 
EXPOSE 8181

#command to run the app 
CMD ["./app"]




