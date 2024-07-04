package mauth

import (
	"context"

	"github.com/coreos/go-oidc"
)

type AuthClient struct {
	AuthEndpoint string
	AuthClientID string
}

type Verifier struct {
	verify oidc.IDTokenVerifier
	ctx    context.Context
}
