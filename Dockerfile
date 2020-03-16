FROM scratch

COPY k8s-graceful-shutdown /
ENTRYPOINT ["/k8s-graceful-shutdown"]
