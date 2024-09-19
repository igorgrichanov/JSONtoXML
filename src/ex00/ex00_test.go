package ex00

import (
	"encoding/xml"
	"reflect"
	"testing"
)

var origDBPath, stolenDBPath = "../original_database.xml", "../stolen_database.json"

var origDB = &Recipes{
	XMLName: xml.Name{
		Space: "",
		Local: "recipes",
	},
	Cakes: []Cake{
		{
			Name:      "Red Velvet Strawberry Cake",
			StoveTime: "40 min",
			Ingredients: []IngredientItem{
				{
					ItemName:  "Flour",
					ItemCount: "3",
					ItemUnit:  "cups",
				},
				{
					ItemName:  "Vanilla extract",
					ItemCount: "1.5",
					ItemUnit:  "tablespoons",
				},
				{
					ItemName:  "Strawberries",
					ItemCount: "7",
					ItemUnit:  "",
				},
				{
					ItemName:  "Cinnamon",
					ItemCount: "1",
					ItemUnit:  "pieces",
				},
			}},
		{
			Name:      "Blueberry Muffin Cake",
			StoveTime: "30 min",
			Ingredients: []IngredientItem{
				{
					ItemName:  "Baking powder",
					ItemCount: "3",
					ItemUnit:  "teaspoons",
				},
				{
					ItemName:  "Brown sugar",
					ItemCount: "0.5",
					ItemUnit:  "cup",
				},
				{
					ItemName:  "Blueberries",
					ItemCount: "1",
					ItemUnit:  "cup",
				},
			},
		},
	},
}

var origDBJSON = []byte(`{
    "cake": [
        {
            "name": "Red Velvet Strawberry Cake",
            "time": "40 min",
            "ingredients": [
                {
                    "ingredient_name": "Flour",
                    "ingredient_count": "3",
                    "ingredient_unit": "cups"
                },
                {
                    "ingredient_name": "Vanilla extract",
                    "ingredient_count": "1.5",
                    "ingredient_unit": "tablespoons"
                },
                {
                    "ingredient_name": "Strawberries",
                    "ingredient_count": "7"
                },
                {
                    "ingredient_name": "Cinnamon",
                    "ingredient_count": "1",
                    "ingredient_unit": "pieces"
                }
            ]
        },
        {
            "name": "Blueberry Muffin Cake",
            "time": "30 min",
            "ingredients": [
                {
                    "ingredient_name": "Baking powder",
                    "ingredient_count": "3",
                    "ingredient_unit": "teaspoons"
                },
                {
                    "ingredient_name": "Brown sugar",
                    "ingredient_count": "0.5",
                    "ingredient_unit": "cup"
                },
                {
                    "ingredient_name": "Blueberries",
                    "ingredient_count": "1",
                    "ingredient_unit": "cup"
                }
            ]
        }
    ]
}`)

var stolenDB = &Recipes{
	XMLName: xml.Name{
		Space: "",
		Local: "",
	},
	Cakes: []Cake{
		{
			Name:      "Red Velvet Strawberry Cake",
			StoveTime: "45 min",
			Ingredients: []IngredientItem{
				{
					ItemName:  "Flour",
					ItemCount: "2",
					ItemUnit:  "mugs",
				},
				{
					ItemName:  "Strawberries",
					ItemCount: "8",
					ItemUnit:  "",
				},
				{
					ItemName:  "Coffee Beans",
					ItemCount: "2.5",
					ItemUnit:  "tablespoons",
				},
				{
					ItemName:  "Cinnamon",
					ItemCount: "1",
				},
			},
		},
		{
			Name:      "Moonshine Muffin",
			StoveTime: "30 min",
			Ingredients: []IngredientItem{
				{
					ItemName:  "Brown sugar",
					ItemCount: "1",
					ItemUnit:  "mug",
				},
				{
					ItemName:  "Blueberries",
					ItemCount: "1",
					ItemUnit:  "mug",
				},
			},
		},
	},
}

var stolenDBXML = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<recipes>
    <cake>
        <name>Red Velvet Strawberry Cake</name>
        <stovetime>45 min</stovetime>
        <ingredients>
            <item>
                <itemname>Flour</itemname>
                <itemcount>2</itemcount>
                <itemunit>mugs</itemunit>
            </item>
            <item>
                <itemname>Strawberries</itemname>
                <itemcount>8</itemcount>
            </item>
            <item>
                <itemname>Coffee Beans</itemname>
                <itemcount>2.5</itemcount>
                <itemunit>tablespoons</itemunit>
            </item>
            <item>
                <itemname>Cinnamon</itemname>
                <itemcount>1</itemcount>
            </item>
        </ingredients>
    </cake>
    <cake>
        <name>Moonshine Muffin</name>
        <stovetime>30 min</stovetime>
        <ingredients>
            <item>
                <itemname>Brown sugar</itemname>
                <itemcount>1</itemcount>
                <itemunit>mug</itemunit>
            </item>
            <item>
                <itemname>Blueberries</itemname>
                <itemcount>1</itemcount>
                <itemunit>mug</itemunit>
            </item>
        </ingredients>
    </cake>
</recipes>`)

func TestParseDB(t *testing.T) {
	type args struct {
		filepath *string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 *Recipes
	}{
		{
			name: "test 1",
			args: args{
				filepath: &origDBPath,
			},
			want:  0,
			want1: origDB,
		},
		{
			name: "test 2",
			args: args{
				filepath: &stolenDBPath,
			},
			want:  1,
			want1: stolenDB,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ParseDB(tt.args.filepath)
			if got != tt.want {
				t.Errorf("ParseDB() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ParseDB() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMarshalDB(t *testing.T) {
	type args struct {
		recipes   *Recipes
		extension int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "test 1",
			args: args{
				recipes:   origDB,
				extension: xmlExtension,
			},
			want: origDBJSON,
		},
		{
			name: "test 2",
			args: args{
				recipes:   stolenDB,
				extension: jsonExtension,
			},
			want: stolenDBXML,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MarshalDB(tt.args.recipes, tt.args.extension); !reflect.DeepEqual(
				got,
				tt.want,
			) {
				t.Errorf("MarshalDB() = %v, want %v", got, tt.want)
			}
		})
	}
}
