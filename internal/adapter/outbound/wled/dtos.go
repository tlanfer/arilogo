package wled

type Info struct {
	Ver  string `json:"ver"`
	Vid  int    `json:"vid"`
	Leds struct {
		Count  int   `json:"count"`
		Pwr    int   `json:"pwr"`
		Fps    int   `json:"fps"`
		Maxpwr int   `json:"maxpwr"`
		Maxseg int   `json:"maxseg"`
		Seglc  []int `json:"seglc"`
		Lc     int   `json:"lc"`
		Rgbw   bool  `json:"rgbw"`
		Wv     int   `json:"wv"`
		Cct    int   `json:"cct"`
	} `json:"leds"`
	Str       bool   `json:"str"`
	Name      string `json:"name"`
	Udpport   int    `json:"udpport"`
	Live      bool   `json:"live"`
	Liveseg   int    `json:"liveseg"`
	Lm        string `json:"lm"`
	Lip       string `json:"lip"`
	Ws        int    `json:"ws"`
	Fxcount   int    `json:"fxcount"`
	Palcount  int    `json:"palcount"`
	Cpalcount int    `json:"cpalcount"`
	Maps      []struct {
		Id int `json:"id"`
	} `json:"maps"`
	Wifi struct {
		Bssid   string `json:"bssid"`
		Rssi    int    `json:"rssi"`
		Signal  int    `json:"signal"`
		Channel int    `json:"channel"`
	} `json:"wifi"`
	Fs struct {
		U   int `json:"u"`
		T   int `json:"t"`
		Pmt int `json:"pmt"`
	} `json:"fs"`
	Ndc      int    `json:"ndc"`
	Arch     string `json:"arch"`
	Core     string `json:"core"`
	Lwip     int    `json:"lwip"`
	Freeheap int    `json:"freeheap"`
	Uptime   int    `json:"uptime"`
	Time     string `json:"time"`
	Opt      int    `json:"opt"`
	Brand    string `json:"brand"`
	Product  string `json:"product"`
	Mac      string `json:"mac"`
	Ip       string `json:"ip"`
}
