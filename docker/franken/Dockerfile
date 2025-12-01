FROM dunglas/frankenphp:builder-php8.4.15 AS builder

ARG USER_ID
ARG GROUP_ID

ENV GO_VERSION="1.25.0"

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
      wget \
      cmake && \
      rm -rf /var/lib/apt/lists/*

RUN rm -rf /usr/local/sconcur
RUN mkdir -p -m 777 /usr/local/sconcur

# go
RUN wget -P /tmp "https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz"

RUN tar -C /usr/local -xzf "/tmp/go${GO_VERSION}.linux-amd64.tar.gz"
RUN rm "/tmp/go${GO_VERSION}.linux-amd64.tar.gz"

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH" && mkdir -p /.cache && chmod -R 777 /.cache

RUN mkdir -p -m 777 /go/pkg/mod/cache/download
RUN chown $USER_ID:$GROUP_ID /go/pkg/mod/cache/download
RUN chmod -R 777 /go/

# watcher
WORKDIR /usr/local/sconcur
RUN rm -rf watcher && \
  git clone https://github.com/e-dant/watcher.git && \
  cd watcher && \
  cmake -B build && \
  cmake --build build && \
  cmake --install build && \
  echo "/usr/local/lib" | tee /etc/ld.so.conf.d/watcher.conf && \
  ldconfig

# xcaddy
RUN wget https://github.com/caddyserver/xcaddy/releases/download/v0.4.5/xcaddy_0.4.5_linux_amd64.deb && \
    dpkg -i xcaddy_0.4.5_linux_amd64.deb && \
    rm xcaddy_0.4.5_linux_amd64.deb

# composer
RUN curl -fsSL https://getcomposer.org/download/latest-stable/composer.phar -o /usr/local/bin/composer && \
    chmod +x /usr/local/bin/composer

RUN getent group "$GROUP_ID" || addgroup --gid "$GROUP_ID" user
RUN getent passwd "$USER_ID" || adduser --disabled-password --gecos '' --uid "$USER_ID" --gid "$GROUP_ID" user

USER "$USER_ID"

WORKDIR /app