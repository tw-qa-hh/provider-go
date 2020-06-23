#!/bin/sh

curl -LO https://github.com/pact-foundation/pact-provider-verifier/releases/download/v1.30.1-1/pact-provider-verifier-1.30.1-1-linux-x86_64.tar.gz
tar xzf pact-provider-verifier-1.30.1-1-linux-x86_64.tar.gz
cd pact-provider-verifier-1.30.1-1-linux-x86_64/bin
./pact-provider-verifier help

