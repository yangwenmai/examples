package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/http2"
)

const writeTimeout = 2 * time.Second

func TestConnWriteTimeout(t *testing.T) {
	stuckLn := startBlockingTCP(t)
	stuckURL := fmt.Sprintf("http://%v", stuckLn.Addr())

	t.Logf("stuckURL: %v", stuckURL)
	dialFunc := func(network, addr string) (net.Conn, error) {
		conn, err := net.Dial(network, addr)
		return wrappedConn{
			Conn:         conn,
			writeTimeout: writeTimeout,
		}, err
	}

	client := newHTTP2ClientWithDialFunc(dialFunc)
	beginCallTime := time.Now()
	_, err := client.Post(stuckURL, "application/raw", &infiniteReader{})
	require.Error(t, err)
	assert.True(t, time.Since(beginCallTime) > writeTimeout)
	t.Logf("got err: %v", err)
	for i := 0; i < 3; i++ {
		beginCallTime = time.Now()
		_, err = client.Post(stuckURL, "application/raw", &infiniteReader{})
		assert.Error(t, err)
		t.Logf("got err: %v", err)
		rtt := time.Since(beginCallTime)
		assert.True(t, rtt > writeTimeout, "expect response time to be larger than %v, actual: %v", writeTimeout, rtt)
	}
}

type wrappedConn struct {
	net.Conn

	writeTimeout time.Duration
}

func (wc wrappedConn) Write(b []byte) (n int, err error) {
	if conn, ok := wc.Conn.(*net.TCPConn); ok {
		if err := conn.SetWriteDeadline(time.Now().Add(wc.writeTimeout)); err != nil {
			panic(err)
		}
	}

	return wc.Conn.Write(b)
}

type infiniteReader struct {
	read int64
}

func (r *infiniteReader) Read(p []byte) (int, error) {
	atomic.AddInt64(&r.read, int64(len(p)))
	return len(p), nil
}

func newHTTP2ClientWithDialFunc(dialFunc func(network, addr string) (net.Conn, error)) *http.Client {
	return &http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, _ *tls.Config) (net.Conn, error) {
				conn, err := dialFunc(network, addr)
				if wrappedConn, ok := conn.(wrappedConn); ok {
					if tcpconn, ok := wrappedConn.Conn.(*net.TCPConn); ok {
						if err := tcpconn.SetWriteBuffer(100); err != nil {
							panic(err)
						} else {
							fmt.Println("successfully set TCP write buffer to be 100")
						}
					}
				}
				return conn, err
			},
		},
	}
}

func startBlockingTCP(t *testing.T) net.Listener {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err, "listen failed")

	go func() {
		for {
			conn, _ := ln.Accept()
			if conn, ok := conn.(*net.TCPConn); ok {
				// Set a small read buffer so we're more likely to hit the write buffer limit on the client.
				conn.SetReadBuffer(10)
			}
		}
	}()

	return ln
}
