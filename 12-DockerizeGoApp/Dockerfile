FROM golang:alpine3.20 as build
RUN  mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .

FROM scratch
COPY --from=build /app/main /app/main
CMD ["/app/main"] 