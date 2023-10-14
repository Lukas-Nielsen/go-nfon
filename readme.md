# go-nfon

A go libary for the [NFON](https://nfon.com)-API

## Installation

```bash
go get github.com/Lukas-Nielsen/go-nfon
```

## Usage

### import

```go
import "github.com/Lukas-Nielsen/go-nfon"
```

### perform request

```go
// new client
client := nfon.NewClient(<api key>, <api secret>, <api url>, <debug>)

// new request
req := client.NewRequest()

// add options to request
req.AddLink(<href>, <rel>)
req.AddData(<name>, <value>)

// send request
req.Send(<GET|POST|DELETE|PUT>, <api path>, <pointer to result or nil>)
```

## API Reference

[NFON](https://cdn.cloudya.com/API_Documentation.zip)

## Authors

- [@Lukas-Nielsen](https://github.com/Lukas-Nielsen)

## License

[MIT](LICENSE)
