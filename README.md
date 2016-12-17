# travishook
Listen for [Travis CI](https://travis-ci.org/)'s [Webhooks](https://docs.travis-ci.com/user/notifications/#Webhook-notifications) in [Go](https://golang.org).

This repository contains:

* Go `structs` for unmarshaling JSON-encoded data sent by Travis ([src/travishook/structs.go])
* `CheckTravisSignature` to check the signature of the payload sent by Travis ([src/travishook/travishook.go])
* `GetPubKey` to retrieve Travis's public key from the API server ([src/travishook/travishook.go])
* an example implementation of what an httpHandler in Go might look like for Travis webhooks ([src/main/example_test.go])
