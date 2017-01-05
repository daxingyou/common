package loginService

import (
	"casino_common/proto/ddproto"
	"casino_common/common/userService"
	"errors"
	"casino_common/common/service/signService"
	"casino_common/common/log"
	"github.com/golang/protobuf/proto"
)

//做登录的操作...
func DoLogin(weixin *ddproto.WeixinInfo, userId uint32) (*ddproto.User, error) {
	var user *ddproto.User
	//不是初次登录,需要用userId 来登录
	if weixin == nil {
		//判断uerId
		user := userService.GetUserById(userId)
		return user, nil
	} else {
		//微信不等于空的情况,表示是新用户
		//1,首先通过weixinInfo 在数据库中查找 用户是否存在，如果用户存在，则表示，登陆成功
		user = userService.GetUserByOpenId(weixin.GetUnionId())
		if user == nil {
			//表示数据库中不存在次用户，新增加一个人后返回
			if weixin.GetOpenId() == "" || weixin.GetHeadUrl() == "" || weixin.GetNickName() == "" {
				return nil, errors.New("新增用户失败...")
			}
			//如果数据库中不存在用户，那么重新生成一个user
			return userService.NewUserAndSave(weixin.GetUnionId(), weixin.GetOpenId(), weixin.GetNickName(), weixin.GetHeadUrl(), weixin.GetSex(), weixin.GetCity())
		}
	}

	//最终表示登录失败...
	return nil, nil
}

func DoLoginSuccess(userId uint32) error {
	error := signService.DoSignLottery(userId)
	if error != nil {
		log.T("签到失败, 错误信息[%v]", error)
		return error
	}
	return nil
}

//游客注册
func TouristReg() *ddproto.CommonAckReg {
	//func NewUserAndSave(unionId, openId, wxNickName, headUrl string, sex int32, city string) (*ddproto.User, error) {
	user, err := userService.NewUserAndSave("", "", "", "", 1, "")
	if err != nil || user == nil {
		log.E("注册用户的时候失败...")
		return nil
	}

	//返回结果
	ack := new(ddproto.CommonAckReg)
	ack.UserId = proto.Uint32(user.GetId())
	return ack
}

//微信注册
func WxReg(weixin *ddproto.WeixinInfo) *ddproto.CommonAckReg {
	//检测参数
	if weixin.GetOpenId() == "" || weixin.GetHeadUrl() == "" || weixin.GetNickName() == "" {
		log.E("玩家注册的时候失败，因为微信的信息[%v]不够...", weixin)
		return nil
	}

	//微信注册的时候需要先判断是否已经注册过了，如果注册过了直接返回userId ,否则注册
	user := userService.GetUserByOpenId(weixin.GetUnionId())
	if user == nil {
		user, _ = userService.NewUserAndSave(weixin.GetUnionId(), weixin.GetOpenId(), weixin.GetNickName(), weixin.GetHeadUrl(), weixin.GetSex(), weixin.GetCity())
	}

	if user == nil {
		log.E("玩家注册的时候失败，注册的时候失败", weixin)
		return nil
	}

	//返回结果
	ack := new(ddproto.CommonAckReg)
	ack.UserId = proto.Uint32(user.GetId())
	return ack

}