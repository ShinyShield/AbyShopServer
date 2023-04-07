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

func Exec() (Output, error) {

	var output Output

	output.HotSales.FirstPic = "p30babL.png"
	output.HotSales.SecondPic = "OoBzmPm.jpg"

	output.SeasonAd.FirstPic = "6OydL5n.jpg"
	output.SeasonAd.SecondPic = "XGkBM7R.jpg"
	output.SeasonAd.ThirdPic = "wzPOMHe.jpg"
	output.SeasonAd.FourPic = "K6Ec069.jpg"

	output.SeasonAd.FirstText = "Item1 $ 12,999"
	output.SeasonAd.SecondText = "Item2 $ 16,999"
	output.SeasonAd.ThirdText = "Item3 $ 21,000"
	output.SeasonAd.FourText = "Item4 $ 300"

	output.LimitedProduct.FirstPic = "h2o29dy.jpg"
	output.LimitedProduct.SecondPic = "hpWHcv2.jpg"
	output.LimitedProduct.ThirdPic = "51Vatnn.jpg"
	output.LimitedProduct.FourPic = "CKPuzhJ.jpg"

	output.LimitedProduct.FirstText = "limItem1 $ 299"
	output.LimitedProduct.SecondText = "limItem2 $ 1,130"
	output.LimitedProduct.ThirdText = "limItem3 $ 79,000"
	output.LimitedProduct.FourText = "limItem4 $ 3,300"

	return output, nil
}
