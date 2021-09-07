# GO API starter kit

This is a bare bones, no frills, just "git'er done" webserver/api starter kit.

It is NOT production ready, but as a bootstrap it's acceptable.

There is zero security. Bake it in here as middleware or add a front end ([AWS API Gateway](https://aws.amazon.com/api-gateway/), for example), but you've been warned.

## Fiber

This is a complete rewrite from using [Gin](https://github.com/gin-gonic/gin) to using [Fiber](https://github.com/gofiber/fiber). If you are looking for the `gin` version, please see the old commits.

### routing

`routes` and `routesv1` show two different methods of defining routes. This also shows how to easily version routes.

### person

`person` is an over simplified package for processing a person structure.
