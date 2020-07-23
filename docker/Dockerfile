FROM alpine:3.12.0

# Copy the executable
RUN mkdir /FLEcli
COPY FLEcli /FLEcli/
RUN chmod +x /FLEcli/FLEcli

# Create the data directory and set it as the defautl directory
RUN mkdir /FLEcli_data
WORKDIR /FLEcli_data

ENTRYPOINT ["/FLEcli/FLEcli"]