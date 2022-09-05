
FROM golang:1.19.0
RUN mkdir /shares
WORKDIR /usr/src/app
COPY . .
RUN go build -v -o /usr/local/bin/app ./...

ENTRYPOINT ["app"]
CMD  ["-p 8182"]
EXPOSE 8182
# This isnt working , i dont know
#  flag provided but not defined: -p 8182
# but If I hop inside the docker image and run app ,it works - dont have time to fix ..if someone can - thanks