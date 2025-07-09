this repo is used as a workshop to learning microservices, dockerization and much more

The project can be ran through the make commands

1. Change directory into /project
2. Run make up_build
  - This will build the project and start it
3. Run the front end through make start
4. Go to localhost on your browser and start clicking :)
5. To stop the frontend use make stop
6. To stop the services use make down
7. To access the mailhog first start the project than go to localhost:8025


If the rabbitMQ isn't able to connect, and it displays the following error message: 
``` 
Cookie file /var/lib/rabbitmq/.erlang.cookie must be accessible by owner only
``` 

Use the following bash command to remove RabbitMQ data from the project:
```
rm -rf ./db-data/rabbitmq/
```
And then try to build docker compose again, it will most probably solve it.


Same goes for Mongo data. 
It can happen that Mongo crashes after some seconds being up. 
It usually helps to remove the db data. 
Maybe removing whole db-data folder is a better option if you're starting the service after a while, so when it's being built, it will add all the necessary files. 

So, similarly it goes like: 
```
rm -rf ./db-data/
```

*NOTE*: 
Be aware that if you're starting the service for the first time, you should execute the SQL queries located in the authentication-service/scripts/userScript.sql file. 

This will create your users table with one user entry which is used for testing the authentication service. 


In Go, if you're to use RPC, it has to be both applications written in Go. 
For example, we would want to implement the RPC protocol to be used between the Broker and the Logger service. This would not work if the Broker is written in Go, and then Logger to be written in Java, or Rust or anything else. 

Luckily for us, we have everything written in Go :)

Now with gRPC that is not the case.
You can have services written in various languages but have them communicate via gRPC. 