# amqp-quacker

Send mock data to your AMQP exchange.

## Usage docker
Direct to a queue  
```
docker run -e QUACKER_HOST=amqp.host.com
 -e QUACKER_PORT=5672 
 -e QUACKER_USERNAME=amqp-username 
 -e QUACKER_PASSWORD=amqp-password
 -e QUACKER_TOPIC=my-queue-name 
 -v /home/zgldh/my-project/data.json:/data.json 
 zgldh/amqp-quacker:1.1
```  

Via an exchange  
```
docker run -e QUACKER_HOST=amqp.host.com
 -e QUACKER_PORT=5672 
 -e QUACKER_USERNAME=amqp-username 
 -e QUACKER_PASSWORD=amqp-password 
 -e QUACKER_EXCHANGE=my-exchange
 -e QUACKER_TOPIC=my-topic/telemetry 
 -v /home/zgldh/my-project/data.json:/data.json 
 zgldh/amqp-quacker:1.1
```

## Usage docker-compose

Edit the docker-compose.yml  
```
docker-compose up 
```


## Variables

name| descrpition | sample
----|-------------|---------
QUACKER_HOST| The host to your AMQP server. | `amqp.host.com`
QUACKER_PORT| The AMQP server port. |`1883`
QUACKER_USERNAME| For AMQP server auth. |`amqp-username`
QUACKER_PASSWORD| For AMQP server auth. |`amqp-password`
QUACKER_EXCHANGE| Which exchange do you want the mock data send to? (Optional)|(empty)
QUACKER_TOPIC|If `QUACKER_EXCHANGE` is set, this is topic, else this is queue name|`your/topic/name`
QUACKER_INTERVAL| Time interval between two data sending. (in ms) |`1000`
QUACKER_DATAFILE| The mock data template. |`./data.json`
QUACKER_DRYRUN| Dont push to server, just output payload. |`true` or keep it empty

## Custom Data
Please edit the file `data.json` to any text you want. It supports following placeholders:
- `q:float:{min},{max}` to generate a float number between [min, max).
- `q:int:{min},{max}` to generate an integer number between [min, max).
- `q:string:{str1},{str2},{str3},...,{strN}` to get one string from n strings randomly.
- `q:timestamp` to get a current timestamp like `1620904435`.
- `q:datetime` to get a current datetime string. You can custom the date format: `q:datetime:Mon Jan _2 15:04:05 MST 2006` Please refer to: https://golang.org/src/time/format.go

Currently, no more placeholders supported.

## Update Logs
1.2
- Break changes:  
    Use `{q:datetime}` instead of `{q:timestamp}` 
    The `{q:timestamp}` is changed to output real timestamp long integer like `1620904435`
    
1.1
- Added `{q:timestamp}` token.

1.0
- The beginning.