package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var productsCsv_file = "./products.csv "

func main() {
	// dir, err := os.Getwd()
	// if err != nil {
	//	log.Fatal(err)
	// 	fmt.Println(err)
	//  fmt.Println(err.Error())
	// }
	// fmt.Println(dir)

	app()
	// executeShellCommand()
}

func app() {
	// Remove all products
	// removeAllProducts()

	// csvReader("./products.csv")
	res := csvReader(productsCsv_file)

	fmt.Println("Total rown: ", len(res))
	fmt.Println("Total columns: ", len(res[0]))
	fmt.Println(res[1][0])

	// Create head map
	head := make(map[int]string)
	for i, val := range res[0] {
		head[i] = val
		// fmt.Println(len(val), "|", val, "|", i)
		// fmt.Println(i, "|", head[i])
	}

	createProduct_template := createProduct_processor(res, head, createProduct_template)
	fmt.Println(createProduct_template)
	executeShellCommand(createProduct_template)

}

func removeUnattachedImages() {
}

func createProduct_processor(res [][]string, head map[int]string, template string) string {
	// # Create single product string
	x := createProduct_template
	for i, val := range res[1] {
		// head[val] = i
		// fmt.Println(len(val), "|", val, "|", i)
		// fmt.Println(val, "|", head[val])

		// fmt.Println(i, "|", head[i], "|", val)
		x = strings.Replace(x, "{{"+head[i]+"}}", val, 10)
		// strings.Replace(x, old, new)
	}
	// # Clean up remaining strings
	//  Note1:  In order to update clearing fields in a product,  ....="" should remain (eg: --slug="")
	//  Note2:  --status, should be one of: draft, pending, private, publish, future
	//reg := regexp.MustCompile(`^.*{{.*}}.*$`)			// Full line
	reg := regexp.MustCompile(`{{.*}}`) // Only {{...}}
	out := reg.ReplaceAllString(x, "")

	return out
}

// --status, should be one of: draft, pending, private, publish, future
var createProduct_template string = `
	wp wc product create  --user=1 --path="/home/beanstalkweb/public_html/porto" \
		--name="{{Name}}" \
		--type="{{Type}}" \
		--sku="{{SKU}}" \
		--slug="{{}}" \
		
		--status="publish" \

		--featured="{{Is featured?}}" \
		--catalog_visibility="{{Visibility in catalog}}"\
		--description="{{Description}}" \
		--short_description="{{Short description}}" \
		--regular_price="{{Regular price}}" \
		--sale_price="{{Sale price}}" \
		--date_on_sale_from="{{Date sale price starts}}" \
		--date_on_sale_from_gmt="{{}}" \
		--date_on_sale_to="{{Date sale price ends}}" \
		--date_on_sale_to_gmt="{{}}" \
		--virtual="{{}}" \
		--downloadable="{{}}" \
		--downloads="{{}}" \
		--download_limit="{{}}" \
		--download_expiry="{{}}" \
		--external_url="{{}}" \
		--button_text="{{}}" \
		--tax_status="{{Tax status}}" \
		--tax_class="{{Tax class}}" \
		--manage_stock="{{}}" \
		--stock_quantity="{{}}" \
		--in_stock="{{In stock?}}" \
		--backorders="{{Backorders allowed?}}" \
		--sold_individually="{{Sold individually?}}" \
		--weight="{{Weight (kg)}}" \
		--dimensions="{{}}" \
		--shipping_class="{{Shipping class}}" \
		--reviews_allowed="{{Allow customer reviews?}}" \
		--upsell_ids="{{Upsells}}" \
		--cross_sell_ids="{{Cross-sells}}" \
		--parent_id="{{Parent}}" \
		--purchase_note="{{}}" \
		--categories="{{Categories}}" \
		--tags="{{Tags}}" \
		--images="{{Images}}" \
		--attributes="{{}}" \
		--default_attributes="{{}}" \
		--grouped_products="{{}}" \
		--menu_order="{{}}" \
		--meta_data="{{}}" \
		--wpml_language="{{}}" \
		--porcelain

`

func removeAllProducts() {
	executeShellCommand(removeAllProducts_template)
}

var removeAllProducts_template string = `
	wp post delete $(wp wc product list --user=1 --format=ids) --force
`

func executeShellCommand(cmdLine string) {
	// cmdLine := "ll"
	cmd := exec.Command(cmdLine)

	// app := "echo"
	// arg0 := "-e"
	// arg1 := "Hello world"
	// arg2 := "\n\tfrom"
	// arg3 := "golang"
	// cmd := exec.Command(app, arg0, arg1, arg2, arg3)

	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))
}

func csvReader(file string) [][]string {
	// 1. Open the file
	recordFile, err := os.Open(file)
	if err != nil {
		fmt.Println("An error encountered ::", err)
	}
	// 2. Initialize the reader
	reader := csv.NewReader(recordFile)
	// 3. Read all the records
	records, _ := reader.ReadAll()

	// 4. Iterate through the records as you wish
	// fmt.Println("Total rown: ", len(records))
	// fmt.Println("Total columns: ", len(records[0]))
	// fmt.Println(records[1][0])
	// for i, val := range records[0] {
	// 	fmt.Println(val, "|", i)
	// }

	return records
}

// x := map[string]int {
//   "steve": 12000,
//   "jamie": 15000,
// }

// for i, val := range x { }  // /range returns index and value
