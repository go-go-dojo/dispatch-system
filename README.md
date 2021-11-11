## Dispatch System

Uber-like service which dispatches drivers according to client trip requests.

### Requirements

- Drivers send periodic updates on their position and status.
- Clients request trips for a given location and must be assigned the closest available driver
- Clients with ongoing trip requests receive periodic updates on their driver's location and status
- Do not use swagger

### Next steps (in no particular order)

- Replace the switch case in driverRepository.handleRequestChannel with a Registry
- Implement periodic updates to the client upon a successful trip request

### 11/11/2021

- Implemented registry pattern in driverRepository.handleRequestChannel