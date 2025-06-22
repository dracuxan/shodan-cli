# shodan-cli

A HTTP client that interacts with Shodan.io [Basically a shodan-cli clone written in Go]

[![Go Report Card](https://goreportcard.com/badge/github.com/dracuxan/shodan-cli)](https://goreportcard.com/report/github.com/dracuxan/shodan-cli)

## Install using `go install`

```sh
go install github.com/dracuxan/shodan-cli@latest
```

> [!NOTE]
> You need a `.shodan-cli.conf` file with your Shodan API key in your home directory to run the tool.

## 3 Easy Steps to Get Started

1. Clone the repository

```sh
git clone https://github.com/dracuxan/shodan-cli.git

```

2. Set your Shodan API key in the `.shodan-cli.conf` file

```sh
echo "SHODAN_API_KEY=[api_key]" > /home/$USER/.shodan-cli.conf
```

3. Run the tool

```sh
make
```
