package service

import (
	"context"
	"lottery-api/internal/model"
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func (s *Svc) GetPrizeList(c context.Context) (data *model.PrizeListReply, err error) {
	data = new(model.PrizeListReply)
	data.List, err = s.dao.FetchPrizes(c)
	return
}

func (s *Svc) DrawPrize(c context.Context, phone int64) (drawReply *model.DrawReply, err error) {
	// default
	drawReply = new(model.DrawReply)
	drawReply.Prize = model.NoPrize
	drawReply.Msg = model.PrizeNames[model.NoPrize] // 提示用
	defer s.asyncSaveRecords(drawReply, phone)      // 退出时异步保存记录

	curPrize, curPrizeName := s.prizeMatching(getMyPrizeCode()) // 奖品匹配

	lockName, curLockId := curPrize.GetLockName(), getLockId() // 根据奖品类型独立设锁

	defer s.dao.UnLock(lockName, curLockId)                       // 释放 只有自己可以解锁自己 其他等解锁或超时
	if !s.dao.SetLock(lockName, curLockId, model.HandleTimeOut) { // 加锁 更新数据库
		drawReply.Msg = "请求频繁，请稍后再试"
		s.dao.Logger.Println("s.DrawPrize Get： ", drawReply.Msg)
		return
	}

	verify := VerifyUserWinOver(s, c, phone)
	if verify {
		drawReply.Msg = "当日已达到最大抽奖次数，请明天再来哦"
		return
	}

	if curPrize.IsUnlimited() { // 无限奖品 贴纸
		drawReply.IsWin = true
		drawReply.Prize = curPrize
		drawReply.Msg = curPrizeName
		return
	}

	curPrizeInfo, err := s.dao.FindOnePrize(c, curPrize.Where()) // 以下为有限奖品逻辑
	if err != nil {
		s.dao.Logger.Printf("s.DrawPrize FindOnePrize prize(%s) err(%v)", curPrizeName, err)
		return
	}

	if curPrizeInfo.Stock == 0 { // 有限且没有剩余奖品，无法发奖
		s.dao.Logger.Println("s.DrawPrize out of stock： ", drawReply.Msg)
		return
	}

	row, err := s.dao.UpdatePrize(c, curPrize) // 有限的 还有剩余奖品 扣库存
	if err != nil {
		s.dao.Logger.Printf("s.DrawPrize UpdatePrize prize(%s) err(%v)", curPrizeName, err) // 提示谢谢参与

		return
	}

	if row != model.StockDeductSuccess { // 更新库存失败： 无库存了 未能更新任何记录
		s.dao.Logger.Println("s.DrawPrize Get： ", drawReply.Msg) // 提示谢谢参与
		return
	}

	// 成功
	drawReply.IsWin = true
	drawReply.Prize = curPrize
	drawReply.Msg = curPrizeName
	s.dao.Logger.Println("s.DrawPrize Get： ", drawReply.Msg, "----", curPrize)
	return
}

func VerifyUserWinOver(s *Svc, c context.Context, phone int64) (isOk bool) {
	t := time.Now()
	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	tm2 := tm1.AddDate(0, 0, 1)
	start := tm1.Format("2006-01-02 15:04:05")
	end := tm2.Format("2006-01-02 15:04:05")
	where := " where phone = " + strconv.FormatInt(phone, 10) + " and create_time > " + start + " and create_time <= " + end

	userPrizes, err := s.dao.FindOnePrize(c, where)
	if err != nil {
		s.dao.Logger.Println("s.FindOnePrize err", err)
		return true
	}
	if userPrizes != nil && userPrizes.Id > 0 {
		return true
	}
	return false
}

func (s *Svc) prizeMatching(code int) (curPrize model.PrizeKey, curPrizeName string) {
	curPrize = model.Paster
	curPrizeName = model.PrizeNames[curPrize]

	for i, prize := range model.PrizeList { // 从奖品列表中匹配，是否中奖
		rate := &model.RateList[i]
		if code >= rate.Start && code <= rate.End {
			// 满足中奖条件
			curPrize, curPrizeName = prize, model.PrizeNames[prize]
			break
		}
	}
	return
}

func getMyPrizeCode() int {
	seed := time.Now().UnixNano() // 第一步，抽奖，根据随机数匹配奖品
	r := rand.New(rand.NewSource(seed))
	return r.Intn(model.BaseNum) // 得到个人的抽奖编码
}

func getLockId() string {
	uid, _ := uuid.NewUUID()
	return uid.String()
}
