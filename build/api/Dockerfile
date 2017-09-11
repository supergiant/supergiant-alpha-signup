from alpine
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
ADD ./alpha-api-linux-amd64 /supergiant-alpha-api
CMD ["/supergiant-alpha-api"]
