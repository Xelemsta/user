## Prerequisites

* docker-compose

## About application

### Remaining tasks undone

* Spawn a swagger UI for API from docker-compose
* No test on kafka
* listing users does not return meta object (this + total)

### Build application

  ```sh
  make build
  ```

### Start application

  ```sh
  make run
  ```

### Test application

  ```sh
  make test
  ```

### Technical environment

* docker
* go-swagger
* apache kafka
* postgresql

### Choices/Assumptions

First of all, all my choices are mostly going towards stabilized and widely used technologies because it mostly offers a large community that can help on that matters if necessary and an exhaustive documentation.

About the API, i decided to choose go-swagger to build it because it is "contract-first" and brings mostly all the common features we would need to build that application including skeleton generation + interactive documentation. From that we will be able to focus mostly on business code (implemeting handlers) and routes definitions.

To store users, i rely on postgresql which allow us to set up efficiently control on data.

To notify the whole IS about a change on users, i decided to go with kafka to handle events producing.

### Extensions/Improvements

* Storing configuration/sensible secrets with a secret management solution (Vault for example).

* Sending application logs inside ElasticSearch/OpenSearch solution (for example) to be able to monitor the system (then crafting monitoring application that request our elasticsearch/opensearch cluster or using application like ElastAlert). It could be usefull to compute response time for example.

* Using a context environment where you can set a unique id that could be forwarded to each microservices communicating between each other. It will helps to easier track what's happened on a specific workflow or action (by getting all microservices involved).