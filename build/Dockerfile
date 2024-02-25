FROM golang:1.18-alpine as base
WORKDIR /app

# Install dependencies not included with alpine
RUN apk update && apk add --no-cache ca-certificates git curl make jq moreutils rsync && update-ca-certificates

# Configure go
ENV GO111MODULE on
ENV GOOS linux
ENV CGO_ENABLED 0
ENV GOROOT /usr/local/go
ENV GOPATH /gopath
ENV GOBIN ${GOPATH}/bin
ENV PATH /usr/local/go/bin:/gopath/go/bin/:/gopath/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin

# Install golangci-lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.41.1

FROM base as devcontainer
WORKDIR /app

# Devcontainer git setup 
# ARG EVT_GITHUB_EMAIL
# ARG EVT_GITHUB_NAME
# RUN if [[ "${EVT_GITHUB_NAME}" ]] ; then git config --global user.name ${EVT_GITHUB_NAME}; fi
# RUN if [[ "${EVT_GITHUB_EMAIL}" ]] ; then git config --global user.email ${EVT_GITHUB_EMAIL}; fi

# Install go tools
RUN go install golang.org/x/tools/gopls@latest

# Install zsh
RUN apk add zsh
# Run the installation script  
RUN wget https://github.com/robbyrussell/oh-my-zsh/raw/master/tools/install.sh -O - | zsh || true
# Install delve and air
RUN go install github.com/go-delve/delve/cmd/dlv@latest && go install github.com/cosmtrek/air@latest
# Copy delve dlv directory to delve for VSCode debugger
RUN cp -r ${GOBIN}/dlv ${GOBIN}/dlv-dap


FROM base as build
WORKDIR /app

# Download go modules
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download && go mod verify

# Copy the rest of the project in now that we've cached that layer
COPY . .

# Build our production release
RUN go build -o release/go-solid-design main.go

FROM alpine as production
WORKDIR /app

# Copy web folder to load templates
# COPY --from=build app/web web
# Copy executable files 
COPY --from=build app/release/go-solid-design app/release/go-solid-design

EXPOSE 80

ENTRYPOINT app/release/go-solid-design
