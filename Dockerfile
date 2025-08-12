FROM scratch

ENV PATH "/bin"
COPY easy-ca-cli /bin/easy-ca-cli

ENTRYPOINT ["easy-ca-cli"]
