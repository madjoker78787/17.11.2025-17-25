package models

import (
	"sync"
)

type DataResponse struct {
	Links []string `json:"links"`
}

type DataRequest struct {
	Links    map[string]string `json:"links"`
	LinksNum int               `json:"links_num"`
}

type Cache struct {
	Mu   sync.Mutex
	List map[int][]string `json:"list"`
}

func NewCache() *Cache {
	cache := Cache{
		List: make(map[int][]string),
	}
	return &cache
}

func (c *Cache) Set(id int, link []string) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.List[id] = link
}

func (c *Cache) Get() map[int][]string {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	return c.List
}

//

type Storage struct {
	Cached      *Cache
	Mu          sync.Mutex
	NextID      int
	StorageList map[int]Item
}

type Item struct {
	Links map[string]string
}

func NewStorage(c *Cache) *Storage {
	data := Storage{
		NextID:      0,
		Cached:      c,
		StorageList: make(map[int]Item),
	}
	return &data
}

func (s *Storage) Set(link Item) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.StorageList[s.NextID] = link

}

func (s *Storage) SetNextID() {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.NextID++
}

//func (s *Storage) Get(id int) (*Item, bool) {
//	value, ok := s.StorageList[id]
//	if !ok {
//		return nil, false
//	}
//	return &value, false
//}

//func (s *Storage) GetLastID() int {
//	s.Mu.Lock()
//	defer s.Mu.Unlock()
//
//	keys := make([]int, 0, len(s.StorageList))
//	for k, _ := range s.StorageList {
//		keys = append(keys, k)
//	}
//	sort.Ints(keys)
//
//	return len(keys)
//}
