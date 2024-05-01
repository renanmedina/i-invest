#!/bin/bash
[[ -s "$GVM_ROOT/scripts/gvm" ]] && source "$GVM_ROOT/scripts/gvm"
gvm use go1.22
gow run cmd/webserver.go 2>> logs.log