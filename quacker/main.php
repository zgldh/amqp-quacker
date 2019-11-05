<?php
require_once __DIR__ . '/vendor/autoload.php';

use PhpAmqpLib\Connection\AMQPStreamConnection;
use PhpAmqpLib\Message\AMQPMessage;

$connection;
$channel;

try {
  $connection = autoConnect(
    $_ENV['RABBITMQ_HOST'],
    $_ENV['RABBITMQ_PORT'],
    $_ENV['RABBITMQ_USER'],
    $_ENV['RABBITMQ_PASS']);
  $channel = $connection->channel();
  $channel->exchange_declare($_ENV['RABBITMQ_EXCHANGE_NAME'], 'fanout', false, false, false);
  output("Declared a fanout exchange: ". $_ENV['RABBITMQ_EXCHANGE_NAME']);

  while (true) {
    $time = time();
    $data = json_encode([
      'value' => sin($time *  M_PI / 45) * rand(100, 110) / 100,
      'ts' => $time
    ]);
    $msg = new AMQPMessage($data);
    $channel->basic_publish($msg, $_ENV['RABBITMQ_EXCHANGE_NAME']);
    output("Published: ".$data);
    sleep(1);
  }
} catch (\Exception $e) {
  output($e->getTrace());
  output($e->getMessage());
  whenShutdown();
}

function autoConnect($host = 'rabbitmq', $port = 5672, $user = 'guest', $pass = 'guest')
{
  global $connection;
  while (!$connection) {
    try {
      $connection = new AMQPStreamConnection($host, $port, $user, $pass);
    } catch (\Exception $e) {
      output($e->getMessage());
      output('Retrying...');
      $connection = null;
      sleep(2);
    }
  }
  return $connection;
}

function output($msg)
{
  $msg = is_string($msg) ? $msg : json_encode($msg);
  echo $msg;
  echo "\n";
}


function whenShutdown()
{
  output("Shutting down...");
  global $connection;
  global $channel;
  if ($channel) {
    $channel->close();
  }
  if ($connection) {
    $connection->close();
  }
}

register_shutdown_function('whenShutdown');
