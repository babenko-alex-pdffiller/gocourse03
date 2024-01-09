package main

import (
	"fmt"
)

type HospitalItem struct {
	Name string
	Type string
}

type HospitalInventory struct {
	Items []HospitalItem
}

func NewHospitalInventory(capacity int) *HospitalInventory {
	return &HospitalInventory{
		Items: make([]HospitalItem, 0, capacity),
	}
}

// Додає новий елемент до інвентарю
func (h *HospitalInventory) AddItem(item HospitalItem) {
	h.Items = append(h.Items, item)
	fmt.Printf("Added %v. Capacity after append: %d\n", item, cap(h.Items))
}

// Видаляє елемент за індексом
func (h *HospitalInventory) RemoveItem(index int) {
	if index < 0 || index >= len(h.Items) {
		return
	}
	h.Items = append(h.Items[:index], h.Items[index+1:]...)
}

// Оновлює елемент за індексом
func (h *HospitalInventory) UpdateItem(index int, newItem HospitalItem) {
	if index < 0 || index >= len(h.Items) {
		return
	}
	h.Items[index] = newItem
}

// Демонструє операції із підкладковим масивом
func (h *HospitalInventory) DemonstrateSubsliceOperations() {
	sharedSlice := h.Items[:1]
	sharedSlice[0].Name = "Chair"
	fmt.Println("Original inventory after modification in sharedSlice:", h.Items[0].Name)
}

// Копіює слайс
func (h *HospitalInventory) CopyInventory() []HospitalItem {
	independentSlice := make([]HospitalItem, len(h.Items))
	copy(independentSlice, h.Items)
	return independentSlice
}

// Виводить всі елементи інвентарю
func (h *HospitalInventory) PrintItems() {
	fmt.Println("Hospital Inventory Items:")
	for i, item := range h.Items {
		fmt.Printf("%d: %s - %s\n", i, item.Name, item.Type)
	}
}

func main() {
	hospitalItems := [3]HospitalItem{
		{Name: "Bed", Type: "Furniture"},
		{Name: "Bed 2", Type: "Furniture"},
		{Name: "Bed 3", Type: "Furniture"},
	}

	inventory := NewHospitalInventory(0)
	inventory.AddItem(hospitalItems[0])
	inventory.AddItem(hospitalItems[1])
	inventory.AddItem(hospitalItems[2])
	inventory.UpdateItem(0, HospitalItem{Name: "Big Bed", Type: "Furniture"})
	inventory.DemonstrateSubsliceOperations()

	// Копіювання та робота з незалежним слайсом
	independentSlice := inventory.CopyInventory()
	independentSlice[0] = HospitalItem{Name: "Couch", Type: "Furniture"}
	fmt.Println("Original slice:")
	inventory.PrintItems()
	fmt.Println("Independent slice:", independentSlice)

	// Видалення елементу
	inventory.RemoveItem(1)
	fmt.Println("Inventory after removing:", inventory.Items)

	// вплив винесеного слайсу на оригінальний
	outerItems := inventory.Items
	outerItems[0].Type = `medicine`
	fmt.Println("Original slice after updeting outer one")
	inventory.PrintItems()

	// Проблеми із витоком пам'яті
	leakBySubslicing()
}

// Приклад функції, яка може спричинити витік пам'яті
func leakBySubslicing() {
	fmt.Printf("\nMemory leak example\n")
	originalSlice := make([]HospitalItem, 1000000)
	_ = originalSlice[:10] // Невеликий підслайс із великого масиву
	fmt.Printf("Inventory after slicing: len: %d cap %d\n", len(originalSlice), cap(originalSlice))
	fixedSlice := originalSlice[0:10:15]
	fmt.Printf("Inventory after fix: len: %d cap %d\n", len(fixedSlice), cap(fixedSlice))
}
