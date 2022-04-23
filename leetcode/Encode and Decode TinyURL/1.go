type Codec struct {
	urlMap map[string]string
}

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	prefix      = "http://tinyurl.com/"
)

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Constructor() Codec {
	rand.Seed(time.Now().UnixNano())
	return Codec{
		urlMap: make(map[string]string),
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	var shortURL string
	for {
		shortURL = RandStringRunes(5)
		_, ok := this.urlMap[shortURL]
		if !ok {
			break
		}
	}
	this.urlMap[shortURL] = longUrl
	return prefix + shortURL
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return this.urlMap[shortUrl[len(shortUrl)-5:]]
}


/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
