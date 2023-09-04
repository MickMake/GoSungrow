## [3.0.7] - 2023-05-10
### Features

- Alpha support for Modbus, (direct connect to your Sungrow inverter).


### Fixes

- Fixup ResultData.result_data.org_id error.


## [3.0.6] - 2023-05-10
### Features

- Added extra sub-commands:
```
show ps save
show device save
show template save
show point save
show point ps-save
show point device-save
show point template-save
```

### Fixes

- HA binary_sensor not being set correctly.


## [3.0.5] - Not released


## [3.0.4] - 2023-01-10
### Features

- Now supports the Energy Dashboard of HA!


## [3.0.3] - 2022-12-23
### Features

- Now supports the Energy Dashboard of HA!

### Fixes

- Fix double virtual.virtual entries.
- Fixup float value determination logic.
- CacheTimeout matches fetch schedule.
- Fixed unit type guessing - "Template error: float got invalid input"
- Fix battery class - correct point used for HA.
- Added more DeviceClass types for HA - will affect long term stats as units will change on some points.


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
