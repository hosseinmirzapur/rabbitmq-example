# RabbitMQ example using Go

## Usage

Using this CLI, you can create a consumer and a producer of RabbitMQ service to communicate with eachother using AMQP (Advanced Message Queuing Protocol)

## Using docker compose

```bash
docker-compose up --build    # for "permission-denied" errors run the command as "sudo"
```

## Convenient Commands using the Makefile

Be sure to have `buildtools` installed on your machine before running these commands:

- Create a RabbitMQ producer:

```bash
make producer
```

- Create a RabbitMQ consumer:

```bash
make consumer
```

## Contribution

Any contribution on this project is appreciated, also If you liked this code be sure to give it a simple star, thanks ^_^ 