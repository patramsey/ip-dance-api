package ipdance

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetIp(t *testing.T) {

	// Create an HTTP mock server
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Define a response for the IP request
	httpmock.RegisterResponder("GET", "https://api.ip.dance/",
		httpmock.NewStringResponder(200, `{
			"ip": "192.168.1.10", 
			"regionCode": "CO",
			"asn": "209, CenturyLink",
			"httpProtocol": "HTTP/3",
			"location": "Denver, Colorado, US"
		}`))

	// Call the getIp function with the mock server as the client
	ip, err := GetIp()

	// Assert that the IP address is correct
	assert.Equal(t, "192.168.1.10", ip)

	// Assert that there is no error
	assert.NoError(t, err)

	// Print the retrieved IP address
	fmt.Println("Retrieved IP:", ip)

}

func TestGetHttpVersion(t *testing.T) {

	// Create an HTTP mock server
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Define a response for the IP request
	httpmock.RegisterResponder("GET", "https://api.ip.dance/",
		httpmock.NewStringResponder(200, `{
			"ip": "192.168.1.10", 
			"regionCode": "CO",
			"asn": "209, CenturyLink",
			"httpProtocol": "HTTP/3",
			"location": "Denver, Colorado, US"
		}`))

	httpVersion, err := GetHttpVersion()

	// Assert that the HTTP version is correct
	assert.Equal(t, "HTTP/3", httpVersion)

	// Assert that there is no error
	assert.NoError(t, err)

	// Print the retrieved HTTP version
	fmt.Println("Retrieved HTTP Version:", httpVersion)

}

func TestGetLocation(t *testing.T) {

	// Create an HTTP mock server
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Define a response for the location request
	httpmock.RegisterResponder("GET", "https://api.ip.dance/",
		httpmock.NewStringResponder(200, `{
			"ip": "192.168.1.10", 
			"regionCode": "CO",
			"asn": "209, CenturyLink",
			"httpProtocol": "HTTP/3",
			"location": "Denver, Colorado, US"
		}`))

	location, err := GetLocation()

	// Assert that the location is correct
	assert.Equal(t, "Denver, Colorado, US", location)

	// Assert that there is no error
	assert.NoError(t, err)

	// Print the retrieved location
	fmt.Println("Retrieved Location:", location)

}

func TestGetAsn(t *testing.T) {

	// Create an HTTP mock server
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Define a response for the ASN request
	httpmock.RegisterResponder("GET", "https://api.ip.dance/",
		httpmock.NewStringResponder(200, `{
			"ip": "192.168.1.10", 
			"regionCode": "CO",
			"asn": "209, CenturyLink",
			"httpProtocol": "HTTP/3",
			"location": "Denver, Colorado, US"
		}`))

	asn, err := GetAsn()

	// Assert that the ASN is correct
	assert.Equal(t, "209, CenturyLink", asn)

	// Assert that there is no error
	assert.NoError(t, err)

	// Print the retrieved ASN
	fmt.Println("Retrieved ASN:", asn)

}

func TestGetRegionCode(t *testing.T) {

	// Create an HTTP mock server
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Define a response for the region code request
	httpmock.RegisterResponder("GET", "https://api.ip.dance/",
		httpmock.NewStringResponder(200, `{
			"ip": "192.168.1.10", 
			"regionCode": "CO",
			"asn": "209, CenturyLink",
			"httpProtocol": "HTTP/3",
			"location": "Denver, Colorado, US"
		}`))

	regionCode, err := GetRegionCode()

	// Assert that the region code is correct
	assert.Equal(t, "CO", regionCode)

	// Assert that there is no error
	assert.NoError(t, err)

	// Print the retrieved region code
	fmt.Println("Retrieved Region Code:", regionCode)

}

func TestGetIpWithError(t *testing.T) {

	// Create an HTTP mock server
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Define a response for the IP request with a 404 status code
	httpmock.RegisterResponder("GET", "https://api.ip.dance/",
		httpmock.NewStringResponder(http.StatusNotFound, ``))

	// Call the getIp function with the mock server as the client
	_, err := getIpAddressInfo()

	// Assert that an error is returned
	assert.Error(t, err)
}
