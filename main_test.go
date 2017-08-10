package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/rms1000watt/degeneres-test/ballpark"
	"github.com/rms1000watt/degeneres-test/data"
	"github.com/rms1000watt/degeneres-test/server"

	log "github.com/sirupsen/logrus"
)

const (
	host      = "localhost"
	port      = 8080
	certsPath = "./certs"
	keyName   = "server.key"
	certName  = "server.cer"

	serverProtocol = "https" // TODO: Set https if secure middleware is added
)

var x509CertPool *x509.CertPool

func TestMain(m *testing.M) {
	// This is supposed to be the CA Cert, but the key/cert is self signed
	// so passing in the server cert instead as single node chain of trust
	certBytes, err := ioutil.ReadFile(filepath.Join(certsPath, certName)) // TODO: Only read file if secure middleware is added
	if err != nil {
		fmt.Println("Failed reading cert:", err)
		return
	}
	x509CertPool = x509.NewCertPool()
	x509CertPool.AppendCertsFromPEM(certBytes) // TODO: Only add if secure middleware is added

	log.SetLevel(log.DebugLevel)

	go server.Ballpark(server.Config{
		Host:      host,
		Port:      port,
		CertsPath: certsPath,
		KeyName:   keyName,
		CertName:  certName,
	}, ballpark.Config{})

	time.Sleep(500 * time.Millisecond)

	os.Exit(m.Run())
}

func TestTicket(t *testing.T) {
	url := fmt.Sprintf("%s://%s:%d/ticket", serverProtocol, host, port)
	jsonBytes := []byte(`{"id":"912391-123-123-8182123"}`)

	resp := doReq(t, url, jsonBytes)

	ticketOut := &data.TicketOut{}
	if err := json.NewDecoder(resp.Body).Decode(ticketOut); err != nil {
		t.Error("Failed to decode response:", err)
		t.FailNow()
	}

	fmt.Println("Ticket Response:", *ticketOut)
}

func doReq(t *testing.T, url string, jsonBytes []byte) *http.Response {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	if err != nil {
		t.Error("Failed creating request:", err)
		t.FailNow()
	}

	req.Header.Add("Origin", fmt.Sprintf("%s://localhost", serverProtocol))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: x509CertPool,
		},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Do(req)
	if err != nil {
		t.Error("Failed doing req:", err)
		t.FailNow()
	}

	if resp.StatusCode != http.StatusOK {
		t.Error("Bad status code:", resp.StatusCode)
		t.FailNow()
	}

	return resp
}
