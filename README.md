# Starter Code For Go API services
This repo has all you need to get you started

### What's included?
- Dependancy Injection of all packages
- Router (Chi Router)
- Kafka Consumer
- Responder package for encapsulating the logic of responding to http request
- Errors package
- Configurations with environment variables loaded in go structs on runtime
- Logging with zaplogger
- Example on how you can write a new function/service/handler (look for "dosomething package" in the internal directoy)
- Example API client you can follow to consume an external API (look for "apiclient" package in the internal directoy)
- Unit tests
- Mocking for unit tests
- Validator which you can use to validate incoming requests against different rules
