# TCP Chat

TCP Chat is a command-line utility written in Go that allows users to communicate with each other over a TCP connection. It provides a simple chat room environment where multiple clients can connect to a server and exchange messages.

Features
Server-Client Architecture: TCP Chat operates in a server-client architecture, allowing multiple clients to connect to a central server.

Name Requirement: Clients are required to provide a name upon connection, which is used to identify them in the chat.

Controlled Connections: The server controls the number of connections, with a maximum limit of 10 connections.

Message Broadcasting: Clients can send messages to the chat, which are broadcasted to all other connected clients.

Message Formatting: Messages sent are identified by the time they were sent and the name of the user who sent them, in the format [timestamp][username]:[message].

Handling Empty Messages: Empty messages sent by clients are not broadcasted to other clients.

Join and Leave Notifications: When a client joins or leaves the chat, the server notifies all other clients.

Message History: When a new client joins the chat, they receive the message history from the server.

Port Specification: The server listens on a specified port, defaulting to port 8989 if no port is specified.

Usage
To run the TCP Chat server:

bash
Copy code
`$ go run .
Listening on the port :8989`
To specify a custom port:

bash
Copy code
`$ go run . 2525
Listening on the port :2525`
If the usage is incorrect:

bash
Copy code
`$ go run . 2525 localhost
[USAGE]: ./TCPChat $port`

## Client Connection

When a client connects to the server, they are greeted with a welcome message and prompted to enter their name:

sql
Copy code

After entering their name, the client can start sending messages, which will be displayed in the chat with the appropriate timestamp and username.

### Dependencies

fmt: Provides formatted I/O functions.
net: Provides support for TCP and UDP networking.
bufio: Implements buffered I/O operations.
time: Provides functionality for measuring and displaying time.
sync: Provides synchronization primitives.
log: Implements a simple logging package.
Bonus Features
Terminal UI: TCP Chat supports a terminal-based user interface using the gocui package.

Message Logging: All chat messages can be saved to a log file for later reference.

Contributing
Contributions to TCP Chat are welcome! Feel free to submit bug reports, feature requests, or pull requests on the GitHub repository.

License
TCP Chat is licensed under the MIT License. See the LICENSE file for more information.