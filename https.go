package helper

import (
	"crypto/tls"
	"errors"
	"net/http"
)

// HttpsExpireTimestamp https证书过期时间
func HttpsExpireTimestamp(URL string) (timestamp int64, err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(URL)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.TLS == nil {
		return 0, errors.New("invalid url")
	}
	certInfo := resp.TLS.PeerCertificates
	if len(certInfo) <= 0 {
		return 0, errors.New("invalid url")
	}
	return certInfo[0].NotAfter.Unix(), nil
}
