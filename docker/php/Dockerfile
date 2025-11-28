FROM ubuntu:24.04

WORKDIR /app

ENV COMPOSER_ALLOW_SUPERUSER=1
ENV COMPOSER_PHP="frankenphp php-cli"

# Container user & group
ARG USER_ID
ARG GROUP_ID

# Install required packages
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    ca-certificates \
    curl \
    wget && \
    rm -rf /var/lib/apt/lists/*

# Download and install xcaddy
RUN wget https://github.com/caddyserver/xcaddy/releases/download/v0.4.5/xcaddy_0.4.5_linux_amd64.deb && \
    dpkg -i xcaddy_0.4.5_linux_amd64.deb && \
    rm xcaddy_0.4.5_linux_amd64.deb

# Download and install FrankenPHP
RUN wget https://github.com/php/frankenphp/releases/download/v1.10.1/frankenphp-linux-x86_64 && \
    chmod +x frankenphp-linux-x86_64 && \
    mv frankenphp-linux-x86_64 /usr/local/bin/frankenphp

# Composer
RUN curl -fsSL https://getcomposer.org/download/latest-stable/composer.phar -o /usr/local/bin/composer && \
    chmod +x /usr/local/bin/composer

# Check if user and group doesn't exist before creating
RUN getent group "$GROUP_ID" || addgroup --gid "$GROUP_ID" user
RUN getent passwd "$USER_ID" || adduser --disabled-password --gecos '' --uid "$USER_ID" --gid "$GROUP_ID" user

USER "$USER_ID"
