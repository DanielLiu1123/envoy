# docker run -it --rm --name=envoy -p 10000:10000 -p 9901:9901 -v $(pwd)/cluster.yaml:/cluster.yaml -v $(pwd)/listener.yaml:/listener.yaml -v $(pwd)/envoy.yaml:/etc/envoy/envoy.yaml -v $(pwd)/upstream/proto.pb:/upstream/proto.pb envoyproxy/envoy:v1.21.1

FROM envoyproxy/envoy:v1.21.1

COPY envoy.yaml /etc/envoy/envoy.yaml
COPY cluster.yaml /cluster.yaml
COPY listener.yaml /listener.yaml
