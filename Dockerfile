
FROM golang:1.19.0-alpine
RUN mkdir -p /shares
RUN chmod 777 /shares

WORKDIR /usr/src/app
COPY . .

RUN chmod +x ./run.sh
RUN go build -v -o /usr/local/bin/app ./...

RUN adduser -D -u 3977 nj
USER nj
ENV port_number=8182
ENV shared_folder=/shares

ENTRYPOINT ["./run.sh"]

EXPOSE 8182
