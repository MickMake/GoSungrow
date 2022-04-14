#!/usr/bin/env bashio
set -e

CONFIG_PATH="/data/options.json"

bashio::log.info "Setting up GoSungrow config ..."

export SUNGROW_HOST=$(bashio::config 'sungrow_host')
export SUNGROW_USER=$(bashio::config 'sungrow_user')
export SUNGROW_PASSWORD=$(bashio::config 'sungrow_password')
export SUNGROW_APPKEY=$(bashio::config 'sungrow_appkey')
export SUNGROW_DEBUG=$(bashio::config 'sungrow_debug')
export SUNGROW_TIMEOUT=$(bashio::config 'sungrow_timeout')

export SUNGROW_MQTT_HOST=$(bashio::config 'sungrow_mqtt_host')
export SUNGROW_MQTT_PORT=$(bashio::config 'sungrow_mqtt_port')
export SUNGROW_MQTT_USER=$(bashio::config 'sungrow_mqtt_user')
export SUNGROW_MQTT_PASSWORD=$(bashio::config 'sungrow_mqtt_password')

bashio::log.info "Writing GoSungrow config ..."

/usr/local/bin/GoSungrow config write \
	--host="${SUNGROW_HOST}" \
	--user="${SUNGROW_USER}" \
	--password="${SUNGROW_PASSWORD}" \
	--appkey="${SUNGROW_APPKEY}" \
	--timeout="${SUNGROW_TIMEOUT}s" \
	--mqtt-host="${SUNGROW_MQTT_HOST}" \
	--mqtt-port="${SUNGROW_MQTT_PORT}" \
	--mqtt-user="${SUNGROW_MQTT_USER}" \
	--mqtt-password="${SUNGROW_MQTT_PASSWORD}" \
	--debug="${SUNGROW_DEBUG}"

bashio::log.info "Login to iSolarCloud using gateway ${SUNGROW_HOST} ..."

/usr/local/bin/GoSungrow api login

bashio::log.info "Syncing data from gateway ${SUNGROW_HOST} ..."

/usr/local/bin/GoSungrow mqtt sync

