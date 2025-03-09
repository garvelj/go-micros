FROM alpine:latest

RUN mkdir /app

# da li ovo zapravo kopira samo exe u /app?
COPY mailerServiceApp /app
# a gdje je zapravo ovaj /app i /templates?
# da bi templates bio dostupan i kad se pokrene kontejner treba da ima templates 
COPY templates /templates 

CMD [ "/app/mailerServiceApp" ]