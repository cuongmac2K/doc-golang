package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strconv"
)

var cfgPath = "letcode/money.yaml"

func main() {
	cfg, err := NewConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	l, _ := strconv.Atoi(cfg.Linh.TienMua)
	t, _ := strconv.Atoi(cfg.TaAnh.TienMua)
	c, _ := strconv.Atoi(cfg.Cuong.TienMua)
	TinhTien(l, t, c)

}
func TinhTien(linh, TaAnh, Cuong int) {
	sum := linh + TaAnh + Cuong

	TienMotNguoi := sum / 3
	fmt.Println("nedd ", TienMotNguoi)
	fmt.Println("moi linh can dong ", linh-TienMotNguoi)
	fmt.Println("moi linh can dong ", TaAnh-TienMotNguoi)
	fmt.Println("moi linh can dong ", Cuong-TienMotNguoi)
}

type Config struct {
	Linh struct {
		TienMua string `yaml:"tienMua"`
	} `yaml:"linh"`
	Cuong struct {
		TienMua string `yaml:"tienMua"`
	} `yaml:"cuong"`
	TaAnh struct {
		TienMua string `yaml:"tienMua"`
	} `yaml:"taanh"`
}

func NewConfig(configPath string) (*Config, error) {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
