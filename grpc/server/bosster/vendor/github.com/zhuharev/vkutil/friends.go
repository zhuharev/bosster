package vkutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/zhuharev/vk"
)

func (api *Api) FriendsGetIds(args ...url.Values) ([]int, int, error) {
	resp, err := api.VkApi.Request(vk.METHOD_FRIENDS_GET, args...)
	if err != nil {
		return nil, 0, err
	}
	var r ResponseIds
	err = json.Unmarshal(resp, &r)
	if err != nil {
		return nil, 0, err
	}
	return r.Resp.Items, r.Resp.Count, nil
}

func (api *Api) FriendsGet(args ...url.Values) ([]User, int, error) {
	arg := url.Values{"fields": {"id"}}
	if len(args) > 0 {
		arg = args[0]
		if arg.Get("fields") == "" {
			arg.Set("fields", "id")
		}
	}

	resp, err := api.VkApi.Request(vk.METHOD_FRIENDS_GET, arg)
	if err != nil {
		return nil, 0, err
	}
	var r ResponseUserWithCount
	err = json.Unmarshal(resp, &r)
	if err != nil {
		return nil, 0, err
	}
	return r.Response.Items, r.Response.Count, nil
}

func (api *Api) FriendsGetRequests(args ...url.Values) ([]int, error) {
	resp, err := api.VkApi.Request(vk.METHOD_FRIENDS_GET_REQUESTS, args...)
	if err != nil {
		return nil, err
	}
	var r ResponseIds
	err = json.Unmarshal(resp, &r)
	if err != nil {
		return nil, err
	}
	return r.Resp.Items, nil
}

func (api *Api) FriendsGetAllFollowers() ([]int, error) {
	ids, _, err := api.FriendsGetIds()
	if err != nil {
		return nil, err
	}
	followers, err := api.UsersGetFollowers()
	if err != nil {
		return nil, err
	}
	followers = append(followers, ids...)
	return followers, nil
}

func (api *Api) FriendsGetAllFollowing(withFriendss ...bool) ([]int, error) {
	code := `var a = API.friends.getRequests({"out":1,count:1000});
var b = a.items;
var cnt = 0;
while(b.length<a.count){
cnt=cnt+1000;
b=b+API.friends.getRequests({"out":1,count:1000,offset:cnt}).items;
}
return b;`
	resp, err := api.Execute(code)
	var r struct {
		Response []int `json:"response"`
	}
	err = json.Unmarshal(resp, &r)
	if err != nil {
		fmt.Println(string(resp))
		return nil, err
	}
	if withFriendss != nil && withFriendss[0] {
		ids, _, err := api.FriendsGetIds()
		if err != nil {
			return nil, err
		}
		r.Response = append(r.Response, ids...)
	}
	return r.Response, nil
}

type ResponseAreFriends struct {
	Response []AreFriends `json:"response"`
	ResponseError
}

type AreFriends struct {
	Id           int    `json:"user_id"`
	FriendStatus int    `json:"friend_status"`
	Sign         string `json:"sign"`
}

func (api *Api) FriendsAreFriends(ids []int,
	args ...url.Values) ([]AreFriends, error) {
	resp, err := api.VkApi.Request(vk.METHOD_FRIENDS_ARE_FRIENDS, url.Values{
		"user_ids": {strings.Join(arrIntToStr(ids), ",")},
	})
	if err != nil {
		return nil, err
	}
	var r ResponseAreFriends
	err = json.Unmarshal(resp, &r)
	if err != nil {
		return nil, err
	}
	return r.Response, nil
}

//TODO
func (api *Api) FriendsGetSuggestions() {

}

func (api *Api) FriendsAdd(userId int, args ...url.Values) (int, error) {
	resp, err := api.VkApi.Request(vk.METHOD_FRIENDS_ADD, url.Values{
		"user_id": {fmt.Sprint(userId)},
	})
	if err != nil {
		return 0, err
	}
	var r ResponseInt
	err = json.Unmarshal(resp, &r)
	if err != nil {
		return 0, err
	}
	if r.Error.Code != 0 {
		return 0, errors.New(r.Error.Msg)
	}
	return r.Response, nil
}
func (api *Api) FriendsDelete(userId int, args ...url.Values) error {
	resp, err := api.VkApi.Request(vk.METHOD_FRIENDS_DELETE, url.Values{
		"user_id": {fmt.Sprint(userId)},
	})
	if err != nil {
		return err
	}
	var r OutRequestDeletedResponse
	fmt.Printf("%s", resp)
	err = json.Unmarshal(resp, &r)
	if err != nil {
		return err
	}
	if r.Error.Code != 0 {
		return errors.New(r.Error.Msg)
	}
	if r.Success != 1 {
		return errors.New("Some error on friends delete")
	}
	return nil
}

func (api *Api) FriendsGetMutual(sourceId int, targetUids []int, args ...url.Values) ([]RespFriendsGetMutual, error) {
	resp, err := api.VkApi.Request(vk.METHOD_FRIENDS_GET_MUTUAL, url.Values{
		"user_id":     {fmt.Sprint(sourceId)},
		"target_uids": {strings.Join(arrIntToStr(targetUids), ",")},
	})
	if err != nil {
		return nil, err
	}
	var r ResponseFriendsGetMutual
	err = json.Unmarshal(resp, &r)
	if err != nil {
		return nil, err
	}
	if r.Error.Code != 0 {
		return nil, errors.New(r.Error.Msg)
	}
	return r.Response, nil
}

//utils

func (api *Api) UtilsFriendsGetOne(userId int) (friends []int, total int, e error) {
	vals := url.Values{
		"user_id": {fmt.Sprint(userId)},
		"count":   {"5000"},
	}
	friends, total, e = api.FriendsGetIds(vals)
	if e != nil {
		return
	}
	if total > 5000 {
		vals.Set("offset", "5000")
	}
	f, _, e := api.FriendsGetIds(vals)
	if e != nil {
		return
	}
	friends = append(friends, f...)

	return
}

//todo handle it execute
func (api *Api) UtilsFriendsOfFriends(targetId int) (friendsOfFriends []int, err error) {
	friends, _, err := api.FriendsGetIds(url.Values{"user_id": {
		fmt.Sprint(targetId),
	}})
	if err != nil {
		return nil, err
	}
	/*
		for _, v := range friends {
			friendsOfFriend, err := api.FriendsGet(url.Values{"user_id": {
				fmt.Sprint(v),
			}})
			if err != nil {
				return nil, err
			}
			friendsOfFriends = append(friendsOfFriends, friendsOfFriend...)
		}*/
	m, e := api.UtilsFriendsGet(friends...)
	if e != nil {
		return nil, e
	}

	var dup = map[int]struct{}{}

	for _, arr := range m {
		for _, id := range arr {
			if _, ok := dup[id]; !ok {
				friendsOfFriends = append(friendsOfFriends, id)
				dup[id] = struct{}{}
			}
		}

	}
	return
}

func (api *Api) UtilsFriendsGet(ids ...int) (m map[int][]int, e error) {
	m = map[int][]int{}
	rch, done, _ := api.GoUtilsFriendsGet(ids...)
	for {
		select {
		case nm := <-rch:
			for k, v := range nm {
				m[k] = v
			}
		case <-done:
			return m, nil
		}
	}
}

func (api *Api) GoUtilsFriendsGet(ids ...int) (m chan map[int][]int, done chan struct{}, chError chan error) {
	var (
		iters  = 0
		offset = 0
		count  = 25
		dones  = make(chan struct{})
		tcode  = `var b = [%s];
var i = %d;
var a = [API.friends.get({user_id:b[i],limit:5000}).items];
i = i+1;
while(i<%d){
a = a+[API.friends.get({user_id:b[i],limit:5000}).items];
i = i + 1;
};
return a;`
	)

	m = make(chan map[int][]int)
	done = make(chan struct{})
	chError = make(chan error)

	gf := func(a []int, ch chan map[int][]int, count int, counter int) {
		code := fmt.Sprintf(tcode, strings.Join(arrIntToStr(a), ","),
			0, count)

		resp, err := api.Execute(code)
		if err != nil {
			chError <- err
			dones <- struct{}{}
			return
		}
		var r struct {
			Response [][]int `json:"response"`
			ResponseError
		}
		if DEBUG {
			log.Println(string(resp))
		}
		err = json.Unmarshal(resp, &r)
		if err != nil {
			chError <- err
			dones <- struct{}{}
			return
		}
		if r.ResponseError.Error.Code != 0 {
			chError <- errors.New(r.ResponseError.Error.Msg)
			dones <- struct{}{}
			return
		}
		m := map[int][]int{}
		for i, j := range r.Response {
			m[a[i]] = j
		}

		ch <- m
		dones <- struct{}{}

	}

	tcount := count
	for i := 0; i < len(ids); i = i + tcount {
		iters++
		if len(ids)-i <= tcount {
			tcount = len(ids) - i
		}
	}

	go func() {
		for i := 0; i < len(ids); i = i + count {

			if len(ids)-i <= count {
				count = len(ids) - i
			}

			gf(ids[offset:count+offset], m, count, i)

			offset += count
		}
	}()

	go func() {
		for i := 0; i < iters; i++ {
			<-dones
			if DEBUG {
				log.Printf("done %d/%d", i, iters)
			}
		}
		done <- struct{}{}
	}()

	return
}

func (api *Api) UtilsGetMutual(sourceId int, ids ...int) (result map[int]int, e error) {
	result = map[int]int{}
	var tmpArr []int
	for i, v := range ids {
		if !inArr(v, tmpArr) && sourceId != v {
			tmpArr = append(tmpArr, v)
		}
		if len(tmpArr) == 1000 || i+1 == len(ids) {
			arr, e := api.utilsGetThousandMutual(sourceId, tmpArr...)
			if e != nil {
				return nil, e
			}
			for k, v := range arr {
				result[k] = v
			}
			tmpArr = nil
		}
	}
	return
}

func (api *Api) utilsGetThousandMutual(sourceId int, ids ...int) (result map[int]int, e error) {
	res := []int{}
	codefmt := `var a = API.friends.getMutual({source_uid:%d,target_uids:[%s]})@.common_count;
return a;`
	code := fmt.Sprintf(codefmt, sourceId, strings.Join(arrIntToStr(ids), ","))
	resp, err := api.Execute(code)
	var r struct {
		Response []int `json:"response"`
	}
	err = json.Unmarshal(resp, &r)
	if err != nil {
		fmt.Println(string(resp))
		return nil, err
	}
	//arr := arrStrToInt(strings.Split(r.Response, ","))
	for _, j := range r.Response {
		res = append(res, j)
	}
	if len(res) != len(ids) {
		return nil, fmt.Errorf("%d != %d", len(res), len(ids))
	}
	result = map[int]int{}
	for i, v := range ids {
		result[v] = res[i]
	}
	return
}

/*func (api *Api) UtilsFriendsGetFull(ids ...int) (res []User, error) {
	var tcode = `var b = [%s];
var i = %d;
var a = API.friends.get({user_id:b[i],fields:"online,bdate,education"}).items;
i = i+1;
while(i<%d){
a = a+","+API.friends.get({user_id:b[i],fields:"online,bdate,education"}).items;
i = i + 1;
};
return a;`

	offset := 0
	count := 25
	for i := 0; i < len(ids); i = i + count {
		resp, err := api.Execute(fmt.Sprintf(tcode, strings.Join(arrIntToStr(ids), ","),
			offset, offset+count))
		var r struct {
			Response string `response`
		}
		err = json.Unmarshal(resp, &r)
		if err != nil {
			fmt.Println(string(resp))
			return nil, err
		}
		arr := arrStrToInt(strings.Split(r.Response, ","))
		for _, j := range arr {
			res = append(res, j)
		}
	}
	return
}*/

func inArr(exp int, arr []int) bool {
	for _, v := range arr {
		if v == exp {
			return true
		}
	}
	return false
}
