# go-client-template

Go Client Template for internal services

> Note: original credit for this should go to the permissions team and the SecOps team for their collective work on the
> permissions-client - https://github.com/circleci/permissions-client


## TODO

- [ ] Observability integration
- [ ] Testing
- [ ] ??

## Usage

In order to create a new go-client-template, you can use the `NewClient` constructor function within this package.

You'll need to pass:

- **baseURL** - this is the base URL for the permissions service.
- **userAgent** - this is the string that identifies your service to the permissions service.

```go
import example "github.com/circleci/go-client-template"


func main() {
    // define this value within your service's values.yml
    serviceURL := os.Getenv("MY_SERIVCE_URL")
    client, err := example.NewClient(serviceURL, "my-super-awesome-service")
    if err != nil {
        log.Error(err.Error())
        return &ecs.Client{}, err
    }

    err := client.CallSomeAPIEndpoint(context.Background(), "additional argument")
    // handle err
}
```

## Tests

```
task test
```

or

```
go test ./... -race -count 1
```
