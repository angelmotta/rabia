FROM ubuntu:18.04
RUN apt-get update && apt-get install -y \
  git \
  vim \
  wget \
  sudo \
  inetutils-ping \
  iproute2 \
  openssh-server \
  net-tools \
  telnet
# Download/Copy Rabia project to the default path
ENV HOME /root
RUN mkdir -p $HOME/go/src
WORKDIR $HOME/go/src
# RUN git clone https://github.com/haochenpan/rabia.git
# COPY rabia code to the container
COPY . rabia/
# Install Rabia dependencies
# WORKDIR $HOME/go/src/rabia/deployment
RUN sed -i '/go_tar=go1.20.2.linux-amd64.tar.gz/c\go_tar=go1.20.2.linux-arm64.tar.gz # the version of Golang to be downloaded in install_go' rabia/deployment/install/install.sh
RUN chmod +x rabia/deployment/install/install.sh
RUN /bin/bash -c 'rabia/deployment/install/install.sh'
# Post requirements rabia
ENTRYPOINT service ssh start && /bin/bash