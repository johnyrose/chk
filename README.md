# CHK

`chk` is a small CLI tool to check connection from a local machine to a remote target in various protocols. 

The following protocols are currently supported:
 * ICMP (Ping)
 * SSH
 * TCP
 * HTTP
 
`chk` is indended to help with basic testing of the connection. It can be used in place of combining other tools and does not have any dependencies other than the executable. 

With every connection error message, `chk` will attempt to show an indicative message that will show exactly what the error is and why the connection fails.

## Installation

chk can be used by simply downloading the latest Golang executable from the releases page. For comfortable CLI usage, put the executable in a place that is under your `PATH` environment variable.


## Getting Started

Basic help message:

```bash
>>> chk 

NAME:
   chk - A simple command line tool to check connection between devices.

USAGE:
   chk [command flags]

COMMANDS:
   ping, icmp  Check if the target address is accessible via ICMP protocol.
   http        Check if the target address is accessible via HTTP.
   tcp         Check if the target address is accessible via TCP.
   ssh         Check if the target address is accessible via SSH.
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

Every subcommand's help page be accessed via calling with with the `-h` attribute:

```bash
>>> chk ping -h

NAME:
   c ping - Check if the target address is accessible via ICMP protocol.

USAGE:
   c ping [command options] [arguments...]

OPTIONS:
   --count value, -c value     Amount of ping requests that will be sent. (default: 1)
   --timeout value, -t value   Time to wait for a response in seconds. (default: 1)
   --interval value, -i value  Time to wait between ping requests in seconds if the count is more than 1. (default: 1)
   --verbose, -v               Specifies whether to show verbose information about the ping result. (default: false)
   --help, -h                  show help (default: false)
```
**Examples**:

Ping:
```bash
>>> chk ping google.com

{
        "PacketsRecv": 1,
        "PacketsSent": 1,
        "PacketLoss": 0,
        "IPAddr": {
                "IP": "216.58.207.78",
                "Zone": ""
        },
        "Addr": "google.com",
        "Rtts": [
                58238468
        ],
        "MinRtt": 58238468,
        "MaxRtt": 58238468,
        "AvgRtt": 58238468,
        "StdDevRtt": 0
}
Successful ping connection to google.com
```
SSH:
```bash
>>> chk ssh localhost

Successful SSH connection to localhost:22.

>>> chk ssh google.com

SSH connection to google.com:22 failed. The following error was received: dial tcp 216.58.207.78:22: i/o timeout
```
TCP:
```
>>> chk tcp google.com:443

Successful TCP connection to google.com:443.

>>> chk tcp localhost:443

TCP Connection to localhost:443 failed. Error that was received: 'dial tcp [::1]:443: connect: connection refused
```

HTTP:
```
>>> chk http google.com

Successful http connection to http://google.com. Response code that was received: 200

>>> chk http localhost:8080

Got connection refused from http://localhost:8080. Target is reachable but does not listen in the specified port.
```