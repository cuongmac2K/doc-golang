package main

import (
	"fmt"
	"sort"
)

// Định nghĩa struct dataResource
type dataResource struct {
	Patterns []string `json:"patterns"`
	Type     string   `json:"type"`
	Endpoint string   `json:"endpoint"`
	Url      string   `json:"url"`
}

// Hàm để sắp xếp mảng dataResource dựa trên độ dài của Url
func sortDataResourcesByUrlLength(resources []dataResource) {
	sort.Slice(resources, func(i, j int) bool {
		return len(resources[i].Url) > len(resources[j].Url)
	})
}

func main() {
	// Tạo một số ví dụ về dataResource
	resources := []dataResource{
		{Url: "https://example.com/very/long/url"},
		{Url: "https://short.com"},
		{Url: "https://medium.com/some/path"},
	}

	// In ra trước khi sắp xếp
	fmt.Println("Before sorting:")
	for _, resource := range resources {
		fmt.Println(resource.Url)
	}

	// Sắp xếp mảng dựa trên độ dài của Url
	sortDataResourcesByUrlLength(resources)

	// In ra sau khi sắp xếp
	fmt.Println("After sorting:")
	for _, resource := range resources {
		fmt.Println(resource.Url)
	}
}
