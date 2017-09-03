package anaconda

import (
	"net/url"
	"strconv"
)

// CreateList implements /lists/create.json
func (a TwitterApi) CreateList(name, description string, v url.Values) (list List, err error) {
	v = cleanValues(v)
	v.Set("name", name)
	v.Set("description", description)

	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/lists/create.json", v, &list, _POST, response_ch}
	return list, (<-response_ch).err
}

func (a TwitterApi) ShowList(name, description string, v url.Values) (list List, err error) {
	if v == nil {
		v = url.Values{}
	}
	v.Set("name", name)
	v.Set("description", description)

	response_ch := make(chan response)
	a.queryQueue <- query{BaseUrl + "/lists/create.json", v, &list, _POST, response_ch}
	return list, (<-response_ch).err
}

// AddUserToList implements /lists/members/create.json
func (a TwitterApi) AddListUser(userID int64, listID int64, v url.Values) (users []User, err error) {
//func (a TwitterApi) AddUserToList(screenName string, listID int64, v url.Values) (users []User, err error) {
	if v == nil {
		v = url.Values{}
	}
	v.Set("list_id", strconv.FormatInt(listID, 10))
	v.Set("user_id", strconv.FormatInt(userID, 10))

	var addUserToListResponse AddUserToListResponse

	response_ch := make(chan response)
	a.queryQueue <- query{BaseUrl + "/lists/members/create.json", v, &addUserToListResponse, _POST, response_ch}
	return addUserToListResponse.Users, (<-response_ch).err
}

func (a TwitterApi) RemoveListUser(usrID int64, listID int64, v url.Values) (users []User, err error) {
	if v == nil {
		v = url.Values{}
	}
	v.Set("list_id", strconv.FormatInt(listID, 10))
	v.Set("user_id", strconv.FormatInt(usrID, 10))

	var addUserToListResponse AddUserToListResponse

	response_ch := make(chan response)
	a.queryQueue <- query{BaseUrl + "/lists/members/destroy.json", v, &addUserToListResponse, _POST, response_ch}
	return addUserToListResponse.Users, (<-response_ch).err
}

// GetListsOwnedBy implements /lists/ownerships.json
// screen_name, count, and cursor are all optional values
func (a TwitterApi) GetListsOwnedBy(userID int64, v url.Values) (lists []List, err error) {
	v = cleanValues(v)
	v.Set("user_id", strconv.FormatInt(userID, 10))

	var listResponse ListResponse

	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/lists/ownerships.json", v, &listResponse, _GET, response_ch}
	return listResponse.Lists, (<-response_ch).err
}

func (a TwitterApi) GetListTweets(listID int64, includeRTs bool, v url.Values) (tweets []Tweet, err error) {
	v = cleanValues(v)
	v.Set("list_id", strconv.FormatInt(listID, 10))
	v.Set("include_rts", strconv.FormatBool(includeRTs))

	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/lists/statuses.json", v, &tweets, _GET, response_ch}
	return tweets, (<-response_ch).err
}

func (a TwitterApi) GetListUsers(v url.Values) (c UserCursor, err error) {
	err = a.apiGet(BaseUrl+"/lists/members.json", v, &c)
	return
}