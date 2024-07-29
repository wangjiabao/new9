// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.7
// source: app/app/api/app.proto

package api

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

const (
	App_EthAuthorize_FullMethodName        = "/api.App/EthAuthorize"
	App_RecommendUpdate_FullMethodName     = "/api.App/RecommendUpdate"
	App_UserInfo_FullMethodName            = "/api.App/UserInfo"
	App_UserArea_FullMethodName            = "/api.App/UserArea"
	App_RewardList_FullMethodName          = "/api.App/RewardList"
	App_RecommendRewardList_FullMethodName = "/api.App/RecommendRewardList"
	App_FeeRewardList_FullMethodName       = "/api.App/FeeRewardList"
	App_WithdrawList_FullMethodName        = "/api.App/WithdrawList"
	App_TradeList_FullMethodName           = "/api.App/TradeList"
	App_TranList_FullMethodName            = "/api.App/TranList"
	App_RecommendList_FullMethodName       = "/api.App/RecommendList"
	App_PasswordChange_FullMethodName      = "/api.App/PasswordChange"
	App_Withdraw_FullMethodName            = "/api.App/Withdraw"
	App_Buy_FullMethodName                 = "/api.App/Buy"
	App_Exchange_FullMethodName            = "/api.App/Exchange"
	App_Trade_FullMethodName               = "/api.App/Trade"
	App_Tran_FullMethodName                = "/api.App/Tran"
	App_GetTrade_FullMethodName            = "/api.App/GetTrade"
	App_SetBalanceReward_FullMethodName    = "/api.App/SetBalanceReward"
	App_DeleteBalanceReward_FullMethodName = "/api.App/DeleteBalanceReward"
	App_Deposit_FullMethodName             = "/api.App/Deposit"
	App_AdminWithdraw_FullMethodName       = "/api.App/AdminWithdraw"
	App_AdminWithdrawEth_FullMethodName    = "/api.App/AdminWithdrawEth"
	App_AdminFee_FullMethodName            = "/api.App/AdminFee"
	App_TokenWithdraw_FullMethodName       = "/api.App/TokenWithdraw"
)

// AppClient is the client API for App service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppClient interface {
	EthAuthorize(ctx context.Context, in *EthAuthorizeRequest, opts ...grpc.CallOption) (*EthAuthorizeReply, error)
	RecommendUpdate(ctx context.Context, in *RecommendUpdateRequest, opts ...grpc.CallOption) (*RecommendUpdateReply, error)
	UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoReply, error)
	UserArea(ctx context.Context, in *UserAreaRequest, opts ...grpc.CallOption) (*UserAreaReply, error)
	RewardList(ctx context.Context, in *RewardListRequest, opts ...grpc.CallOption) (*RewardListReply, error)
	RecommendRewardList(ctx context.Context, in *RecommendRewardListRequest, opts ...grpc.CallOption) (*RecommendRewardListReply, error)
	FeeRewardList(ctx context.Context, in *FeeRewardListRequest, opts ...grpc.CallOption) (*FeeRewardListReply, error)
	WithdrawList(ctx context.Context, in *WithdrawListRequest, opts ...grpc.CallOption) (*WithdrawListReply, error)
	TradeList(ctx context.Context, in *TradeListRequest, opts ...grpc.CallOption) (*TradeListReply, error)
	TranList(ctx context.Context, in *TranListRequest, opts ...grpc.CallOption) (*TranListReply, error)
	RecommendList(ctx context.Context, in *RecommendListRequest, opts ...grpc.CallOption) (*RecommendListReply, error)
	PasswordChange(ctx context.Context, in *PasswordChangeRequest, opts ...grpc.CallOption) (*PasswordChangeReply, error)
	Withdraw(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*WithdrawReply, error)
	Buy(ctx context.Context, in *BuyRequest, opts ...grpc.CallOption) (*BuyReply, error)
	Exchange(ctx context.Context, in *ExchangeRequest, opts ...grpc.CallOption) (*ExchangeReply, error)
	Trade(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*WithdrawReply, error)
	Tran(ctx context.Context, in *TranRequest, opts ...grpc.CallOption) (*TranReply, error)
	GetTrade(ctx context.Context, in *GetTradeRequest, opts ...grpc.CallOption) (*GetTradeReply, error)
	SetBalanceReward(ctx context.Context, in *SetBalanceRewardRequest, opts ...grpc.CallOption) (*SetBalanceRewardReply, error)
	DeleteBalanceReward(ctx context.Context, in *DeleteBalanceRewardRequest, opts ...grpc.CallOption) (*DeleteBalanceRewardReply, error)
	Deposit(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*DepositReply, error)
	//	rpc AdminRewardList (AdminRewardListRequest) returns (AdminRewardListReply) {
	//		option (google.api.http) = {
	//			get: "/api/admin_dhb/reward_list"
	//		};
	//	};
	//
	//	rpc AdminUserList (AdminUserListRequest) returns (AdminUserListReply) {
	//		option (google.api.http) = {
	//			get: "/api/admin_dhb/user_list"
	//		};
	//	};
	//
	//	rpc AdminLocationList (AdminLocationListRequest) returns (AdminLocationListReply) {
	//		option (google.api.http) = {
	//			get: "/api/admin_dhb/location_list"
	//		};
	//	};
	//
	//	rpc AdminWithdrawList (AdminWithdrawListRequest) returns (AdminWithdrawListReply) {
	//		option (google.api.http) = {
	//			get: "/api/admin_dhb/withdraw_list"
	//		};
	//	};
	AdminWithdraw(ctx context.Context, in *AdminWithdrawRequest, opts ...grpc.CallOption) (*AdminWithdrawReply, error)
	AdminWithdrawEth(ctx context.Context, in *AdminWithdrawEthRequest, opts ...grpc.CallOption) (*AdminWithdrawEthReply, error)
	AdminFee(ctx context.Context, in *AdminFeeRequest, opts ...grpc.CallOption) (*AdminFeeReply, error)
	TokenWithdraw(ctx context.Context, in *TokenWithdrawRequest, opts ...grpc.CallOption) (*TokenWithdrawReply, error)
}

type appClient struct {
	cc grpc.ClientConnInterface
}

func NewAppClient(cc grpc.ClientConnInterface) AppClient {
	return &appClient{cc}
}

func (c *appClient) EthAuthorize(ctx context.Context, in *EthAuthorizeRequest, opts ...grpc.CallOption) (*EthAuthorizeReply, error) {
	out := new(EthAuthorizeReply)
	err := c.cc.Invoke(ctx, App_EthAuthorize_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) RecommendUpdate(ctx context.Context, in *RecommendUpdateRequest, opts ...grpc.CallOption) (*RecommendUpdateReply, error) {
	out := new(RecommendUpdateReply)
	err := c.cc.Invoke(ctx, App_RecommendUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoReply, error) {
	out := new(UserInfoReply)
	err := c.cc.Invoke(ctx, App_UserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) UserArea(ctx context.Context, in *UserAreaRequest, opts ...grpc.CallOption) (*UserAreaReply, error) {
	out := new(UserAreaReply)
	err := c.cc.Invoke(ctx, App_UserArea_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) RewardList(ctx context.Context, in *RewardListRequest, opts ...grpc.CallOption) (*RewardListReply, error) {
	out := new(RewardListReply)
	err := c.cc.Invoke(ctx, App_RewardList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) RecommendRewardList(ctx context.Context, in *RecommendRewardListRequest, opts ...grpc.CallOption) (*RecommendRewardListReply, error) {
	out := new(RecommendRewardListReply)
	err := c.cc.Invoke(ctx, App_RecommendRewardList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) FeeRewardList(ctx context.Context, in *FeeRewardListRequest, opts ...grpc.CallOption) (*FeeRewardListReply, error) {
	out := new(FeeRewardListReply)
	err := c.cc.Invoke(ctx, App_FeeRewardList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) WithdrawList(ctx context.Context, in *WithdrawListRequest, opts ...grpc.CallOption) (*WithdrawListReply, error) {
	out := new(WithdrawListReply)
	err := c.cc.Invoke(ctx, App_WithdrawList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) TradeList(ctx context.Context, in *TradeListRequest, opts ...grpc.CallOption) (*TradeListReply, error) {
	out := new(TradeListReply)
	err := c.cc.Invoke(ctx, App_TradeList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) TranList(ctx context.Context, in *TranListRequest, opts ...grpc.CallOption) (*TranListReply, error) {
	out := new(TranListReply)
	err := c.cc.Invoke(ctx, App_TranList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) RecommendList(ctx context.Context, in *RecommendListRequest, opts ...grpc.CallOption) (*RecommendListReply, error) {
	out := new(RecommendListReply)
	err := c.cc.Invoke(ctx, App_RecommendList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) PasswordChange(ctx context.Context, in *PasswordChangeRequest, opts ...grpc.CallOption) (*PasswordChangeReply, error) {
	out := new(PasswordChangeReply)
	err := c.cc.Invoke(ctx, App_PasswordChange_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) Withdraw(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*WithdrawReply, error) {
	out := new(WithdrawReply)
	err := c.cc.Invoke(ctx, App_Withdraw_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) Buy(ctx context.Context, in *BuyRequest, opts ...grpc.CallOption) (*BuyReply, error) {
	out := new(BuyReply)
	err := c.cc.Invoke(ctx, App_Buy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) Exchange(ctx context.Context, in *ExchangeRequest, opts ...grpc.CallOption) (*ExchangeReply, error) {
	out := new(ExchangeReply)
	err := c.cc.Invoke(ctx, App_Exchange_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) Trade(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*WithdrawReply, error) {
	out := new(WithdrawReply)
	err := c.cc.Invoke(ctx, App_Trade_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) Tran(ctx context.Context, in *TranRequest, opts ...grpc.CallOption) (*TranReply, error) {
	out := new(TranReply)
	err := c.cc.Invoke(ctx, App_Tran_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetTrade(ctx context.Context, in *GetTradeRequest, opts ...grpc.CallOption) (*GetTradeReply, error) {
	out := new(GetTradeReply)
	err := c.cc.Invoke(ctx, App_GetTrade_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) SetBalanceReward(ctx context.Context, in *SetBalanceRewardRequest, opts ...grpc.CallOption) (*SetBalanceRewardReply, error) {
	out := new(SetBalanceRewardReply)
	err := c.cc.Invoke(ctx, App_SetBalanceReward_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) DeleteBalanceReward(ctx context.Context, in *DeleteBalanceRewardRequest, opts ...grpc.CallOption) (*DeleteBalanceRewardReply, error) {
	out := new(DeleteBalanceRewardReply)
	err := c.cc.Invoke(ctx, App_DeleteBalanceReward_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) Deposit(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*DepositReply, error) {
	out := new(DepositReply)
	err := c.cc.Invoke(ctx, App_Deposit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) AdminWithdraw(ctx context.Context, in *AdminWithdrawRequest, opts ...grpc.CallOption) (*AdminWithdrawReply, error) {
	out := new(AdminWithdrawReply)
	err := c.cc.Invoke(ctx, App_AdminWithdraw_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) AdminWithdrawEth(ctx context.Context, in *AdminWithdrawEthRequest, opts ...grpc.CallOption) (*AdminWithdrawEthReply, error) {
	out := new(AdminWithdrawEthReply)
	err := c.cc.Invoke(ctx, App_AdminWithdrawEth_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) AdminFee(ctx context.Context, in *AdminFeeRequest, opts ...grpc.CallOption) (*AdminFeeReply, error) {
	out := new(AdminFeeReply)
	err := c.cc.Invoke(ctx, App_AdminFee_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) TokenWithdraw(ctx context.Context, in *TokenWithdrawRequest, opts ...grpc.CallOption) (*TokenWithdrawReply, error) {
	out := new(TokenWithdrawReply)
	err := c.cc.Invoke(ctx, App_TokenWithdraw_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppServer is the server API for App service.
// All implementations must embed UnimplementedAppServer
// for forward compatibility
type AppServer interface {
	EthAuthorize(context.Context, *EthAuthorizeRequest) (*EthAuthorizeReply, error)
	RecommendUpdate(context.Context, *RecommendUpdateRequest) (*RecommendUpdateReply, error)
	UserInfo(context.Context, *UserInfoRequest) (*UserInfoReply, error)
	UserArea(context.Context, *UserAreaRequest) (*UserAreaReply, error)
	RewardList(context.Context, *RewardListRequest) (*RewardListReply, error)
	RecommendRewardList(context.Context, *RecommendRewardListRequest) (*RecommendRewardListReply, error)
	FeeRewardList(context.Context, *FeeRewardListRequest) (*FeeRewardListReply, error)
	WithdrawList(context.Context, *WithdrawListRequest) (*WithdrawListReply, error)
	TradeList(context.Context, *TradeListRequest) (*TradeListReply, error)
	TranList(context.Context, *TranListRequest) (*TranListReply, error)
	RecommendList(context.Context, *RecommendListRequest) (*RecommendListReply, error)
	PasswordChange(context.Context, *PasswordChangeRequest) (*PasswordChangeReply, error)
	Withdraw(context.Context, *WithdrawRequest) (*WithdrawReply, error)
	Buy(context.Context, *BuyRequest) (*BuyReply, error)
	Exchange(context.Context, *ExchangeRequest) (*ExchangeReply, error)
	Trade(context.Context, *WithdrawRequest) (*WithdrawReply, error)
	Tran(context.Context, *TranRequest) (*TranReply, error)
	GetTrade(context.Context, *GetTradeRequest) (*GetTradeReply, error)
	SetBalanceReward(context.Context, *SetBalanceRewardRequest) (*SetBalanceRewardReply, error)
	DeleteBalanceReward(context.Context, *DeleteBalanceRewardRequest) (*DeleteBalanceRewardReply, error)
	Deposit(context.Context, *DepositRequest) (*DepositReply, error)
	//	rpc AdminRewardList (AdminRewardListRequest) returns (AdminRewardListReply) {
	//		option (google.api.http) = {
	//			get: "/api/admin_dhb/reward_list"
	//		};
	//	};
	//
	//	rpc AdminUserList (AdminUserListRequest) returns (AdminUserListReply) {
	//		option (google.api.http) = {
	//			get: "/api/admin_dhb/user_list"
	//		};
	//	};
	//
	//	rpc AdminLocationList (AdminLocationListRequest) returns (AdminLocationListReply) {
	//		option (google.api.http) = {
	//			get: "/api/admin_dhb/location_list"
	//		};
	//	};
	//
	//	rpc AdminWithdrawList (AdminWithdrawListRequest) returns (AdminWithdrawListReply) {
	//		option (google.api.http) = {
	//			get: "/api/admin_dhb/withdraw_list"
	//		};
	//	};
	AdminWithdraw(context.Context, *AdminWithdrawRequest) (*AdminWithdrawReply, error)
	AdminWithdrawEth(context.Context, *AdminWithdrawEthRequest) (*AdminWithdrawEthReply, error)
	AdminFee(context.Context, *AdminFeeRequest) (*AdminFeeReply, error)
	TokenWithdraw(context.Context, *TokenWithdrawRequest) (*TokenWithdrawReply, error)
	mustEmbedUnimplementedAppServer()
}

// UnimplementedAppServer must be embedded to have forward compatible implementations.
type UnimplementedAppServer struct {
}

func (UnimplementedAppServer) EthAuthorize(context.Context, *EthAuthorizeRequest) (*EthAuthorizeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EthAuthorize not implemented")
}
func (UnimplementedAppServer) RecommendUpdate(context.Context, *RecommendUpdateRequest) (*RecommendUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecommendUpdate not implemented")
}
func (UnimplementedAppServer) UserInfo(context.Context, *UserInfoRequest) (*UserInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (UnimplementedAppServer) UserArea(context.Context, *UserAreaRequest) (*UserAreaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserArea not implemented")
}
func (UnimplementedAppServer) RewardList(context.Context, *RewardListRequest) (*RewardListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RewardList not implemented")
}
func (UnimplementedAppServer) RecommendRewardList(context.Context, *RecommendRewardListRequest) (*RecommendRewardListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecommendRewardList not implemented")
}
func (UnimplementedAppServer) FeeRewardList(context.Context, *FeeRewardListRequest) (*FeeRewardListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeeRewardList not implemented")
}
func (UnimplementedAppServer) WithdrawList(context.Context, *WithdrawListRequest) (*WithdrawListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawList not implemented")
}
func (UnimplementedAppServer) TradeList(context.Context, *TradeListRequest) (*TradeListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradeList not implemented")
}
func (UnimplementedAppServer) TranList(context.Context, *TranListRequest) (*TranListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TranList not implemented")
}
func (UnimplementedAppServer) RecommendList(context.Context, *RecommendListRequest) (*RecommendListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecommendList not implemented")
}
func (UnimplementedAppServer) PasswordChange(context.Context, *PasswordChangeRequest) (*PasswordChangeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PasswordChange not implemented")
}
func (UnimplementedAppServer) Withdraw(context.Context, *WithdrawRequest) (*WithdrawReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Withdraw not implemented")
}
func (UnimplementedAppServer) Buy(context.Context, *BuyRequest) (*BuyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Buy not implemented")
}
func (UnimplementedAppServer) Exchange(context.Context, *ExchangeRequest) (*ExchangeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exchange not implemented")
}
func (UnimplementedAppServer) Trade(context.Context, *WithdrawRequest) (*WithdrawReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Trade not implemented")
}
func (UnimplementedAppServer) Tran(context.Context, *TranRequest) (*TranReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Tran not implemented")
}
func (UnimplementedAppServer) GetTrade(context.Context, *GetTradeRequest) (*GetTradeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrade not implemented")
}
func (UnimplementedAppServer) SetBalanceReward(context.Context, *SetBalanceRewardRequest) (*SetBalanceRewardReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetBalanceReward not implemented")
}
func (UnimplementedAppServer) DeleteBalanceReward(context.Context, *DeleteBalanceRewardRequest) (*DeleteBalanceRewardReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBalanceReward not implemented")
}
func (UnimplementedAppServer) Deposit(context.Context, *DepositRequest) (*DepositReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
}
func (UnimplementedAppServer) AdminWithdraw(context.Context, *AdminWithdrawRequest) (*AdminWithdrawReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminWithdraw not implemented")
}
func (UnimplementedAppServer) AdminWithdrawEth(context.Context, *AdminWithdrawEthRequest) (*AdminWithdrawEthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminWithdrawEth not implemented")
}
func (UnimplementedAppServer) AdminFee(context.Context, *AdminFeeRequest) (*AdminFeeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminFee not implemented")
}
func (UnimplementedAppServer) TokenWithdraw(context.Context, *TokenWithdrawRequest) (*TokenWithdrawReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TokenWithdraw not implemented")
}
func (UnimplementedAppServer) mustEmbedUnimplementedAppServer() {}

// UnsafeAppServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppServer will
// result in compilation errors.
type UnsafeAppServer interface {
	mustEmbedUnimplementedAppServer()
}

func RegisterAppServer(s grpc.ServiceRegistrar, srv AppServer) {
	s.RegisterService(&App_ServiceDesc, srv)
}

func _App_EthAuthorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EthAuthorizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).EthAuthorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_EthAuthorize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).EthAuthorize(ctx, req.(*EthAuthorizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_RecommendUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecommendUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).RecommendUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_RecommendUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).RecommendUpdate(ctx, req.(*RecommendUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_UserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).UserInfo(ctx, req.(*UserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_UserArea_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAreaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).UserArea(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_UserArea_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).UserArea(ctx, req.(*UserAreaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_RewardList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RewardListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).RewardList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_RewardList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).RewardList(ctx, req.(*RewardListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_RecommendRewardList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecommendRewardListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).RecommendRewardList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_RecommendRewardList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).RecommendRewardList(ctx, req.(*RecommendRewardListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_FeeRewardList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeeRewardListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).FeeRewardList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_FeeRewardList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).FeeRewardList(ctx, req.(*FeeRewardListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_WithdrawList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).WithdrawList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_WithdrawList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).WithdrawList(ctx, req.(*WithdrawListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_TradeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradeListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).TradeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_TradeList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).TradeList(ctx, req.(*TradeListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_TranList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).TranList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_TranList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).TranList(ctx, req.(*TranListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_RecommendList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecommendListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).RecommendList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_RecommendList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).RecommendList(ctx, req.(*RecommendListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_PasswordChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).PasswordChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_PasswordChange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).PasswordChange(ctx, req.(*PasswordChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_Withdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Withdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_Withdraw_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Withdraw(ctx, req.(*WithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_Buy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Buy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_Buy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Buy(ctx, req.(*BuyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_Exchange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExchangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Exchange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_Exchange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Exchange(ctx, req.(*ExchangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_Trade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Trade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_Trade_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Trade(ctx, req.(*WithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_Tran_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Tran(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_Tran_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Tran(ctx, req.(*TranRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetTrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetTrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_GetTrade_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetTrade(ctx, req.(*GetTradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_SetBalanceReward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetBalanceRewardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).SetBalanceReward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_SetBalanceReward_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).SetBalanceReward(ctx, req.(*SetBalanceRewardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_DeleteBalanceReward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBalanceRewardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).DeleteBalanceReward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_DeleteBalanceReward_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).DeleteBalanceReward(ctx, req.(*DeleteBalanceRewardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_Deposit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Deposit(ctx, req.(*DepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_AdminWithdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminWithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).AdminWithdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_AdminWithdraw_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).AdminWithdraw(ctx, req.(*AdminWithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_AdminWithdrawEth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminWithdrawEthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).AdminWithdrawEth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_AdminWithdrawEth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).AdminWithdrawEth(ctx, req.(*AdminWithdrawEthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_AdminFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).AdminFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_AdminFee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).AdminFee(ctx, req.(*AdminFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_TokenWithdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenWithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).TokenWithdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_TokenWithdraw_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).TokenWithdraw(ctx, req.(*TokenWithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// App_ServiceDesc is the grpc.ServiceDesc for App service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var App_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.App",
	HandlerType: (*AppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EthAuthorize",
			Handler:    _App_EthAuthorize_Handler,
		},
		{
			MethodName: "RecommendUpdate",
			Handler:    _App_RecommendUpdate_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _App_UserInfo_Handler,
		},
		{
			MethodName: "UserArea",
			Handler:    _App_UserArea_Handler,
		},
		{
			MethodName: "RewardList",
			Handler:    _App_RewardList_Handler,
		},
		{
			MethodName: "RecommendRewardList",
			Handler:    _App_RecommendRewardList_Handler,
		},
		{
			MethodName: "FeeRewardList",
			Handler:    _App_FeeRewardList_Handler,
		},
		{
			MethodName: "WithdrawList",
			Handler:    _App_WithdrawList_Handler,
		},
		{
			MethodName: "TradeList",
			Handler:    _App_TradeList_Handler,
		},
		{
			MethodName: "TranList",
			Handler:    _App_TranList_Handler,
		},
		{
			MethodName: "RecommendList",
			Handler:    _App_RecommendList_Handler,
		},
		{
			MethodName: "PasswordChange",
			Handler:    _App_PasswordChange_Handler,
		},
		{
			MethodName: "Withdraw",
			Handler:    _App_Withdraw_Handler,
		},
		{
			MethodName: "Buy",
			Handler:    _App_Buy_Handler,
		},
		{
			MethodName: "Exchange",
			Handler:    _App_Exchange_Handler,
		},
		{
			MethodName: "Trade",
			Handler:    _App_Trade_Handler,
		},
		{
			MethodName: "Tran",
			Handler:    _App_Tran_Handler,
		},
		{
			MethodName: "GetTrade",
			Handler:    _App_GetTrade_Handler,
		},
		{
			MethodName: "SetBalanceReward",
			Handler:    _App_SetBalanceReward_Handler,
		},
		{
			MethodName: "DeleteBalanceReward",
			Handler:    _App_DeleteBalanceReward_Handler,
		},
		{
			MethodName: "Deposit",
			Handler:    _App_Deposit_Handler,
		},
		{
			MethodName: "AdminWithdraw",
			Handler:    _App_AdminWithdraw_Handler,
		},
		{
			MethodName: "AdminWithdrawEth",
			Handler:    _App_AdminWithdrawEth_Handler,
		},
		{
			MethodName: "AdminFee",
			Handler:    _App_AdminFee_Handler,
		},
		{
			MethodName: "TokenWithdraw",
			Handler:    _App_TokenWithdraw_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/app/api/app.proto",
}
