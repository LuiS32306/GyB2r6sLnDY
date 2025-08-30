// 代码生成时间: 2025-08-30 21:16:35
package main

import (
    "context"
    "errors"
    "fmt"
    "image"
    "image/jpeg"
    "io/ioutil"
    "log"
    "net"
    "os"
    "path/filepath"
    "strings"

    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

// ImageResizeService is the server API for resizing images.
type ImageResizeService struct {
    // Add any service-specific fields here
}

// ResizeImageRequest is the request for resizing an image.
type ResizeImageRequest struct {
    ImagePath string `protobuf:"bytes,1,opt,name=image_path,json=indexPath,proto3"` // Path to the image file
    NewWidth  int    `protobuf:"varint,2,opt,name=new_width,json=newWidth,proto3"`   // New width in pixels
    NewHeight int    `protobuf:"varint,3,opt,name=new_height,json=newHeight,proto3"` // New height in pixels
}

// ResizeImageResponse is the response for resizing an image.
type ResizeImageResponse struct {
    NewImagePath string `protobuf:"bytes,1,opt,name=new_image_path,json=newImagePath,proto3"` // Path to the resized image file
}

// Resize resizes an image to the specified dimensions.
func (s *ImageResizeService) Resize(ctx context.Context, req *ResizeImageRequest) (*ResizeImageResponse, error) {
    // Check for valid path and dimensions
    if _, err := os.Stat(req.ImagePath); os.IsNotExist(err) {
        return nil, status.Error(codes.NotFound, "Image file not found")
    }
    if req.NewWidth <= 0 || req.NewHeight <= 0 {
        return nil, status.Error(codes.InvalidArgument, "Invalid dimensions")
    }

    // Open the source image
    srcImage, err := loadImage(req.ImagePath)
    if err != nil {
        return nil, status.Error(codes.Internal, "Failed to load image")
    }
    defer srcImage.Close()

    // Create a new image with the desired dimensions
    dstImage := image.NewRGBA(image.Rect(0, 0, req.NewWidth, req.NewHeight))

    // Resize the image
    if err := resizeImage(srcImage, dstImage); err != nil {
        return nil, status.Error(codes.Internal, "Failed to resize image")
    }

    // Save the resized image to a new file
    newImagePath := req.ImagePath + "_resized.jpg"
    if err := saveImage(newImagePath, dstImage); err != nil {
        return nil, status.Error(codes.Internal, "Failed to save resized image")
    }

    // Return the path to the new image
    return &ResizeImageResponse{NewImagePath: newImagePath}, nil
}

// loadImage loads an image from a file path.
func loadImage(path string) (image.Image, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
