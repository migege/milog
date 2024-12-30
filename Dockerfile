ARG HTTP_PROXY
ARG HTTPS_PROXY
ARG NO_PROXY

FROM library/golang AS builder

ENV APP_DIR $GOPATH/home/aib/projects/github.com/migege/milog
RUN mkdir -p $APP_DIR
WORKDIR $APP_DIR

ADD . $APP_DIR

ENV HTTP_PROXY $HTTP_PROXY
ENV HTTPS_PROXY $HTTPS_PROXY
ENV NO_PROXY $NO_PROXY

# Recompile the standard library without CGO
RUN CGO_ENABLED=0 go install -a std

# Compile the binary and statically link
RUN cd $APP_DIR && \
  CGO_ENABLED=0 \
  go build -ldflags '-d -w -s'

ENV HTTP_PORT=
ENV HTTPS_PORT=
ENV NO_PROXY=

FROM alpine:3.21.0 as prod
ENV TZ Asia/Shanghai
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.cloud.tencent.com/g' /etc/apk/repositories \
  && apk add ca-certificates \
  && apk add tzdata && cp /usr/share/zoneinfo/${TZ} /etc/localtime && echo ${TZ} > /etc/localtime \
  && rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=builder /go/home/aib/projects/github.com/migege/milog/milog .
COPY --from=builder /go/home/aib/projects/github.com/migege/milog/static ./static
COPY --from=builder /go/home/aib/projects/github.com/migege/milog/views ./views

EXPOSE 9900
CMD ["./milog"]
