FROM golang:1.24-bullseye

# Install common tools (git, curl, etc.)
RUN apt-get update && apt-get install -y --no-install-recommends \
    git curl ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Set workdir
WORKDIR /workspace
