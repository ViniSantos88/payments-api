# payments-api
Aplicação para armazenar as transações bancárias

h1. DEVELOPMENT

##Requirements:
- go 1.16
- mockgen:
  - GO111MODULE=on go get github.com/golang/mock/mockgen@v1.6.0
- testify:
  - go get github.com/stretchr/testify
- Swagger
  - go get -u github.com/go-openapi/runtime/middleware
  - go get -u github.com/go-swagger/go-swagger/cmd/swagger  
- docker:
  - https://www.docker.com/get-started  
- docker-compose:
  - https://docs.docker.com/compose/install/
- sqlMock
  - go get github.com/DATA-DOG/go-sqlmock  
- database driver sql
  - go get github.com/jmoiron/sqlx  
- database(postgresql)
  - go get -u github.com/lib/pq   

##Unit testing
- HOW TO GENERATE MOCKS
  -INSTALL MOCK
  - GO111MODULE=on go get github.com/golang/mock/mockgen@v1.4.4
    - Infra Layer:   
    - Repositories Layer:
      mockgen -source=application/repositories/payments.go -destination=__test__/mocks/repositories/payments_mock.go -package=repositoriesmocks
    - Services Layer:
      mockgen -source=application/services/payments.go -destination=__test__/mocks/services/payments_mock.go -package=servicesmocks


HOW TO ANALYZE WHAT IS MISSING FROM THE TEST COVERAGE 
    - t="/tmp/go-cover.$$.tmp"
    - go test ./... -coverprofile=$t $@ && go tool cover -html=$t && unlink $t

HOW TO GENERATE AND TO VALIDATE THE SWAGGER
    - swagger generate spec -o ./swagger.yaml --scan-models
    - swagger validate ./swagger.yaml

HOW TO RUN UNIT TEST
- Run
    - ./RunTests.sh

HOW TO RUN DOCKER LOCALLY
docker-compose up -d   

HOW TO REMOVE IMAGE
docker-compose down --rmi all 



