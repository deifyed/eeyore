# Eeyore

## Motivation

I wanted to access ChatGPT through my terminal

## Usage

```shell
# Ask a single question
echo "What is the meaning of life?" | eeyore single
eeyore single "What is the meaning of life?"

# Start a conversation
eeyore thread
```

## Installation

### Prerequisites

- [Go 1.19](https://golang.org/doc/install)

### Build

```shell
# Build the application
make build
```

### Install

```shell
# Install the application
make install
```

This will install the application to `~/.local/bin/eeyore`. Remember to add this to your `PATH`.

To change install location, set the `PREFIX` environment variable.

```shell
PREFIX=/usr/local make install
```

## Configuration

The application can be configured either using environment variables:

| Variable              | Description                                    | Default               |
| --------------------- | ---------------------------------------------- | --------------------- |
| `EEYORE_OPENAI_TOKEN` | The API token to use                           | `""`                  |
| `EEYORE_MAX_TOKENS`   | The maximum amount of tokens GPT should return | `1024`                |
| `EEYORE_MODEL`        | The model to use                               | `text-davinci-003ada` |
| `EEYORE_TEMPERATURE`  | Configure how creative the output is           | `0.5`                 |

Or using a configuration file located at `~/.config/eeyore/config.yaml`:

```yaml
openai-token: "YOUR_TOKEN"
max-tokens: 1024
model: text-davinci-003
temperature: 0.5
```

## FAQ

- How do I acquire an API token? Log into [OpenAI](https://openai.com/api/login), click "Personal" up in the upper right corner and then "View API keys"
