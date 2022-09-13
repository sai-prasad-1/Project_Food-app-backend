# FROM golang

# RUN mkdir /build
# WORKDIR /build

# RUN export GO111MODULE=on
# RUN go get https://github.com/sai-prasad-1/Project_Food-app-backend/main
# RUN cd  /build && git clone https://github.com/sai-prasad-1/Project_Food-app-backend.git

# RUN cd /build/Project_Food-app-backend && go build -o /build/Project_Food-app-backend/

# ENTRYPOINT ["/build/Project_Food-app-backend/main"]

# EXPOSE 8080


FROM golang:latest

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main .

EXPOSE 8080

CMD [ "/app/main" ]