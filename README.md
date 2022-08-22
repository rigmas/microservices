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
3. Wait until docker image has built
4. Run example request from `e-commerce.postman_collection.json` on your postman app or open your chrome browser with this url `http://localhost:8989/` then copy the example request into graphql playground schema query
5. To stop this application, run command `make down` or `docker-compose down`
   <br/><br/>

# How to Test

This project is consists of unit tests, to run the test you have to install `mockery` first. See https://github.com/vektra/mockery. <br>
Then, you can go into unit test files `*_test.go` on your vscode and click on `run package tests` <br>

Results:
<br>

1. `Gateway`
   <br><br>
   <img width="1252" alt="gateway" src="https://user-images.githubusercontent.com/62687089/185966183-d23c89c8-97f4-45fd-9cbd-57ab8ac5457b.png">

2. `Customer`
   <br><br>
   <img width="1255" alt="customer" src="https://user-images.githubusercontent.com/62687089/185966903-e65df9d3-80ff-4f01-87a0-bc9112555b73.png">

3. `Product`
   <br><br>
   <img width="1252" alt="product" src="https://user-images.githubusercontent.com/62687089/185967060-81407d24-a96b-4f57-8f1c-3c4a0fd61661.png">

4. `Order`
   <br><br>
   <img width="1239" alt="order" src="https://user-images.githubusercontent.com/62687089/185967196-49e4cd1a-4348-4640-92b3-2d6f6953b588.png">
