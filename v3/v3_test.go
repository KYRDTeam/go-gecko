package coingecko

import (
	"context"
	"io/ioutil"
	"os"
	"testing"

	gock "gopkg.in/h2non/gock.v1"
)

func init() {
	defer gock.Off()
}

var c = NewClient(nil, "")
var mockURL = "https://api.coingecko.com/api/v3"

func TestPing(t *testing.T) {
	err := setupGock("json/ping.json", "/ping")
	ping, err := c.Ping(context.Background())
	if err != nil {
		t.FailNow()
	}
	if ping.GeckoSays != "(V3) To the Moon!" {
		t.FailNow()
	}
}

func TestSimpleSinglePrice(t *testing.T) {
	err := setupGock("json/simple_single_price.json", "/simple/price")
	if err != nil {
		t.FailNow()
	}
	simplePrice, err := c.SimpleSinglePrice(context.Background(), "bitcoin", "usd", true, true)
	if err != nil {
		t.FailNow()
	}
	if simplePrice.ID != "bitcoin" || simplePrice.Currency != "usd" || simplePrice.MarketPrice != float32(5013.61) {
		t.FailNow()
	}
}

func TestSimplePrice(t *testing.T) {
	err := setupGock("json/simple_price.json", "/simple/price")
	if err != nil {
		t.FailNow()
	}
	ids := []string{"bitcoin", "ethereum"}
	vc := []string{"usd", "myr"}
	sp, err := c.SimplePrice(context.Background(), ids, vc, true, true)
	if err != nil {
		t.FailNow()
	}
	bitcoin := (*sp)["bitcoin"]
	eth := (*sp)["ethereum"]

	if bitcoin["usd"] != 5005.73 || bitcoin["myr"] != 20474 {
		t.FailNow()
	}
	if eth["usd"] != 163.58 || eth["myr"] != 669.07 {
		t.FailNow()
	}
}

func TestSimpleSupportedVSCurrencies(t *testing.T) {
	err := setupGock("json/simple_supported_vs_currencies.json", "/simple/supported_vs_currencies")
	s, err := c.SimpleSupportedVSCurrencies(context.Background())
	if err != nil {
		t.FailNow()
	}
	if len(*s) != 54 {
		t.FailNow()
	}
}

func TestCoinsList(t *testing.T) {
	err := setupGock("json/coins_list.json", "/coins/list")
	list, err := c.CoinsList(context.Background(), false)
	if err != nil {
		t.FailNow()
	}
	item := (*list)[0]
	if item.ID != "01coin" {
		t.FailNow()
	}
}

func TestCoinsID(t *testing.T) {
	err := setupGock("json/coins_id.json", "/coins/tether")
	coins, err := c.CoinsID(context.Background(), "tether", false, false, false, false, false, false)
	if err != nil {
		t.FailNow()
	}
	if len(coins.Platforms) == 0 {
		t.FailNow()
	}
	if len(coins.DetailPlatformsItem) == 0 {
		t.FailNow()
	}
}

func TestAssetPlatform(t *testing.T) {
	err := setupGock("json/asset_platforms.json", "/asset_platforms")
	platforms, err := c.AssetPlatforms(context.Background())
	if err != nil {
		t.FailNow()
	}
	if len(platforms) != 148 {
		t.FailNow()
	}
}

// Util: Setup Gock
func setupGock(filename string, url string) error {
	testJSON, err := os.Open(filename)
	defer testJSON.Close()
	if err != nil {
		return err
	}
	testByte, err := ioutil.ReadAll(testJSON)
	if err != nil {
		return err
	}
	gock.New(mockURL).
		Get(url).
		Reply(200).
		JSON(testByte)

	return nil
}
