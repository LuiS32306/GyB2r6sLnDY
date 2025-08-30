// 代码生成时间: 2025-08-30 16:06:53
 * This service provides a gRPC endpoint to perform sorting algorithms on a given list of integers.
 */

package main

import (
    "fmt"
    "log"
    "net"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "math/rand"
    "time"
)

// SortingAlgorithmService defines the service which will provide sorting capabilities
type SortingAlgorithmService struct {}

// Sort requests a sorting operation on a given slice of integers
type SortRequest struct {
    Numbers []int32 `protobuf:"varint,1,rep,packed,name=numbers"`
}

// SortResponse provides the sorted list of integers
type SortResponse struct {
   _sortedNumbers []int32 `protobuf:"varint,1,rep,packed,name=sorted_numbers"`
}

// Sort performs a sorting operation on the given slice of integers using quicksort algorithm
func (s *SortingAlgorithmService) Sort(ctx context.Context, req *SortRequest) (*SortResponse, error) {
    // Shallow copy the numbers to avoid modifying the original slice
    numbers := make([]int32, len(req.Numbers))
    copy(numbers, req.Numbers)

    // Perform quicksort
    quicksort(numbers, 0, len(numbers)-1)

    // Create and return the response
    return &SortResponse{_sortedNumbers: numbers}, nil
}

// quicksort sorts the slice of integers using the quicksort algorithm
func quicksort(slice []int32, left int, right int) {
    if left < right {
        pivotIndex := partition(slice, left, right)
        quicksort(slice, left, pivotIndex-1)
        quicksort(slice, pivotIndex+1, right)
    }
}

// partition rearranges the elements in the slice based on the pivot value
func partition(slice []int32, left int, right int) int {
    pivot := slice[right]
    i := left - 1
    for j := left; j < right; j++ {
        if slice[j] < pivot {
            i++
            slice[i], slice[j] = slice[j], slice[i]
        }
    }
    slice[i+1], slice[right] = slice[right], slice[i+1]
    return i + 1
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Listening on port 50051")
    defer lis.Close()

    grpcServer := grpc.NewServer()
    defer grpcServer.GracefulStop()

    // Register the service with the gRPC server
    sortingService := &SortingAlgorithmService{}
    RegisterSortingAlgorithmServiceServer(grpcServer, sortingService)

    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// RegisterSortingAlgorithmServiceServer registers the service with the gRPC server
func RegisterSortingAlgorithmServiceServer(s *grpc.Server, srv *SortingAlgorithmService) {
    RegisterSortingAlgorithmServiceServer(s, srv)
}
