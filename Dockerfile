FROM ubuntu:xenial

RUN  sed -i s@/archive.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list \
     && apt-get clean \
     && apt-get update --fix-missing

RUN \
    DEBIAN_FRONTEND=noninteractive \
    apt-get upgrade -y \
    && apt-get install --fix-missing -y \
            python3-pip \
            git \
            locales \
            libreoffice-common \
            libreoffice-writer \
            libreoffice-calc \
            libreoffice-impress \
    && ln -s /usr/bin/python3 /usr/bin/python \
    && pip3 install git+https://github.com/Supervisor/supervisor.git@master \
    && git clone --branch=0.8 --depth=1 https://github.com/unoconv/unoconv.git \
    && cp unoconv/unoconv /bin/unoconv \
    && chmod +x /bin/unoconv \
    && rm -rf unoconv \
    && apt-get -q -y remove libreoffice-gnome make git python3-pip \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*


RUN locale-gen de_DE.UTF-8
ENV LANG=de_DE.UTF-8 \
    LANGUAGE=de_DE:de \
    LC_ALL=de_DE.UTF-8

COPY gouno /go/bin/gouno
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf
COPY boot.sh ./
RUN chmod 777 boot.sh

EXPOSE 3000

ENTRYPOINT ["./boot.sh"]
