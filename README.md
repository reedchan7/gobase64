# gobase64

gobase64 is a CLI for encoding or decoding a string by `Base64` algorithm.


## Quick Start

- encode

    ```text
    $ gobase64 encode hello-world
    aGVsbG8td29ybGQ=
    ```


- decode

    ```text
    $ gobase64 decode aGVsbG8td29ybGQ= 
    hello-world
    ```

- help

    ```text
    $ gobase64 help                    
    gobase64 is a CLI for encoding or decoding string by Base64
    
    Usage:
      gobase64 [command]
    
    Available Commands:
      decode      Use Base64 algorithm to decode a string
      encode      Use Base64 algorithm to encode a string
      help        Help about any command
      version     Print the version number of gobase64
    
    Flags:
      -h, --help   help for gobase64
    
    Use "gobase64 [command] --help" for more information about a command.
    ```
