// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             v3.21.7
// source: app/app/api/app.proto

package api

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationAppAdminFee = "/api.App/AdminFee"
const OperationAppAdminWithdraw = "/api.App/AdminWithdraw"
const OperationAppAdminWithdrawEth = "/api.App/AdminWithdrawEth"
const OperationAppDeleteBalanceReward = "/api.App/DeleteBalanceReward"
const OperationAppDeposit = "/api.App/Deposit"
const OperationAppEthAuthorize = "/api.App/EthAuthorize"
const OperationAppExchange = "/api.App/Exchange"
const OperationAppFeeRewardList = "/api.App/FeeRewardList"
const OperationAppGetTrade = "/api.App/GetTrade"
const OperationAppPasswordChange = "/api.App/PasswordChange"
const OperationAppRecommendList = "/api.App/RecommendList"
const OperationAppRecommendRewardList = "/api.App/RecommendRewardList"
const OperationAppRecommendUpdate = "/api.App/RecommendUpdate"
const OperationAppRewardList = "/api.App/RewardList"
const OperationAppSetBalanceReward = "/api.App/SetBalanceReward"
const OperationAppTokenWithdraw = "/api.App/TokenWithdraw"
const OperationAppTrade = "/api.App/Trade"
const OperationAppTradeList = "/api.App/TradeList"
const OperationAppTran = "/api.App/Tran"
const OperationAppTranList = "/api.App/TranList"
const OperationAppUserArea = "/api.App/UserArea"
const OperationAppUserInfo = "/api.App/UserInfo"
const OperationAppWithdraw = "/api.App/Withdraw"
const OperationAppWithdrawList = "/api.App/WithdrawList"

type AppHTTPServer interface {
	AdminFee(context.Context, *AdminFeeRequest) (*AdminFeeReply, error)
	// AdminWithdraw
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
	//
	AdminWithdraw(context.Context, *AdminWithdrawRequest) (*AdminWithdrawReply, error)
	AdminWithdrawEth(context.Context, *AdminWithdrawEthRequest) (*AdminWithdrawEthReply, error)
	DeleteBalanceReward(context.Context, *DeleteBalanceRewardRequest) (*DeleteBalanceRewardReply, error)
	Deposit(context.Context, *DepositRequest) (*DepositReply, error)
	EthAuthorize(context.Context, *EthAuthorizeRequest) (*EthAuthorizeReply, error)
	Exchange(context.Context, *ExchangeRequest) (*ExchangeReply, error)
	FeeRewardList(context.Context, *FeeRewardListRequest) (*FeeRewardListReply, error)
	GetTrade(context.Context, *GetTradeRequest) (*GetTradeReply, error)
	PasswordChange(context.Context, *PasswordChangeRequest) (*PasswordChangeReply, error)
	RecommendList(context.Context, *RecommendListRequest) (*RecommendListReply, error)
	RecommendRewardList(context.Context, *RecommendRewardListRequest) (*RecommendRewardListReply, error)
	RecommendUpdate(context.Context, *RecommendUpdateRequest) (*RecommendUpdateReply, error)
	RewardList(context.Context, *RewardListRequest) (*RewardListReply, error)
	SetBalanceReward(context.Context, *SetBalanceRewardRequest) (*SetBalanceRewardReply, error)
	TokenWithdraw(context.Context, *TokenWithdrawRequest) (*TokenWithdrawReply, error)
	Trade(context.Context, *WithdrawRequest) (*WithdrawReply, error)
	TradeList(context.Context, *TradeListRequest) (*TradeListReply, error)
	Tran(context.Context, *TranRequest) (*TranReply, error)
	TranList(context.Context, *TranListRequest) (*TranListReply, error)
	UserArea(context.Context, *UserAreaRequest) (*UserAreaReply, error)
	UserInfo(context.Context, *UserInfoRequest) (*UserInfoReply, error)
	Withdraw(context.Context, *WithdrawRequest) (*WithdrawReply, error)
	WithdrawList(context.Context, *WithdrawListRequest) (*WithdrawListReply, error)
}

func RegisterAppHTTPServer(s *http.Server, srv AppHTTPServer) {
	r := s.Route("/")
	r.POST("/api/app_server/eth_authorize", _App_EthAuthorize0_HTTP_Handler(srv))
	r.POST("/api/app_server/recommend_update", _App_RecommendUpdate0_HTTP_Handler(srv))
	r.GET("/api/app_server/user_info", _App_UserInfo0_HTTP_Handler(srv))
	r.GET("/api/app_server/user_area", _App_UserArea0_HTTP_Handler(srv))
	r.GET("/api/app_server/reward_list", _App_RewardList0_HTTP_Handler(srv))
	r.GET("/api/app_server/recommend_reward_list", _App_RecommendRewardList0_HTTP_Handler(srv))
	r.GET("/api/app_server/fee_reward_list", _App_FeeRewardList0_HTTP_Handler(srv))
	r.GET("/api/app_server/withdraw_list", _App_WithdrawList0_HTTP_Handler(srv))
	r.GET("/api/app_server/trade_list", _App_TradeList0_HTTP_Handler(srv))
	r.GET("/api/app_server/tran_list", _App_TranList0_HTTP_Handler(srv))
	r.GET("/api/app_server/recommend_list", _App_RecommendList0_HTTP_Handler(srv))
	r.POST("/api/app_server/password_change", _App_PasswordChange0_HTTP_Handler(srv))
	r.POST("/api/app_server/withdraw", _App_Withdraw0_HTTP_Handler(srv))
	r.POST("/api/app_server/exchange", _App_Exchange0_HTTP_Handler(srv))
	r.POST("/api/app_server/trade", _App_Trade0_HTTP_Handler(srv))
	r.POST("/api/app_server/tran", _App_Tran0_HTTP_Handler(srv))
	r.POST("/api/app_server/get_trade", _App_GetTrade0_HTTP_Handler(srv))
	r.POST("/api/app_server/set_balance_reward", _App_SetBalanceReward0_HTTP_Handler(srv))
	r.POST("/api/app_server/delete_balance_reward", _App_DeleteBalanceReward0_HTTP_Handler(srv))
	r.GET("/api/admin_dhb/deposit", _App_Deposit0_HTTP_Handler(srv))
	r.GET("/api/admin_dhb/withdraw", _App_AdminWithdraw0_HTTP_Handler(srv))
	r.GET("/api/admin_dhb/withdraw_eth", _App_AdminWithdrawEth0_HTTP_Handler(srv))
	r.GET("/api/admin_dhb/fee", _App_AdminFee0_HTTP_Handler(srv))
	r.GET("/api/app_server/token_withdraw", _App_TokenWithdraw0_HTTP_Handler(srv))
}

func _App_EthAuthorize0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in EthAuthorizeRequest
		if err := ctx.Bind(&in.SendBody); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppEthAuthorize)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.EthAuthorize(ctx, req.(*EthAuthorizeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*EthAuthorizeReply)
		return ctx.Result(200, reply)
	}
}

func _App_RecommendUpdate0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RecommendUpdateRequest
		if err := ctx.Bind(&in.SendBody); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppRecommendUpdate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RecommendUpdate(ctx, req.(*RecommendUpdateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RecommendUpdateReply)
		return ctx.Result(200, reply)
	}
}

func _App_UserInfo0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserInfoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppUserInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserInfo(ctx, req.(*UserInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserInfoReply)
		return ctx.Result(200, reply)
	}
}

func _App_UserArea0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserAreaRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppUserArea)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserArea(ctx, req.(*UserAreaRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserAreaReply)
		return ctx.Result(200, reply)
	}
}

func _App_RewardList0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RewardListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppRewardList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RewardList(ctx, req.(*RewardListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RewardListReply)
		return ctx.Result(200, reply)
	}
}

func _App_RecommendRewardList0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RecommendRewardListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppRecommendRewardList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RecommendRewardList(ctx, req.(*RecommendRewardListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RecommendRewardListReply)
		return ctx.Result(200, reply)
	}
}

func _App_FeeRewardList0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FeeRewardListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppFeeRewardList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FeeRewardList(ctx, req.(*FeeRewardListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FeeRewardListReply)
		return ctx.Result(200, reply)
	}
}

func _App_WithdrawList0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in WithdrawListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppWithdrawList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.WithdrawList(ctx, req.(*WithdrawListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*WithdrawListReply)
		return ctx.Result(200, reply)
	}
}

func _App_TradeList0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TradeListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppTradeList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.TradeList(ctx, req.(*TradeListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TradeListReply)
		return ctx.Result(200, reply)
	}
}

func _App_TranList0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TranListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppTranList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.TranList(ctx, req.(*TranListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TranListReply)
		return ctx.Result(200, reply)
	}
}

func _App_RecommendList0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RecommendListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppRecommendList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RecommendList(ctx, req.(*RecommendListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RecommendListReply)
		return ctx.Result(200, reply)
	}
}

func _App_PasswordChange0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PasswordChangeRequest
		if err := ctx.Bind(&in.SendBody); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppPasswordChange)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PasswordChange(ctx, req.(*PasswordChangeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PasswordChangeReply)
		return ctx.Result(200, reply)
	}
}

func _App_Withdraw0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in WithdrawRequest
		if err := ctx.Bind(&in.SendBody); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppWithdraw)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Withdraw(ctx, req.(*WithdrawRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*WithdrawReply)
		return ctx.Result(200, reply)
	}
}

func _App_Exchange0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ExchangeRequest
		if err := ctx.Bind(&in.SendBody); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppExchange)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Exchange(ctx, req.(*ExchangeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ExchangeReply)
		return ctx.Result(200, reply)
	}
}

func _App_Trade0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in WithdrawRequest
		if err := ctx.Bind(&in.SendBody); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppTrade)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Trade(ctx, req.(*WithdrawRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*WithdrawReply)
		return ctx.Result(200, reply)
	}
}

func _App_Tran0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TranRequest
		if err := ctx.Bind(&in.SendBody); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppTran)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Tran(ctx, req.(*TranRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TranReply)
		return ctx.Result(200, reply)
	}
}

func _App_GetTrade0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTradeRequest
		if err := ctx.Bind(&in.SendBody); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppGetTrade)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTrade(ctx, req.(*GetTradeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetTradeReply)
		return ctx.Result(200, reply)
	}
}

func _App_SetBalanceReward0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetBalanceRewardRequest
		if err := ctx.Bind(&in.SendBody); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppSetBalanceReward)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetBalanceReward(ctx, req.(*SetBalanceRewardRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SetBalanceRewardReply)
		return ctx.Result(200, reply)
	}
}

func _App_DeleteBalanceReward0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteBalanceRewardRequest
		if err := ctx.Bind(&in.SendBody); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppDeleteBalanceReward)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteBalanceReward(ctx, req.(*DeleteBalanceRewardRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteBalanceRewardReply)
		return ctx.Result(200, reply)
	}
}

func _App_Deposit0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DepositRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppDeposit)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Deposit(ctx, req.(*DepositRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DepositReply)
		return ctx.Result(200, reply)
	}
}

func _App_AdminWithdraw0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AdminWithdrawRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppAdminWithdraw)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AdminWithdraw(ctx, req.(*AdminWithdrawRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AdminWithdrawReply)
		return ctx.Result(200, reply)
	}
}

func _App_AdminWithdrawEth0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AdminWithdrawEthRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppAdminWithdrawEth)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AdminWithdrawEth(ctx, req.(*AdminWithdrawEthRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AdminWithdrawEthReply)
		return ctx.Result(200, reply)
	}
}

func _App_AdminFee0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AdminFeeRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppAdminFee)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AdminFee(ctx, req.(*AdminFeeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AdminFeeReply)
		return ctx.Result(200, reply)
	}
}

func _App_TokenWithdraw0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TokenWithdrawRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAppTokenWithdraw)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.TokenWithdraw(ctx, req.(*TokenWithdrawRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TokenWithdrawReply)
		return ctx.Result(200, reply)
	}
}

type AppHTTPClient interface {
	AdminFee(ctx context.Context, req *AdminFeeRequest, opts ...http.CallOption) (rsp *AdminFeeReply, err error)
	AdminWithdraw(ctx context.Context, req *AdminWithdrawRequest, opts ...http.CallOption) (rsp *AdminWithdrawReply, err error)
	AdminWithdrawEth(ctx context.Context, req *AdminWithdrawEthRequest, opts ...http.CallOption) (rsp *AdminWithdrawEthReply, err error)
	DeleteBalanceReward(ctx context.Context, req *DeleteBalanceRewardRequest, opts ...http.CallOption) (rsp *DeleteBalanceRewardReply, err error)
	Deposit(ctx context.Context, req *DepositRequest, opts ...http.CallOption) (rsp *DepositReply, err error)
	EthAuthorize(ctx context.Context, req *EthAuthorizeRequest, opts ...http.CallOption) (rsp *EthAuthorizeReply, err error)
	Exchange(ctx context.Context, req *ExchangeRequest, opts ...http.CallOption) (rsp *ExchangeReply, err error)
	FeeRewardList(ctx context.Context, req *FeeRewardListRequest, opts ...http.CallOption) (rsp *FeeRewardListReply, err error)
	GetTrade(ctx context.Context, req *GetTradeRequest, opts ...http.CallOption) (rsp *GetTradeReply, err error)
	PasswordChange(ctx context.Context, req *PasswordChangeRequest, opts ...http.CallOption) (rsp *PasswordChangeReply, err error)
	RecommendList(ctx context.Context, req *RecommendListRequest, opts ...http.CallOption) (rsp *RecommendListReply, err error)
	RecommendRewardList(ctx context.Context, req *RecommendRewardListRequest, opts ...http.CallOption) (rsp *RecommendRewardListReply, err error)
	RecommendUpdate(ctx context.Context, req *RecommendUpdateRequest, opts ...http.CallOption) (rsp *RecommendUpdateReply, err error)
	RewardList(ctx context.Context, req *RewardListRequest, opts ...http.CallOption) (rsp *RewardListReply, err error)
	SetBalanceReward(ctx context.Context, req *SetBalanceRewardRequest, opts ...http.CallOption) (rsp *SetBalanceRewardReply, err error)
	TokenWithdraw(ctx context.Context, req *TokenWithdrawRequest, opts ...http.CallOption) (rsp *TokenWithdrawReply, err error)
	Trade(ctx context.Context, req *WithdrawRequest, opts ...http.CallOption) (rsp *WithdrawReply, err error)
	TradeList(ctx context.Context, req *TradeListRequest, opts ...http.CallOption) (rsp *TradeListReply, err error)
	Tran(ctx context.Context, req *TranRequest, opts ...http.CallOption) (rsp *TranReply, err error)
	TranList(ctx context.Context, req *TranListRequest, opts ...http.CallOption) (rsp *TranListReply, err error)
	UserArea(ctx context.Context, req *UserAreaRequest, opts ...http.CallOption) (rsp *UserAreaReply, err error)
	UserInfo(ctx context.Context, req *UserInfoRequest, opts ...http.CallOption) (rsp *UserInfoReply, err error)
	Withdraw(ctx context.Context, req *WithdrawRequest, opts ...http.CallOption) (rsp *WithdrawReply, err error)
	WithdrawList(ctx context.Context, req *WithdrawListRequest, opts ...http.CallOption) (rsp *WithdrawListReply, err error)
}

type AppHTTPClientImpl struct {
	cc *http.Client
}

func NewAppHTTPClient(client *http.Client) AppHTTPClient {
	return &AppHTTPClientImpl{client}
}

func (c *AppHTTPClientImpl) AdminFee(ctx context.Context, in *AdminFeeRequest, opts ...http.CallOption) (*AdminFeeReply, error) {
	var out AdminFeeReply
	pattern := "/api/admin_dhb/fee"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppAdminFee))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) AdminWithdraw(ctx context.Context, in *AdminWithdrawRequest, opts ...http.CallOption) (*AdminWithdrawReply, error) {
	var out AdminWithdrawReply
	pattern := "/api/admin_dhb/withdraw"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppAdminWithdraw))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) AdminWithdrawEth(ctx context.Context, in *AdminWithdrawEthRequest, opts ...http.CallOption) (*AdminWithdrawEthReply, error) {
	var out AdminWithdrawEthReply
	pattern := "/api/admin_dhb/withdraw_eth"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppAdminWithdrawEth))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) DeleteBalanceReward(ctx context.Context, in *DeleteBalanceRewardRequest, opts ...http.CallOption) (*DeleteBalanceRewardReply, error) {
	var out DeleteBalanceRewardReply
	pattern := "/api/app_server/delete_balance_reward"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAppDeleteBalanceReward))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.SendBody, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) Deposit(ctx context.Context, in *DepositRequest, opts ...http.CallOption) (*DepositReply, error) {
	var out DepositReply
	pattern := "/api/admin_dhb/deposit"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppDeposit))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) EthAuthorize(ctx context.Context, in *EthAuthorizeRequest, opts ...http.CallOption) (*EthAuthorizeReply, error) {
	var out EthAuthorizeReply
	pattern := "/api/app_server/eth_authorize"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAppEthAuthorize))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.SendBody, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) Exchange(ctx context.Context, in *ExchangeRequest, opts ...http.CallOption) (*ExchangeReply, error) {
	var out ExchangeReply
	pattern := "/api/app_server/exchange"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAppExchange))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.SendBody, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) FeeRewardList(ctx context.Context, in *FeeRewardListRequest, opts ...http.CallOption) (*FeeRewardListReply, error) {
	var out FeeRewardListReply
	pattern := "/api/app_server/fee_reward_list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppFeeRewardList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetTrade(ctx context.Context, in *GetTradeRequest, opts ...http.CallOption) (*GetTradeReply, error) {
	var out GetTradeReply
	pattern := "/api/app_server/get_trade"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAppGetTrade))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.SendBody, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) PasswordChange(ctx context.Context, in *PasswordChangeRequest, opts ...http.CallOption) (*PasswordChangeReply, error) {
	var out PasswordChangeReply
	pattern := "/api/app_server/password_change"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAppPasswordChange))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.SendBody, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) RecommendList(ctx context.Context, in *RecommendListRequest, opts ...http.CallOption) (*RecommendListReply, error) {
	var out RecommendListReply
	pattern := "/api/app_server/recommend_list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppRecommendList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) RecommendRewardList(ctx context.Context, in *RecommendRewardListRequest, opts ...http.CallOption) (*RecommendRewardListReply, error) {
	var out RecommendRewardListReply
	pattern := "/api/app_server/recommend_reward_list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppRecommendRewardList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) RecommendUpdate(ctx context.Context, in *RecommendUpdateRequest, opts ...http.CallOption) (*RecommendUpdateReply, error) {
	var out RecommendUpdateReply
	pattern := "/api/app_server/recommend_update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAppRecommendUpdate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.SendBody, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) RewardList(ctx context.Context, in *RewardListRequest, opts ...http.CallOption) (*RewardListReply, error) {
	var out RewardListReply
	pattern := "/api/app_server/reward_list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppRewardList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) SetBalanceReward(ctx context.Context, in *SetBalanceRewardRequest, opts ...http.CallOption) (*SetBalanceRewardReply, error) {
	var out SetBalanceRewardReply
	pattern := "/api/app_server/set_balance_reward"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAppSetBalanceReward))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.SendBody, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) TokenWithdraw(ctx context.Context, in *TokenWithdrawRequest, opts ...http.CallOption) (*TokenWithdrawReply, error) {
	var out TokenWithdrawReply
	pattern := "/api/app_server/token_withdraw"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppTokenWithdraw))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) Trade(ctx context.Context, in *WithdrawRequest, opts ...http.CallOption) (*WithdrawReply, error) {
	var out WithdrawReply
	pattern := "/api/app_server/trade"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAppTrade))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.SendBody, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) TradeList(ctx context.Context, in *TradeListRequest, opts ...http.CallOption) (*TradeListReply, error) {
	var out TradeListReply
	pattern := "/api/app_server/trade_list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppTradeList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) Tran(ctx context.Context, in *TranRequest, opts ...http.CallOption) (*TranReply, error) {
	var out TranReply
	pattern := "/api/app_server/tran"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAppTran))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.SendBody, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) TranList(ctx context.Context, in *TranListRequest, opts ...http.CallOption) (*TranListReply, error) {
	var out TranListReply
	pattern := "/api/app_server/tran_list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppTranList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) UserArea(ctx context.Context, in *UserAreaRequest, opts ...http.CallOption) (*UserAreaReply, error) {
	var out UserAreaReply
	pattern := "/api/app_server/user_area"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppUserArea))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...http.CallOption) (*UserInfoReply, error) {
	var out UserInfoReply
	pattern := "/api/app_server/user_info"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppUserInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) Withdraw(ctx context.Context, in *WithdrawRequest, opts ...http.CallOption) (*WithdrawReply, error) {
	var out WithdrawReply
	pattern := "/api/app_server/withdraw"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAppWithdraw))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.SendBody, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) WithdrawList(ctx context.Context, in *WithdrawListRequest, opts ...http.CallOption) (*WithdrawListReply, error) {
	var out WithdrawListReply
	pattern := "/api/app_server/withdraw_list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAppWithdrawList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
