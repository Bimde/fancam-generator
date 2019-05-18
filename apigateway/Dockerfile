# Telling to use Docker's golang ready image
FROM golang
# Create app folder 
RUN mkdir /app
# Copy our file in the host contianer to our contianer
ADD . /app
# Set /app to the go folder as workdir
WORKDIR /app
# Generate binary file from our /app
RUN go build
# Expose the port 3000
EXPOSE 3000:3000
# Run the app binarry file 
CMD ["./apigateway"]

