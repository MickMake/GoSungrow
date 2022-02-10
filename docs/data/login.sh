#!/bin/bash

USER_AGENT="$(cat user_agent.data)"

curl \
	-X POST -H 'Content-Type: application/json' -A "${USER_AGENT}" \
	--data-binary "@login.post" \
	https://augateway.isolarcloud.com/v1/userService/login | tee login.response

