# booking-calendar-be

docker compose up -d --build

# /src/internal/modules
## **MODULES:**
- constant: define enum, variable common
- controllers: receive request
- dto: data object transfer from controller -> service -> repository
- filters: filter input for model query
- mappers: mapping data model to object response
- entities: define entity mapping document NoSQL
- models: define modeling mapping table RDS (MySQL)
- repositories: define repository layer, it is operation with model
- requests: data send from client to server
- responses: data send from server to client
- errors: error code response to client
- scopes: build query scope
- routers: define router for module
- service: handle business layer
- formatters: format request params before handle on service

**You can make module struct with command:**
- cd src/internal/modules/
- mkdir -p user && find store -mindepth 1 -maxdepth 1 -type d | while read dir; do mkdir -p "user/$(basename "$dir")"; done
  (sample: make module user from module store)

**How to register service:
- you can register service: src/pkg/boostrap/reverse_proxy.go
- add config port to config/local.yml
- modify: src/internal/core/constant/service_register.go