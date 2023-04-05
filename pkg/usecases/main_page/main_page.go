package main_page

type Output struct {
	HotSales       hotSales       `json:"hot_sales"`
	SeasonAd       seasonAd       `json:"season_ad"`
	LimitedProduct limitedProduct `json:"limited_product"`
}

type hotSales struct {
	FirstPic  string `json:"first_pic"`
	SecondPic string `json:"second_pic"`
}

type seasonAd struct {
	FirstPic  string `json:"first_pic"`
	SecondPic string `json:"second_pic"`
	ThirdPic  string `json:"third_pic"`
	FourPic   string `json:"four_pic"`

	FirstText  string `json:"first_text"`
	SecondText string `json:"second_text"`
	ThirdText  string `json:"third_text"`
	FourText   string `json:"four_text"`
}

type limitedProduct struct {
	FirstPic  string `json:"first_pic"`
	SecondPic string `json:"second_pic"`
	ThirdPic  string `json:"third_pic"`
	FourPic   string `json:"four_pic"`

	FirstText  string `json:"first_text"`
	SecondText string `json:"second_text"`
	ThirdText  string `json:"third_text"`
	FourText   string `json:"four_text"`
}

// 返回主頁需要的資訊
func Exec() (Output, error) {

	var output Output

	output.HotSales.FirstPic = "p30babL.png"
	output.HotSales.SecondPic = "OoBzmPm.jpg"

	output.SeasonAd.FirstPic = "6OydL5n.jpg"
	output.SeasonAd.SecondPic = "XGkBM7R.jpg"
	output.SeasonAd.ThirdPic = "wzPOMHe.jpg"
	output.SeasonAd.FourPic = "K6Ec069.jpg"

	output.SeasonAd.FirstText = "下殺 $ 12,999"
	output.SeasonAd.SecondText = "特價 $ 16,999"
	output.SeasonAd.ThirdText = "出清 $ 21,000"
	output.SeasonAd.FourText = "史無前例 $ 300"

	output.LimitedProduct.FirstPic = "h2o29dy.jpg"
	output.LimitedProduct.SecondPic = "hpWHcv2.jpg"
	output.LimitedProduct.ThirdPic = "51Vatnn.jpg"
	output.LimitedProduct.FourPic = "CKPuzhJ.jpg"

	output.LimitedProduct.FirstText = "限時下殺 $ 299"
	output.LimitedProduct.SecondText = "限時特價 $ 1,130"
	output.LimitedProduct.ThirdText = "限時出清 $ 79,000"
	output.LimitedProduct.FourText = "限時出售 $ 3,300"

	return output, nil
}
