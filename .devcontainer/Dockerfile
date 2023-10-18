FROM golang:1.20 as builder

# Install Dependencies
WORKDIR /tmp
RUN apt-get update && apt-get install -y \
  autoconf \
  bison \
  build-essential \
  curl \
  flex \
  libreadline-dev \
  libncurses5-dev \
  m4

# Install BIRD from sources
RUN curl -SLO https://gitlab.nic.cz/labs/bird/-/archive/master/bird-master.tar.gz \
  && tar -zxvf bird-master.tar.gz \
  && rm bird-master.tar.gz

WORKDIR /tmp/bird-master
RUN autoreconf --install \
  && ./configure \
  && make \
  && make install
RUN mkdir -pv /etc/bird \
  && curl -sSLo /etc/bird/bird.conf https://gitlab.nic.cz/labs/bird/-/raw/master/doc/bird.conf.example

USER root
