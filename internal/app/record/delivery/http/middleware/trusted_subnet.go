package middleware

import (
	"net/http"
	"net/netip"
)

// TrustedSubnetChecker проверяет входит ли IP-адрес клиента в доверенную подсеть
type TrustedSubnetChecker struct {
	trustedSubnet string
}

// NewTrustedSubnetChecker создание нового TrustedSubnetChecker
func NewTrustedSubnetChecker(trustedSubnet string) *TrustedSubnetChecker {
	return &TrustedSubnetChecker{
		trustedSubnet: trustedSubnet,
	}
}

// Handle обработка Middleware
func (c TrustedSubnetChecker) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if c.trustedSubnet == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		network, err := netip.ParsePrefix(c.trustedSubnet)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ipStr := r.Header.Get("X-Real-IP")

		ip, err := netip.ParseAddr(ipStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !network.Contains(ip) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
