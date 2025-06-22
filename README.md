# shodan-cli

A HTTP client that interacts with Shodan.io [Basically a shodan-cli clone written in Go]

[![Go Report Card](https://goreportcard.com/badge/github.com/dracuxan/shodan-cli)](https://goreportcard.com/report/github.com/dracuxan/shodan-cli)

## Install using `go install`

```sh
go install github.com/dracuxan/shodan-cli@latest
```

> [!NOTE]
> You need a .env file with your Shodan API key to run the tool.

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

4. Run the tool

```sh
make run
```
