package main

import (
	"MyGoCrawler/myhttp"
	"MyGoCrawler/utils"
)

func main(){
	//fmt.Println("Hello MyGoCrawler!")

	body, ok := myhttp.Get("https://vol.moe")

	if ok {
		utils.WriteFile("output/result.html", body )
	}

}
