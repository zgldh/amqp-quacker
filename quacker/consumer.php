<?php

require_once __DIR__ . '/vendor/autoload.php';

use PhpAmqpLib\Connection\AMQPStreamConnection;

$connection = new AMQPStreamConnection(
  $_ENV['RABBITMQ_HOST'],
  $_ENV['RABBITMQ_PORT'],
  $_ENV['RABBITMQ_USER'],
  $_ENV['RABBITMQ_PASS']
);
$channel = $connection->channel();

$channel->exchange_declare($_ENV['RABBITMQ_EXCHANGE_NAME'], 'fanout', false, false, false);

list($queue_name,,) = $channel->queue_declare("", false, false, true, false);

$channel->queue_bind($queue_name, $_ENV['RABBITMQ_EXCHANGE_NAME']);

echo " [*] Waiting for quacks. To exit press CTRL+C\n";

$callback = function ($msg) {
  echo ' [x] ', $msg->body, "\n";
};

$channel->basic_consume($queue_name, '', false, true, false, false, $callback);

while ($channel->is_consuming()) {
  $channel->wait();
}

$channel->close();
$connection->close();
