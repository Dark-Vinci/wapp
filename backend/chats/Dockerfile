FROM golang:1.22-alpine

WORKDIR /go/src/github.com/dark-vinci/wapp/chats

ARG port=5050
ARG app_env=development

ENV PORT="$port"
ENV APP_ENV="$app_env"

COPY sdk /go/src/github.com/dark-vinci/wapp/sdk
COPY chats /go/src/github.com/dark-vinci/wapp/chats

RUN go mod tidy
RUN go get ./...
RUN go build
RUN go get github.com/pilu/fresh

CMD if [[ "$APP_ENV" -eq production ]]; then \
        chats \
    else \
        fresh \
    fi

EXPOSE "$PORT"