## [3.0.2] - 2022-12-21
### Features

- Control of GoSungrow via MQTT; gosungrow_option_fetchschedule, gosungrow_option_loglevel, gosungrow_option_sleepdelay & gosungrow_option_servicestate
- Updated MQTT message types.
- Generate Lovelace dashboards for your ps_ids - `GoSungrow ha lovelace`

### Fixes

- Incorrect units on points p13119 & p13149.


## [3.0.1] - 2022-12-18
### Features

- Skip version number, to align with HA addon.

### Fixes

- Skip version number, to align with HA addon.


## [3.0.0] - 2022-12-15
### Features

- Support better api changes - This allows a much quicker change in my code when the api changes.
- Support multiple devices - Previous versions only allowed query of one device.
- Improve cli for queries - Now can fetch endpoints in either "list" or "table" formats, the latter showing point data in a similar fashion to pivot tables in spreadsheets.
- Multiple output formats supported - not only csv, png and ascii, but markdown, xml, HTML, sql and plain MQTT.

### Fixes

- Lots of bug fixes. :-).
