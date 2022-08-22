# Introduction

`SIMPLE E-COMMERCE` <br/>
This is prototype of e-commerce microservices backend system. <br/>
This microservices backend system built with Go programming language with graphQL as a gateway server and gRPC as a services.
<br/>

System design and usecase diagram of this backend system is provided on `docs` folder.
<br><br>

# Getting Started

Prerequisites:

1. Make sure you have `Make` in your local environment to run app using `Makefile`
2. Make sure you have `docker` in your local environment
3. Make sure you're not using port `8989` and `9000` on your other application while running this project
   <br/><br/>

# How to Run

1. Move to base directory project on top of all services (align with docker-compose.yml)
2. If you have `Make` run command `make up`, if not run command `docker-compose up --build`
3. Wait until docker image has been built
4. Run example request from `e-commerce.postman_collection.json` on your postman app or open your chrome browser with this url `http://localhost:8989/` then copy the example request into graphql playground schema query
5. To stop this application, run command `make down` or `docker-compose down`
   <br/><br/>

# How to Test

This project is consists of unit tests, to run the test you can go into unit test files `*_test.go` on your vscode and click on `run package tests`
