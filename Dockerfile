FROM golang:latest
MAINTAINER Sergey Skoredin [sergey@skoredin.pro]
RUN go version
WORKDIR /app/iiko_report
COPY . .
RUN go build -o proc .
CMD ["/app/iiko_report/proc"]