# Device Microservice
Build Status: ![](https://github.com/UdamLiyanage/device-service/workflows/Go/badge.svg)
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
