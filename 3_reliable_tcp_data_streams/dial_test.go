package ch03

import(
	"fmt"
	"io"
	"net"
	"testing"
)

func TestDial(t *testing.T) {
	// Create a listener on a random port.
	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Listener Address is: %s", listener.Addr().String())

	// Make a channel called 'done'.
	done := make(chan struct{})

	// Anonymous go function execution.
	go func() {
		// Add empty struct into 'done' channel.
		defer func() { done <- struct{}{} }()

		for {
			// An implementation of Listing 3-2: a for loop to accept
			// and handle incoming tcp connection requests.
			conn, err := listener.Accept()
			if err != nil {
				t.Log(err)
				return
			}

			// Anonymous function takes a connection interface.
			go func(c net.Conn) {
				// Close the connection and add empty struct into 'done' channel.
				defer func() {
					c.Close()
					done <- struct{}{}
				}()

				// Create byte slice called 'buf'
				buf := make([]byte, 1024)
				for {
					// Use net.Conn's Read() to read data from the connection.
					n, err := c.Read(buf)
					if err != nil {
						if err != io.EOF {
							t.Error(err)
						}
						return
					}

					t.Logf("received: %q", buf[:n])
				}
			}(conn)
		}

	}()

	// Dial to connect to the listener.
	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}

	conn.Close()
	<-done
	listener.Close()
	<-done
}