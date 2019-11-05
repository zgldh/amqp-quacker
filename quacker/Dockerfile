FROM php:7.2-cli

COPY "./" "/var/php"

WORKDIR /var/php

RUN apt-get update && apt-get install -y \
  git zip \
  && docker-php-ext-install -j$(nproc) bcmath \
  && docker-php-ext-install -j$(nproc) sockets

COPY --from=composer /usr/bin/composer /usr/bin/composer

# Un-comment following line if you prefer Aliyun composer mirrors.
# RUN composer config -g repo.packagist composer https://mirrors.aliyun.com/composer/
RUN composer install
