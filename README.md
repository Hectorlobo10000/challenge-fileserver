# challenge-fileserver

we have solution folder inside we have three folders

1. client
2. frontend
3. server

## Description

- client: it contains the logic to connect to the server tcp
- server: it contains the logic publisher/subscriber
- fronted: it contains the logic to connect to the server tcp as well it serve the single page with Vuejs that it shows the recieved files.

In each folder can you run de main func with the command go run .

## clients' commands

- -subscribe [channel]
- -listend
- -send [file.extension]

## Frontend
The server that it's listening the tcp server it's subscrib to #tree channel
I recommed that the clients are subscribing to the #tree channel


:)