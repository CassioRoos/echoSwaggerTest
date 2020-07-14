# Echo Swagger

## [Read the Docs - Swag](https://github.com/swaggo/swag)

#### How to run the application

```shell script
docker-compose -f docker/docker-compose.yml up --build 
```

After that, just access the [swagger - http://localhost/swagger/index.html](http://localhost:7777/swagger/index.html)

While doing this POC I made some comments in the code that might help who is looking to implement Swagger

Want to get in touch, give me a shout cassio.roos@gmail.com

## Also using nginx to loadbalance the requests.

   > The load balance chosen is Round Robin, each request will be attended for one container at a time. It will respect the order in which they were configured.
    