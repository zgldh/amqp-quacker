# rabbitmq-quacker

Send mock data to your AMQP(RabbitMQ) exchange.

## Usage docker
`
docker run -e QUACKER_HOST=rabbitmq.host.com
 -e QUACKER_PORT=5672 
 -e QUACKER_USERNAME=rabbitmq-username 
 -e QUACKER_PASSWORD=rabbitmq-password 
 -e QUACKER_EXCHANGE=amq.topic
 -e QUACKER_TOPIC=my-topic/telemetry 
 -v /home/zgldh/my-project/data.json:/data.json 
 zgldh/rabbitmq-quacker
`

## Usage docker-compose

Edit the docker-compose.yml  
```
docker-compose up 
```


## Variables

name| descrpition | sample
----|-------------|---------
QUACKER_HOST| The host to your RabbitMQ server. | "rabbitmq.host.com"
QUACKER_PORT| The RabbitMQ server port. |"1883"
QUACKER_USERNAME| For RabbitMQ server auth. |"rabbitmq-username"
QUACKER_PASSWORD| For RabbitMQ server auth. |"rabbitmq-password"
QUACKER_EXCHANGE| Which exchange do you want the mock data send to? |"amq.topic"
QUACKER_TOPIC|Which topic do you want the data to|"your/topic/name"
QUACKER_INTERVAL| Time interval between two data sending. (in ms) |"1000"
QUACKER_DATAFILE| The mock data template. |"/data.json"

## Custom Data
Please edit the file `data.json` to any text you want. It support following placeholders:
- `q:float:{min}:{max}` to generate a float number between [min, max).
- `q:int:{min}:{max}` to generate an integer number between [min, max).

Currently, no more placeholders supported.

