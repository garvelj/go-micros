FROM alpine:latest

RUN mkdir /app

COPY loggerServiceApp /app

CMD [ "/app/loggerServiceApp" ]

# ono sto ce ovo da uradi jeste da napravi docker image
# od naseg kompletnog koda za brokera
# onda ce da napravi od executable-a novi mini image 