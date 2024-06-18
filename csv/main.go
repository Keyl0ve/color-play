// package main

// import (
// 	"encoding/csv"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// )

// type Menu struct {
// 	RecipeID                  string
// 	CompanyID                 string
// 	RecipeName                string
// 	RecipeType                int
// 	IsSoup                    bool
// 	Ingredients               string
// 	SaltScore                 float64
// 	VegetableScore            float64
// 	ProteinScore              float64
// 	SaturatedFattyAcidScore   float64
// 	TotalScore                float64
// 	HasSaltMark               bool
// 	HasVegetableMark          bool
// 	HasProteinMark            bool
// 	HasSaturatedFattyAcidMark bool
// 	HasGoldenDishMark         bool
// 	Energy                    float64
// 	Protein                   float64
// 	Fat                       float64
// 	Salt                      float64
// 	SaturatedFattyAcid        float64
// 	Vegetable                 float64
// 	Discription               string
// 	Genres                    string
// 	ImageURL                  string
// 	DetailPageURL             string
// }

// func main() {
// 	// CSVファイルを開く
// 	f, err := os.Open("test.csv")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()

// 	// CSVファイルの内容を読み込む
// 	r := csv.NewReader(f)
// 	records, err := r.ReadAll()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var menus []Menu
// 	for i, record := range records {
// 		if i < 3 {
// 			continue
// 		}

// 		recipeID := record[0]
// 		companyID := record[1]
// 		recipeName := record[2]
// 		recipeType, _ := strconv.Atoi(record[3])
// 		isSoup, _ := strconv.ParseBool(record[4])
// 		ingredients := record[5]
// 		saltScore, _ := strconv.ParseFloat(record[6], 64)
// 		vegetableScore, _ := strconv.ParseFloat(record[7], 64)
// 		proteinScore, _ := strconv.ParseFloat(record[8], 64)
// 		saturatedFattyAcidScore, _ := strconv.ParseFloat(record[9], 64)
// 		totalScore, _ := strconv.ParseFloat(record[10], 64)
// 		hasSaltMark, _ := strconv.ParseBool(record[11])
// 		hasVegetableMark, _ := strconv.ParseBool(record[12])
// 		hasProteinMark, _ := strconv.ParseBool(record[13])
// 		hasSaturatedFattyAcidMark, _ := strconv.ParseBool(record[14])
// 		hasGoldenDishMark, _ := strconv.ParseBool(record[15])
// 		energy, _ := strconv.ParseFloat(record[16], 64)
// 		protein, _ := strconv.ParseFloat(record[17], 64)
// 		fat, _ := strconv.ParseFloat(record[18], 64)
// 		salt, _ := strconv.ParseFloat(record[19], 64)
// 		saturatedFattyAcid, _ := strconv.ParseFloat(record[20], 64)
// 		vegetable, _ := strconv.ParseFloat(record[21], 64)
// 		discription := record[22]
// 		genres := record[23]
// 		imageURL := record[24]
// 		detailPageURL := record[25]

// 		menu := Menu{
// 			RecipeID:                  recipeID,
// 			CompanyID:                 companyID,
// 			RecipeName:                recipeName,
// 			RecipeType:                recipeType,
// 			IsSoup:                    isSoup,
// 			Ingredients:               ingredients,
// 			SaltScore:                 saltScore,
// 			VegetableScore:            vegetableScore,
// 			ProteinScore:              proteinScore,
// 			SaturatedFattyAcidScore:   saturatedFattyAcidScore,
// 			TotalScore:                totalScore,
// 			HasSaltMark:               hasSaltMark,
// 			HasVegetableMark:          hasVegetableMark,
// 			HasProteinMark:            hasProteinMark,
// 			HasSaturatedFattyAcidMark: hasSaturatedFattyAcidMark,
// 			HasGoldenDishMark:         hasGoldenDishMark,
// 			Energy:                    energy,
// 			Protein:                   protein,
// 			Fat:                       fat,
// 			Salt:                      salt,
// 			SaturatedFattyAcid:        saturatedFattyAcid,
// 			Vegetable:                 vegetable,
// 			Discription:               discription,
// 			Genres:                    genres,
// 			ImageURL:                  imageURL,
// 			DetailPageURL:             detailPageURL,
// 		}
// 		// menuの要素と値を一行ずつ表示
// 		fmt.Println("RecipeID:", recipeID)
// 		fmt.Println("CompanyID:", companyID)
// 		fmt.Println("RecipeName:", recipeName)
// 		fmt.Println("RecipeType:", recipeType)
// 		fmt.Println("IsSoup:", isSoup)
// 		fmt.Println("Ingredients:", ingredients)
// 		fmt.Println("SaltScore:", saltScore)
// 		fmt.Println("VegetableScore:", vegetableScore)
// 		fmt.Println("ProteinScore:", proteinScore)
// 		fmt.Println("SaturatedFattyAcidScore:", saturatedFattyAcidScore)
// 		fmt.Println("TotalScore:", totalScore)
// 		fmt.Println("HasSaltMark:", hasSaltMark)
// 		fmt.Println("HasVegetableMark:", hasVegetableMark)
// 		fmt.Println("HasProteinMark:", hasProteinMark)
// 		fmt.Println("HasSaturatedFattyAcidMark:", hasSaturatedFattyAcidMark)
// 		fmt.Println("HasGoldenDishMark:", hasGoldenDishMark)
// 		fmt.Println("Energy:", energy)
// 		fmt.Println("Protein:", protein)
// 		fmt.Println("Fat:", fat)
// 		fmt.Println("Salt:", salt)
// 		fmt.Println("SaturatedFattyAcid:", saturatedFattyAcid)
// 		fmt.Println("Vegetable:", vegetable)
// 		fmt.Println("Discription:", discription)
// 		fmt.Println("Genres:", genres)
// 		fmt.Println("ImageURL:", imageURL)
// 		fmt.Println("DetailPageURL:", detailPageURL)
// 		menus = append(menus, menu)
// 	}

// 	fmt.Println(menus)

// }
