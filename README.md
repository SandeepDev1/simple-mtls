# simple-mtls

### Steps to setup server:
1. run the `keys.sh` file which will generate certificates for the CA ( Certificate Authority ) and the Server
2. run `go run server.go` -> It starts the server on port `8443`
3. run `go run client.go` -> It'll run the client which will use the CA cert from the CA directory.
4. you can also run this via python which takes the cert `python client.py`

Currently the CA is locked to localhost, if you want to use any domain, you can change in keys.sh file
and generate again by running that file
