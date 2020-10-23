# amqp-quacker

Send mock data to your AMQP exchange.

## Usage docker
`
docker run -e QUACKER_HOST=amqp.host.com
 -e QUACKER_PORT=5672 
 -e QUACKER_USERNAME=amqp-username 
 -e QUACKER_PASSWORD=amqp-password 
 -e QUACKER_EXCHANGE=amq.topic
 -e QUACKER_TOPIC=my-topic/telemetry 
 -v /home/zgldh/my-project/data.json:/data.json 
 zgldh/amqp-quacker
`

## Usage docker-compose

Edit the docker-compose.yml  
```
docker-compose up 
```


## Variables

name| descrpition | sample
----|-------------|---------
QUACKER_HOST| The host to your AMQP server. | "amqp.host.com"
QUACKER_PORT| The AMQP server port. |"1883"
QUACKER_USERNAME| For AMQP server auth. |"amqp-username"
QUACKER_PASSWORD| For AMQP server auth. |"amqp-password"
QUACKER_EXCHANGE| Which exchange do you want the mock data send to? |(empty)
QUACKER_TOPIC|If `QUACKER_EXCHANGE` is set, this is topic, else this is queue name|"your/topic/name"
QUACKER_INTERVAL| Time interval between two data sending. (in ms) |"1000"
QUACKER_DATAFILE| The mock data template. |"/data.json"

## Custom Data
Please edit the file `data.json` to any text you want. It support following placeholders:
- `q:float:{min}:{max}` to generate a float number between [min, max).
- `q:int:{min}:{max}` to generate an integer number between [min, max).

Currently, no more placeholders supported.

