FROM python:3.6.4-alpine3.6

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

ENV LC_ALL=en_US.UTF-8 \
	LANG=en_US.UTF-8 \
	LANGUAGE=en_US.UTF-8

RUN apk add --no-cache \
        --virtual .build-deps \
        gcc \
        g++ \
        linux-headers \
        libc-dev \
    && apk add --no-cache \
        libreoffice-common \
        libreoffice-writer \
        libreoffice-calc \
        libreoffice-impress \
    && pip install circus \
    && rm -rf /var/cache/apk/* \
    && rm -rf /root/.cache/ \
    && apk del .build-deps

COPY unoconv /bin/unoconv
COPY circus.ini /etc/circus/circus.ini
COPY boot.sh ./
COPY gouno /go/bin/gouno
RUN chmod 777 boot.sh \
    && chmod +x /go/bin/gouno \
    && chmod +x /bin/unoconv \
    && ln -s /usr/bin/python3 /usr/bin/python

EXPOSE 3000

ENTRYPOINT ["./boot.sh"]
