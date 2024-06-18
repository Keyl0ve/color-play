package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"

	"github.com/gocarina/gocsv"
)

type Menu struct {
	RecipeID                  string  `csv:"レシピ番号"`
	CompanyID                 string  `csv:"企業名"`
	RecipeName                string  `csv:"レシピ名"`
	RecipeType                int     `csv:"レシピタイプ"`
	Ingredients               string  `csv:"材料"`
	SaltScore                 float64 `csv:"塩分スコア"`
	VegetableScore            float64 `csv:"野菜スコア"`
	ProteinScore              float64 `csv:"たんぱく質スコア"`
	SaturatedFattyAcidScore   float64 `csv:"飽和脂肪酸スコア"`
	TotalScore                float64 `csv:"総合スコア"`
	HasSaltMark               bool    `csv:"塩分よしフラグ"`
	HasVegetableMark          bool    `csv:"野菜よしフラグ"`
	HasProteinMark            bool    `csv:"たんぱく質よしフラグ"`
	HasSaturatedFattyAcidMark bool    `csv:"飽和脂肪酸よしフラグ"`
	HasGoldenDishMark         bool    `csv:"ゴールデンDishフラグ"`
	Energy                    float64 `csv:"エネルギー"`
	Protein                   float64 `csv:"たんぱく質"`
	Fat                       float64 `csv:"脂質"`
	Salt                      float64 `csv:"塩分"`
	SaturatedFattyAcid        float64 `csv:"飽和脂肪酸"`
	Vegetable                 float64 `csv:"野菜"`
	Discription               string  `csv:"ディスクリプション"`
	Genres                    string  `csv:"ジャンル"`
	ImageURL                  string  `csv:"画像URL"`
	DetailPageURL             string  `csv:"ページURL"`
}

func main() {
	file := deleteLine()
	readCSV(file)
}

func readCSV(file string) {
	in, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	menu := []*Menu{}

	if err := gocsv.UnmarshalFile(in, &menu); err != nil {
		panic(err)
	}
	for _, m := range menu {
		fmt.Println("RecipeID: ", m.RecipeID)
		fmt.Println("CompanyID: ", m.CompanyID)
		fmt.Println("RecipeName: ", m.RecipeName)
		fmt.Println("RecipeType: ", m.RecipeType)
		fmt.Println("Ingredients: ", m.Ingredients)
		fmt.Println("SaltScore: ", m.SaltScore)
		fmt.Println("VegetableScore: ", m.VegetableScore)
		fmt.Println("ProteinScore: ", m.ProteinScore)
		fmt.Println("SaturatedFattyAcidScore: ", m.SaturatedFattyAcidScore)
		fmt.Println("TotalScore: ", m.TotalScore)
		fmt.Println("HasSaltMark: ", m.HasSaltMark)
		fmt.Println("HasVegetableMark: ", m.HasVegetableMark)
		fmt.Println("HasProteinMark: ", m.HasProteinMark)
		fmt.Println("HasSaturatedFattyAcidMark: ", m.HasSaturatedFattyAcidMark)
		fmt.Println("HasGoldenDishMark: ", m.HasGoldenDishMark)
		fmt.Println("Energy: ", m.Energy)
		fmt.Println("Protein: ", m.Protein)
		fmt.Println("Fat: ", m.Fat)
		fmt.Println("Salt: ", m.Salt)
		fmt.Println("SaturatedFattyAcid: ", m.SaturatedFattyAcid)
		fmt.Println("Vegetable: ", m.Vegetable)
		fmt.Println("Discription: ", m.Discription)
		fmt.Println("Genres: ", m.Genres)
		fmt.Println("ImageURL: ", m.ImageURL)
		fmt.Println("DetailPageURL: ", m.DetailPageURL)
		fmt.Println()
	}
}

func deleteLine() string {
	inputFile, err := os.Open("test.csv")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	// CSVリーダーを作成
	reader := csv.NewReader(inputFile)

	// 全ての行を読み込む
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// filenameにランダムな数字を5桁つける
	filename := fmt.Sprintf("tmp_%05d.csv", rand.Intn(100000))

	// 出力CSVファイルを作成
	outputFile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// CSVライターを作成
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// 削除する行番号を指定
	rowsToDelete := map[int]bool{1: true, 3: true}

	// 指定された行を除いて新しいCSVファイルに書き込む
	for i, record := range records {
		if !rowsToDelete[i+1] { // 行番号は1から始まるため、i+1を使用
			err := writer.Write(record)
			if err != nil {
				panic(err)
			}
		}
	}

	return filename
}
