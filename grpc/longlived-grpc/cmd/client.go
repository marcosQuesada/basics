package cmd

import (
	"fmt"

	"context"
	"github.com/holdedhub/longlived-grpc/pkg/protos"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"sync"
	"time"
)

var total int

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("starting %d clients", total)

		var wg sync.WaitGroup

		for i := 1; i <= total; i++ {
			wg.Add(1)
			conn, err := grpc.Dial("127.0.0.1:7070", []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}...)
			if err != nil {
				log.Fatal(err)
			}
			client, err := newLonglivedClient(i, conn)
			if err != nil {
				log.Fatal(err)
			}

			go client.start()
		}

		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
	clientCmd.PersistentFlags().IntVar(&total, "total", 1, "total clients")
}

type longlivedClient struct {
	client protos.LonglivedClient
	conn   *grpc.ClientConn
	id     int32
}

func newLonglivedClient(id int, conn *grpc.ClientConn) (*longlivedClient, error) {
	return &longlivedClient{
		client: protos.NewLonglivedClient(conn),
		conn:   conn,
		id:     int32(id),
	}, nil
}

func (c *longlivedClient) close() {
	if err := c.conn.Close(); err != nil {
		log.Fatal(err)
	}
}

func (c *longlivedClient) subscribe() (protos.Longlived_SubscribeClient, error) {
	log.Printf("Subscribing client ID: %d", c.id)
	return c.client.Subscribe(context.Background(), &protos.Request{Id: c.id})
}

func (c *longlivedClient) unsubscribe() error {
	log.Printf("Unsubscribing client ID %d", c.id)
	_, err := c.client.Unsubscribe(context.Background(), &protos.Request{Id: c.id})
	return err
}

func (c *longlivedClient) start() {
	var err error
	var stream protos.Longlived_SubscribeClient
	for {
		if stream == nil {
			if stream, err = c.subscribe(); err != nil {
				log.Errorf("Failed to subscribe: %v", err)
				c.sleep()
				continue
			}
		}
		response, err := stream.Recv()
		if err != nil {
			log.Errorf("Failed to receive message: %v", err)
			stream = nil
			c.sleep()
			continue
		}
		log.Infof("Client ID %d got response: %q", c.id, response.Data)
	}
}

func (c *longlivedClient) sleep() {
	time.Sleep(time.Second * 5)
}
