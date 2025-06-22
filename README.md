# A HTTP client that interacts with Shodan.io [Basically a shodan-cli clone written in Go]

[![Go Report Card](https://goreportcard.com/badge/github.com/dracuxan/shodan-cli)](https://goreportcard.com/report/github.com/dracuxan/shodan-cli)
[![GoDoc](https://godoc.org/github.com/dracuxan/shodan-cli?status.svg)](https://godoc.org/github.com/dracuxan/shodan-cli)

## Install using `go install`

```sh
go install github.com/dracuxan/shodan-cli@latest
```

## Setup

1. Clone the repository

```sh
git clone https://github.com/dracuxan/shodan-cli.git

```

2. Install the required dependencies

```sh
make deps
```

3. Set your Shodan API key in the `.env` file

```sh
echo "SHODAN_API_KEY=your_api_key_here" > .env
```

4. Source the `.env` file to load the environment variables

```sh
source .env
```

5. Run the tool

```sh
make run
```
