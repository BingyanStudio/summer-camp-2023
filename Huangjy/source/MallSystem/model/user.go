package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
 *	后台管理部分应该和前台的用户部分完全分离，独立存储管理员的相关信息
 */

/*
 *	用户信息的结构体
 *	用户ID（在表单中不传递）
 *	用户姓名和昵称（必填项）
 *	用户密码（在解码时不解析）（必填项）
 *	用户电话号码和邮箱
 *	用户收藏数和被浏览数
 */
type UserInfo struct {
	ID            primitive.ObjectID `json:"id" form:"-" binding:"-" bson:"_id,omitempty"`
	Username      string             `json:"username" form:"username" binding:"required" bson:"username"`
	Password      string             `json:"-" form:"password" binding:"required" bson:"password"`
	Nickname      string             `json:"nickname" form:"nickname" binding:"-" bson:"nickname"`
	Mobile        string             `json:"mobile" form:"mobile" binding:"required" bson:"mobile"`
	Email         string             `json:"email" form:"email" binding:"required" bson:"email"`
	CollectCount  int                `json:"collectCount" form:"-" binding:"-" bson:"collectCount"`
	BeViewedCount int                `json:"beViewedCount" form:"-" binding:"-" bson:"beViewedCount"`
}

/*****************************************/

/*
 *	登陆时，用户的信息
 *	用户的姓名
 *	用户的密码
 */
type Login struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
