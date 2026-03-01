FROM golang:1.25.3-trixie AS builder

RUN apt-get update --fix-missing && \
    apt-get install -y --no-install-recommends \
    zip \
    vim \
    unzip \
    curl \
    bash \
    tzdata \
    libmagickwand-dev \
    qpdf \
    wget \
    less \
    default-mysql-client && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

COPY --from=node:22.14.0-bookworm-slim /usr/local/ /usr/local/

RUN npm install -g \
    @dotenvx/dotenvx \
    @commitlint/cli \
    @commitlint/config-conventional \
    semantic-release \
    @semantic-release/gitlab \
    @semantic-release/commit-analyzer \
    @semantic-release/release-notes-generator \
    @semantic-release/exec @semantic-release/git \
    conventional-changelog-conventionalcommits && \
    sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin && \
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.5.0

RUN go install github.com/air-verse/air@latest

FROM builder

RUN \
    set -xe \
    && delgroup dialout

# Set working directory
WORKDIR /var/www/

ARG DOCKER_HOST_UID
ARG DOCKER_HOST_GID

RUN \
    set -xe \
    && groupadd --gid "$DOCKER_HOST_GID" "app" \
    && useradd --uid "$DOCKER_HOST_UID" --gid "$DOCKER_HOST_GID" --create-home --shell /bin/bash "app"

USER $DOCKER_HOST_UID:$DOCKER_HOST_GID

ARG HTTP_PORT=9000
# Expose port
EXPOSE $HTTP_PORT
