package geoip

import (
	//_ "embed" //diy
	"os" //diy
	"errors"
	"net"
	"strings"
	"sync"

	maxminddb "github.com/oschwald/maxminddb-golang"
)

/*//go:embed geoip.db diy*/
//var db []byte

var (
	/*dbOnce = sync.OnceValues(func() (*maxminddb.Reader, error) {
		db, err := maxminddb.FromBytes(db)
		if err != nil {
			return nil, err
		}
		return db, nil
	})//diy*/
	// Init 初始化 GeoIP 数据库
	// 优先级:
	// 1 flag
	// 2 env
	// 3 config.yml
	// 4 默认 ./data/geoip.db
	
	reader     *maxminddb.Reader //diy
	readerErr  error //diy
	readerOnce sync.Once //diy
)
//diy
func Init(path string) error {

	if path == "" {
		path = "./data/geoip.db"
	}

	readerOnce.Do(func() {

		if _, err := os.Stat(path); err != nil {
			readerErr = err
			return
		}

		reader, readerErr = maxminddb.Open(path)

	})

	return readerErr
}
//diy Close 释放资源（可选）
func Close() {
	if reader != nil {
		reader.Close()
	}
}
type IPInfo struct {
	/*Country       string `maxminddb:"country"`
	CountryName   string `maxminddb:"country_name"`
	Continent     string `maxminddb:"continent"`
	ContinentName string `maxminddb:"continent_name"`*/
	Country struct { //diy 兼容 GeoLite2
		ISOCode string `maxminddb:"iso_code"`
	} `maxminddb:"country"`

	Continent struct {
		Code string `maxminddb:"code"`
	} `maxminddb:"continent"`
}

func Lookup(ip net.IP) (string, error) {
	/*db, err := dbOnce()
	if err != nil {
		return "", err
	}diy*/
	if reader == nil {
		return "", errors.New("geoip database not initialized")
	}

	var record IPInfo
	//err = db.Lookup(ip, &record) //diy
	err := reader.Lookup(ip, &record)
	if err != nil {
		return "", err
	}

	/*if record.Country != "" {
		return strings.ToLower(record.Country), nil
	} else if record.Continent != "" {
		return strings.ToLower(record.Continent), nil
	}diy*/
	if record.Country.ISOCode != "" { //diy 兼容 GeoLite2
		return strings.ToLower(record.Country.ISOCode), nil
	}
	if record.Continent.Code != "" {
		return strings.ToLower(record.Continent.Code), nil
	}

	return "", errors.New("IP not found")
}
