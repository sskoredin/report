FROM golang:latest
MAINTAINER Sergey Skoredin [sergey@skoredin.pro]
RUN go version
WORKDIR /app/iiko_report
COPY . .
RUN go build -o /app/proc .
RUN rm -rf /app/iiko_report
COPY ./public /app
CMD ["/app/proc"]