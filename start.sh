#!/bin/sh

go generate && \
  templ generate && \
  echo "Running server..." && \
  go run .