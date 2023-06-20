package main

import (
	"reflect"
	"testing"
)

func getOriginalFlavors() []string {
	return []string{
		"Banana Nut Fudge",
		"Black Walnut",
		"Burgundy Cherry",
		"Butterscotch Ribbon",
		"Cherry Macaron",
		"Chocolate",
		"Chocolate Almond",
		"Chocolate Chip",
		"Chocolate Fudge",
		"Chocolate Mint",
		"Chocolate Ribbon",
		"Coffee",
		"Coffee Candy",
		"Date Nut",
		"Eggnog",
		"French Vanilla",
		"Green Mint Stick",
		"Lemon Crisp",
		"Lemon Custard",
		"Lemon Sherbet",
		"Maple Nut",
		"Orange Sherbet",
		"Peach",
		"Peppermint Fudge Ribbon",
		"Peppermint Stick",
		"Pineapple Sherbet",
		"Raspberry Sherbet",
		"Rocky Road",
		"Strawberry",
		"Vanilla",
		"Vanilla Burnt Almond",
	}
}

func TestFooFunction(t *testing.T) {
	result := foo()
	if result != "bar" {
		t.Errorf("fooFunction failed, expected 'bar' but got '%s'", result)
	}
}

func TestCopy(t *testing.T) {
	originalFlavors := getOriginalFlavors()
	result := copySlice(originalFlavors)
	want := getOriginalFlavors()
	if !reflect.DeepEqual(result, want) {
		t.Errorf("copy failed, expected %v but got %v", originalFlavors, result)
	}
}

func TestIs31Flavors(t *testing.T) {
	originalFlavors := getOriginalFlavors()
	result := is31Flavors(originalFlavors)
	if !result {
		t.Errorf("is31Flavors failed, expected true but got false")
	}
}

func TestAddFlavor(t *testing.T) {
	expected := []string{
		"Rainbow Sherbert",
		"Banana Nut Fudge",
		"Black Walnut",
		"Burgundy Cherry",
		"Butterscotch Ribbon",
		"Cherry Macaron",
		"Chocolate",
		"Chocolate Almond",
		"Chocolate Chip",
		"Chocolate Fudge",
		"Chocolate Mint",
		"Chocolate Ribbon",
		"Coffee",
		"Coffee Candy",
		"Date Nut",
		"Eggnog",
		"French Vanilla",
		"Green Mint Stick",
		"Lemon Crisp",
		"Lemon Custard",
		"Lemon Sherbet",
		"Maple Nut",
		"Orange Sherbet",
		"Peach",
		"Peppermint Fudge Ribbon",
		"Peppermint Stick",
		"Pineapple Sherbet",
		"Raspberry Sherbet",
		"Rocky Road",
		"Strawberry",
		"Vanilla",
		"Vanilla Burnt Almond",
	}

	originalFlavors := getOriginalFlavors()
	result := addFlavor(originalFlavors, "Rainbow Sherbert")
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("addFlavor failed, expected %v but got %v", expected, result)
	}
}

func TestRemoveLastFlavor(t *testing.T) {
	expected := []string{
		"Banana Nut Fudge",
		"Black Walnut",
		"Burgundy Cherry",
		"Butterscotch Ribbon",
		"Cherry Macaron",
		"Chocolate",
		"Chocolate Almond",
		"Chocolate Chip",
		"Chocolate Fudge",
		"Chocolate Mint",
		"Chocolate Ribbon",
		"Coffee",
		"Coffee Candy",
		"Date Nut",
		"Eggnog",
		"French Vanilla",
		"Green Mint Stick",
		"Lemon Crisp",
		"Lemon Custard",
		"Lemon Sherbet",
		"Maple Nut",
		"Orange Sherbet",
		"Peach",
		"Peppermint Fudge Ribbon",
		"Peppermint Stick",
		"Pineapple Sherbet",
		"Raspberry Sherbet",
		"Rocky Road",
		"Strawberry",
		"Vanilla",
	}

	originalFlavors := getOriginalFlavors()
	result := removeLastFlavor(originalFlavors)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("removeLastFlavor failed, expected %v but got %v", expected, result)
	}
}

func TestGetFlavorByIndex(t *testing.T) {
	originalFlavors := getOriginalFlavors()
	expected := "Burgundy Cherry"
	result := getFlavorByIndex(originalFlavors, 2)
	if result != expected {
		t.Errorf("getFlavorByIndex failed, expected '%s' but got '%s'", expected, result)
	}
}

func TestRemoveFlavorByName(t *testing.T) {
	expected := []string{
		"Banana Nut Fudge",
		"Black Walnut",
		"Burgundy Cherry",
		"Butterscotch Ribbon",
		"Cherry Macaron",
		"Chocolate",
		"Chocolate Almond",
		"Chocolate Chip",
		"Chocolate Fudge",
		"Chocolate Mint",
		"Chocolate Ribbon",
		"Coffee",
		"Coffee Candy",
		"Date Nut",
		"Eggnog",
		"French Vanilla",
		"Green Mint Stick",
		"Lemon Crisp",
		"Lemon Custard",
		"Lemon Sherbet",
		"Maple Nut",
		"Orange Sherbet",
		"Peach",
		"Peppermint Fudge Ribbon",
		"Peppermint Stick",
		"Pineapple Sherbet",
		"Raspberry Sherbet",
		"Rocky Road",
		"Strawberry",
		"Vanilla Burnt Almond",
	}

	originalFlavors := getOriginalFlavors()
	result := removeFlavorByName(originalFlavors, "Vanilla")
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("removeFlavorByName failed, expected %v but got %v", expected, result)
	}
}

func TestFilterByWord(t *testing.T) {
	expected := []string{
		"Chocolate",
		"Chocolate Almond",
		"Chocolate Chip",
		"Chocolate Fudge",
		"Chocolate Mint",
		"Chocolate Ribbon",
	}

	originalFlavors := getOriginalFlavors()
	result := filterByWord(originalFlavors, "Chocolate")
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("filterByWord failed, expected %v but got %v", expected, result)
	}
}
