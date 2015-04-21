Scut
====

### Manage your configs Like a Boss!

Scut it's package that providee config and server for your config, and with
scut-server you can get or set values of your config.

Very useful for daemons which use some configurations.

## Installation

`go get github.com/kovetskiy/scut`

## Usage

```go
// first, we need to create Config structure
config, err := scut.NewConfig("config.toml")
if err != nil {
	panic(err)
}

// after that, you can use config just like as https://github.com/zazab/zhash
someHost, err := config.GetString("mysql", "host")
if err != nil {
	panic(err)
}

somePort, err := config.GetInt("mysql", "port")
if err != nil {
	panic(err)
}

fmt.Printf("%s:%d\n", someHost, somePort)

// Okay, then we started our daemon and wants to change stored data in
// initialized config structure,
// We can create ConfigServer, and give him our config
server, err := scut.NewConfigServer(config)
if err != nil {
	panic(err)
}

// Good, we have a server, but we should start him, for example on 8080 port
go server.Listen(":8080")
```


For example we have config.toml:

```toml
simpleKey="simpleValue"

[array_values]
keyA="valueA"
keyB="valueB"
```

And started our app with scut.ConfigServer on 8080 port.

Get all data with GET method:

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

Get item:

```
$ curl http://localhost:8080/simpleKey
"simpleValue"
```

Get nested item:
```
$ curl http://localhost:8080/array_values/keyA
"valueA"
```

or Get all nested items in `array_values`:
```
$ curl http://localhost:8080/array_values/
{
  "keyA": "valueA",
  "keyB": "valueB"
}
```

Hmm, I want to change stored data in `array_values.keyA`
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
