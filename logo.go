package main

import (
	"strings"
)

func fetchLogo(domain string) string {
	//fmt.Print("Fetching logo for " + domain + "...")
	//client := &http.Client{}
	//
	//req, err := http.NewRequest("GET", "https://logo.clearbit.com/"+domain+"?size=600&format=png", nil)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//res, err := client.Do(req)
	//
	//if err != nil {
	//	fmt.Println("Failed to fetch image for " + domain)
	//
	//	return ""
	//}
	//
	//defer res.Body.Close()
	//
	//imageData, err := io.ReadAll(res.Body)
	//catch(err, 0, "Failed to decode image for "+domain)
	//
	//// write the image into a png file with the domain name as a filename but with dashes instead of dots
	//file, err := os.Create("logos/" + strings.ReplaceAll(domain, ".", "-") + ".png")
	//catch(err, 1, "Failed to create file")
	//defer file.Close()
	//
	//_, err = file.Write(imageData)
	//catch(err, 1, "Failed to write image data to file")
	//
	//fmt.Println("Done\n")

	return "https://cdn.unistudents.app/assets/institutions/logos/" + strings.ReplaceAll(domain, ".", "-") + ".png"
}
