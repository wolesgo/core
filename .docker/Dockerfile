FROM alpine:3.13.6
LABEL org.opencontainers.image.authors="waldiirawan127@gmail.com"

ARG architecture=amd64
ARG version=1.18
ENV ARCHITECTURE $architecture
ENV VERSION $version

RUN apk add --no-cache wget

RUN wget https://go.dev/dl/go${VERSION}.linux-${ARCHITECTURE}.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go${VERSION}.linux-${ARCHITECTURE}.tar.gz

ENV PATH="$PATH:/usr/local/go/bin"

CMD ["go", "test", "."]
