Scut
====

### Manage your configs Like a Boss!

Scut is a package that provide config and server for your config, and with
scut-server you can get or set values of your config.

Very useful for daemons which use some configurations.

## Installation

`go get github.com/kovetskiy/scut`

## Usage

```go
// First, we need to create Config structure..
config, err := scut.NewConfig("config.toml")
if err != nil {
	panic(err)
}

// See https://github.com/zazab/zhash
someHost, err := config.GetString("mysql", "host")
if err != nil {
	panic(err)
}

somePort, err := config.GetInt("mysql", "port")
if err != nil {
	panic(err)
}

fmt.Printf("%s:%d\n", someHost, somePort)

// Okay, then we will start our daemon and we want to change stored data in
// initialized config structure,
//
// We are creating ConfigServer, and giving to it our config.
server, err := scut.NewConfigServer(config)
if err != nil {
	panic(err)
}

// Good, we have a server, then we should start it, for example on 8080 port.
go server.Listen(":8080")
```

For example we have `config.toml`:

```toml
simpleKey="simpleValue"

[array_values]
keyA="valueA"
keyB="valueB"
```

App with scut.ConfigServer is started on 8080 port.

### Operations

* Get all data with GET method:
  ```
  $ curl http://localhost:8080/
  {
    "array_values": {
      "keyA": "valueA",
      "keyB": "valueB"
    },
    "simpleKey": "simpleValue"
  }
  ```

* Get item:
  ```
  $ curl http://localhost:8080/simpleKey
  "simpleValue"
  ```

* Get nested item:
  ```
  $ curl http://localhost:8080/array_values/keyA
  "valueA"
  ```

* Get all nested items in `array_values`:
  ```
  $ curl http://localhost:8080/array_values/
  {
    "keyA": "valueA",
    "keyB": "valueB"
  }
  ```

* Change stored data in `array_values.keyA`:
  ```
  $ curl -X PATCH --data '"changedValueA"' \
      http://localhost:8080/array_values/keyA
  ```

  Gotcha!

  ```
  $ curl http://localhost:8080/array_values/
  {
    "keyA": "changedValueA",
    "keyB": "valueB"
  }
  ```

  ```
  $ curl http://localhost:8080/
  {
    "array_values": {
      "keyA": "changedValueA",
      "keyB": "valueB"
    },
    "simpleKey": "simpleValue"
  }
  ```
