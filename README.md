# Device Microservice
Build Status: ![](https://github.com/UdamLiyanage/device-service/workflows/Go/badge.svg)
[![Travis](https://travis-ci.com/UdamLiyanage/device-service.svg?branch=master)](https://travis-ci.com/UdamLiyanage/device-service)
***
Device Configuration Microservice for IoT Platform. This service is responsible for the following features:
*Create devices
*Read devices
*Update devices
*Delete devices

***
## Document Structure for Device - Revision 1
Below is the structure of the JSON document that holds device data
```
{
	"name": "Device Name",
	"serial": "Device Serial",
	"configurations": [{
			"configuration_id": "Configuration ID",
			"configuration_type": "Type"
		},
		{
			"configuration_id": "Configuration ID",
			"configuration_type": "Type"
		},
		{
			"configuration_id": "Configuration ID",
			"configuration_type": "Type"
		}
	]
}
```
## Document Structure for Device - Revision 2
```
{
	"_id": "objectID",
	"serial": "device_serial",
	"uuid": "UUID",
	"device_definition_id": "device_definition_id",
	"name": "device_name",
	"firmware_version": "firmware_version",
	"last_firmware_update": "last_firmware_update_datetime",
	"created_at": "created_datetime",
	"updated_at": "updated_datetime"
}
```
