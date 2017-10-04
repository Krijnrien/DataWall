package cache

import (
	"sync"
	"time"
)

// TODO Add class wide inner func code comments

/** Page struct
 * Stores content of page along with expiration time.
 */
type Page struct {
	Content    []byte // Content of page serialized into byte array.
	Expiration int64  // Time expiration measured. Must be int64 as time.Now().UnixNano() requires int64 to compare with.
}

/** Expired
 * Checks if the cached page has expired
 * @Param Page:
 * @return bool: true if cache has expired, otherwise false.
 */
func (page Page) Expired() bool {
	if page.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > page.Expiration
}

/** MemCache struct
 * Stores map of Pages struct and mutex to prevent reading & writing conflicts.
 * As this cache type is a struct this represent memory cache since struct is kept in RAM.
 */
type MemCache struct {
	pages map[string]Page
	mutex *sync.RWMutex
}

//NewMemCache creates a new in memory storage
/** NewMemCache
 * Sets type of cache
 * @param
 * @return pointer of new initialized MemCache struct.
 */
func NewMemCache() *MemCache {
	return &MemCache{
		pages: make(map[string]Page),
		mutex: &sync.RWMutex{},
	}
}

/** Get
 * Gets the content of cache with the request endpoint url as key.
 * @param key: Endpoint url as key
 * @return []Byte: Content of page in byte array representation
 */
func (memCache MemCache) Get(key string) []byte {
	memCache.mutex.RLock()
	defer memCache.mutex.RUnlock()

	item := memCache.pages[key]
	if item.Expired() {
		delete(memCache.pages, key)
		return nil
	}
	return item.Content
}

/** Set
 * Sets the cache with endpoint url as key, page content as byte array and duration of cache given in seconds.
 * @param key: Endpoint url as key.
 * @param content: page content serialized into byte array.
 * @param duration: Time for cache to expire
 */
func (memCache MemCache) Set(key string, content []byte, duration time.Duration) {
	// Lock
	memCache.mutex.Lock()
	defer memCache.mutex.Unlock()

	memCache.pages[key] = Page{
		Content:    content,
		Expiration: time.Now().Add(duration).UnixNano(),
	}
}
