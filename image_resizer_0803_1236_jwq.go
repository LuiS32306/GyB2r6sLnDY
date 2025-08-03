// 代码生成时间: 2025-08-03 12:36:16
package main

import (
    "fmt"
    "google.golang.org/grpc"
    "google.golang.org/protobuf/proto"
    "log"
    "net"

    "image"
    "image/jpeg"
    "os"
    "path/filepath"

    // 引入自定义的proto文件
    "imageresizer/resizerproto"
)

// ImageResizerService 定义了图片尺寸批量调整器服务接口
type ImageResizerService struct{
    resizerproto.UnimplementedImageResizerServer
}
# 增强安全性

// ResizeImages 实现proto文件中定义的RPC方法，用于批量调整图片尺寸
func (s *ImageResizerService) ResizeImages(ctx context.Context, req *resizerproto.ResizeImagesRequest) (*resizerproto.ResizeImagesResponse, error) {
    var response resizerproto.ResizeImagesResponse
    for _, imageRequest := range req.GetImageRequests() {
# 扩展功能模块
        // 根据请求中的图片路径读取图片
# TODO: 优化性能
        imgFile, err := os.Open(imageRequest.GetImagePath())
        if err != nil {
# TODO: 优化性能
            return nil, err
        }
        defer imgFile.Close()

        // 读取图片内容
        img, _, err := image.Decode(imgFile)
# FIXME: 处理边界情况
        if err != nil {
            return nil, err
# 扩展功能模块
        }
# 优化算法效率

        // 调整图片尺寸
# NOTE: 重要实现细节
        resizedImg := resizeImage(img, imageRequest.GetTargetWidth(), imageRequest.GetTargetHeight())
# 扩展功能模块

        // 保存调整后尺寸的图片
        if err := saveImage(resizedImg, imageRequest.GetTargetImagePath()); err != nil {
            return nil, err
        }

        // 将成功调整尺寸的图片信息添加到响应中
        response.ResizedImages = append(response.ResizedImages, &resizerproto.ResizedImageInfo{
            ImagePath: imageRequest.GetTargetImagePath(),
            Width:    int32(resizedImg.Bounds().Dx()),
            Height:   int32(resizedImg.Bounds().Dy()),
        })
    }
    return &response, nil
}
# 添加错误处理

// resizeImage 调整图片尺寸的辅助函数
# TODO: 优化性能
func resizeImage(img image.Image, targetWidth, targetHeight int32) image.Image {
    // 创建一个新的图片缓冲区，并设置图片尺寸
# 增强安全性
    resizedImg := image.NewRGBA(image.Rect(0, 0, int(targetWidth), int(targetHeight)))
# TODO: 优化性能
    // 将原图片绘制到新的缓冲区，实现尺寸调整
    draw.Draw(resizedImg, resizedImg.Bounds(), img, image.Pt(0, 0), draw.Src)
    return resizedImg
# FIXME: 处理边界情况
}

// saveImage 保存图片到指定路径的辅助函数
func saveImage(img image.Image, path string) error {
    // 创建文件
# 增强安全性
    file, err := os.Create(path)
    if err != nil {
        return err
    }
    defer file.Close()
# 改进用户体验

    // 保存为JPEG格式
    return jpeg.Encode(file, img, nil)
}
# 增强安全性

func main() {
    lis, err := net.Listen("tcp", ":50051")
# TODO: 优化性能
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
# 扩展功能模块
    fmt.Println("Server listening on port 50051")

    // 创建gRPC服务器
    s := grpc.NewServer()
    resizerproto.RegisterImageResizerServer(s, &ImageResizerService{})

    // 启动服务器
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
