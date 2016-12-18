# travishook
Listen for [Travis CI](https://travis-ci.org/)'s [Webhooks](https://docs.travis-ci.com/user/notifications/#Webhook-notifications) in [Go](https://golang.org).

This repository contains:

* Go `structs` for unmarshaling JSON-encoded data sent by Travis ((structs.go)[structs.go])
* `CheckTravisSignature` to check the signature of the payload sent by Travis 
((travishook.go)[travishook.go])
* `GetPubKey` to retrieve Travis's public key from the API server 
((travishook.go)[travishook.go])
* an example implementation of what an httpHandler in Go might look like for Travis webhooks 
((example_test.go)[example_test.go])
