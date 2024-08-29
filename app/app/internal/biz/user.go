package biz

import (
	"context"
	"crypto/ecdsa"
	v1 "dhb/app/app/api"
	"encoding/base64"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"strconv"
	"strings"
	"time"
)

type User struct {
	ID         int64
	Address    string
	AddressTwo string
	PrivateKey string
	Password   string
	Undo       int64
	Amount     uint64
	Total      uint64
	Kkdt       int64
	Uudt       int64
	Last       uint64
	TotalA     int64
	TotalB     int64
	TotalC     int64
	TotalD     int64
	TotalF     int64
	CreatedAt  time.Time
}

type UserInfo struct {
	ID               int64
	UserId           int64
	Vip              int64
	UseVip           int64
	HistoryRecommend int64
	TeamCsdBalance   int64
}

type UserRecommend struct {
	ID            int64
	UserId        int64
	RecommendCode string
	CreatedAt     time.Time
}

type UserRecommendArea struct {
	ID            int64
	RecommendCode string
	Num           int64
	CreatedAt     time.Time
}

type Trade struct {
	ID           int64
	UserId       int64
	AmountCsd    int64
	RelAmountCsd int64
	AmountHbs    int64
	RelAmountHbs int64
	Status       string
	CreatedAt    time.Time
}

type UserArea struct {
	ID         int64
	UserId     int64
	Amount     int64
	SelfAmount int64
	Level      int64
}

type UserCurrentMonthRecommend struct {
	ID              int64
	UserId          int64
	RecommendUserId int64
	Date            time.Time
}

type Config struct {
	ID      int64
	KeyName string
	Name    string
	Value   string
}

type UserBalance struct {
	ID             int64
	UserId         int64
	BalanceUsdt    int64
	BalanceUsdt2   int64
	BalanceDhb     float64
	RecommendTotal int64
	AreaTotal      int64
	FourTotal      int64
	LocationTotal  int64
}

type Withdraw struct {
	ID              int64
	UserId          int64
	Amount          int64
	RelAmount       int64
	BalanceRecordId int64
	Status          string
	Type            string
	AmountNew       float64
	AmountNewRel    float64
	CreatedAt       time.Time
}

type UserSortRecommendReward struct {
	UserId int64
	Total  int64
}

type UserUseCase struct {
	repo                          UserRepo
	urRepo                        UserRecommendRepo
	configRepo                    ConfigRepo
	uiRepo                        UserInfoRepo
	ubRepo                        UserBalanceRepo
	locationRepo                  LocationRepo
	userCurrentMonthRecommendRepo UserCurrentMonthRecommendRepo
	tx                            Transaction
	log                           *log.Helper
}

type LocationNew struct {
	ID                int64
	UserId            int64
	Num               int64
	Status            string
	Current           int64
	CurrentMax        int64
	StopLocationAgain int64
	StopCoin          int64
	CurrentMaxNew     int64
	Term              int64
	Usdt              int64
	Biw               int64
	Total             int64
	TotalTwo          int64
	TotalThree        int64
	LastLevel         int64
	StopDate          time.Time
	CreatedAt         time.Time
}

type UserBalanceRecord struct {
	ID        int64
	UserId    int64
	Amount    int64
	CoinType  string
	Balance   int64
	Type      string
	CreatedAt time.Time
}

type BalanceReward struct {
	ID        int64
	UserId    int64
	Status    int64
	Amount    int64
	SetDate   time.Time
	UpdatedAt time.Time
	CreatedAt time.Time
}

type Reward struct {
	ID               int64
	UserId           int64
	Amount           int64
	AmountB          int64
	BalanceRecordId  int64
	Type             string
	TypeRecordId     int64
	Reason           string
	ReasonLocationId int64
	LocationType     string
	Address          string
	AmountNew        float64
	CreatedAt        time.Time
}

type Pagination struct {
	PageNum  int
	PageSize int
}

type ConfigRepo interface {
	GetConfigByKeys(ctx context.Context, keys ...string) ([]*Config, error)
	GetConfigs(ctx context.Context) ([]*Config, error)
	UpdateConfig(ctx context.Context, id int64, value string) (bool, error)
}

type UserBalanceRepo interface {
	CreateUserBalance(ctx context.Context, u *User) error
	CreateUserBalanceLock(ctx context.Context, u *User) (*UserBalance, error)
	LocationReward(ctx context.Context, userId int64, amount int64, locationId int64, myLocationId int64, locationType string) (int64, error)
	WithdrawReward(ctx context.Context, userId int64, amount int64, locationId int64, myLocationId int64, locationType string) (int64, error)
	RecommendReward(ctx context.Context, userId int64, amount int64, locationId int64) (int64, error)
	SystemWithdrawReward(ctx context.Context, amount int64, locationId int64) error
	SystemReward(ctx context.Context, amount int64, locationId int64) error
	SystemFee(ctx context.Context, amount int64, locationId int64) error
	GetSystemYesterdayDailyReward(ctx context.Context) (*Reward, error)
	UserFee(ctx context.Context, userId int64, amount int64) (int64, error)
	RecommendWithdrawReward(ctx context.Context, userId int64, amount int64, locationId int64) (int64, error)
	NormalRecommendReward(ctx context.Context, userId int64, amount int64, locationId int64) (int64, error)
	NormalWithdrawRecommendReward(ctx context.Context, userId int64, amount int64, locationId int64) (int64, error)
	Deposit(ctx context.Context, userId int64, amount int64) (int64, error)
	DepositLast(ctx context.Context, userId int64, lastAmount int64, locationId int64) (int64, error)
	DepositDhb(ctx context.Context, userId int64, amount int64) (int64, error)
	GetUserBalance(ctx context.Context, userId int64) (*UserBalance, error)
	GetUserRewardByUserId(ctx context.Context, userId int64) ([]*Reward, error)
	GetUserRewardTodayBuy(ctx context.Context) ([]*Reward, error)
	GetLocationsToday(ctx context.Context) ([]*LocationNew, error)
	GetUserRewardByUserIds(ctx context.Context, userIds ...int64) (map[int64]*UserSortRecommendReward, error)
	GetUserRewards(ctx context.Context, b *Pagination, userId int64) ([]*Reward, error, int64)
	GetUserRewardsLastMonthFee(ctx context.Context) ([]*Reward, error)
	GetUserBalanceByUserIds(ctx context.Context, userIds ...int64) (map[int64]*UserBalance, error)
	GetUserBalanceUsdtTotal(ctx context.Context) (int64, error)
	GreateWithdraw(ctx context.Context, userId int64, relAmount float64, amount float64, coinType string) (*Withdraw, error)
	WithdrawUsdt(ctx context.Context, userId int64, amount int64, tmpRecommendUserIdsInt []int64) error
	WithdrawUsdt2(ctx context.Context, userId int64, amount float64) error
	Exchange(ctx context.Context, userId int64, amount int64, amountUsdtSubFee int64, amountUsdt int64, locationId int64) error
	WithdrawUsdt3(ctx context.Context, userId int64, amount int64) error
	TranUsdt(ctx context.Context, userId int64, toUserId int64, amount int64, tmpRecommendUserIdsInt []int64, tmpRecommendUserIdsInt2 []int64) error
	WithdrawDhb(ctx context.Context, userId int64, amount int64) error
	TranDhb(ctx context.Context, userId int64, toUserId int64, amount int64) error
	GetWithdrawByUserId(ctx context.Context, userId int64, typeCoin string) ([]*Withdraw, error)
	GetWithdrawByUserId2(ctx context.Context, userId int64) ([]*Withdraw, error)
	GetUserBalanceRecordByUserId(ctx context.Context, userId int64, typeCoin string, tran string) ([]*UserBalanceRecord, error)
	GetUserBalanceRecordsByUserId(ctx context.Context, userId int64) ([]*UserBalanceRecord, error)
	GetTradeByUserId(ctx context.Context, userId int64) ([]*Trade, error)
	GetWithdraws(ctx context.Context, b *Pagination, userId int64) ([]*Withdraw, error, int64)
	GetWithdrawPassOrRewarded(ctx context.Context) ([]*Withdraw, error)
	UpdateWithdraw(ctx context.Context, id int64, status string) (*Withdraw, error)
	GetWithdrawById(ctx context.Context, id int64) (*Withdraw, error)
	GetWithdrawNotDeal(ctx context.Context) ([]*Withdraw, error)
	GetUserBalanceRecordUserUsdtTotal(ctx context.Context, userId int64) (int64, error)
	GetUserBalanceRecordUsdtTotal(ctx context.Context) (int64, error)
	GetUserBalanceRecordUsdtTotalToday(ctx context.Context) (int64, error)
	GetUserWithdrawUsdtTotalToday(ctx context.Context) (int64, error)
	GetUserWithdrawUsdtTotal(ctx context.Context) (int64, error)
	GetUserRewardUsdtTotal(ctx context.Context) (int64, error)
	GetSystemRewardUsdtTotal(ctx context.Context) (int64, error)
	UpdateWithdrawAmount(ctx context.Context, id int64, status string, amount int64) (*Withdraw, error)
	GetUserRewardRecommendSort(ctx context.Context) ([]*UserSortRecommendReward, error)
	GetUserRewardTodayTotalByUserId(ctx context.Context, userId int64) (*UserSortRecommendReward, error)

	SetBalanceReward(ctx context.Context, userId int64, amount int64) error
	UpdateBalanceReward(ctx context.Context, userId int64, id int64, amount int64, status int64) error
	GetBalanceRewardByUserId(ctx context.Context, userId int64) ([]*BalanceReward, error)

	GetUserBalanceLock(ctx context.Context, userId int64) (*UserBalance, error)
	Trade(ctx context.Context, userId int64, amount int64, amountB int64, amountRel int64, amountBRel int64, tmpRecommendUserIdsInt []int64, amount2 int64) error
}

type UserRecommendRepo interface {
	GetUserRecommendByUserId(ctx context.Context, userId int64) (*UserRecommend, error)
	GetUserRecommendsFour(ctx context.Context) ([]*UserRecommend, error)
	CreateUserRecommend(ctx context.Context, u *User, recommendUser *UserRecommend) (*UserRecommend, error)
	UpdateUserRecommend(ctx context.Context, u *User, recommendUser *UserRecommend) (bool, error)
	GetUserRecommendByCode(ctx context.Context, code string) ([]*UserRecommend, error)
	GetUserRecommendLikeCode(ctx context.Context, code string) ([]*UserRecommend, error)
	CreateUserRecommendArea(ctx context.Context, u *User, recommendUser *UserRecommend) (bool, error)
	DeleteOrOriginUserRecommendArea(ctx context.Context, code string, originCode string) (bool, error)
	GetUserRecommendLowArea(ctx context.Context, code string) ([]*UserRecommendArea, error)
	GetUserAreas(ctx context.Context, userIds []int64) ([]*UserArea, error)
	CreateUserArea(ctx context.Context, u *User) (bool, error)
	GetUserArea(ctx context.Context, userId int64) (*UserArea, error)
}

type UserCurrentMonthRecommendRepo interface {
	GetUserCurrentMonthRecommendByUserId(ctx context.Context, userId int64) ([]*UserCurrentMonthRecommend, error)
	GetUserCurrentMonthRecommendGroupByUserId(ctx context.Context, b *Pagination, userId int64) ([]*UserCurrentMonthRecommend, error, int64)
	CreateUserCurrentMonthRecommend(ctx context.Context, u *UserCurrentMonthRecommend) (*UserCurrentMonthRecommend, error)
	GetUserCurrentMonthRecommendCountByUserIds(ctx context.Context, userIds ...int64) (map[int64]int64, error)
	GetUserLastMonthRecommend(ctx context.Context) ([]int64, error)
}

type UserInfoRepo interface {
	CreateUserInfo(ctx context.Context, u *User) (*UserInfo, error)
	GetUserInfoByUserId(ctx context.Context, userId int64) (*UserInfo, error)
	UpdateUserInfo(ctx context.Context, u *UserInfo) (*UserInfo, error)
	GetUserInfoByUserIds(ctx context.Context, userIds ...int64) (map[int64]*UserInfo, error)
}

type UserRepo interface {
	GetUserById(ctx context.Context, Id int64) (*User, error)
	GetUserByAddresses(ctx context.Context, Addresses ...string) (map[string]*User, error)
	GetUserByAddress(ctx context.Context, address string) (*User, error)
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByUserIds(ctx context.Context, userIds ...int64) (map[int64]*User, error)
	GetUsers(ctx context.Context, b *Pagination, address string) ([]*User, error, int64)
	GetUserCount(ctx context.Context) (int64, error)
	GetUserCountToday(ctx context.Context) (int64, error)
	UpdateUserNewTwoNew(ctx context.Context, userId int64, amount uint64, strUpdate string, uudt int64, kkdt int64) error
	InRecordNew(ctx context.Context, userId int64, address string, amount int64, originTotal int64) error
	GetEthUserRecordListByUserId(ctx context.Context, userId int64) ([]*EthUserRecord, error)
	GetUsersNew(ctx context.Context) ([]*User, error)
}

func NewUserUseCase(repo UserRepo, tx Transaction, configRepo ConfigRepo, uiRepo UserInfoRepo, urRepo UserRecommendRepo, locationRepo LocationRepo, userCurrentMonthRecommendRepo UserCurrentMonthRecommendRepo, ubRepo UserBalanceRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo:                          repo,
		tx:                            tx,
		configRepo:                    configRepo,
		locationRepo:                  locationRepo,
		userCurrentMonthRecommendRepo: userCurrentMonthRecommendRepo,
		uiRepo:                        uiRepo,
		urRepo:                        urRepo,
		ubRepo:                        ubRepo,
		log:                           log.NewHelper(logger),
	}
}

func (uuc *UserUseCase) GetUserByAddress(ctx context.Context, Addresses ...string) (map[string]*User, error) {
	return uuc.repo.GetUserByAddresses(ctx, Addresses...)
}

func (uuc *UserUseCase) GetDhbConfig(ctx context.Context) ([]*Config, error) {
	return uuc.configRepo.GetConfigByKeys(ctx, "level1Dhb", "level2Dhb", "level3Dhb")
}

func (uuc *UserUseCase) GetExistUserByAddressOrCreate(ctx context.Context, u *User, req *v1.EthAuthorizeRequest) (*User, error) {
	var (
		user          *User
		recommendUser *UserRecommend
		err           error
		userId        int64
		decodeBytes   []byte
	)

	user, err = uuc.repo.GetUserByAddress(ctx, u.Address) // 查询用户
	if nil == user || nil != err {
		code := req.SendBody.Code // 查询推荐码 abf00dd52c08a9213f225827bc3fb100 md5 dhbmachinefirst
		if "abf00dd52c08a9213f225827bc3fb100" != code {
			decodeBytes, err = base64.StdEncoding.DecodeString(code)
			code = string(decodeBytes)
			if 1 >= len(code) {
				return nil, errors.New(500, "USER_ERROR", "无效的推荐码1")
			}
			if userId, err = strconv.ParseInt(code[1:], 10, 64); 0 >= userId || nil != err {
				return nil, errors.New(500, "USER_ERROR", "无效的推荐码2")
			}

			var (
				locationNew *User
			)
			locationNew, err = uuc.repo.GetUserById(ctx, userId)
			if nil != err || nil == locationNew {
				return nil, errors.New(500, "USER_ERROR", "无效的推荐码3")
			}

			//if 0 >= locationNew.Total {
			//	return nil, errors.New(500, "USER_ERROR", "推荐人未入金，无效的推荐码")
			//}

			// 查询推荐人的相关信息
			recommendUser, err = uuc.urRepo.GetUserRecommendByUserId(ctx, userId)
			if err != nil {
				return nil, errors.New(500, "USER_ERROR", "无效的推荐码4")
			}
		}

		// 创建私钥
		var (
			address    string
			privateKey string
		)
		address, privateKey, err = generateKey()
		if 0 >= len(address) || 0 >= len(privateKey) || err != nil {
			return nil, errors.New(500, "USER_ERROR", "生成地址错误")
		}

		u.PrivateKey = privateKey
		u.AddressTwo = address

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			user, err = uuc.repo.CreateUser(ctx, u) // 用户创建
			if err != nil {
				return err
			}

			_, err = uuc.uiRepo.CreateUserInfo(ctx, user) // 创建用户信息
			if err != nil {
				return err
			}

			_, err = uuc.urRepo.CreateUserRecommend(ctx, user, recommendUser) // 创建用户推荐信息
			if err != nil {
				return err
			}

			err = uuc.ubRepo.CreateUserBalance(ctx, user) // 创建余额信息
			if err != nil {
				return err
			}

			return nil
		}); err != nil {
			return nil, err
		}
	}

	return user, nil
}

func generateKey() (string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	//fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", "", nil
	}

	//publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	//fmt.Println(hexutil.Encode(publicKeyBytes)[4:])

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	//fmt.Println(address)

	return address, hexutil.Encode(privateKeyBytes)[2:], nil
}

func (uuc *UserUseCase) UpdateUserRecommend(ctx context.Context, u *User, req *v1.RecommendUpdateRequest) (*v1.RecommendUpdateReply, error) {
	var (
		err                   error
		userId                int64
		recommendUser         *UserRecommend
		userRecommend         *UserRecommend
		myRecommendUser       *User
		myUserRecommendUserId int64
		Address               string
		decodeBytes           []byte
	)

	code := req.SendBody.Code // 查询推荐码 abf00dd52c08a9213f225827bc3fb100 md5 dhbmachinefirst
	if "abf00dd52c08a9213f225827bc3fb100" != code {
		decodeBytes, err = base64.StdEncoding.DecodeString(code)
		code = string(decodeBytes)
		if 1 >= len(code) {
			return nil, errors.New(500, "USER_ERROR", "无效的推荐码")
		}
		if userId, err = strconv.ParseInt(code[1:], 10, 64); 0 >= userId || nil != err {
			return nil, errors.New(500, "USER_ERROR", "无效的推荐码")
		}

		// 现有推荐人信息，判断推荐人是否改变
		userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, u.ID)
		if nil == userRecommend {
			return nil, err
		}
		if "" != userRecommend.RecommendCode {
			tmpRecommendUserIds := strings.Split(userRecommend.RecommendCode, "D")
			if 2 <= len(tmpRecommendUserIds) {
				myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
			}
			myRecommendUser, err = uuc.repo.GetUserById(ctx, myUserRecommendUserId)
			if nil != err {
				return nil, err
			}
		}

		if nil == myRecommendUser {
			return &v1.RecommendUpdateReply{InviteUserAddress: ""}, nil
		}

		if myRecommendUser.ID == userId {
			return &v1.RecommendUpdateReply{InviteUserAddress: myRecommendUser.Address}, nil
		}

		if u.ID == userId {
			return &v1.RecommendUpdateReply{InviteUserAddress: myRecommendUser.Address}, nil
		}

		// 我的占位信息
		var (
			user *User
		)
		user, err = uuc.repo.GetUserById(ctx, myUserRecommendUserId)
		if nil != err {
			return nil, err
		}
		if user.Total > 0 {
			return &v1.RecommendUpdateReply{InviteUserAddress: myRecommendUser.Address}, nil
		}

		// 查询推荐人的相关信息
		recommendUser, err = uuc.urRepo.GetUserRecommendByUserId(ctx, userId)
		if err != nil {
			return nil, errors.New(500, "USER_ERROR", "无效的推荐码")
		}

		// 推荐人信息
		myRecommendUser, err = uuc.repo.GetUserById(ctx, userId)
		if err != nil {
			return nil, err
		}

		// 更新
		_, err = uuc.urRepo.UpdateUserRecommend(ctx, u, recommendUser)
		if err != nil {
			return nil, err
		}
		Address = myRecommendUser.Address
	}

	return &v1.RecommendUpdateReply{InviteUserAddress: Address}, err
}

func (uuc *UserUseCase) UserInfo(ctx context.Context, user *User) (*v1.UserInfoReply, error) {
	var (
		err                   error
		myUser                *User
		userRecommend         *UserRecommend
		myCode                string
		encodeString          string
		myUserRecommendUserId int64
		inviteUserAddress     string
		myRecommendUser       *User
		configs               []*Config
		userBalance           *UserBalance
		buyLimit              int64
		withdrawMin           float64
		one                   float64
		two                   float64
		three                 float64
		four                  float64
		five                  float64
		six                   float64
		oneTwo                float64
		twoTwo                float64
		threeTwo              float64
		fiveThree             float64
		fourThree             float64
	)

	// 配置
	configs, err = uuc.configRepo.GetConfigByKeys(ctx,
		"withdraw_min", "one", "two", "three", "four", "five", "six",
		"one_two", "two_two", "three_two", "four_two", "five_two", "four_three", "five_three",
	)
	if nil != configs {
		for _, vConfig := range configs {
			if "withdraw_min" == vConfig.KeyName {
				withdrawMin, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "one" == vConfig.KeyName {
				one, err = strconv.ParseFloat(vConfig.Value, 10)
				if nil != err {
					return nil, err
				}
			} else if "two" == vConfig.KeyName {
				two, err = strconv.ParseFloat(vConfig.Value, 10)
				if nil != err {
					return nil, err
				}
			} else if "three" == vConfig.KeyName {
				three, err = strconv.ParseFloat(vConfig.Value, 10)
				if nil != err {
					return nil, err
				}
			} else if "four" == vConfig.KeyName {
				four, err = strconv.ParseFloat(vConfig.Value, 10)
				if nil != err {
					return nil, err
				}
			} else if "five" == vConfig.KeyName {
				five, err = strconv.ParseFloat(vConfig.Value, 10)
				if nil != err {
					return nil, err
				}
			} else if "six" == vConfig.KeyName {
				six, err = strconv.ParseFloat(vConfig.Value, 10)
				if nil != err {
					return nil, err
				}
			} else if "one_two" == vConfig.KeyName {
				oneTwo, err = strconv.ParseFloat(vConfig.Value, 10)
				if nil != err {
					return nil, err
				}
			} else if "two_two" == vConfig.KeyName {
				twoTwo, err = strconv.ParseFloat(vConfig.Value, 10)
				if nil != err {
					return nil, err
				}
			} else if "three_two" == vConfig.KeyName {
				threeTwo, err = strconv.ParseFloat(vConfig.Value, 10)
				if nil != err {
					return nil, err
				}
			} else if "four_three" == vConfig.KeyName {
				fourThree, err = strconv.ParseFloat(vConfig.Value, 10)
				if nil != err {
					return nil, err
				}
			} else if "five_three" == vConfig.KeyName {
				fiveThree, err = strconv.ParseFloat(vConfig.Value, 10)
				if nil != err {
					return nil, err
				}
			}
		}
	}

	var (
		users []*User
	)
	users, err = uuc.repo.GetUsersNew(ctx)
	if nil != err {
		fmt.Println("分红", err)
		return nil, nil
	}

	var (
		amountSecond float64 // 节点超级节点奖励额度
		amounts      float64
	)
	for _, vUsers := range users {
		// 超级节点，节点
		number := vUsers.Total
		if 15000 <= number {
			for number > 0 {
				if 1000000 <= number {
					number -= 1000000
					amounts += six
				} else if 30000 <= number {
					number -= 30000
					amounts += five
				} else if 15000 <= number {
					number -= 15000
					amounts += four
				} else if 5000 <= number {
					number -= 5000
					amounts += three
					amountSecond += threeTwo
				} else if 3000 <= number {
					number -= 3000
					amounts += two
					amountSecond += twoTwo
				} else if 1000 <= number {
					number -= 1000
					amounts += one
					amountSecond += oneTwo
				} else {
					break
				}
			}
		} else {
			// 普通
			for i := 0; i < int(vUsers.TotalA); i++ {
				amounts += one
				amountSecond += oneTwo
			}
			for i := 0; i < int(vUsers.TotalB); i++ {
				amounts += two
				amountSecond += twoTwo
			}
			for i := 0; i < int(vUsers.TotalC); i++ {
				amounts += three
				amountSecond += threeTwo
			}
		}
	}

	fourAmounts := amountSecond * fourThree / 100
	fiveAmounts := amountSecond * fiveThree / 100

	myUser, err = uuc.repo.GetUserById(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	// 余额，收益总数
	userBalance, err = uuc.ubRepo.GetUserBalance(ctx, myUser.ID)
	if nil != err {
		return nil, err
	}

	// 推荐
	userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, myUser.ID)
	if nil == userRecommend {
		return nil, err
	}

	myCode = "D" + strconv.FormatInt(myUser.ID, 10)
	codeByte := []byte(myCode)
	encodeString = base64.StdEncoding.EncodeToString(codeByte)
	if "" != userRecommend.RecommendCode {
		tmpRecommendUserIds := strings.Split(userRecommend.RecommendCode, "D")
		if 2 <= len(tmpRecommendUserIds) {
			myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
		}
		myRecommendUser, err = uuc.repo.GetUserById(ctx, myUserRecommendUserId)
		if nil != err {
			return nil, err
		}
		inviteUserAddress = myRecommendUser.Address
		myCode = userRecommend.RecommendCode + myCode
	}

	var (
		myUserRecommend []*UserRecommend
		recommendTotal  int64
	)
	myUserRecommend, err = uuc.urRepo.GetUserRecommendByCode(ctx, myCode)
	if nil != err {
		return nil, err
	}
	myRecommendList := make([]*v1.UserInfoReply_ListRecommend, 0)
	if nil != myUserRecommend {
		for _, vMyUserRecommend := range myUserRecommend {
			var (
				myAllRecommendUser *User
			)
			myAllRecommendUser, err = uuc.repo.GetUserById(ctx, vMyUserRecommend.UserId)
			if nil != err {
				return nil, err
			}

			if nil == myAllRecommendUser {
				continue
			}

			recommendTotal += int64(myAllRecommendUser.Total)
			myRecommendList = append(myRecommendList, &v1.UserInfoReply_ListRecommend{Address: myAllRecommendUser.Address})
		}
	}

	// 提现
	var (
		withdraws      []*Withdraw
		withdrawAmount float64
		withdrawList   []*v1.UserInfoReply_ListWithdraw
	)

	withdraws, err = uuc.ubRepo.GetWithdrawByUserId2(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	withdrawList = make([]*v1.UserInfoReply_ListWithdraw, 0)
	for _, v := range withdraws {
		withdrawAmount += v.AmountNew
		withdrawList = append(withdrawList, &v1.UserInfoReply_ListWithdraw{
			Amount:   v.AmountNew,
			CreateAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
		})
	}

	// 充值
	var (
		userEth []*EthUserRecord
	)
	userEth, err = uuc.repo.GetEthUserRecordListByUserId(ctx, myUser.ID)
	if nil != err {
		return nil, err
	}
	listUserEth := make([]*v1.UserInfoReply_ListEthRecord, 0)
	for _, vUserEth := range userEth {
		listUserEth = append(listUserEth, &v1.UserInfoReply_ListEthRecord{
			Amount:    vUserEth.AmountTwo,
			CreatedAt: vUserEth.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
		})
	}

	// 分红
	var (
		RewardFirst  float64
		RewardSecond float64
		RewardThird  float64
		userRewards  []*Reward
		rewardsBuy   []*Reward
	)

	listReward := make([]*v1.UserInfoReply_ListReward, 0)
	LocationList := make([]*v1.UserInfoReply_List, 0)
	userRewards, err = uuc.ubRepo.GetUserRewardByUserId(ctx, myUser.ID)
	if nil != userRewards {
		for _, vUserReward := range userRewards {
			if "reward_first" == vUserReward.Reason {
				RewardFirst = vUserReward.AmountNew
				listReward = append(listReward, &v1.UserInfoReply_ListReward{
					CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
					Reward:    vUserReward.AmountNew,
					Type:      1,
				})
			} else if "reward_second" == vUserReward.Reason {
				RewardSecond = vUserReward.AmountNew
				listReward = append(listReward, &v1.UserInfoReply_ListReward{
					CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
					Reward:    vUserReward.AmountNew,
					Type:      2,
				})
			} else if "reward_third" == vUserReward.Reason {
				RewardThird = vUserReward.AmountNew
				listReward = append(listReward, &v1.UserInfoReply_ListReward{
					CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
					Reward:    vUserReward.AmountNew,
					Type:      3,
				})
			} else if "buy" == vUserReward.Reason {
				LocationList = append(LocationList, &v1.UserInfoReply_List{
					Amount:    uint64(vUserReward.Amount),
					CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
				})
			} else {
				continue
			}
		}
	}

	first := make([]*v1.UserInfoReply_ListFirst, 0)
	second := make([]*v1.UserInfoReply_ListSecond, 0)
	rewardsBuy, err = uuc.ubRepo.GetUserRewardTodayBuy(ctx)
	if nil != err {
		return nil, err
	}

	for _, vRewardsBuy := range rewardsBuy {
		if vRewardsBuy.AmountB < vRewardsBuy.Amount {
			continue
		}

		// 超级
		if 30000 <= vRewardsBuy.AmountB {
			if 30000 <= vRewardsBuy.AmountB-vRewardsBuy.Amount { // 已经是超级节点
				continue
			}

			second = append(second, &v1.UserInfoReply_ListSecond{Address: vRewardsBuy.Address})
		} else if 15000 <= vRewardsBuy.AmountB {
			if 15000 <= vRewardsBuy.AmountB-vRewardsBuy.Amount { // 已经是节点
				continue
			}

			first = append(first, &v1.UserInfoReply_ListFirst{Address: vRewardsBuy.Address})
		}
	}

	return &v1.UserInfoReply{
		Address:           myUser.AddressTwo,
		WithdrawTotal:     withdrawAmount,
		BalanceBiw:        userBalance.BalanceDhb,
		RecommendNum:      int64(len(myUserRecommend)),
		Time:              time.Now().Unix(),
		WithdrawList:      withdrawList,
		InviteUserAddress: inviteUserAddress,
		InviteUrl:         encodeString,
		RecommendTotal:    recommendTotal,
		LocationUsdt:      int64(myUser.Total),
		ListReward:        listReward,
		ListRecommend:     myRecommendList,
		WithdrawMin:       withdrawMin,
		BuyLimit:          buyLimit,
		RewardFirst:       RewardFirst,
		RewardSecond:      RewardSecond,
		RewardThird:       RewardThird,
		First:             first,
		Second:            second,
		ListEth:           listUserEth,
		LocationList:      LocationList,
		Amount:            myUser.Amount,
		Kkdt:              myUser.Kkdt,
		Uudt:              myUser.Uudt,
		FourAmounts:       fourAmounts,
		FiveAmounts:       fiveAmounts,
	}, nil
}
func (uuc *UserUseCase) UserArea(ctx context.Context, req *v1.UserAreaRequest, user *User) (*v1.UserAreaReply, error) {
	var (
		err           error
		userRecommend *UserRecommend
		locationId    = req.LocationId
		myCode        string
	)

	res := make([]*v1.UserAreaReply_List, 0)
	if 0 >= locationId {
		locationId = user.ID
	}
	// 推荐
	userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, locationId)
	if nil == userRecommend {
		return nil, err
	}

	myCode = "D" + strconv.FormatInt(locationId, 10)
	if "" != userRecommend.RecommendCode {
		myCode = userRecommend.RecommendCode + myCode
	}

	if 0 >= len(myCode) {
		return nil, nil
	}

	var (
		myUserRecommend []*UserRecommend
	)
	myUserRecommend, err = uuc.urRepo.GetUserRecommendByCode(ctx, myCode)
	if nil != err {
		return nil, err
	}

	if nil != myUserRecommend {
		for _, vMyUserRecommend := range myUserRecommend {
			var (
				myAllRecommendUser *User
			)
			myAllRecommendUser, err = uuc.repo.GetUserById(ctx, vMyUserRecommend.UserId)
			if nil != err {
				return nil, err
			}

			if nil == myAllRecommendUser {
				continue
			}

			var address1 string
			if 20 <= len(myAllRecommendUser.Address) {
				address1 = myAllRecommendUser.Address[:6] + "..." + myAllRecommendUser.Address[len(myAllRecommendUser.Address)-4:]
			}
			res = append(res, &v1.UserAreaReply_List{
				Address:    address1,
				LocationId: myAllRecommendUser.ID,
				CountLow:   0,
			})
		}
	}

	return &v1.UserAreaReply{Area: res}, nil
}

func (uuc *UserUseCase) UserInfoOld(ctx context.Context, user *User) (*v1.UserInfoReply, error) {
	//var (
	//	myUser                  *User
	//	userInfo                *UserInfo
	//	locations               []*LocationNew
	//	myLastStopLocations     []*LocationNew
	//	userBalance             *UserBalance
	//	userRecommend           *UserRecommend
	//	userRecommends          []*UserRecommend
	//	userRewards             []*Reward
	//	userRewardTotal         int64
	//	userRewardWithdrawTotal int64
	//	encodeString            string
	//	recommendTeamNum        int64
	//	myCode                  string
	//	amount                  = "0"
	//	amount2                 = "0"
	//	configs                 []*Config
	//	myLastLocationCurrent   int64
	//	withdrawAmount          int64
	//	withdrawAmount2         int64
	//	myUserRecommendUserId   int64
	//	inviteUserAddress       string
	//	myRecommendUser         *User
	//	withdrawRate            int64
	//	myLocations             []*v1.UserInfoReply_List
	//	myLocations2            []*v1.UserInfoReply_List22
	//	allRewardList           []*v1.UserInfoReply_List9
	//	timeAgain               int64
	//	err                     error
	//)
	//
	//// 配置
	//configs, err = uuc.configRepo.GetConfigByKeys(ctx,
	//	"time_again",
	//)
	//if nil != configs {
	//	for _, vConfig := range configs {
	//		if "time_again" == vConfig.KeyName {
	//			timeAgain, _ = strconv.ParseInt(vConfig.Value, 10, 64)
	//		}
	//	}
	//}
	//
	//myUser, err = uuc.repo.GetUserById(ctx, user.ID)
	//if nil != err {
	//	return nil, err
	//}
	//userInfo, err = uuc.uiRepo.GetUserInfoByUserId(ctx, myUser.ID)
	//if nil != err {
	//	return nil, err
	//}
	//locations, err = uuc.locationRepo.GetLocationsByUserId(ctx, myUser.ID)
	//if nil != locations && 0 < len(locations) {
	//	for _, v := range locations {
	//		var tmp int64
	//		if v.Current <= v.CurrentMax {
	//			tmp = v.CurrentMax - v.Current
	//		}
	//		if "running" == v.Status {
	//			amount = fmt.Sprintf("%.4f", float64(tmp)/float64(100000))
	//		}
	//
	//		myLocations = append(myLocations, &v1.UserInfoReply_List{
	//			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
	//			Amount:    fmt.Sprintf("%.2f", float64(v.Usdt)/float64(100000)),
	//			Num:       v.Num,
	//			AmountMax: fmt.Sprintf("%.2f", float64(v.CurrentMax)/float64(100000)),
	//		})
	//	}
	//
	//}
	//
	//// 冻结
	//myLastStopLocations, err = uuc.locationRepo.GetMyStopLocationsLast(ctx, myUser.ID)
	//now := time.Now().UTC()
	//tmpNow := now.Add(8 * time.Hour)
	//if nil != myLastStopLocations {
	//	for _, vMyLastStopLocations := range myLastStopLocations {
	//		if tmpNow.Before(vMyLastStopLocations.StopDate.Add(time.Duration(timeAgain) * time.Minute)) {
	//			myLastLocationCurrent += vMyLastStopLocations.Current - vMyLastStopLocations.CurrentMax
	//		}
	//	}
	//}
	//
	//var (
	//	locations2 []*LocationNew
	//)
	//locations2, err = uuc.locationRepo.GetLocationsByUserId2(ctx, myUser.ID)
	//if nil != locations2 && 0 < len(locations2) {
	//	for _, v := range locations2 {
	//		var tmp int64
	//		if v.Current <= v.CurrentMax {
	//			tmp = v.CurrentMax - v.Current
	//		}
	//
	//		if "running" == v.Status {
	//			amount2 = fmt.Sprintf("%.4f", float64(tmp)/float64(100000))
	//		}
	//
	//		myLocations2 = append(myLocations2, &v1.UserInfoReply_List22{
	//			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
	//			Amount:    fmt.Sprintf("%.2f", float64(v.Usdt)/float64(100000)),
	//			AmountMax: fmt.Sprintf("%.2f", float64(v.CurrentMax)/float64(100000)),
	//		})
	//	}
	//
	//}
	//
	//userBalance, err = uuc.ubRepo.GetUserBalance(ctx, myUser.ID)
	//if nil != err {
	//	return nil, err
	//}
	//
	//userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, myUser.ID)
	//if nil == userRecommend {
	//	return nil, err
	//}
	//
	//myCode = "D" + strconv.FormatInt(myUser.ID, 10)
	//codeByte := []byte(myCode)
	//encodeString = base64.StdEncoding.EncodeToString(codeByte)
	//if "" != userRecommend.RecommendCode {
	//	tmpRecommendUserIds := strings.Split(userRecommend.RecommendCode, "D")
	//	if 2 <= len(tmpRecommendUserIds) {
	//		myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
	//	}
	//	myRecommendUser, err = uuc.repo.GetUserById(ctx, myUserRecommendUserId)
	//	if nil != err {
	//		return nil, err
	//	}
	//	inviteUserAddress = myRecommendUser.Address
	//	myCode = userRecommend.RecommendCode + myCode
	//}
	//
	//// 团队
	//var (
	//	teamUserIds        []int64
	//	teamUsers          map[int64]*User
	//	teamUserInfos      map[int64]*UserInfo
	//	teamUserAddresses  []*v1.UserInfoReply_List7
	//	recommendAddresses []*v1.UserInfoReply_List11
	//	teamLocations      map[int64][]*Location
	//	recommendUserIds   map[int64]int64
	//	userBalanceMap     map[int64]*UserBalance
	//)
	//recommendUserIds = make(map[int64]int64, 0)
	//userRecommends, err = uuc.urRepo.GetUserRecommendLikeCode(ctx, myCode)
	//if nil != userRecommends {
	//	for _, vUserRecommends := range userRecommends {
	//		if myCode == vUserRecommends.RecommendCode {
	//			recommendUserIds[vUserRecommends.UserId] = vUserRecommends.UserId
	//		}
	//		teamUserIds = append(teamUserIds, vUserRecommends.UserId)
	//	}
	//
	//	// 用户信息
	//	recommendTeamNum = int64(len(userRecommends))
	//	teamUsers, _ = uuc.repo.GetUserByUserIds(ctx, teamUserIds...)
	//	teamUserInfos, _ = uuc.uiRepo.GetUserInfoByUserIds(ctx, teamUserIds...)
	//	teamLocations, _ = uuc.locationRepo.GetLocationMapByIds(ctx, teamUserIds...)
	//	userBalanceMap, _ = uuc.ubRepo.GetUserBalanceByUserIds(ctx, teamUserIds...)
	//	if nil != teamUsers {
	//		for _, vTeamUsers := range teamUsers {
	//			var locationAmount int64
	//			if _, ok := teamUserInfos[vTeamUsers.ID]; !ok {
	//				continue
	//			}
	//
	//			if _, ok := userBalanceMap[vTeamUsers.ID]; !ok {
	//				continue
	//			}
	//
	//			if _, ok := teamLocations[vTeamUsers.ID]; ok {
	//
	//				for _, vTeamLocations := range teamLocations[vTeamUsers.ID] {
	//					locationAmount += vTeamLocations.Usdt
	//				}
	//			}
	//			if _, ok := recommendUserIds[vTeamUsers.ID]; ok {
	//				recommendAddresses = append(recommendAddresses, &v1.UserInfoReply_List11{
	//					Address: vTeamUsers.Address,
	//					Usdt:    fmt.Sprintf("%.2f", float64(locationAmount)/float64(100000)),
	//					Vip:     teamUserInfos[vTeamUsers.ID].Vip,
	//				})
	//			}
	//
	//			teamUserAddresses = append(teamUserAddresses, &v1.UserInfoReply_List7{
	//				Address: vTeamUsers.Address,
	//				Usdt:    fmt.Sprintf("%.2f", float64(locationAmount)/float64(100000)),
	//				Vip:     teamUserInfos[vTeamUsers.ID].Vip,
	//			})
	//		}
	//	}
	//}
	//
	//var (
	//	totalDeposit      int64
	//	userBalanceRecord []*UserBalanceRecord
	//	depositList       []*v1.UserInfoReply_List13
	//)
	//
	//userBalanceRecord, _ = uuc.ubRepo.GetUserBalanceRecordsByUserId(ctx, myUser.ID)
	//for _, vUserBalanceRecord := range userBalanceRecord {
	//	totalDeposit += vUserBalanceRecord.Amount
	//	depositList = append(depositList, &v1.UserInfoReply_List13{
	//		CreatedAt: vUserBalanceRecord.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
	//		Amount:    fmt.Sprintf("%.4f", float64(vUserBalanceRecord.Amount)/float64(100000)),
	//	})
	//}
	//
	//// 累计奖励
	//var (
	//	locationRewardList            []*v1.UserInfoReply_List2
	//	recommendRewardList           []*v1.UserInfoReply_List5
	//	locationTotal                 int64
	//	yesterdayLocationTotal        int64
	//	recommendRewardTotal          int64
	//	recommendRewardDhbTotal       int64
	//	yesterdayRecommendRewardTotal int64
	//)
	//
	//var startDate time.Time
	//var endDate time.Time
	//if 16 <= now.Hour() {
	//	startDate = now.AddDate(0, 0, -1)
	//	endDate = startDate.AddDate(0, 0, 1)
	//} else {
	//	startDate = now.AddDate(0, 0, -2)
	//	endDate = startDate.AddDate(0, 0, 1)
	//}
	//yesterdayStart := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 16, 0, 0, 0, time.UTC)
	//yesterdayEnd := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 16, 0, 0, 0, time.UTC)
	//
	//fmt.Println(now, yesterdayStart, yesterdayEnd)
	//userRewards, err = uuc.ubRepo.GetUserRewardByUserId(ctx, myUser.ID)
	//if nil != userRewards {
	//	for _, vUserReward := range userRewards {
	//		if "location" == vUserReward.Reason {
	//			locationTotal += vUserReward.Amount
	//			if vUserReward.CreatedAt.Before(yesterdayEnd) && vUserReward.CreatedAt.After(yesterdayStart) {
	//				yesterdayLocationTotal += vUserReward.Amount
	//			}
	//			locationRewardList = append(locationRewardList, &v1.UserInfoReply_List2{
	//				CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
	//				Amount:    fmt.Sprintf("%.2f", float64(vUserReward.Amount)/float64(100000)),
	//			})
	//
	//			userRewardTotal += vUserReward.Amount
	//			allRewardList = append(allRewardList, &v1.UserInfoReply_List9{
	//				CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
	//				Amount:    fmt.Sprintf("%.2f", float64(vUserReward.Amount)/float64(100000)),
	//			})
	//		} else if "recommend" == vUserReward.Reason {
	//
	//			recommendRewardList = append(recommendRewardList, &v1.UserInfoReply_List5{
	//				CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
	//				Amount:    fmt.Sprintf("%.2f", float64(vUserReward.Amount)/float64(100000)),
	//			})
	//
	//			recommendRewardTotal += vUserReward.Amount
	//			if vUserReward.CreatedAt.Before(yesterdayEnd) && vUserReward.CreatedAt.After(yesterdayStart) {
	//				yesterdayRecommendRewardTotal += vUserReward.Amount
	//			}
	//
	//			userRewardTotal += vUserReward.Amount
	//			allRewardList = append(allRewardList, &v1.UserInfoReply_List9{
	//				CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
	//				Amount:    fmt.Sprintf("%.2f", float64(vUserReward.Amount)/float64(100000)),
	//			})
	//		} else if "reward_withdraw" == vUserReward.Reason {
	//			userRewardTotal += vUserReward.Amount
	//			userRewardWithdrawTotal += vUserReward.Amount
	//		} else if "recommend_token" == vUserReward.Reason {
	//			recommendRewardDhbTotal += vUserReward.Amount
	//		}
	//	}
	//}
	//
	//var (
	//	withdraws []*Withdraw
	//)
	//withdraws, err = uuc.ubRepo.GetWithdrawByUserId2(ctx, user.ID)
	//if nil != err {
	//	return nil, err
	//}
	//for _, v := range withdraws {
	//	if "usdt" == v.Type {
	//		withdrawAmount += v.RelAmount
	//	}
	//	if "usdt_2" == v.Type {
	//		withdrawAmount2 += v.RelAmount
	//	}
	//}
	//
	//return &v1.UserInfoReply{
	//	Address:                 myUser.Address,
	//	Level:                   userInfo.Vip,
	//	Amount:                  amount,
	//	Amount2:                 amount2,
	//	LocationList2:           myLocations2,
	//	AmountB:                 fmt.Sprintf("%.2f", float64(myLastLocationCurrent)/float64(100000)),
	//	DepositList:             depositList,
	//	BalanceUsdt:             fmt.Sprintf("%.2f", float64(userBalance.BalanceUsdt)/float64(100000)),
	//	BalanceUsdt2:            fmt.Sprintf("%.2f", float64(userBalance.BalanceUsdt2)/float64(100000)),
	//	BalanceDhb:              fmt.Sprintf("%.2f", float64(userBalance.BalanceDhb)/float64(100000)),
	//	InviteUrl:               encodeString,
	//	RecommendNum:            userInfo.HistoryRecommend,
	//	RecommendTeamNum:        recommendTeamNum,
	//	Total:                   fmt.Sprintf("%.2f", float64(userRewardTotal)/float64(100000)),
	//	RewardWithdraw:          fmt.Sprintf("%.2f", float64(userRewardWithdrawTotal)/float64(100000)),
	//	WithdrawAmount:          fmt.Sprintf("%.2f", float64(withdrawAmount)/float64(100000)),
	//	WithdrawAmount2:         fmt.Sprintf("%.2f", float64(withdrawAmount2)/float64(100000)),
	//	LocationTotal:           fmt.Sprintf("%.2f", float64(locationTotal)/float64(100000)),
	//	Account:                 "0xAfC39F3592A1024144D1ba6DC256397F4DbfbE2f",
	//	LocationList:            myLocations,
	//	TeamAddressList:         teamUserAddresses,
	//	AllRewardList:           allRewardList,
	//	InviteUserAddress:       inviteUserAddress,
	//	RecommendAddressList:    recommendAddresses,
	//	WithdrawRate:            withdrawRate,
	//	RecommendRewardTotal:    fmt.Sprintf("%.2f", float64(recommendRewardTotal)/float64(100000)),
	//	RecommendRewardDhbTotal: fmt.Sprintf("%.2f", float64(recommendRewardDhbTotal)/float64(100000)),
	//	TotalDeposit:            fmt.Sprintf("%.2f", float64(totalDeposit)/float64(100000)),
	//	WithdrawAll:             fmt.Sprintf("%.2f", float64(withdrawAmount+withdrawAmount2)/float64(100000)),
	//}, nil
	return nil, nil

}

func (uuc *UserUseCase) RewardList(ctx context.Context, req *v1.RewardListRequest, user *User) (*v1.RewardListReply, error) {

	res := &v1.RewardListReply{
		Rewards: make([]*v1.RewardListReply_List, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) RecommendRewardList(ctx context.Context, user *User) (*v1.RecommendRewardListReply, error) {

	res := &v1.RecommendRewardListReply{
		Rewards: make([]*v1.RecommendRewardListReply_List, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) FeeRewardList(ctx context.Context, user *User) (*v1.FeeRewardListReply, error) {
	res := &v1.FeeRewardListReply{
		Rewards: make([]*v1.FeeRewardListReply_List, 0),
	}
	return res, nil
}

func (uuc *UserUseCase) TranList(ctx context.Context, user *User, reqTypeCoin string, reqTran string) (*v1.TranListReply, error) {

	var (
		userBalanceRecord []*UserBalanceRecord
		typeCoin          = "usdt"
		tran              = "tran"
		err               error
	)

	res := &v1.TranListReply{
		Tran: make([]*v1.TranListReply_List, 0),
	}

	if "" != reqTypeCoin {
		typeCoin = reqTypeCoin
	}

	if "tran_to" == reqTran {
		tran = reqTran
	}

	userBalanceRecord, err = uuc.ubRepo.GetUserBalanceRecordByUserId(ctx, user.ID, typeCoin, tran)
	if nil != err {
		return res, err
	}

	for _, v := range userBalanceRecord {
		res.Tran = append(res.Tran, &v1.TranListReply_List{
			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Amount:    fmt.Sprintf("%.2f", float64(v.Amount)/float64(100000)),
		})
	}

	return res, nil
}

func (uuc *UserUseCase) WithdrawList(ctx context.Context, user *User, reqTypeCoin string) (*v1.WithdrawListReply, error) {

	var (
		withdraws []*Withdraw
		typeCoin  = "usdt"
		err       error
	)

	res := &v1.WithdrawListReply{
		Withdraw: make([]*v1.WithdrawListReply_List, 0),
	}

	if "" != reqTypeCoin {
		typeCoin = reqTypeCoin
	}

	withdraws, err = uuc.ubRepo.GetWithdrawByUserId(ctx, user.ID, typeCoin)
	if nil != err {
		return res, err
	}

	for _, v := range withdraws {
		res.Withdraw = append(res.Withdraw, &v1.WithdrawListReply_List{
			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Amount:    fmt.Sprintf("%.2f", float64(v.Amount)/float64(100000)),
			Status:    v.Status,
			Type:      v.Type,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) TradeList(ctx context.Context, user *User) (*v1.TradeListReply, error) {

	var (
		trades []*Trade
		err    error
	)

	res := &v1.TradeListReply{
		Trade: make([]*v1.TradeListReply_List, 0),
	}

	trades, err = uuc.ubRepo.GetTradeByUserId(ctx, user.ID)
	if nil != err {
		return res, err
	}

	for _, v := range trades {
		res.Trade = append(res.Trade, &v1.TradeListReply_List{
			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			AmountCsd: fmt.Sprintf("%.2f", float64(v.AmountCsd)/float64(100000)),
			AmountHbs: fmt.Sprintf("%.2f", float64(v.AmountHbs)/float64(100000)),
			Status:    v.Status,
		})
	}

	return res, nil
}

// Exchange Exchange.
func (uuc *UserUseCase) Exchange(ctx context.Context, req *v1.ExchangeRequest, user *User) (*v1.ExchangeReply, error) {
	return &v1.ExchangeReply{
		Status: "ok",
	}, nil

}

func (uuc *UserUseCase) Buy(ctx context.Context, req *v1.BuyRequest, user *User) (*v1.BuyReply, error) {
	var (
		amount    = req.SendBody.Amount
		strUpdate string
		kkdt      uint64
		uudt      int64
		err       error
	)

	var (
		configs []*Config
		kkdtA   uint64
		kkdtB   uint64
		kkdtC   uint64
		kkdtD   uint64
		kkdtE   uint64
		kkdtF   uint64
		kkdtG   uint64
		kkdtH   uint64
	)
	configs, err = uuc.configRepo.GetConfigByKeys(ctx,
		"kkdt_a", "kkdt_b", "kkdt_c", "kkdt_d", "kkdt_e", "kkdt_f", "kkdt_g", "kkdt_h",
	)
	if nil != configs {
		for _, vConfig := range configs {
			if "kkdt_a" == vConfig.KeyName {
				kkdtA, _ = strconv.ParseUint(vConfig.Value, 10, 64)
			}
			if "kkdt_b" == vConfig.KeyName {
				kkdtB, _ = strconv.ParseUint(vConfig.Value, 10, 64)
			}
			if "kkdt_c" == vConfig.KeyName {
				kkdtC, _ = strconv.ParseUint(vConfig.Value, 10, 64)
			}
			if "kkdt_d" == vConfig.KeyName {
				kkdtD, _ = strconv.ParseUint(vConfig.Value, 10, 64)
			}
			if "kkdt_e" == vConfig.KeyName {
				kkdtE, _ = strconv.ParseUint(vConfig.Value, 10, 64)
			}
			if "kkdt_f" == vConfig.KeyName {
				kkdtF, _ = strconv.ParseUint(vConfig.Value, 10, 64)
			}
			if "kkdt_g" == vConfig.KeyName {
				kkdtG, _ = strconv.ParseUint(vConfig.Value, 10, 64)
			}
			if "kkdt_h" == vConfig.KeyName {
				kkdtH, _ = strconv.ParseUint(vConfig.Value, 10, 64)
			}

		}
	}

	if 30000 <= amount {
		strUpdate = "total_f"
		amount = 30000
		kkdt = kkdtA
		uudt = 30000
	} else if 15000 <= amount {
		strUpdate = "total_d"
		amount = 15000
		kkdt = kkdtB
		uudt = 15000
	} else if 5000 <= amount {
		strUpdate = "total_c"
		amount = 5000
		kkdt = kkdtC
		uudt = 5000
	} else if 3000 <= amount {
		strUpdate = "total_b"
		amount = 3000
		kkdt = kkdtD
		uudt = 3000
	} else if 1000 <= amount {
		strUpdate = "total_a"
		amount = 1000
		kkdt = kkdtE
		uudt = 1000
	} else if 500 <= amount {
		strUpdate = "total_h"
		amount = 500
		kkdt = kkdtF
	} else if 300 <= amount {
		strUpdate = "total_i"
		amount = 300
		kkdt = kkdtG
	} else if 100 <= amount {
		strUpdate = "total_j"
		amount = 100
		kkdt = kkdtH
	} else {
		return &v1.BuyReply{
			Status: "无效金额",
		}, nil
	}

	if amount > user.Amount {
		return &v1.BuyReply{
			Status: "余额不足",
		}, nil
	}

	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = uuc.repo.UpdateUserNewTwoNew(ctx, user.ID, amount, strUpdate, uudt, int64(kkdt))
		if nil != err {
			return err
		}

		// 充值记录
		err = uuc.repo.InRecordNew(ctx, user.ID, user.Address, int64(amount), int64(user.Total))
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		return nil, err
	}

	return &v1.BuyReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) Withdraw(ctx context.Context, req *v1.WithdrawRequest, user *User, password string) (*v1.WithdrawReply, error) {
	var (
		//u           *User
		err         error
		userBalance *UserBalance
	)

	//if "2" == req.SendBody.Type {
	//	req.SendBody.Type = "usdt"
	//}
	//u, _ = uuc.repo.GetUserById(ctx, user.ID)
	//if nil != err {
	//	return nil, err
	//}

	//if "" == u.Password || 6 > len(u.Password) {
	//	return nil, errors.New(500, "ERROR_TOKEN", "未设置密码，联系管理员")
	//}
	//
	//if u.Password != user.Password {
	//	return nil, errors.New(403, "ERROR_TOKEN", "无效TOKEN")
	//}

	//if password != u.Password {
	//	return nil, errors.New(500, "密码错误", "密码错误")
	//}

	//if "usdt" != req.SendBody.Type {
	//	return &v1.WithdrawReply{
	//		Status: "fail",
	//	}, nil
	//}

	userBalance, err = uuc.ubRepo.GetUserBalance(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	amountFloat, _ := strconv.ParseFloat(req.SendBody.Amount, 10)
	if 0 >= amountFloat {
		return &v1.WithdrawReply{
			Status: "无效金额",
		}, nil
	} else if userBalance.BalanceDhb < amountFloat {
		return &v1.WithdrawReply{
			Status: "余额不足",
		}, nil
	}
	//if "dhb" == req.SendBody.Type {
	//	if userBalance.BalanceDhb < amount {
	//		return &v1.WithdrawReply{
	//			Status: "fail",
	//		}, nil
	//	}
	//
	//	if 10000000 > amount {
	//		return &v1.WithdrawReply{
	//			Status: "fail",
	//		}, nil
	//	}
	//}

	// 配置
	var (
		configs      []*Config
		withdrawRate float64
		withdrawMin  float64
	)
	configs, err = uuc.configRepo.GetConfigByKeys(ctx,
		"withdraw_rate",
		"withdraw_min",
	)
	if nil != configs {
		for _, vConfig := range configs {
			if "withdraw_rate" == vConfig.KeyName {
				withdrawRate, _ = strconv.ParseFloat(vConfig.Value, 10)
			}

			if "withdraw_win" == vConfig.KeyName {
				withdrawMin, _ = strconv.ParseFloat(vConfig.Value, 10)
			}
		}
	}

	if amountFloat < withdrawMin {
		return &v1.WithdrawReply{
			Status: "提现不足最小值",
		}, nil
	}

	//if "usdt_2" == req.SendBody.Type {
	//	if userBalance.BalanceUsdt2 < amount {
	//		return &v1.WithdrawReply{
	//			Status: "fail",
	//		}, nil
	//	}
	//
	//	if 1000000 > amount {
	//		return &v1.WithdrawReply{
	//			Status: "fail",
	//		}, nil
	//	}
	//}

	//var userRecommend *UserRecommend
	//userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, user.ID)
	//if nil == userRecommend {
	//	return &v1.WithdrawReply{
	//		Status: "信息错误",
	//	}, nil
	//}
	//
	//var (
	//	tmpRecommendUserIds    []string
	//	tmpRecommendUserIdsInt []int64
	//)
	//if "" != userRecommend.RecommendCode {
	//	tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
	//}
	//lastKey := len(tmpRecommendUserIds) - 1
	//if 1 <= lastKey {
	//	for i := 0; i <= lastKey; i++ {
	//		// 有占位信息，推荐人推荐人的上一代
	//		if lastKey-i <= 0 {
	//			break
	//		}
	//
	//		tmpMyTopUserRecommendUserId, _ := strconv.ParseInt(tmpRecommendUserIds[lastKey-i], 10, 64) // 最后一位是直推人
	//		tmpRecommendUserIdsInt = append(tmpRecommendUserIdsInt, tmpMyTopUserRecommendUserId)
	//	}
	//}

	// 配置
	//var (
	//	configs      []*Config
	//	withdrawRate int64
	//)
	//configs, err = uuc.configRepo.GetConfigByKeys(ctx,
	//	"withdraw_rate",
	//)
	//if nil != configs {
	//	for _, vConfig := range configs {
	//		if "withdraw_rate" == vConfig.KeyName {
	//			withdrawRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
	//		}
	//	}
	//}

	relAmount := amountFloat - amountFloat*withdrawRate/100
	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

		err = uuc.ubRepo.WithdrawUsdt2(ctx, user.ID, amountFloat) // 提现
		if nil != err {
			return err
		}

		_, err = uuc.ubRepo.GreateWithdraw(ctx, user.ID, relAmount, amountFloat, "BIW")
		if nil != err {
			return err
		}
		//else if "usdt_2" == req.SendBody.Type {
		//	err = uuc.ubRepo.WithdrawUsdt3(ctx, user.ID, amount) // 提现
		//	if nil != err {
		//		return err
		//	}
		//	_, err = uuc.ubRepo.GreateWithdraw(ctx, user.ID, amount, amount-amount*withdrawRate/100, amount*withdrawRate/100, req.SendBody.Type)
		//	if nil != err {
		//		return err
		//	}
		//
		//}
		//else if "dhb" == req.SendBody.Type {
		//	err = uuc.ubRepo.WithdrawDhb(ctx, user.ID, amount) // 提现
		//	if nil != err {
		//		return err
		//	}
		//	_, err = uuc.ubRepo.GreateWithdraw(ctx, user.ID, amount-1000000, 1000000, req.SendBody.Type)
		//	if nil != err {
		//		return err
		//	}
		//}

		return nil
	}); nil != err {
		return nil, err
	}

	return &v1.WithdrawReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) Tran(ctx context.Context, req *v1.TranRequest, user *User, password string) (*v1.TranReply, error) {
	return &v1.TranReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) Trade(ctx context.Context, req *v1.WithdrawRequest, user *User, amount int64, amountB int64, amount2 int64, password string) (*v1.WithdrawReply, error) {
	return &v1.WithdrawReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) SetBalanceReward(ctx context.Context, req *v1.SetBalanceRewardRequest, user *User) (*v1.SetBalanceRewardReply, error) {
	var (
		err         error
		userBalance *UserBalance
	)

	amountFloat, _ := strconv.ParseFloat(req.SendBody.Amount, 10)
	amountFloat *= 100000
	amount, _ := strconv.ParseInt(strconv.FormatFloat(amountFloat, 'f', -1, 64), 10, 64)
	if 0 >= amount {
		return &v1.SetBalanceRewardReply{
			Status: "fail",
		}, nil
	}

	userBalance, err = uuc.ubRepo.GetUserBalance(ctx, user.ID)
	if nil != err {
		return nil, err
	}

	if userBalance.BalanceUsdt < amount {
		return &v1.SetBalanceRewardReply{
			Status: "fail",
		}, nil
	}

	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

		err = uuc.ubRepo.SetBalanceReward(ctx, user.ID, amount) // 提现
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		return nil, err
	}

	return &v1.SetBalanceRewardReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) DeleteBalanceReward(ctx context.Context, req *v1.DeleteBalanceRewardRequest, user *User) (*v1.DeleteBalanceRewardReply, error) {
	var (
		err            error
		balanceRewards []*BalanceReward
	)

	amountFloat, _ := strconv.ParseFloat(req.SendBody.Amount, 10)
	amountFloat *= 100000
	amount, _ := strconv.ParseInt(strconv.FormatFloat(amountFloat, 'f', -1, 64), 10, 64)
	if 0 >= amount {
		return &v1.DeleteBalanceRewardReply{
			Status: "fail",
		}, nil
	}

	balanceRewards, err = uuc.ubRepo.GetBalanceRewardByUserId(ctx, user.ID)
	if nil != err {
		return &v1.DeleteBalanceRewardReply{
			Status: "fail",
		}, nil
	}

	var totalBalanceRewardAmount int64
	for _, vBalanceReward := range balanceRewards {
		totalBalanceRewardAmount += vBalanceReward.Amount
	}

	if totalBalanceRewardAmount < amount {
		return &v1.DeleteBalanceRewardReply{
			Status: "fail",
		}, nil
	}

	for _, vBalanceReward := range balanceRewards {
		tmpAmount := int64(0)
		Status := int64(1)

		if amount-vBalanceReward.Amount < 0 {
			tmpAmount = amount
		} else {
			tmpAmount = vBalanceReward.Amount
			Status = 2
		}

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = uuc.ubRepo.UpdateBalanceReward(ctx, user.ID, vBalanceReward.ID, tmpAmount, Status) // 提现
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			return nil, err
		}
		amount -= tmpAmount

		if amount <= 0 {
			break
		}
	}

	return &v1.DeleteBalanceRewardReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) AdminRewardList(ctx context.Context, req *v1.AdminRewardListRequest) (*v1.AdminRewardListReply, error) {
	res := &v1.AdminRewardListReply{
		Rewards: make([]*v1.AdminRewardListReply_List, 0),
	}
	return res, nil
}

func (uuc *UserUseCase) AdminUserList(ctx context.Context, req *v1.AdminUserListRequest) (*v1.AdminUserListReply, error) {

	res := &v1.AdminUserListReply{
		Users: make([]*v1.AdminUserListReply_UserList, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) GetUserByUserIds(ctx context.Context, userIds ...int64) (map[int64]*User, error) {
	return uuc.repo.GetUserByUserIds(ctx, userIds...)
}

func (uuc *UserUseCase) GetUserByUserId(ctx context.Context, userId int64) (*User, error) {
	return uuc.repo.GetUserById(ctx, userId)
}

func (uuc *UserUseCase) AdminLocationList(ctx context.Context, req *v1.AdminLocationListRequest) (*v1.AdminLocationListReply, error) {
	res := &v1.AdminLocationListReply{
		Locations: make([]*v1.AdminLocationListReply_LocationList, 0),
	}
	return res, nil

}

func (uuc *UserUseCase) AdminRecommendList(ctx context.Context, req *v1.AdminUserRecommendRequest) (*v1.AdminUserRecommendReply, error) {
	res := &v1.AdminUserRecommendReply{
		Users: make([]*v1.AdminUserRecommendReply_List, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) AdminMonthRecommend(ctx context.Context, req *v1.AdminMonthRecommendRequest) (*v1.AdminMonthRecommendReply, error) {

	res := &v1.AdminMonthRecommendReply{
		Users: make([]*v1.AdminMonthRecommendReply_List, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) AdminConfig(ctx context.Context, req *v1.AdminConfigRequest) (*v1.AdminConfigReply, error) {
	res := &v1.AdminConfigReply{
		Config: make([]*v1.AdminConfigReply_List, 0),
	}
	return res, nil
}

func (uuc *UserUseCase) AdminConfigUpdate(ctx context.Context, req *v1.AdminConfigUpdateRequest) (*v1.AdminConfigUpdateReply, error) {
	res := &v1.AdminConfigUpdateReply{}
	return res, nil
}

func (uuc *UserUseCase) GetWithdrawPassOrRewardedList(ctx context.Context) ([]*Withdraw, error) {
	return uuc.ubRepo.GetWithdrawPassOrRewarded(ctx)
}

func (uuc *UserUseCase) UpdateWithdrawDoing(ctx context.Context, id int64) (*Withdraw, error) {
	return uuc.ubRepo.UpdateWithdraw(ctx, id, "doing")
}

func (uuc *UserUseCase) UpdateWithdrawSuccess(ctx context.Context, id int64) (*Withdraw, error) {
	return uuc.ubRepo.UpdateWithdraw(ctx, id, "success")
}

func (uuc *UserUseCase) AdminWithdrawList(ctx context.Context, req *v1.AdminWithdrawListRequest) (*v1.AdminWithdrawListReply, error) {
	res := &v1.AdminWithdrawListReply{
		Withdraw: make([]*v1.AdminWithdrawListReply_List, 0),
	}

	return res, nil

}

func (uuc *UserUseCase) AdminFee(ctx context.Context, req *v1.AdminFeeRequest) (*v1.AdminFeeReply, error) {
	return &v1.AdminFeeReply{}, nil
}

func (uuc *UserUseCase) AdminAll(ctx context.Context, req *v1.AdminAllRequest) (*v1.AdminAllReply, error) {

	return &v1.AdminAllReply{}, nil
}

func (uuc *UserUseCase) AdminWithdraw(ctx context.Context, req *v1.AdminWithdrawRequest) (*v1.AdminWithdrawReply, error) {
	return &v1.AdminWithdrawReply{}, nil
}
