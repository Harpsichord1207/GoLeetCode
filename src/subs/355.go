package subs

import (
	"fmt"
	"sort"
)

type Twitter struct {
	follows  map[int]map[int]bool
	twitters map[int][]int
	idMap    map[int]int
	id       int
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	t := Twitter{}
	t.follows = make(map[int]map[int]bool)
	t.twitters = make(map[int][]int)
	t.idMap = make(map[int]int)
	t.id = 0
	return t
}

/** Compose a new tweet. */
func (t *Twitter) PostTweet(userId int, tweetId int) {
	twitters := t.twitters[userId]
	twitters = append(twitters, tweetId)
	t.twitters[userId] = twitters
	t.id++
	t.idMap[tweetId] = t.id
}

func mergeSlice(t *Twitter, slice1 []int, slice2 []int) []int {
	slice := append(slice1, slice2...)
	sort.Slice(slice, func(i, j int) bool {
		return t.idMap[slice[i]] > t.idMap[slice[j]]
	})
	if len(slice) > 10 {
		return slice[:10]
	}
	return slice
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (t *Twitter) GetNewsFeed(userId int) []int {
	users := []int{userId}

	follows := t.follows[userId]
	for k, v := range follows {
		if v {
			users = append(users, k)
		}
	}
	var feed []int

	for _, id := range users {
		twitters := t.twitters[id]
		feed = mergeSlice(t, feed, twitters)
	}
	return feed
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Follow(followerId int, followeeId int) {
	follows := t.follows[followerId]
	if follows == nil {
		follows = make(map[int]bool)
	}
	follows[followeeId] = true
	t.follows[followerId] = follows
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Unfollow(followerId int, followeeId int) {
	follows := t.follows[followerId]
	if follows == nil {
		follows = make(map[int]bool)
	}
	delete(follows, followeeId)
	t.follows[followerId] = follows
}

func Test355() {
	t := Constructor()
	t.PostTweet(2, 5)
	t.PostTweet(1, 3)
	t.PostTweet(1, 101)
	t.PostTweet(2, 13)
	t.PostTweet(2, 10)
	t.PostTweet(1, 2)
	t.PostTweet(2, 94)
	t.PostTweet(2, 505)
	t.PostTweet(1, 333)
	t.PostTweet(1, 22)
	fmt.Println(t.GetNewsFeed(2))
	t.Follow(2, 1)
	fmt.Println(t.GetNewsFeed(2))
}
