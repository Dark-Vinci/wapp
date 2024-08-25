#!/bin/sh
set -e
set -x

if [[ -z $1 && -z $2 ]]; then
    echo >&2 "Service name and migration name was not provided"
    exit
fi

if [[ -z $1 ]]; then 
    echo >&2 "Service name was not provided"
fi

if [[ -z $2 ]]; then 
    echo >&2 "Service migration was not provided"
fi

case $1 in
  "account")
    echo "about to cd into account service to generate migration"
    cd ../backend/account/migrations
    goose create "$2" go;;

  "chats")
    echo "about to cd into reactions service to generate migration"
    cd ../backend/chats/migrations
    goose create "$2" go
    go mod tidy;;

  "posts")
    echo "about to cd into posts service to generate migration"
    cd ../backend/posts/migrations
    goose create "$2" go
    create "$2" go;;

  "media")
    echo "about to cd into reactions service to generate migration"
    cd ../backend/reactions/migrations
    goose create "$2" go
    create "$2" go;;

  *)
    echo "No known service was chosen";;
esac