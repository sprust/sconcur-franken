FROM ubuntu:24.04

# Container user & group
ARG USER_ID
ARG GROUP_ID

# Install required packages
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    ca-certificates \
    curl \
    build-essential \
    pkg-config \
    libxml2-dev \
    libsqlite3-dev \
    php-dev \
    php-cli \
    php-pear \
    make \
    wget && \
    rm -rf /var/lib/apt/lists/*

# Download and install FrankenPHP
RUN wget https://github.com/php/frankenphp/releases/download/v1.10.1/frankenphp-linux-x86_64 && \
    chmod +x frankenphp-linux-x86_64 && \
    mv frankenphp-linux-x86_64 /usr/local/bin/frankenphp

# Composer
RUN curl -fsSL https://getcomposer.org/download/latest-stable/composer.phar -o /usr/local/bin/composer && \
    chmod +x /usr/local/bin/composer

# PHP source
RUN rm -rf /usr/local/sconcur
RUN mkdir -p -m 777 /usr/local/sconcur
WORKDIR /usr/local/sconcur

RUN rm -rf php-src && \
  mkdir -p php-src && \
  wget -P /tmp https://www.php.net/distributions/php-8.4.15.tar.gz && \
  tar -xf /tmp/php-8.4.15.tar.gz -C php-src --strip-components=1 && \
  rm /tmp/php-8.4.15.tar.gz && \
  chmod -R 777 php-src && \
  cd php-src && \
  ./configure \
    --enable-embed \
    --enable-zts \
    --disable-zend-signals \
    --enable-zend-max-execution-timers && \
  make -j"$(getconf _NPROCESSORS_ONLN)" && \
  make install

WORKDIR /app

# Check if user and group doesn't exist before creating
RUN getent group "$GROUP_ID" || addgroup --gid "$GROUP_ID" user
RUN getent passwd "$USER_ID" || adduser --disabled-password --gecos '' --uid "$USER_ID" --gid "$GROUP_ID" user

USER "$USER_ID"

WORKDIR /app
