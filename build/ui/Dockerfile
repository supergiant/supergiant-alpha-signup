from alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
ADD ./alpha-ui-linux-amd64 /supergiant-alpha-ui
CMD ["/supergiant-alpha-ui"]
