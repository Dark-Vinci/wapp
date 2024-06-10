#!/bin/sh

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

  "reactions")
    echo "about to cd into reactions service to generate migration"
    cd ../backend/reactions/migrations
    goose create "$2" go
    create "$2" go;;

  "utils")
      echo "about to cd into utils service to generate migration"
      cd ../backend/utils/migrations
      goose create "$2" go
      create "$2" go;;

  *)
    echo "No known service was chosen";;
esac