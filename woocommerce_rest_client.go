package woocommerce_rest_client

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

var wpBaseUrl string
var authorization string

func init() {

	consumerKey := os.Getenv("WC_CONSUMER_KEY")
	consumerSecret := os.Getenv("WC_CONSUMER_SECRET")
	wpBaseUrl = os.Getenv("WP_BASE_URL")

	apiToken := consumerKey + ":" + consumerSecret
	authorization = "Basic " + base64.StdEncoding.EncodeToString([]byte(apiToken))

}

type WooCommerceProduct struct {
	ProductId uint
	Permalink string
	Name      string
	Image     string
}

func GetProduct(productId uint) (*WooCommerceProduct, error) {

	url := wpBaseUrl + "/wp-json/wc/v2/products/" + strconv.Itoa(int(productId))
	method := "GET"

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", authorization)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))

	var bodyMap map[string]interface{}
	err = json.Unmarshal(body, &bodyMap)

	if err != nil {
		return nil, err
	}

	var result WooCommerceProduct
	result.ProductId = uint(bodyMap["id"].(float64))
	result.Permalink = bodyMap["permalink"].(string)
	result.Name = bodyMap["name"].(string)

	images := bodyMap["images"].([]interface{})
	imageSrc := images[0].(map[string]interface{})["src"].(string)
	result.Image = imageSrc

	return &result, nil

}
