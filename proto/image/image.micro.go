// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/image/image.proto

package go_micro_srv_image

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Image service

type ImageService interface {
	CreateAlbum(ctx context.Context, in *CreateAlbumReq, opts ...client.CallOption) (*CreateAlbumResp, error)
	DeleteAlbum(ctx context.Context, in *DeleteAlbumReq, opts ...client.CallOption) (*DeleteAlbumResp, error)
	Upload(ctx context.Context, opts ...client.CallOption) (Image_UploadService, error)
	Download(ctx context.Context, in *DownloadImageReq, opts ...client.CallOption) (Image_DownloadService, error)
	Delete(ctx context.Context, in *DeleteImageReq, opts ...client.CallOption) (*DeleteImageResp, error)
	List(ctx context.Context, in *ListImageReq, opts ...client.CallOption) (*ListImageResp, error)
}

type imageService struct {
	c    client.Client
	name string
}

func NewImageService(name string, c client.Client) ImageService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.image"
	}
	return &imageService{
		c:    c,
		name: name,
	}
}

func (c *imageService) CreateAlbum(ctx context.Context, in *CreateAlbumReq, opts ...client.CallOption) (*CreateAlbumResp, error) {
	req := c.c.NewRequest(c.name, "Image.CreateAlbum", in)
	out := new(CreateAlbumResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageService) DeleteAlbum(ctx context.Context, in *DeleteAlbumReq, opts ...client.CallOption) (*DeleteAlbumResp, error) {
	req := c.c.NewRequest(c.name, "Image.DeleteAlbum", in)
	out := new(DeleteAlbumResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageService) Upload(ctx context.Context, opts ...client.CallOption) (Image_UploadService, error) {
	req := c.c.NewRequest(c.name, "Image.Upload", &UploadImageReq{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &imageServiceUpload{stream}, nil
}

type Image_UploadService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*UploadImageReq) error
}

type imageServiceUpload struct {
	stream client.Stream
}

func (x *imageServiceUpload) Close() error {
	return x.stream.Close()
}

func (x *imageServiceUpload) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *imageServiceUpload) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *imageServiceUpload) Send(m *UploadImageReq) error {
	return x.stream.Send(m)
}

func (c *imageService) Download(ctx context.Context, in *DownloadImageReq, opts ...client.CallOption) (Image_DownloadService, error) {
	req := c.c.NewRequest(c.name, "Image.Download", &DownloadImageReq{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &imageServiceDownload{stream}, nil
}

type Image_DownloadService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*DownloadImageResp, error)
}

type imageServiceDownload struct {
	stream client.Stream
}

func (x *imageServiceDownload) Close() error {
	return x.stream.Close()
}

func (x *imageServiceDownload) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *imageServiceDownload) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *imageServiceDownload) Recv() (*DownloadImageResp, error) {
	m := new(DownloadImageResp)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *imageService) Delete(ctx context.Context, in *DeleteImageReq, opts ...client.CallOption) (*DeleteImageResp, error) {
	req := c.c.NewRequest(c.name, "Image.Delete", in)
	out := new(DeleteImageResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageService) List(ctx context.Context, in *ListImageReq, opts ...client.CallOption) (*ListImageResp, error) {
	req := c.c.NewRequest(c.name, "Image.List", in)
	out := new(ListImageResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Image service

type ImageHandler interface {
	CreateAlbum(context.Context, *CreateAlbumReq, *CreateAlbumResp) error
	DeleteAlbum(context.Context, *DeleteAlbumReq, *DeleteAlbumResp) error
	Upload(context.Context, Image_UploadStream) error
	Download(context.Context, *DownloadImageReq, Image_DownloadStream) error
	Delete(context.Context, *DeleteImageReq, *DeleteImageResp) error
	List(context.Context, *ListImageReq, *ListImageResp) error
}

func RegisterImageHandler(s server.Server, hdlr ImageHandler, opts ...server.HandlerOption) error {
	type image interface {
		CreateAlbum(ctx context.Context, in *CreateAlbumReq, out *CreateAlbumResp) error
		DeleteAlbum(ctx context.Context, in *DeleteAlbumReq, out *DeleteAlbumResp) error
		Upload(ctx context.Context, stream server.Stream) error
		Download(ctx context.Context, stream server.Stream) error
		Delete(ctx context.Context, in *DeleteImageReq, out *DeleteImageResp) error
		List(ctx context.Context, in *ListImageReq, out *ListImageResp) error
	}
	type Image struct {
		image
	}
	h := &imageHandler{hdlr}
	return s.Handle(s.NewHandler(&Image{h}, opts...))
}

type imageHandler struct {
	ImageHandler
}

func (h *imageHandler) CreateAlbum(ctx context.Context, in *CreateAlbumReq, out *CreateAlbumResp) error {
	return h.ImageHandler.CreateAlbum(ctx, in, out)
}

func (h *imageHandler) DeleteAlbum(ctx context.Context, in *DeleteAlbumReq, out *DeleteAlbumResp) error {
	return h.ImageHandler.DeleteAlbum(ctx, in, out)
}

func (h *imageHandler) Upload(ctx context.Context, stream server.Stream) error {
	return h.ImageHandler.Upload(ctx, &imageUploadStream{stream})
}

type Image_UploadStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*UploadImageReq, error)
}

type imageUploadStream struct {
	stream server.Stream
}

func (x *imageUploadStream) Close() error {
	return x.stream.Close()
}

func (x *imageUploadStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *imageUploadStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *imageUploadStream) Recv() (*UploadImageReq, error) {
	m := new(UploadImageReq)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *imageHandler) Download(ctx context.Context, stream server.Stream) error {
	m := new(DownloadImageReq)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.ImageHandler.Download(ctx, m, &imageDownloadStream{stream})
}

type Image_DownloadStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*DownloadImageResp) error
}

type imageDownloadStream struct {
	stream server.Stream
}

func (x *imageDownloadStream) Close() error {
	return x.stream.Close()
}

func (x *imageDownloadStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *imageDownloadStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *imageDownloadStream) Send(m *DownloadImageResp) error {
	return x.stream.Send(m)
}

func (h *imageHandler) Delete(ctx context.Context, in *DeleteImageReq, out *DeleteImageResp) error {
	return h.ImageHandler.Delete(ctx, in, out)
}

func (h *imageHandler) List(ctx context.Context, in *ListImageReq, out *ListImageResp) error {
	return h.ImageHandler.List(ctx, in, out)
}
