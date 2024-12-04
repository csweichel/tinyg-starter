FROM --platform=amd64 docker.io/tinygo/tinygo:latest@sha256:79c73c4246a1079e648d6e8425bae4654f26d2baed24c96537a8893b41c04d26

# install jq
USER root
RUN apt-get update && apt-get install -y jq && apt-get clean && rm -rf /var/lib/apt/lists/*

RUN echo 'eval "$(tinygo-autocmpl --completion-script-bash)"' >> ~/.bashrc
