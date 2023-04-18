package bc

import (
	"os"
)

const (
	DefaultApiHost = "https://bandu-api.songy.info/v1/credit"
)

func UseApiHost(host string) {
	httpClient.HostURL = host
}

func init() {

	if host, ok := os.LookupEnv("BANDU_CREDIT_SDK_API_HOST"); ok && host != "" {
		UseApiHost(host)
	}

}
