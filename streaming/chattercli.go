package chatter

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/signal"

	chattersvc "goa.design/examples/streaming/gen/chatter"
)

//InteractWithStreams performs the sample operations on the stream
func InteractWithStreams(data interface{}) {
	if data != nil {
		switch stream := data.(type) {
		case chattersvc.EchoerClientStream:
			// bidirectional streaming
			trapCtrlC(stream)
			fmt.Println("Press Ctrl+D to stop chatting.")
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				p := scanner.Text()
				if err := stream.Send(p); err != nil {
					fmt.Println(fmt.Errorf("error sending into stream: %s", err))
					os.Exit(1)
				}
				d, err := stream.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					fmt.Println(fmt.Errorf("error reading from stream: %s", err))
				}
				prettyPrint(d)
			}
			if err := stream.Close(); err != nil {
				fmt.Println(fmt.Errorf("error closing stream: %s", err))
			}
		case chattersvc.ListenerClientStream:
			// payload streaming (no server response)
			trapCtrlC(stream)
			fmt.Println("Press Ctrl+D to stop chatting.")
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				p := scanner.Text()
				if err := stream.Send(p); err != nil {
					fmt.Println(fmt.Errorf("error sending into stream: %s", err))
					os.Exit(1)
				}
			}
			if err := stream.Close(); err != nil {
				fmt.Println(fmt.Errorf("error closing stream: %s", err))
			}
		case chattersvc.SummaryClientStream:
			// payload streaming (server responds with a result type)
			trapCtrlC(stream)
			fmt.Println("Press Ctrl+D to stop chatting.")
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				p := scanner.Text()
				if err := stream.Send(p); err != nil {
					fmt.Println(fmt.Errorf("error sending into stream: %s", err))
					os.Exit(1)
				}
			}
			if p, err := stream.CloseAndRecv(); err != nil {
				fmt.Println(fmt.Errorf("error closing stream: %s", err))
			} else {
				prettyPrint(p)
			}
		case chattersvc.SubscribeClientStream:
			// result streaming
			for {
				p, err := stream.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					fmt.Println(fmt.Errorf("error reading from stream: %v", err))
					break
				}
				prettyPrint(p)
			}
		case chattersvc.HistoryClientStream:
			// result streaming with views
			for {
				p, err := stream.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					fmt.Println(fmt.Errorf("error reading from stream: %v", err))
				}
				prettyPrint(p)
			}
		default:
			prettyPrint(data)
		}
	}
	if data != nil {
		m, _ := json.MarshalIndent(data, "", "    ")
		fmt.Println(string(m))
	}
}

func prettyPrint(s interface{}) {
	m, _ := json.MarshalIndent(s, "", "    ")
	fmt.Println(string(m))
}

// Trap Ctrl+C to gracefully exit the client.
func trapCtrlC(stream interface{}) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	go func(stream interface{}) {
		for range ch {
			fmt.Println("\nexiting")
			if s, ok := stream.(chattersvc.EchoerClientStream); ok {
				s.Close()
			} else if s, ok := stream.(chattersvc.ListenerClientStream); ok {
				s.Close()
			} else if s, ok := stream.(chattersvc.SummaryClientStream); ok {
				s.CloseAndRecv()
			}
			os.Exit(0)
		}
	}(stream)
}
