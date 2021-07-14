package nekosAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func DoRequest(target string) (res *Response, err error) {
	reqURL := fmt.Sprintf("%s/%s", baseUrl, target)
	resp, err := http.Get(reqURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := new(Response)
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

var EndPoints = []string{"solog", "smug", "feet", "smallboobs", "lewdkemo", "woof", "gasm", "solo", "8ball", "goose", "cuddle", "avatar", "cum", "slap", "les", "v3", "erokemo", "bj", "pwankg", "nekoapi_v3.1", "ero", "hololewd", "pat", "gecg", "holo", "poke", "feed", "fox_girl", "tits", "nsfw_neko_gif", "eroyuri", "holoero", "pussy", "Random_hentai_gif", "lizard", "yuri", "keta", "neko", "hentai", "feetg", "eron", "erok", "baka", "kemonomimi", "hug", "cum_jpg", "nsfw_avatar", "erofeet", "meow", "kiss", "wallpaper", "tickle", "blowjob", "spank", "kuni", "classic", "waifu", "femdom", "boobs", "trap", "lewd", "pussy_jpg", "anal", "futanari", "ngif", "lewdk"}
