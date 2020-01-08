#This is a builder image and is destroyed after we build the binary

# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.13 as builder
SHELL ["/bin/bash", "-c"]

# BEGIN private repo
# If supplied, should be a base64 encoded string containing an ssh private key
#ARG SSH_PRIVATE_KEY
## If supplied, should be in the format `http[s]://githost.com`.  Will be replaced with `git@githost.com`
#ARG GIT_CONFIG_HOST
#ARG GIT_CONFIG_USER=git
## This will load the SSH key to memory, reconfigure git to use SSH for the specified hosts and finally verify that connections to that host succeed over SSH
#RUN if [ ! -z "$SSH_PRIVATE_KEY" ];then mkdir -p ~/.ssh && echo ${SSH_PRIVATE_KEY} | base64 -d > ~/.ssh/id_rsa && chmod 400 ~/.ssh/id_rsa;fi
#RUN if [ ! -z "$GIT_CONFIG_HOST" ];then export MODIFIED=$(echo ${GIT_CONFIG_HOST} | sed 's%http[s]://%'${GIT_CONFIG_USER}'@%g; s%/%:%g;') && git config --global url."${MODIFIED}".insteadOf "${GIT_CONFIG_HOST}" && ssh -T -o "StrictHostKeyChecking=no" "${MODIFIED/:/}";fi
# END private repo

ENV PATH /go/bin:$PATH
WORKDIR /go/src/github.com/gamezop
ADD . .
ENV GO111MODULE on

RUN go build -o ./bin/queue -i ./cmd/queue/


# final image
FROM ubuntu as queue-server
RUN apt-get update && apt-get install ca-certificates -y
COPY --from=builder /go/src/github.com/gamezop/bin/queue /usr/bin
CMD ["queue"]