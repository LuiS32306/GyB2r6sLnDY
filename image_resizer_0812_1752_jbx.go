// 代码生成时间: 2025-08-12 17:52:48
package main

import (
    "context"
    "fmt"
    "io"
    "log"
    "net"
    "os"
    "path/filepath"
    "plugin"
    "time"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "google.golang.org/protobuf/types/known/emptypb"
)

// ImageResizerService is the server API for ImageResizer service.
type ImageResizerService struct {
    // TODO: populate any external dependencies
}

// ResizeImageRequest is the request message for resizing an image.
type ResizeImageRequest struct {
    FilePath string `protobuf:"2,opt,name=filePath,proto3" json:"filePath,omitempty"`
    Width    int32  `protobuf:"3,opt,name=width,proto3" json:"width,omitempty"`
    Height   int32  `protobuf:"4,opt,name=height,proto3" json:"height,omitempty"`
}

// ResizeImageResponse is the response message for resizing an image.
type ResizeImageResponse struct {
    Result string `protobuf:"1,opt,name=result,proto3" json:"result,omitempty"`
}

// ResizeImagesRequest is the request message for resizing multiple images.
type ResizeImagesRequest struct {
    Requests []*ResizeImageRequest `protobuf:"1,rep,name=requests,proto3" json:"requests,omitempty"`
}

// ResizeImagesResponse is the response message for resizing multiple images.
type ResizeImagesResponse struct {
    Results []string `protobuf:"1,rep,name=results,proto3" json:"results,omitempty"`
}

// ProtoMessage is a dummy function to satisfy the protobuf.Message interface
func (*ResizeImageRequest) ProtoMessage() {}

// Reset resets the message
func (m *ResizeImageRequest) Reset() { *m = ResizeImageRequest{} }

// String returns the string representation of the message
func (m *ResizeImageRequest) String() string { return proto.CompactTextString(m) }

// ProtoMessage is a dummy function to satisfy the protobuf.Message interface
func (*ResizeImageResponse) ProtoMessage() {}

// Reset resets the message
func (m *ResizeImageResponse) Reset() { *m = ResizeImageResponse{} }

// String returns the string representation of the message
func (m *ResizeImageResponse) String() string { return proto.CompactTextString(m) }

// ProtoMessage is a dummy function to satisfy the protobuf.Message interface
func (*ResizeImagesRequest) ProtoMessage() {}

// Reset resets the message
func (m *ResizeImagesRequest) Reset() { *m = ResizeImagesRequest{} }

// String returns the string representation of the message
func (m *ResizeImagesRequest) String() string { return proto.CompactTextString(m) }

// ProtoMessage is a dummy function to satisfy the protobuf.Message interface
func (*ResizeImagesResponse) ProtoMessage() {}

// Reset resets the message
func (m *ResizeImagesResponse) Reset() { *m = ResizeImagesResponse{} }

// String returns the string representation of the message
func (m *ResizeImagesResponse) String() string { return proto.CompactTextString(m) }

// Server API for ImageResizer service
type ImageResizerServer struct {
    ImageResizerService
}

// ResizeImage resizes an image to the specified dimensions
func (s *ImageResizerServer) ResizeImage(ctx context.Context, req *ResizeImageRequest) (*ResizeImageResponse, error) {
    // TODO: implement resizing logic
    fmt.Printf("Resizing image: %s to width: %d, height: %d
", req.FilePath, req.Width, req.Height)
    // For demonstration purposes, we are simulating a successful resize operation
    return &ResizeImageResponse{Result: "Image resized successfully"}, nil
}

// ResizeImages resizes multiple images to the specified dimensions
func (s *ImageResizerServer) ResizeImages(ctx context.Context, req *ResizeImagesRequest) (*ResizeImagesResponse, error) {
    var results []string
    for _, r := range req.Requests {
        resp, err := s.ResizeImage(ctx, r)
        if err != nil {
            log.Printf("Failed to resize image: %s, error: %v
", r.FilePath, err)
            return nil, err
        }
        results = append(results, resp.Result)
    }
    return &ResizeImagesResponse{Results: results}, nil
}

func main() {
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Starting image resizer gRPC server... on port 50051")

    server := grpc.NewServer()
    pb.RegisterImageResizerServer(server, &ImageResizerServer{ImageResizerService{}})
    reflection.Register(server)
    if err := server.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}