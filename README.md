# Item Sale Application

This is an application that I wrote to practice writing cloud native applications using the microservices architecture and gRPC. Using the UCLA Free and for Sale page as inspiration, this application allows users to register products for sale and pickup. In the future, I will create a REST implementation of this application as well. 


# Technologies

• Go (1.10) 

• MongoDB

• React

• Go-micro

• Protoc

• Docker/Docker-Compose

• Kubernetes

• Jenkins

• JWT Authorization

## Why Go? 

The main reason why GoLang is a great language for building microservices is the small memory footprint it leaves. Go services use 
considerably less RAM compared to Spring-boot based services providing benefits such as lowering cloud provider bills and increasing horizontal scalability. 

## RPC vs REST

### REST

As REST is currently one of the most preferred microservice architectures today, we will start with a definition for REST.

**Definition**

REST stands for Representational State Transfer and is a set of architectural constraints that are widely used for web services.

__Architectural Constraints__

• Client-server architecture: By separating the UI from the concerns of the data storage, components of web services can evolve separately and independently while improving scalability through simplicity.

• Statelessness: No client context is stored on the server between requests. This simplifies code flow and logic. The request from the client contains the current session state.

• Cacheability: Clients and intermediaries may cache responses to improve performance and scalability. However, responses must directly define themselves as cacheable or not.

• Layered System: Intermediary servers may intercept calls, meaning that the client does not know whether they are interacting directly with the end/final server or whether they are interacting with some sort of proxy.

• Uniform Interface: We must use interfaces to decouple classes from the implementation to provide appropriate opaque encapsulation and information hiding. To do this, there are four key standards that we must follow.

1) Identification of resources - normally done through a URI that is exposed to identify a resource.
2) Manipulation of resources - use the HTTP standard to describe communication (GET, POST, etc) 
3) Self-descriptive messages - Make sure that the client doesn;t have to know about the application specific data structures that are used. This is done through using standard MIME types.
4) Hypermedia as the engine of application state (HATEOAS) - Use hyperlinks and URI structures to decouple the client. Annotate these hyperlinks for simplicity and use on the client side.

## Advantages

• Easy to understand and widely used

• Testing tools are widely available and changes are relatively easy due to the architectural constraints (statelessness, loose coupling, etc.)

• Status codes for error and success responses are well defined

## Disadvantages

• Duplex streaming is impossible, tough to manage concurrent requests

• Versioning is necessary whenever an API contract needs to be changed

• Hard to get multiple resources in one request

## Then Why RPC?

Along with the fact that I wanted to practice using RPC's, there are more practical reasons to my choice.

**Protobuf vs JSON/XML**

• Protobuf messages sent through RPC's are generally very efficient and tightly packed, while JSON follows a strict textual format.

• Although you can compress JSON, you will loose the benefit of the textual format it has, while also complicating code.

• Because protobuf provides binary data, it is also extremely lightweight.

**HTTP 1.1 vs HTTP/2**

Overall, HTTP 1.1 used by standard REST calls have massive latency issues while also increasing overhead per request. 

**gRPC**

• This framework for RPC's is language agnostic, meaning that a grpc server written in one language can handle interactions with client calls from different languages. 

• Simplicity: Creating servers with gRPC is extremely simple, as we just need to specify the service and payload definitions in the Protocol Buffer .proto file, generate gRPC code by compiling with `protoc`, and then run the server. 

In the world of microservices, RPC is starting to become more and more dominant due to the numerous advantages it holds over REST.


## Go-Micro

When building microservices, it is important to establish some sort of communication between the different services that are running. Popular choices include RabbitMQ, NATS.io, ArcSight, etc. 

Because these popular third-party services lose the use of protobuf for something like JSON or XML, we lose the ability to communicate using binary streams, increasing the performance overhead vs. gRPC. In our case, we want to take full advantage of the light nature of gRPC calls so we use go-micro's built in pub-sub layer.

Go-micro also has rich support for other queue and pub-sub technologies such as kafka, rabbitmq, sidecar, and nats. All you ahve to do is simply change the environment variable `MICRO_BROKER` to something else and change the import in `main.go`.

## Microservices to think about

1) user microservice 
2) product microservice
3) image microservice
4) file sharing microservice
5) email/messaging microservice
6) geolocation microservice

## Things to implement
1) Frontend (Angular? React?)
2) Authentication with JWT/OAuth?
3) CI/CD pipeline
4) Docker and Kubernetes
5) RPC
6) RabbitMQ asynchronous messaging
7) Batch task handling and parallel computing
8) Health Monitoring and Load Balancing
9) Centralized logging
10) Think of databases to use
11) Testing and Mocking
12) Swagger API Docs
13) Asynchronous task queue -> something like celery? 
14) Connection pooling and DB pooling

## Environment Variables

