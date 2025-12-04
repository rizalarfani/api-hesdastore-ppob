package constants

import "net/textproto"

var (
	XApiKey                = textproto.CanonicalMIMEHeaderKey("x-api-key")
	Authorization          = textproto.CanonicalMIMEHeaderKey("authorization")
	XHubSignatureDigiflazz = textproto.CanonicalMIMEHeaderKey("x-hub-signature")
)
