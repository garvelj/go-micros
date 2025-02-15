#  ovo govori dockeru kako da napravi image

# FROM golang:1.23-alpine as builder

# RUN mkdir /app

# COPY . /app

# WORKDIR /app

# RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api 

# RUN chmod +x brokerApp

# build a tiny img 

FROM alpine:latest

RUN mkdir /app

COPY brokerApp /app

CMD [ "/app/brokerApp" ]

# ono sto ce ovo da uradi jeste da napravi docker image
# od naseg kompletnog koda za brokera
# onda ce da napravi od executable-a novi mini image 