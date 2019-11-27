FROM alpine
WORKDIR /consul-template
RUN apk update && \
    apk add ca-certificates wget && \
    update-ca-certificates && \
    wget -c https://releases.hashicorp.com/consul-template/0.23.0/consul-template_0.23.0_linux_amd64.tgz && \
    tar -xvzf consul-template_0.23.0_linux_amd64.tgz
FROM scratch

COPY --from=0 /consul-template /
ENTRYPOINT ["/consul-template","-config=/etc/config/consul-template-config.hcl","-once"]