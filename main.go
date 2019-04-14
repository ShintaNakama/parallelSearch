package main

import (
	"log"
	"parallelSearch/app/controllers"
	"parallelSearch/config"
	"parallelSearch/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	log.Println(controllers.StartWebServer())
	// c1 := make(chan *googleMapPlace.RecivedResults)
	// c2 := make(chan *googleMapPlace.RecivedResults)
	// c3 := make(chan *googleMapPlace.RecivedResults)
	// // var wg sync.WaitGroup
	// // wg.Add(3)
	// // go googleMapPlace.GooglePlaces("焼肉", "赤坂見附駅", &wg, c1)
	// // go googleMapPlace.GooglePlaces("居酒屋", "赤坂見附駅", &wg, c2)
	// // go googleMapPlace.GooglePlaces("公園", "赤坂見附駅", &wg, c3)
	// go googleMapPlace.GooglePlaces("焼肉", "赤坂見附駅", c1)
	// go googleMapPlace.GooglePlaces("居酒屋", "赤坂見附駅", c2)
	// go googleMapPlace.GooglePlaces("公園", "赤坂見附駅", c3)
	// // go googleMapPlace.GooglePlaces("焼肉", "赤坂見附駅", &wg)
	// // go googleMapPlace.GooglePlaces("居酒屋", "赤坂見附駅", &wg)
	// // go googleMapPlace.GooglePlaces("公園", "赤坂見附駅", &wg)
	// // wg.Wait()
	// x := <-c1
	// fmt.Println(x)
	// y := <-c2
	// fmt.Println(y)
	// z := <-c3
	// fmt.Println(z)
}
