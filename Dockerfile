
FROM golang:1.19.0-alpine
WORKDIR /usr/src/app
COPY . .
RUN go build -v -o /usr/local/bin/app ./main.go

FROM alpine:latest
COPY --from=0 /usr/local/bin/app /usr/local/bin/app
COPY . .
RUN mkdir -p /shares
RUN chmod 777 /shares
RUN chmod +x ./run.sh
RUN adduser -D -u 3977 nj
USER nj
ENV port_number=8182
ENV shared_folder=/shares

ENTRYPOINT ["./run.sh"]

EXPOSE 8182
