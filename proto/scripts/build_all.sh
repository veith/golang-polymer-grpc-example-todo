#! /bin/bash

./scripts/mid2proto.sh auth
./scripts/build.sh auth

./scripts/mid2proto.sh healthcheck
./scripts/build.sh healthcheck

./scripts/mid2proto.sh task
./scripts/build.sh task


./scripts/mid2proto.sh tag
./scripts/build.sh tag

