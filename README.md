# kafka-example

## How to test

- Run kafka
```bash
make up
```
- Wait until kafka is ready
- Open new console and run consumer
```bash
make consumer
```
- Open [kafka-ui](http://localhost:8080/ui/docker-kafka-server/topic/topic/data?sort=Oldest&partition=All)
- Open new console and send message
```bash
make message
```

## Do not forget
- Open console with consumer and stop
- Open console with kafka, stop it and run
```bash
make down && docker volume rm $(docker volume ls -q | grep kafka-example)
```