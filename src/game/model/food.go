package model

type Food struct {
	Type      int
	Subtype   int
	Quantity  int
	Freshness float32
	Taste     float32
	Nutrition float32
}

type NutritionFacts struct {
	Calories int
}

const (
	FoodTypeFruit = 0
)

var FoodTypeNames = map[int]string{
	0: "Fruit",
}

const (
	FoodSubtypeApple = 0
)

var FoodSubtypeNames = map[int]string{
	0: "Pomme",
}

var FoodSubtypeNutritionFacts = map[int]NutritionFacts{
	0: NutritionFacts{100},
}
