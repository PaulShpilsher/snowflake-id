# Snowflake ID

Snowflake ID generator Go package



## What is the Snowflake ID

From [wikipedia](https://en.wikipedia.org/wiki/Snowflake_ID)

Snowflake IDs, or snowflakes, are a form of unique identifier used in distributed computing. The format was created by Twitter and is used for the IDs of tweets. It is popularly believed that every snowflake has a unique structure, so they took the name "snowflake ID". The format has been adopted by other companies, including Discord and Instagram.

Snowflakes are 64 bits in binary. (Only 63 are used to fit in a signed integer.) The first 41 bits are a timestamp, representing milliseconds since the chosen epoch. The next 10 bits represent a machine ID, preventing clashes. Twelve more bits represent a per-machine sequence number, to allow creation of multiple snowflakes in the same millisecond. The final number is generally serialized in decimal.

Snowflakes are sortable by time, because they are based on the time they were created Additionally, the time a snowflake was created can be calculated from the snowflake. This can be used to get snowflakes (and their associated objects) that were created before or after a particular date.

```
    | 1 bit unused | 41 bit timestamp | 10 bit nodeID | 12 bit sequence |
```


# Usage

Create new go project
```shell
mkdir snowflake-example
cd snowflake-example
go mod init snowflake-example
```

Get snowflake package
```shell
go get github.com/PaulShpilsher/snowflake-id/snowflake
```

Create main.go file
```go
package main

import (
    "fmt"

    "github.com/PaulShpilsher/snowflake-id/snowflake"
)

func main() {

    // define a node identifier, aka machine ID
    machineID := 1

    // Create a new snowflake ID generator
    idGenerator, err := snowflake.NewGenerator(machineID)
    if err != nil {
        panic(err)
    }

    // get a couple of snowflake IDs
    id1 := idGenerator.NextID()
    id2 := idGenerator.NextID()

    // print them out
    fmt.Println(id1, id2)
}
```

Run the code 
```bash
go run .
```
