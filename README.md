## Dispatch System

Uber-like service which dispatches drivers according to client trip requests.

### Requirements

- Drivers send periodic updates on their position and status.
- Clients request trips for a given location and must be assigned the closest available driver
- Clients with ongoing trip requests receive periodic updates on their driver's location and status
- Do not use swagger

### Next steps (in no particular order)

- ~~Unit tests~~
- ~~Replace the switch case in driverRepository.handleRequestChannel with a Registry~~
- Periodic updates to the passenger upon a successful trip request
- Mock driver/passenger applications to visualize working application

? Assign a driver to a passenger only upon driver acceptance

### Testing

The content of JSON files can be sent via Curl using the following command:

**Register a new driver:**

`curl -X POST -H 'Content-Type: application/json' -d "@./driver_info_generic.json" "localhost:8089/api/driver/updateInfo"`

**Update driver status and location**

`curl -X POST -H 'Content-Type: application/json' -d "@./update_driver_generic.json" "localhost:8089/api/driver/update"`

**Request trip**

`curl -X POST -H 'Content-Type: application/json' -d "@./trip_request_generic.json" "localhost:8089/api/trips"`

### 11/11/2021

- Implemented registry pattern in driverRepository.handleRequestChannel

### 25/11/2021

- Fixed `TestDriverRepository_ProcessDriverInfo`, other tests are still not working.