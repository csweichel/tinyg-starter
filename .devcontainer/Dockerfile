FROM --platform=amd64 tinygo/tinygo-dev:latest@sha256:be8a85a6f6454240d9ed1bb8f504717020abfe54e96be95b94ec4f00655d6ef1

# install jq
RUN apt-get update && apt-get install -y jq && apt-get clean && rm -rf /var/lib/apt/lists/*

RUN echo 'eval "$(tinygo-autocmpl --completion-script-bash)"' >> ~/.bashrc
