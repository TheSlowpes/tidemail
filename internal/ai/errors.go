package ai

import (
	"errors"
	"fmt"
	"net"
)

func providerRequestError(provider string, err error) error {
	if dnsErr, ok := errors.AsType[*net.DNSError](err); ok {
		msg := fmt.Sprintf("%s: DNS lookup failed for %s", provider, dnsErr.Name)
		if dnsErr.Server != "" {
			msg += " via " + dnsErr.Server
		}
		if dnsErr.Err != "" {
			msg += ": " + dnsErr.Err
		}
		return fmt.Errorf("%s; check your DNS/network connection and try again", msg)
	}
	return fmt.Errorf("%s: request failed: %w", provider, err)
}
