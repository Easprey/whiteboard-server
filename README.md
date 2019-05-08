# whiteboard-server

A Swagger 2.0 compliant web API server generated using go-swagger (https://github.com/go-swagger/go-swagger).
Serves as the back end for the whiteboard Android app (https://github.com/ejweber/whiteboard).

## Installation

```bash
mkdir $GOPATH/github.io/Easprey
cd $GOPATH/github.io/Easprey
git clone https://github.com/Easprey/whiteboard-server.git
cd whiteboard-server
go install ./...
```

## Usage

Set appropriate environment variables using "source <some_config_script>.sh. Then, start the server.

```bash
export DB_CONN="<[username[:password]@][protocol[(address)]]/dbname>"
export PORT=<port>
whiteboard-server
```

OR

Pass variables using command line arguments.

```bash
whiteboard-server --port=<port> --database=<[username[:password]@][protocol[(address)]]/dbname>
```
