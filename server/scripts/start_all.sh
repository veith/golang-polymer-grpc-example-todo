#! /bin/bash

(trap 'kill 0' SIGINT; build/designer & build/transcoder & build/reverseproxy)