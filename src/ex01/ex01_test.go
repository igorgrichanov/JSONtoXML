package ex01

import (
	"JSONtoXML/src/ex00"
	"encoding/xml"
	"reflect"
	"sort"
	"testing"
)

var origDB = &ex00.Recipes{
	XMLName: xml.Name{
		Space: "",
		Local: "recipes",
	},
	Cakes: []ex00.Cake{
		{
			Name:      "Red Velvet Strawberry Cake",
			StoveTime: "40 min",
			Ingredients: []ex00.IngredientItem{
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
			Ingredients: []ex00.IngredientItem{
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

var origCopy = &ex00.Recipes{
	XMLName: xml.Name{
		Space: "",
		Local: "recipes",
	},
	Cakes: []ex00.Cake{
		{
			Name:      "Red Velvet Strawberry Cake",
			StoveTime: "40 min",
			Ingredients: []ex00.IngredientItem{
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
			Ingredients: []ex00.IngredientItem{
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

var origDB2 = &ex00.Recipes{
	XMLName: xml.Name{
		Space: "",
		Local: "recipes",
	},
	Cakes: []ex00.Cake{
		{
			Name:      "Red Velvet Strawberry Cake",
			StoveTime: "45 min",
			Ingredients: []ex00.IngredientItem{
				{
					ItemName:  "Flour",
					ItemCount: "3",
					ItemUnit:  "mugs",
				},
				{
					ItemName:  "Strawberries",
					ItemCount: "8",
					ItemUnit:  "",
				},
				{
					ItemName:  "Cinnamon",
					ItemCount: "1",
					ItemUnit:  "",
				},
				{
					ItemName:  "Coffee Beans",
					ItemCount: "1",
					ItemUnit:  "pieces",
				},
			},
		},
		{
			Name:      "Moonshine Muffin",
			StoveTime: "30 min",
			Ingredients: []ex00.IngredientItem{
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
		{
			Name:      "Blueberry Igor Cake",
			StoveTime: "50 min",
			Ingredients: []ex00.IngredientItem{
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

var stolenDB = &ex00.Recipes{
	XMLName: xml.Name{
		Space: "",
		Local: "",
	},
	Cakes: []ex00.Cake{
		{
			Name:      "Red Velvet Strawberry Cake",
			StoveTime: "45 min",
			Ingredients: []ex00.IngredientItem{
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
			Ingredients: []ex00.IngredientItem{
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

var cake1 = ex00.Cake{
	Name:      "Red Velvet Strawberry Cake",
	StoveTime: "45 min",
	Ingredients: []ex00.IngredientItem{
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
}

var cake2 = ex00.Cake{
	Name:      "Moonshine Muffin",
	StoveTime: "30 min",
	Ingredients: []ex00.IngredientItem{
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
}

var ingredients1 = []ex00.IngredientItem{
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
	{
		ItemName:  "Cinnamon",
		ItemCount: "1",
	},
}

var ingredients2 = []ex00.IngredientItem{
	{
		ItemName:  "Brown powder",
		ItemCount: "3",
		ItemUnit:  "mugs",
	},
	{
		ItemName:  "Blueberries",
		ItemCount: "1",
		ItemUnit:  "mug",
	},
	{
		ItemName:  "Lemon",
		ItemCount: "1",
		ItemUnit:  "",
	},
	{
		ItemName:  "Baking powder",
		ItemCount: "3",
		ItemUnit:  "teaspoons",
	},
	{
		ItemName:  "Cinnamon",
		ItemCount: "1",
	},
}

var ingredients3 = []ex00.IngredientItem{
	{
		ItemName:  "Baking power",
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
	{
		ItemName:  "Cinnamon",
		ItemCount: "1",
	},
}

var ingredients4 = []ex00.IngredientItem{
	{
		ItemName:  "Baking power",
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
	{
		ItemName:  "Cinnamon",
		ItemCount: "1",
	},
}

var ingredients5 = []ex00.IngredientItem{
	{
		ItemName:  "Baking power",
		ItemCount: "3",
		ItemUnit:  "teaspoons",
	},
	{
		ItemName:  "Brown sugar",
		ItemCount: "0.5",
		ItemUnit:  "cup",
	},
}

var ingredients6 = []ex00.IngredientItem{
	{
		ItemName:  "Baking power",
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
	{
		ItemName:  "Cinnamon",
		ItemCount: "1",
	},
}

func Test_compareCakes(t *testing.T) {
	type args struct {
		old *ex00.Recipes
		new *ex00.Recipes
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 []string
	}{
		{
			name: "test 1",
			args: args{
				old: origDB,
				new: origDB2,
			},
			want:  []string{"Blueberry Muffin Cake"},
			want1: []string{"Blueberry Igor Cake", "Moonshine Muffin"},
		},
		{
			name: "test 2",
			args: args{
				old: origCopy,
				new: stolenDB,
			},
			want:  []string{"Blueberry Muffin Cake"},
			want1: []string{"Moonshine Muffin"},
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := compareCakes(tt.args.old, tt.args.new)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compareCakes() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("compareCakes() got1 = %v, want %v", got1, tt.want1)
			}
			if i == 0 {
				if len(origDB.Cakes) != len(origDB2.Cakes) {
					t.Errorf(
						"len(origDB.Cakes) != len(origDB2.Cakes): %d != %d",
						len(origDB.Cakes),
						len(origDB2.Cakes),
					)
				} else {
					for j := 0; j < len(origDB.Cakes); j++ {
						if !reflect.DeepEqual(origDB.Cakes[j].Name, origDB2.Cakes[j].Name) {
							t.Errorf("Different cakes at the same index: %s - %s", origDB.Cakes[j].Name, origDB2.Cakes[j].Name)
						}
					}
				}
			} else if i == 1 {
				if len(origCopy.Cakes) != len(stolenDB.Cakes) {
					t.Errorf(
						"len(origCopyPtr.Cakes) != len(stolenDB.Cakes): %d != %d",
						len(origCopy.Cakes),
						len(stolenDB.Cakes),
					)
				} else {
					for j := 0; j < len(origCopy.Cakes); j++ {
						if !reflect.DeepEqual(origCopy.Cakes[j].Name, stolenDB.Cakes[j].Name) {
							t.Errorf("Different cakes at the same index: %s - %s", origCopy.Cakes[j].Name, stolenDB.Cakes[j].Name)
						}
					}
				}
			}
		})
	}
}

func Test_compareCookingTime(t *testing.T) {
	type args struct {
		old *ex00.Cake
		new *ex00.Cake
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test 1",
			args: args{
				old: &cake1,
				new: &cake2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareCookingTime(tt.args.old, tt.args.new); got != tt.want {
				t.Errorf("compareCookingTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compareIngredients(t *testing.T) {
	type args struct {
		old []ex00.IngredientItem
		new []ex00.IngredientItem
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 []string
		want2 []ex00.IngredientItem
		want3 []ex00.IngredientItem
	}{
		{
			name: "test 1",
			args: args{
				old: ingredients1,
				new: ingredients2,
			},
			want:  []string{"Brown sugar"},
			want1: []string{"Brown powder", "Lemon"},
			want2: []ex00.IngredientItem{
				{
					ItemName:  "Baking powder",
					ItemCount: "3",
					ItemUnit:  "teaspoons",
				},
				{
					ItemName:  "Blueberries",
					ItemCount: "1",
					ItemUnit:  "cup",
				},
				{
					ItemName:  "Cinnamon",
					ItemCount: "1",
				},
			},
			want3: []ex00.IngredientItem{
				{
					ItemName:  "Baking powder",
					ItemCount: "3",
					ItemUnit:  "teaspoons",
				},
				{
					ItemName:  "Blueberries",
					ItemCount: "1",
					ItemUnit:  "mug",
				},
				{
					ItemName:  "Cinnamon",
					ItemCount: "1",
				},
			},
		},
		{
			name: "test 2",
			args: args{
				old: ingredients3,
				new: ingredients4,
			},
			want:  []string{},
			want1: []string{},
			want2: []ex00.IngredientItem{
				{
					ItemName:  "Baking power",
					ItemCount: "3",
					ItemUnit:  "teaspoons",
				},
				{
					ItemName:  "Blueberries",
					ItemCount: "1",
					ItemUnit:  "cup",
				},
				{
					ItemName:  "Brown sugar",
					ItemCount: "0.5",
					ItemUnit:  "cup",
				},
				{
					ItemName:  "Cinnamon",
					ItemCount: "1",
				},
			},
			want3: []ex00.IngredientItem{
				{
					ItemName:  "Baking power",
					ItemCount: "3",
					ItemUnit:  "teaspoons",
				},
				{
					ItemName:  "Blueberries",
					ItemCount: "1",
					ItemUnit:  "cup",
				},
				{
					ItemName:  "Brown sugar",
					ItemCount: "0.5",
					ItemUnit:  "cup",
				},
				{
					ItemName:  "Cinnamon",
					ItemCount: "1",
				},
			},
		},
		{
			name: "test 3",
			args: args{
				old: ingredients5,
				new: ingredients6,
			},
			want:  []string{},
			want1: []string{"Blueberries", "Cinnamon"},
			want2: []ex00.IngredientItem{
				{
					ItemName:  "Baking power",
					ItemCount: "3",
					ItemUnit:  "teaspoons",
				},
				{
					ItemName:  "Brown sugar",
					ItemCount: "0.5",
					ItemUnit:  "cup",
				},
			},
			want3: []ex00.IngredientItem{
				{
					ItemName:  "Baking power",
					ItemCount: "3",
					ItemUnit:  "teaspoons",
				},
				{
					ItemName:  "Brown sugar",
					ItemCount: "0.5",
					ItemUnit:  "cup",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3 := compareIngredients(
				tt.args.old,
				tt.args.new,
			)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compareIngredients() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf(
					"compareIngredients() got1 = %v, want %v",
					got1,
					tt.want1,
				)
			}
			sort.Slice(
				got2,
				func(i, j int) bool { return got2[i].ItemName < got2[j].ItemName },
			)
			sort.Slice(
				got3,
				func(i, j int) bool { return got3[i].ItemName < got3[j].ItemName },
			)
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf(
					"compareIngredients() got2 = %v, want %v",
					got2,
					tt.want2,
				)
			}
			if !reflect.DeepEqual(got3, tt.want3) {
				t.Errorf(
					"compareIngredients() got3 = %v, want %v",
					got3,
					tt.want3,
				)
			}
		})
	}
}

func Test_compareIngredientCount(t *testing.T) {
	type args struct {
		old *ex00.IngredientItem
		new *ex00.IngredientItem
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test 1",
			args: args{
				old: &ex00.IngredientItem{
					ItemName:  "Brown sugar",
					ItemCount: "0.5",
					ItemUnit:  "cup",
				},
				new: &ex00.IngredientItem{
					ItemName:  "Brown sugar",
					ItemCount: "1",
					ItemUnit:  "cup",
				},
			},
			want: false,
		},
		{
			name: "test 2",
			args: args{
				old: &ex00.IngredientItem{
					ItemName:  "Brown sugar",
					ItemCount: "0.5",
					ItemUnit:  "cup",
				},
				new: &ex00.IngredientItem{
					ItemName:  "Brown sugar",
					ItemCount: "0.5",
					ItemUnit:  "cup",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareIngredientCount(tt.args.old, tt.args.new); got != tt.want {
				t.Errorf("compareIngredientCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compareIngredientUnit(t *testing.T) {
	type args struct {
		old *ex00.IngredientItem
		new *ex00.IngredientItem
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test 1",
			args: args{
				old: &ex00.IngredientItem{
					ItemName:  "Brown sugar",
					ItemCount: "0.5",
					ItemUnit:  "cup",
				},
				new: &ex00.IngredientItem{
					ItemName:  "Brown sugar",
					ItemCount: "0.5",
					ItemUnit:  "mug",
				},
			},
			want: false,
		},
		{
			name: "test 2",
			args: args{
				old: &ex00.IngredientItem{
					ItemName:  "Brown sugar",
					ItemCount: "0.5",
					ItemUnit:  "cup",
				},
				new: &ex00.IngredientItem{
					ItemName:  "Brown sugar",
					ItemCount: "0.5",
					ItemUnit:  "cup",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareIngredientUnit(tt.args.old, tt.args.new); got != tt.want {
				t.Errorf("compareIngredientUnit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isIngredientUnitMissed(t *testing.T) {
	type args struct {
		old *ex00.IngredientItem
		new *ex00.IngredientItem
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test 1",
			args: args{
				old: &ex00.IngredientItem{
					ItemName:  "Brown sugar",
					ItemCount: "0.5",
					ItemUnit:  "cup",
				},
				new: &ex00.IngredientItem{
					ItemName:  "Brown sugar",
					ItemCount: "0.5",
					ItemUnit:  "cup",
				},
			},
			want: false,
		},
		{
			name: "test 2",
			args: args{
				old: &ex00.IngredientItem{
					ItemName:  "Brown sugar",
					ItemCount: "0.5",
					ItemUnit:  "cup",
				},
				new: &ex00.IngredientItem{
					ItemName:  "Brown sugar",
					ItemCount: "0.5",
					ItemUnit:  "",
				},
			},
			want: true,
		},
		{
			name: "test 3",
			args: args{
				old: &ex00.IngredientItem{
					ItemName:  "Brown sugar",
					ItemCount: "0.5",
					ItemUnit:  "",
				},
				new: &ex00.IngredientItem{
					ItemName:  "Brown sugar",
					ItemCount: "0.5",
					ItemUnit:  "cup",
				},
			},
			want: false,
		},
		{
			name: "test 4",
			args: args{
				old: &ex00.IngredientItem{
					ItemName:  "Brown sugar",
					ItemCount: "0.5",
					ItemUnit:  "",
				},
				new: &ex00.IngredientItem{
					ItemName:  "Brown sugar",
					ItemCount: "0.5",
					ItemUnit:  "",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIngredientUnitMissed(tt.args.old, tt.args.new); got != tt.want {
				t.Errorf("isIngredientUnitMissed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isIngredientUnitAdded(t *testing.T) {
	type args struct {
		old *ex00.IngredientItem
		new *ex00.IngredientItem
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test 1",
			args: args{
				old: &ex00.IngredientItem{
					ItemName:  "Brown sugar",
					ItemCount: "0.5",
					ItemUnit:  "cup",
				},
				new: &ex00.IngredientItem{
					ItemName:  "Brown sugar",
					ItemCount: "0.5",
					ItemUnit:  "cup",
				},
			},
			want: false,
		},
		{
			name: "test 2",
			args: args{
				old: &ex00.IngredientItem{
					ItemName:  "Brown sugar",
					ItemCount: "0.5",
					ItemUnit:  "cup",
				},
				new: &ex00.IngredientItem{
					ItemName:  "Brown sugar",
					ItemCount: "0.5",
					ItemUnit:  "",
				},
			},
			want: false,
		},
		{
			name: "test 3",
			args: args{
				old: &ex00.IngredientItem{
					ItemName:  "Brown sugar",
					ItemCount: "0.5",
					ItemUnit:  "",
				},
				new: &ex00.IngredientItem{
					ItemName:  "Brown sugar",
					ItemCount: "0.5",
					ItemUnit:  "cup",
				},
			},
			want: true,
		},
		{
			name: "test 4",
			args: args{
				old: &ex00.IngredientItem{
					ItemName:  "Brown sugar",
					ItemCount: "0.5",
					ItemUnit:  "",
				},
				new: &ex00.IngredientItem{
					ItemName:  "Brown sugar",
					ItemCount: "0.5",
					ItemUnit:  "",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIngredientUnitAdded(tt.args.old, tt.args.new); got != tt.want {
				t.Errorf("isIngredientUnitAdded() = %v, want %v", got, tt.want)
			}
		})
	}
}
