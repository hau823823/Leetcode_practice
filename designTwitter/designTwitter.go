package designTwitter

import "container/heap"

type Tweet struct {
	id   int     // 推文的唯一标识符
	at   int     // 推文发布的时间戳
	next *Tweet  // 指向同一用户的上一条推文的指针，形成一个单向链表
	// 推文存储方式：使用单链表按时间逆序存储每个用户的推文，使得每次操作都能快速访问最新的推文。
}

type Twitter struct {
	follows map[int]map[int]bool // 追踪用户间的关注关系
	tweets  map[int]*Tweet       // 每个用户的推文链表头结点
	time    int                  // 模拟全局时间戳
}

// 实现了 heap.Interface，允许将推文按时间戳顺序存储在一个最大堆中。这样，最新的推文总是在堆的顶部。
type TweetHeap []*Tweet

func (h TweetHeap) Len() int           { return len(h) }
func (h TweetHeap) Less(i, j int) bool { return h[i].at > h[j].at }
func (h TweetHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *TweetHeap) Push(x any)        { *h = append(*h, x.(*Tweet)) }
func (h *TweetHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

// 初始化 Twitter 对象
func Constructor() Twitter {
	return Twitter{map[int]map[int]bool{}, map[int]*Tweet{}, 0}
}

// 用户发布推文时，将新推文添加到该用户的推文链表头部，并更新时间戳
func (t *Twitter) PostTweet(userId int, tweetId int) {
	t.tweets[userId] = &Tweet{tweetId, t.time, t.tweets[userId]}
	t.time++
}

// 为指定用户获取新闻推送，即获取该用户和他关注的人最近的推文
func (t *Twitter) GetNewsFeed(userId int) []int {
	tweets := TweetHeap{}
	heap.Init(&tweets)
	if t.tweets[userId] != nil {
		heap.Push(&tweets, t.tweets[userId])
	}
	for following := range t.follows[userId] {
		if t.tweets[following] != nil {
			heap.Push(&tweets, t.tweets[following])
		}
	}

	feed := []int{}
	for len(feed) < 10 && len(tweets) > 0 {
		mostRecent := heap.Pop(&tweets).(*Tweet)
		if mostRecent.next != nil {
			heap.Push(&tweets, mostRecent.next)
		}
		feed = append(feed, mostRecent.id)
	}
	return feed
}

// 添加关注关系
func (t *Twitter) Follow(followerId int, followeeId int) {
	if t.follows[followerId] == nil {
		t.follows[followerId] = make(map[int]bool)
	}
	t.follows[followerId][followeeId] = true
}

// 删除关注关系
func (t *Twitter) Unfollow(followerId int, followeeId int) {
	if t.follows[followerId] == nil {
		t.follows[followerId] = make(map[int]bool)
	}
	delete(t.follows[followerId], followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
