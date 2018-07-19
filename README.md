# least-connections
least-connections is lock-free least-connections balancing algorithm written in golang

## Requrement

Go (>= 1.8)

## Installation

```shell
go get github.com/hlts2/least-connections
```

## Example

### Basic Example
```go
lc, err := leastconnections.New([]string{
		"server-1",
		"server-2",
		"server-3",
    })

// return the server with the least number of connections
server := lc.Next()

```

### Connection Management
```go

// Increment the number of connections of server1
lc.IncrementConnections(server1) // i.e) server1 is "server-1"

// Decrement the number of connections of server1
lc.DecrementConnections(server1) // i.e) server1 is "server-1"

```

## Author
[hlts2](https://github.com/hlts2)

## LICENSE
least-connections released under MIT license, refer [LICENSE](https://github.com/hlts2/least-connections/blob/master/LICENSE) file.
