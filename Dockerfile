FROM golang as build_stage
COPY . /pebble
WORKDIR /pebble
RUN go build -o pebble_bin ./cmd/pebble/

FROM ubuntu
ENV PEBBLE="/root"
WORKDIR /root
COPY --from=build_stage /pebble/pebble_bin pebble
COPY pebble/layers/ ${PEBBLE}/layers
COPY signal_test/ signal_test/
RUN apt update -y && apt install -y python3 gcc
RUN cd signal_test/ && gcc capture.c -o capture && cd ..
ENTRYPOINT ["/root/pebble"]
