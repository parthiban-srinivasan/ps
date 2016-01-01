/*
 * 
 */

package main

import (
        "log"
        "os"
        _ "runtime/trace"
        pb "github.com/golang/demo/proto/perils"
        "golang.org/x/net/context"
        "google.golang.org/grpc"
)
const (
        address     = "localhost:50051"
        defaultInput = "flood"
)

func main() {

        // Set up a connection to the server.
        conn, err := grpc.Dial(address, grpc.WithInsecure())
        if err != nil {
                log.Fatalf("did not connect: %v", err)
        }
        defer conn.Close()

        c := pb.NewPerilSearchClient(conn)
        // Contact the server and print out its response.
        name := defaultInput

        inputlen := len(os.Args) 
 
        if inputlen > 1 {

               log.Printf("length of Args: %s", inputlen)

               for i := 0; i < inputlen; i++ {
                  
                    name = os.Args[i]
                         
                    log.Printf("Value of Args: %s", name)

                    r, err := c.SearchPeril(context.Background(), &pb.SearchRequest{Name: name})
                    if err != nil {
                       log.Fatalf("could not greet: %v", err)
                     }
                     log.Printf("Peril Search: %s", r.Perilinfo)    
               }
             
               return     
        }

        r, err := c.SearchPeril(context.Background(), &pb.SearchRequest{Name: name})
        if err != nil {
                log.Fatalf("could not greet: %v", err)
        }
        log.Printf("Peril Search: %s", r.Perilinfo)
}
