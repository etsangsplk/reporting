package reporting

import (
	"context"
	"crypto/tls"
	"net"
	"testing"
	"time"

	"github.com/cloudflare/cfssl/csr"
	"github.com/gravitational/license/authority"
	"github.com/pborman/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	check "gopkg.in/check.v1"
)

func TestReporting(t *testing.T) { check.TestingT(t) }

type ReportingSuite struct {
	client   Client
	eventsCh chan Event
}

var _ = check.Suite(&ReportingSuite{})

func (r *ReportingSuite) SetUpSuite(c *check.C) {
	r.eventsCh = make(chan Event, 10)
	serverAddr := startTestServer(c, r.eventsCh)
	r.client = getTestClient(c, serverAddr)
}

func (r *ReportingSuite) TestReporting(c *check.C) {
	event := &ServerEvent{
		Action:    EventActionLogin,
		AccountID: uuid.New(),
		ServerID:  uuid.New(),
	}
	r.client.Record(event)
	select {
	case received := <-r.eventsCh:
		c.Assert(event, check.DeepEquals, received)
	case <-time.After(testTimeout):
		c.Fatal("timeout waiting for event")
	}
}

// startTestServer starts gRPC events server that will be submitting events
// into the provided channel, and returns the server address
func startTestServer(c *check.C, ch chan Event) (addr string) {
	// generate certificate authority
	ca, err := authority.GenerateSelfSignedCA(csr.CertificateRequest{CN: "localhost"})
	c.Assert(err, check.IsNil)
	cert, err := tls.X509KeyPair(ca.CertPEM, ca.KeyPEM)
	c.Assert(err, check.IsNil)
	// start gRPC test server
	l, err := net.Listen("tcp", "localhost:0")
	c.Assert(err, check.IsNil)
	server := grpc.NewServer(grpc.Creds(credentials.NewServerTLSFromCert(&cert)))
	RegisterEventsServiceServer(server, NewServer(ServerConfig{
		Sinks: []Sink{NewChannelSink(ch)},
	}))
	go server.Serve(l)
	return l.Addr().String()
}

// getTestClient returns a new gRPC events client for the provided server address
func getTestClient(c *check.C, addr string) Client {
	client, err := NewClient(context.Background(), ClientConfig{
		ServerAddr: addr,
		Insecure:   true,
	})
	c.Assert(err, check.IsNil)
	return client
}
