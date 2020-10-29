package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	//"sync"

	"lottery_backend/src/access/model"
)

func LogIn(phoneNumber uint64) (string, error) {
	data := model.LogInRequest{}
	data.Action = "LogIn"
	data.PhoneNumber = phoneNumber
	data.VerifyCode = 1

	b, err := json.Marshal(&data)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	req, err := http.NewRequest("POST", "http://106.75.224.81:8888", bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var response model.LogInResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}
	if response.RetCode != 0 {
		return "", errors.New(response.Message)
	}

	return response.UserId, nil
}

func SubmitArticle(userId string) error {
	data := model.SubmitArticleRequest{}
	data.Action = "SubmitArticle"
	data.UserId = userId
	data.Text = "good good study, day day up!"

	b, err := json.Marshal(&data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req, err := http.NewRequest("POST", "http://106.75.224.81:8888", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var response model.SubmitArticleResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}
	if response.RetCode != 0 {
		return errors.New(response.Message)
	}

	return nil
}

func SubOneDayForRecord() error {
	data := model.SubOneDayForRecordRequest{}
	data.Action = "SubOneDayForRecord"

	b, err := json.Marshal(&data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req, err := http.NewRequest("POST", "http://106.75.224.81:8888", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var response model.SubOneDayForRecordResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}
	if response.RetCode != 0 {
		return errors.New(response.Message)
	}

	return nil
}

func Lottery(phoneNumber uint64) error {
	data := model.LotteryRequest{}
	data.Action = "Lottery"
	//data.UserId = userId
	data.PhoneNumber = phoneNumber

	b, err := json.Marshal(&data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req, err := http.NewRequest("POST", "http://106.75.224.81:8888", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var response model.LogInResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}
	fmt.Println(response)
	if response.RetCode != 0 {
		return errors.New(response.Message)
	}
	return nil
}

var first = flag.Bool("first", false, "first")

func main() {
	flag.Parse()

	if *first {
		// test-1 and test-2 only need first
		for i := 1; i <= 1000; i++ {
			// test-1: user log in
			userId, err := LogIn(uint64(i))
			if err != nil {
				fmt.Println(err)
				return
			}
			// test-2: submit article
			err = SubmitArticle(userId)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

	// 3 days
	for i := 0; i < 3; i++ {
		// test-3: lottery
		err := SubOneDayForRecord()
		if err != nil {
			fmt.Println(err)
		}

		//wg := sync.WaitGroup{} // too many competition in redis
		//for j := 1; j <= 1000; j++ {
		//	wg.Add(1)
		//	go func(number uint64) {
		//		err := Lottery(number)
		//		if err != nil {
		//			fmt.Println(err)
		//		}
		//		wg.Done()
		//	}(uint64(j))
		//}
		//wg.Wait()

		for j := 1; j <= 1000; j++ {
			err := Lottery(uint64(j))
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
