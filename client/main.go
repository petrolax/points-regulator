package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/asim/go-micro/v3"
	reg "github.com/petrolax/points-regulator/server/proto/regulator"
)

func ParseDataToPoints(data string) []*reg.Point {
	p := new(reg.Point)
	ps := make([]*reg.Point, 0)
	num := ""

	for i := 0; i < len(data); i++ {
		if data[i] < 48 || data[i] > 57 {
			val, _ := strconv.Atoi(num)
			num = ""
			if string(data[i]) == "\n" {
				p.Y = int64(val)
				ps = append(ps, p)
				p = new(reg.Point)
				continue
			}
			p.X = int64(val)
			continue
		}
		num += string(data[i])
	}
	return ps
}

func WriteToFile(ps []*reg.Point) error {
	str := ""
	for i := 0; i < len(ps); i++ {
		str += fmt.Sprintf("%d;%d\n", ps[i].X, ps[i].Y)
	}

	err := ioutil.WriteFile("assets/output.txt", []byte(str), 0644)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	file, err := os.Open("assets/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}
	err = file.Close()
	if err != nil {
		log.Fatalln(err)
	}

	points := ParseDataToPoints(string(data))

	service := micro.NewService()

	service.Init()

	client := reg.NewPointsRegulatorService("points-regulator", service.Client())

	res, err := client.RegulateSetsOfFourPoints(context.Background(), &reg.PointsRequest{
		Points: points,
	})

	if err != nil {
		log.Fatalln(err)
	}

	if err = WriteToFile(res.Points); err != nil {
		log.Fatalln(err)
	}
}
