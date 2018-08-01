# operations-slack-webhooks
Service for sending operations messages to slack incoming webhooks

## Usage
### Server
```
NAME:
   osiw-server

USAGE:
    [global options] command [command options] [arguments...]

DESCRIPTION:
   operational slack incoming webhooks

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --listen value       the address to listen on (default: "127.0.0.1:57438")
   --webhook-url value  the slack incoming webhook url
   --help, -h           show help
   --version, -v        print the version
```

### Client
```
NAME:
   osiw-client

USAGE:
    [global options] command [command options] [arguments...]

DESCRIPTION:
   post operational slack incoming webhooks

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --server value    the address to send requests to (default: "127.0.0.1:57438")
   --hostname value  the hostname sending the requests
   --title value     the title of the message (default: "Status Update")
   --help, -h        show help
   --version, -v     print the version
```

The client reads from stdin and posts the message to the provided
webhook.
