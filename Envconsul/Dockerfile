FROM alpine
WORKDIR /envconsul
RUN wget -c https://releases.hashicorp.com/envconsul/0.9.0/envconsul_0.9.0_linux_amd64.tgz && \
    tar -xvzf envconsul_0.9.0_linux_amd64.tgz
FROM scratch

COPY --from=0 /envconsul /