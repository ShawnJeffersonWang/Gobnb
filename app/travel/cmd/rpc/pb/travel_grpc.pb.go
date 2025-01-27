// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.28.2
// source: travel.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TravelClient is the client API for Travel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TravelClient interface {
	// homestayDetail
	HomestayDetail(ctx context.Context, in *HomestayDetailReq, opts ...grpc.CallOption) (*HomestayDetailResp, error)
	HomestayDetailWithoutLogin(ctx context.Context, in *HomestayDetailReq, opts ...grpc.CallOption) (*HomestayDetailResp, error)
	AddHomestay(ctx context.Context, in *AddHomestayReq, opts ...grpc.CallOption) (*AddHomestayResp, error)
	UpdateHomestay(ctx context.Context, in *UpdateHomestayReq, opts ...grpc.CallOption) (*UpdateHomestayResp, error)
	DeleteHomestay(ctx context.Context, in *DeleteHomestayReq, opts ...grpc.CallOption) (*DeleteHomestayResp, error)
	AdminDeleteHomestay(ctx context.Context, in *AdminDeleteHomestayReq, opts ...grpc.CallOption) (*AdminDeleteHomestayResp, error)
	AddComment(ctx context.Context, in *AddCommentReq, opts ...grpc.CallOption) (*AddCommentResp, error)
	LikeComment(ctx context.Context, in *LikeCommentReq, opts ...grpc.CallOption) (*LikeCommentResp, error)
	WishList(ctx context.Context, in *WishListReq, opts ...grpc.CallOption) (*WishListResp, error)
	AddWishList(ctx context.Context, in *AddWishListReq, opts ...grpc.CallOption) (*AddWishListResp, error)
	RemoveWishList(ctx context.Context, in *RemoveWishListReq, opts ...grpc.CallOption) (*RemoveWishListResp, error)
	AddGuess(ctx context.Context, in *AddGuessReq, opts ...grpc.CallOption) (*AddGuessResp, error)
	HistoryList(ctx context.Context, in *HistoryListReq, opts ...grpc.CallOption) (*HistoryListResp, error)
	RemoveHistory(ctx context.Context, in *RemoveHistoryReq, opts ...grpc.CallOption) (*RemoveHistoryResp, error)
	ClearHistory(ctx context.Context, in *ClearHistoryReq, opts ...grpc.CallOption) (*ClearHistoryResp, error)
	SearchHistory(ctx context.Context, in *SearchHistoryReq, opts ...grpc.CallOption) (*SearchHistoryResp, error)
	SearchByLocation(ctx context.Context, in *SearchByLocationReq, opts ...grpc.CallOption) (*SearchByLocationResp, error)
	QueryShopByType(ctx context.Context, in *QueryShopByTypeRequest, opts ...grpc.CallOption) (*QueryShopByTypeResponse, error)
	SeckillVoucher(ctx context.Context, in *SeckillVoucherRequest, opts ...grpc.CallOption) (*SeckillVoucherResponse, error)
	// 查询热门博客
	QueryHotBlog(ctx context.Context, in *QueryHotBlogRequest, opts ...grpc.CallOption) (*QueryHotBlogResponse, error)
	// 根据ID查询博客
	QueryBlogById(ctx context.Context, in *QueryBlogByIdRequest, opts ...grpc.CallOption) (*QueryBlogByIdResponse, error)
	// 点赞或取消点赞
	LikeBlog(ctx context.Context, in *LikeBlogRequest, opts ...grpc.CallOption) (*LikeBlogResponse, error)
	// 查询某篇博客的点赞用户
	QueryBlogLikes(ctx context.Context, in *QueryBlogLikesRequest, opts ...grpc.CallOption) (*QueryBlogLikesResponse, error)
	// 保存博客
	SaveBlog(ctx context.Context, in *SaveBlogRequest, opts ...grpc.CallOption) (*SaveBlogResponse, error)
	// 查询关注的博客
	QueryBlogOfFollow(ctx context.Context, in *QueryBlogOfFollowRequest, opts ...grpc.CallOption) (*QueryBlogOfFollowResponse, error)
	UpdateShop(ctx context.Context, in *UpdateShopRequest, opts ...grpc.CallOption) (*UpdateShopResponse, error)
}

type travelClient struct {
	cc grpc.ClientConnInterface
}

func NewTravelClient(cc grpc.ClientConnInterface) TravelClient {
	return &travelClient{cc}
}

func (c *travelClient) HomestayDetail(ctx context.Context, in *HomestayDetailReq, opts ...grpc.CallOption) (*HomestayDetailResp, error) {
	out := new(HomestayDetailResp)
	err := c.cc.Invoke(ctx, "/pb.travel/homestayDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) HomestayDetailWithoutLogin(ctx context.Context, in *HomestayDetailReq, opts ...grpc.CallOption) (*HomestayDetailResp, error) {
	out := new(HomestayDetailResp)
	err := c.cc.Invoke(ctx, "/pb.travel/homestayDetailWithoutLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) AddHomestay(ctx context.Context, in *AddHomestayReq, opts ...grpc.CallOption) (*AddHomestayResp, error) {
	out := new(AddHomestayResp)
	err := c.cc.Invoke(ctx, "/pb.travel/addHomestay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) UpdateHomestay(ctx context.Context, in *UpdateHomestayReq, opts ...grpc.CallOption) (*UpdateHomestayResp, error) {
	out := new(UpdateHomestayResp)
	err := c.cc.Invoke(ctx, "/pb.travel/updateHomestay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) DeleteHomestay(ctx context.Context, in *DeleteHomestayReq, opts ...grpc.CallOption) (*DeleteHomestayResp, error) {
	out := new(DeleteHomestayResp)
	err := c.cc.Invoke(ctx, "/pb.travel/deleteHomestay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) AdminDeleteHomestay(ctx context.Context, in *AdminDeleteHomestayReq, opts ...grpc.CallOption) (*AdminDeleteHomestayResp, error) {
	out := new(AdminDeleteHomestayResp)
	err := c.cc.Invoke(ctx, "/pb.travel/adminDeleteHomestay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) AddComment(ctx context.Context, in *AddCommentReq, opts ...grpc.CallOption) (*AddCommentResp, error) {
	out := new(AddCommentResp)
	err := c.cc.Invoke(ctx, "/pb.travel/addComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) LikeComment(ctx context.Context, in *LikeCommentReq, opts ...grpc.CallOption) (*LikeCommentResp, error) {
	out := new(LikeCommentResp)
	err := c.cc.Invoke(ctx, "/pb.travel/likeComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) WishList(ctx context.Context, in *WishListReq, opts ...grpc.CallOption) (*WishListResp, error) {
	out := new(WishListResp)
	err := c.cc.Invoke(ctx, "/pb.travel/wishList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) AddWishList(ctx context.Context, in *AddWishListReq, opts ...grpc.CallOption) (*AddWishListResp, error) {
	out := new(AddWishListResp)
	err := c.cc.Invoke(ctx, "/pb.travel/addWishList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) RemoveWishList(ctx context.Context, in *RemoveWishListReq, opts ...grpc.CallOption) (*RemoveWishListResp, error) {
	out := new(RemoveWishListResp)
	err := c.cc.Invoke(ctx, "/pb.travel/removeWishList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) AddGuess(ctx context.Context, in *AddGuessReq, opts ...grpc.CallOption) (*AddGuessResp, error) {
	out := new(AddGuessResp)
	err := c.cc.Invoke(ctx, "/pb.travel/addGuess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) HistoryList(ctx context.Context, in *HistoryListReq, opts ...grpc.CallOption) (*HistoryListResp, error) {
	out := new(HistoryListResp)
	err := c.cc.Invoke(ctx, "/pb.travel/historyList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) RemoveHistory(ctx context.Context, in *RemoveHistoryReq, opts ...grpc.CallOption) (*RemoveHistoryResp, error) {
	out := new(RemoveHistoryResp)
	err := c.cc.Invoke(ctx, "/pb.travel/removeHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) ClearHistory(ctx context.Context, in *ClearHistoryReq, opts ...grpc.CallOption) (*ClearHistoryResp, error) {
	out := new(ClearHistoryResp)
	err := c.cc.Invoke(ctx, "/pb.travel/clearHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) SearchHistory(ctx context.Context, in *SearchHistoryReq, opts ...grpc.CallOption) (*SearchHistoryResp, error) {
	out := new(SearchHistoryResp)
	err := c.cc.Invoke(ctx, "/pb.travel/searchHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) SearchByLocation(ctx context.Context, in *SearchByLocationReq, opts ...grpc.CallOption) (*SearchByLocationResp, error) {
	out := new(SearchByLocationResp)
	err := c.cc.Invoke(ctx, "/pb.travel/searchByLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) QueryShopByType(ctx context.Context, in *QueryShopByTypeRequest, opts ...grpc.CallOption) (*QueryShopByTypeResponse, error) {
	out := new(QueryShopByTypeResponse)
	err := c.cc.Invoke(ctx, "/pb.travel/queryShopByType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) SeckillVoucher(ctx context.Context, in *SeckillVoucherRequest, opts ...grpc.CallOption) (*SeckillVoucherResponse, error) {
	out := new(SeckillVoucherResponse)
	err := c.cc.Invoke(ctx, "/pb.travel/seckillVoucher", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) QueryHotBlog(ctx context.Context, in *QueryHotBlogRequest, opts ...grpc.CallOption) (*QueryHotBlogResponse, error) {
	out := new(QueryHotBlogResponse)
	err := c.cc.Invoke(ctx, "/pb.travel/QueryHotBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) QueryBlogById(ctx context.Context, in *QueryBlogByIdRequest, opts ...grpc.CallOption) (*QueryBlogByIdResponse, error) {
	out := new(QueryBlogByIdResponse)
	err := c.cc.Invoke(ctx, "/pb.travel/QueryBlogById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) LikeBlog(ctx context.Context, in *LikeBlogRequest, opts ...grpc.CallOption) (*LikeBlogResponse, error) {
	out := new(LikeBlogResponse)
	err := c.cc.Invoke(ctx, "/pb.travel/LikeBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) QueryBlogLikes(ctx context.Context, in *QueryBlogLikesRequest, opts ...grpc.CallOption) (*QueryBlogLikesResponse, error) {
	out := new(QueryBlogLikesResponse)
	err := c.cc.Invoke(ctx, "/pb.travel/QueryBlogLikes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) SaveBlog(ctx context.Context, in *SaveBlogRequest, opts ...grpc.CallOption) (*SaveBlogResponse, error) {
	out := new(SaveBlogResponse)
	err := c.cc.Invoke(ctx, "/pb.travel/SaveBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) QueryBlogOfFollow(ctx context.Context, in *QueryBlogOfFollowRequest, opts ...grpc.CallOption) (*QueryBlogOfFollowResponse, error) {
	out := new(QueryBlogOfFollowResponse)
	err := c.cc.Invoke(ctx, "/pb.travel/QueryBlogOfFollow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelClient) UpdateShop(ctx context.Context, in *UpdateShopRequest, opts ...grpc.CallOption) (*UpdateShopResponse, error) {
	out := new(UpdateShopResponse)
	err := c.cc.Invoke(ctx, "/pb.travel/UpdateShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TravelServer is the server API for Travel service.
// All implementations must embed UnimplementedTravelServer
// for forward compatibility
type TravelServer interface {
	// homestayDetail
	HomestayDetail(context.Context, *HomestayDetailReq) (*HomestayDetailResp, error)
	HomestayDetailWithoutLogin(context.Context, *HomestayDetailReq) (*HomestayDetailResp, error)
	AddHomestay(context.Context, *AddHomestayReq) (*AddHomestayResp, error)
	UpdateHomestay(context.Context, *UpdateHomestayReq) (*UpdateHomestayResp, error)
	DeleteHomestay(context.Context, *DeleteHomestayReq) (*DeleteHomestayResp, error)
	AdminDeleteHomestay(context.Context, *AdminDeleteHomestayReq) (*AdminDeleteHomestayResp, error)
	AddComment(context.Context, *AddCommentReq) (*AddCommentResp, error)
	LikeComment(context.Context, *LikeCommentReq) (*LikeCommentResp, error)
	WishList(context.Context, *WishListReq) (*WishListResp, error)
	AddWishList(context.Context, *AddWishListReq) (*AddWishListResp, error)
	RemoveWishList(context.Context, *RemoveWishListReq) (*RemoveWishListResp, error)
	AddGuess(context.Context, *AddGuessReq) (*AddGuessResp, error)
	HistoryList(context.Context, *HistoryListReq) (*HistoryListResp, error)
	RemoveHistory(context.Context, *RemoveHistoryReq) (*RemoveHistoryResp, error)
	ClearHistory(context.Context, *ClearHistoryReq) (*ClearHistoryResp, error)
	SearchHistory(context.Context, *SearchHistoryReq) (*SearchHistoryResp, error)
	SearchByLocation(context.Context, *SearchByLocationReq) (*SearchByLocationResp, error)
	QueryShopByType(context.Context, *QueryShopByTypeRequest) (*QueryShopByTypeResponse, error)
	SeckillVoucher(context.Context, *SeckillVoucherRequest) (*SeckillVoucherResponse, error)
	// 查询热门博客
	QueryHotBlog(context.Context, *QueryHotBlogRequest) (*QueryHotBlogResponse, error)
	// 根据ID查询博客
	QueryBlogById(context.Context, *QueryBlogByIdRequest) (*QueryBlogByIdResponse, error)
	// 点赞或取消点赞
	LikeBlog(context.Context, *LikeBlogRequest) (*LikeBlogResponse, error)
	// 查询某篇博客的点赞用户
	QueryBlogLikes(context.Context, *QueryBlogLikesRequest) (*QueryBlogLikesResponse, error)
	// 保存博客
	SaveBlog(context.Context, *SaveBlogRequest) (*SaveBlogResponse, error)
	// 查询关注的博客
	QueryBlogOfFollow(context.Context, *QueryBlogOfFollowRequest) (*QueryBlogOfFollowResponse, error)
	UpdateShop(context.Context, *UpdateShopRequest) (*UpdateShopResponse, error)
	mustEmbedUnimplementedTravelServer()
}

// UnimplementedTravelServer must be embedded to have forward compatible implementations.
type UnimplementedTravelServer struct {
}

func (UnimplementedTravelServer) HomestayDetail(context.Context, *HomestayDetailReq) (*HomestayDetailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HomestayDetail not implemented")
}
func (UnimplementedTravelServer) HomestayDetailWithoutLogin(context.Context, *HomestayDetailReq) (*HomestayDetailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HomestayDetailWithoutLogin not implemented")
}
func (UnimplementedTravelServer) AddHomestay(context.Context, *AddHomestayReq) (*AddHomestayResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddHomestay not implemented")
}
func (UnimplementedTravelServer) UpdateHomestay(context.Context, *UpdateHomestayReq) (*UpdateHomestayResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHomestay not implemented")
}
func (UnimplementedTravelServer) DeleteHomestay(context.Context, *DeleteHomestayReq) (*DeleteHomestayResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHomestay not implemented")
}
func (UnimplementedTravelServer) AdminDeleteHomestay(context.Context, *AdminDeleteHomestayReq) (*AdminDeleteHomestayResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminDeleteHomestay not implemented")
}
func (UnimplementedTravelServer) AddComment(context.Context, *AddCommentReq) (*AddCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (UnimplementedTravelServer) LikeComment(context.Context, *LikeCommentReq) (*LikeCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikeComment not implemented")
}
func (UnimplementedTravelServer) WishList(context.Context, *WishListReq) (*WishListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WishList not implemented")
}
func (UnimplementedTravelServer) AddWishList(context.Context, *AddWishListReq) (*AddWishListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddWishList not implemented")
}
func (UnimplementedTravelServer) RemoveWishList(context.Context, *RemoveWishListReq) (*RemoveWishListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveWishList not implemented")
}
func (UnimplementedTravelServer) AddGuess(context.Context, *AddGuessReq) (*AddGuessResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGuess not implemented")
}
func (UnimplementedTravelServer) HistoryList(context.Context, *HistoryListReq) (*HistoryListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HistoryList not implemented")
}
func (UnimplementedTravelServer) RemoveHistory(context.Context, *RemoveHistoryReq) (*RemoveHistoryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveHistory not implemented")
}
func (UnimplementedTravelServer) ClearHistory(context.Context, *ClearHistoryReq) (*ClearHistoryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearHistory not implemented")
}
func (UnimplementedTravelServer) SearchHistory(context.Context, *SearchHistoryReq) (*SearchHistoryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchHistory not implemented")
}
func (UnimplementedTravelServer) SearchByLocation(context.Context, *SearchByLocationReq) (*SearchByLocationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchByLocation not implemented")
}
func (UnimplementedTravelServer) QueryShopByType(context.Context, *QueryShopByTypeRequest) (*QueryShopByTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryShopByType not implemented")
}
func (UnimplementedTravelServer) SeckillVoucher(context.Context, *SeckillVoucherRequest) (*SeckillVoucherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SeckillVoucher not implemented")
}
func (UnimplementedTravelServer) QueryHotBlog(context.Context, *QueryHotBlogRequest) (*QueryHotBlogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryHotBlog not implemented")
}
func (UnimplementedTravelServer) QueryBlogById(context.Context, *QueryBlogByIdRequest) (*QueryBlogByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryBlogById not implemented")
}
func (UnimplementedTravelServer) LikeBlog(context.Context, *LikeBlogRequest) (*LikeBlogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikeBlog not implemented")
}
func (UnimplementedTravelServer) QueryBlogLikes(context.Context, *QueryBlogLikesRequest) (*QueryBlogLikesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryBlogLikes not implemented")
}
func (UnimplementedTravelServer) SaveBlog(context.Context, *SaveBlogRequest) (*SaveBlogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveBlog not implemented")
}
func (UnimplementedTravelServer) QueryBlogOfFollow(context.Context, *QueryBlogOfFollowRequest) (*QueryBlogOfFollowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryBlogOfFollow not implemented")
}
func (UnimplementedTravelServer) UpdateShop(context.Context, *UpdateShopRequest) (*UpdateShopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateShop not implemented")
}
func (UnimplementedTravelServer) mustEmbedUnimplementedTravelServer() {}

// UnsafeTravelServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TravelServer will
// result in compilation errors.
type UnsafeTravelServer interface {
	mustEmbedUnimplementedTravelServer()
}

func RegisterTravelServer(s grpc.ServiceRegistrar, srv TravelServer) {
	s.RegisterService(&Travel_ServiceDesc, srv)
}

func _Travel_HomestayDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HomestayDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).HomestayDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.travel/homestayDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).HomestayDetail(ctx, req.(*HomestayDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_HomestayDetailWithoutLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HomestayDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).HomestayDetailWithoutLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.travel/homestayDetailWithoutLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).HomestayDetailWithoutLogin(ctx, req.(*HomestayDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_AddHomestay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddHomestayReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).AddHomestay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.travel/addHomestay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).AddHomestay(ctx, req.(*AddHomestayReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_UpdateHomestay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHomestayReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).UpdateHomestay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.travel/updateHomestay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).UpdateHomestay(ctx, req.(*UpdateHomestayReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_DeleteHomestay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHomestayReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).DeleteHomestay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.travel/deleteHomestay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).DeleteHomestay(ctx, req.(*DeleteHomestayReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_AdminDeleteHomestay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminDeleteHomestayReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).AdminDeleteHomestay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.travel/adminDeleteHomestay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).AdminDeleteHomestay(ctx, req.(*AdminDeleteHomestayReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.travel/addComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).AddComment(ctx, req.(*AddCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_LikeComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).LikeComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.travel/likeComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).LikeComment(ctx, req.(*LikeCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_WishList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WishListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).WishList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.travel/wishList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).WishList(ctx, req.(*WishListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_AddWishList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddWishListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).AddWishList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.travel/addWishList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).AddWishList(ctx, req.(*AddWishListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_RemoveWishList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveWishListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).RemoveWishList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.travel/removeWishList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).RemoveWishList(ctx, req.(*RemoveWishListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_AddGuess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGuessReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).AddGuess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.travel/addGuess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).AddGuess(ctx, req.(*AddGuessReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_HistoryList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HistoryListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).HistoryList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.travel/historyList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).HistoryList(ctx, req.(*HistoryListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_RemoveHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveHistoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).RemoveHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.travel/removeHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).RemoveHistory(ctx, req.(*RemoveHistoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_ClearHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearHistoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).ClearHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.travel/clearHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).ClearHistory(ctx, req.(*ClearHistoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_SearchHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchHistoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).SearchHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.travel/searchHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).SearchHistory(ctx, req.(*SearchHistoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_SearchByLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchByLocationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).SearchByLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.travel/searchByLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).SearchByLocation(ctx, req.(*SearchByLocationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_QueryShopByType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryShopByTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).QueryShopByType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.travel/queryShopByType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).QueryShopByType(ctx, req.(*QueryShopByTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_SeckillVoucher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeckillVoucherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).SeckillVoucher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.travel/seckillVoucher",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).SeckillVoucher(ctx, req.(*SeckillVoucherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_QueryHotBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryHotBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).QueryHotBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.travel/QueryHotBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).QueryHotBlog(ctx, req.(*QueryHotBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_QueryBlogById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBlogByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).QueryBlogById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.travel/QueryBlogById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).QueryBlogById(ctx, req.(*QueryBlogByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_LikeBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).LikeBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.travel/LikeBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).LikeBlog(ctx, req.(*LikeBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_QueryBlogLikes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBlogLikesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).QueryBlogLikes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.travel/QueryBlogLikes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).QueryBlogLikes(ctx, req.(*QueryBlogLikesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_SaveBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).SaveBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.travel/SaveBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).SaveBlog(ctx, req.(*SaveBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_QueryBlogOfFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBlogOfFollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).QueryBlogOfFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.travel/QueryBlogOfFollow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).QueryBlogOfFollow(ctx, req.(*QueryBlogOfFollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Travel_UpdateShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateShopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).UpdateShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.travel/UpdateShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).UpdateShop(ctx, req.(*UpdateShopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Travel_ServiceDesc is the grpc.ServiceDesc for Travel service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Travel_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.travel",
	HandlerType: (*TravelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "homestayDetail",
			Handler:    _Travel_HomestayDetail_Handler,
		},
		{
			MethodName: "homestayDetailWithoutLogin",
			Handler:    _Travel_HomestayDetailWithoutLogin_Handler,
		},
		{
			MethodName: "addHomestay",
			Handler:    _Travel_AddHomestay_Handler,
		},
		{
			MethodName: "updateHomestay",
			Handler:    _Travel_UpdateHomestay_Handler,
		},
		{
			MethodName: "deleteHomestay",
			Handler:    _Travel_DeleteHomestay_Handler,
		},
		{
			MethodName: "adminDeleteHomestay",
			Handler:    _Travel_AdminDeleteHomestay_Handler,
		},
		{
			MethodName: "addComment",
			Handler:    _Travel_AddComment_Handler,
		},
		{
			MethodName: "likeComment",
			Handler:    _Travel_LikeComment_Handler,
		},
		{
			MethodName: "wishList",
			Handler:    _Travel_WishList_Handler,
		},
		{
			MethodName: "addWishList",
			Handler:    _Travel_AddWishList_Handler,
		},
		{
			MethodName: "removeWishList",
			Handler:    _Travel_RemoveWishList_Handler,
		},
		{
			MethodName: "addGuess",
			Handler:    _Travel_AddGuess_Handler,
		},
		{
			MethodName: "historyList",
			Handler:    _Travel_HistoryList_Handler,
		},
		{
			MethodName: "removeHistory",
			Handler:    _Travel_RemoveHistory_Handler,
		},
		{
			MethodName: "clearHistory",
			Handler:    _Travel_ClearHistory_Handler,
		},
		{
			MethodName: "searchHistory",
			Handler:    _Travel_SearchHistory_Handler,
		},
		{
			MethodName: "searchByLocation",
			Handler:    _Travel_SearchByLocation_Handler,
		},
		{
			MethodName: "queryShopByType",
			Handler:    _Travel_QueryShopByType_Handler,
		},
		{
			MethodName: "seckillVoucher",
			Handler:    _Travel_SeckillVoucher_Handler,
		},
		{
			MethodName: "QueryHotBlog",
			Handler:    _Travel_QueryHotBlog_Handler,
		},
		{
			MethodName: "QueryBlogById",
			Handler:    _Travel_QueryBlogById_Handler,
		},
		{
			MethodName: "LikeBlog",
			Handler:    _Travel_LikeBlog_Handler,
		},
		{
			MethodName: "QueryBlogLikes",
			Handler:    _Travel_QueryBlogLikes_Handler,
		},
		{
			MethodName: "SaveBlog",
			Handler:    _Travel_SaveBlog_Handler,
		},
		{
			MethodName: "QueryBlogOfFollow",
			Handler:    _Travel_QueryBlogOfFollow_Handler,
		},
		{
			MethodName: "UpdateShop",
			Handler:    _Travel_UpdateShop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "travel.proto",
}
