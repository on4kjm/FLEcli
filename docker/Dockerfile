FROM alpine:3.12 AS builder

# Copy the executable
COPY FLEcli /FLEcli_bin/FLEcli
RUN chmod +x /FLEcli_bin/FLEcli


FROM scratch

COPY --from=builder /FLEcli_bin/FLEcli /FLEcli_bin/FLEcli

# Create the data directory and set it as the defautl directory
WORKDIR /FLEcli_data

ENTRYPOINT ["/FLEcli_bin/FLEcli"]
CMD ["help"]