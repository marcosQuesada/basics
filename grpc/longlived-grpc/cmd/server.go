package cmd

import (
	"context"
	"fmt"
	"github.com/holdedhub/longlived-grpc/pkg/protos"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
)


// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
        lis, err := net.Listen("tcp", "127.0.0.1:7070")
        if err != nil {
            log.Fatalf("failed to listen: %v", err)
        }
        grpcServer := grpc.NewServer([]grpc.ServerOption{}...)

        server := &longlivedServer{}

        go server.mockDataGenerator()

        protos.RegisterLonglivedServer(grpcServer, server)

        log.Printf("Starting server on address %s", lis.Addr().String())

        if err := grpcServer.Serve(lis); err != nil {
            log.Fatalf("failed to listen: %v", err)
        }
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

type longlivedServer struct {
	protos.UnimplementedLonglivedServer
	subscribers sync.Map
}

type sub struct {
	stream   protos.Longlived_SubscribeServer
	finished chan<- bool
}

// Subscribe handles a subscribe request from a client
func (s *longlivedServer) Subscribe(request *protos.Request, stream protos.Longlived_SubscribeServer) error {
	log.Printf("Received subscribe request from ID: %d", request.Id)

	fin := make(chan bool)

	s.subscribers.Store(request.Id, sub{stream: stream, finished: fin})

	ctx := stream.Context()
	for {
		select {
		case <-fin:
			log.Printf("Closing stream for client ID: %d", request.Id)
			return nil
		case <- ctx.Done():
			log.Printf("Client ID %d has disconnected", request.Id)
			return nil
		}
	}
}

// Unsubscribe handles a unsubscribe request from a client
func (s *longlivedServer) Unsubscribe(ctx context.Context, request *protos.Request) (*protos.Response, error) {
	v, ok := s.subscribers.Load(request.Id)
	if !ok {
		return nil, fmt.Errorf("failed to load subscriber key: %d", request.Id)
	}
	sub, ok := v.(sub)
	if !ok {
		return nil, fmt.Errorf("failed to cast subscriber value: %T", v)
	}
	select {
	case sub.finished <- true:
		log.Printf("Unsubscribed client: %d", request.Id)
	default:
	}
	s.subscribers.Delete(request.Id)
	return &protos.Response{}, nil
}

func (s *longlivedServer) mockDataGenerator() {
	log.Println("Starting data generation")
	for {
		time.Sleep(time.Second)

		var published int
		var unsubscribe []int32
		s.subscribers.Range(func(k, v interface{}) bool {
			id, ok := k.(int32)
			if !ok {
				log.Printf("Failed to cast subscriber key: %T", k)
				return false
			}
			sub, ok := v.(sub)
			if !ok {
				log.Printf("Failed to cast subscriber value: %T", v)
				return false
			}

			if err := sub.stream.Send(&protos.Response{Data: fmt.Sprintf("data mock for: %d", id)}); err != nil {
				log.Printf("Failed to send data to client: %v", err)
				select {
				case sub.finished <- true:
					log.Printf("Unsubscribed client: %d", id)
				default:
				}

				unsubscribe = append(unsubscribe, id)
			}

			published++
			return true
		})
		if published != 0 {
			log.Infof("publish %d messages", published)
		}

		for _, id := range unsubscribe {
			s.subscribers.Delete(id)
		}
	}
}
