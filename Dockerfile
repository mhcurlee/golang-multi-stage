# builder image
FROM golang as builder
RUN mkdir /build
ADD *.go /build/
WORKDIR /build
RUN go mod init build 
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .

# generate clean, final image for end users
FROM alpine
COPY --from=builder /build/app .
RUN mkdir /html
COPY index.html /html/
CMD [ "/app" ]
