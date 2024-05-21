# Video Annotator

The Video Annotator service facilitates the creation, updating, and deletion of videos. 
It also supports the creation, updating, and deletion of annotations with specified durations.
---

[Go](https://go.dev/)  
[Docker](https://www.docker.com/)

---

**Contents**

1. [Setup](#setup)
1. [Assumptions](#assumptions)
1. [Quality](#quality)
1. [Future Scope](#future-scope)

---

### Setup ###

1. Install Golang and ensure Go project can be run in the system.
1. Install docker to run the application.
1. Clone the repo in local
1. Use the command `docker-compose up` to start the service.
1. Check the successful server initialisation message in docker container logs.
1. Use the postman collections from the file `video_annotator/postman.json`.
1. Use the api-key `secret-api-key` for API authentication.

---


### Assumptions ###
1. Videos with same URL cannot be created.
1. Annotations can exist only with a video.
1. Annotations can overlap with another annotation for a video but cannot exist with another having the same duration.

---

### Quality ###

Unit Test cases are available in `usecase/video_test` file.

Mocks are available in `mocks/video_store` file.

---

### Future Scope ###

1. Proper implementation of authorisation with roles and permissions.
1. Implementation of transactions in repo store. 
1. Documentation of APIs in OpenAPI(Swagger) 
1. Improving unit test coverage.

---
