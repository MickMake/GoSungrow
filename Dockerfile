
FROM alpine:latest

MAINTAINER MickMake <embed@mickmake.com>

USER root
ARG GO_REPO_TOKEN
ENV GO_REPO_TOKEN ${GO_REPO_TOKEN}

# SUNGRO_GIT_TOKEN SUNGRO_HOST SUNGRO_ID SUNGRO_PASSWORD SUNGRO_SECRET SUNGRO_USER SUNGRO_DIFF_CMD
ARG GO_REPO_TOKEN
ENV GO_REPO_TOKEN ${GO_REPO_TOKEN}
ARG SUNGRO_HOST
ENV SUNGRO_HOST ${SUNGRO_HOST}
ARG SUNGRO_USER
ENV SUNGRO_USER ${SUNGRO_USER}
ARG SUNGRO_PASSWORD
ENV SUNGRO_PASSWORD ${SUNGRO_PASSWORD}
ARG SUNGRO_ID
ENV SUNGRO_ID ${SUNGRO_ID}
ARG SUNGRO_SECRET
ENV SUNGRO_SECRET ${SUNGRO_SECRET}
ARG SUNGRO_GIT_REPO
ENV SUNGRO_GIT_REPO ${SUNGRO_GIT_REPO}
ARG SUNGRO_GIT_DIR
ENV SUNGRO_GIT_DIR ${SUNGRO_GIT_DIR}
ARG SUNGRO_GIT_TOKEN
ENV SUNGRO_GIT_TOKEN ${SUNGRO_GIT_TOKEN}
ARG SUNGRO_DIFF_CMD
ENV SUNGRO_DIFF_CMD ${SUNGRO_DIFF_CMD}
ARG TZ
ENV TZ ${TZ}

COPY dist/GoSungro_linux_amd64/GoSungro /usr/local/bin/GoSungro
COPY .ssh/ /root/.ssh/
RUN chmod a+x /usr/local/bin/GoSungro && \
	chmod 500 /root/.ssh && \
	chmod 400 /root/.ssh/gosungro_rsa /root/.ssh/gosungro_rsa.pub && \
	apk add --no-cache colordiff tzdata
#	echo '00 07  *  *  *    /usr/local/bin/GoSungro sync default' > /etc/crontabs/root

#ENTRYPOINT ["/usr/local/bin/GoSungro"]
#CMD ["crond", "-f", "-l", "2", "-L", "/var/log/cronlogs"]
CMD ["/usr/local/bin/GoSungro", "cron", "run", "00", "07", ".", ".", ".", "sync", "default"]

