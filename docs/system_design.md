# System Design

![system_design](https://user-images.githubusercontent.com/62687089/185913618-da411335-e639-4546-92b6-40bcbe006ed7.png)

I use Saga architecture pattern while building this prototype simple e-commerce backend system.
This backend system consists of 1 GraphQL gateway server with 3 gRPC services. <br>
`Gateway` can be use as an interface from client (orchestrator) to connect and centralize each services. `Gateway` also can be use for middleware and load balancer later. <br>
I divided services according to DDD model into 3 services and will connect with each database respectively. <br>

1. Customer <br>
   `Customer` service can be use to register user. <br>
   `Customer` can communicate with `Order` directly through gRPC to see list order that have been made by customer previously.
2. Product <br>
   `Product` service can be use to access any information or data about product (e.g. product title, price, stocks, category, etc.)
3. Order <br>
   `Order` service can be use for customer that want to buy products. <br>
   `Order` can communicate to `Product` through gRPC to calculate the amount of payment and provide any detail information of product that has bought by customer. <br>
   `Order` also can connect to `Customer` through gRPC to validate which customer that made the order and also can be use to see order that has been made by customer previously.
