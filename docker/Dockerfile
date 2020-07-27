FROM alpine:3.12 AS builder

# Copy the executable and make it executable
COPY FLEcli /temp_bin/FLEcli
RUN chmod +x /temp_bin/FLEcli

#we switch to the bare image template as we want the lightest image as possible
FROM scratch

#We retrieve the prepared binary and place it in the appropriate location
COPY --from=builder /temp_bin/FLEcli /FLEcli_bin/FLEcli

# Create the data directory and set it as the defautl directory
WORKDIR /FLEcli_data

ENTRYPOINT ["/FLEcli_bin/FLEcli"]
CMD ["help"]