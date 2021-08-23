package main

import (
	"context"
	"errors"
	"log"

	micro "github.com/asim/go-micro/v3"
	reg "github.com/petrolax/points-regulator/server/proto/regulator"
)

type RegulatorServer struct{}

func (rs *RegulatorServer) intersect(points []*reg.Point) bool {
	a_dx := float64(points[1].X - points[0].X)
	a_dy := float64(points[1].Y - points[0].Y)
	b_dx := float64(points[3].X - points[2].X)
	b_dy := float64(points[3].Y - points[2].Y)
	s := (-(a_dy)*float64(points[0].X-points[2].X) + a_dx*float64(points[0].Y-points[2].Y)) / (-b_dx*a_dy + a_dx*b_dy)
	t := (b_dx*float64(points[0].Y-points[2].Y) - b_dy*float64(points[0].X-points[2].X)) / (-b_dx*a_dy + a_dx*b_dy)
	if s >= 0 && s <= 1 && t >= 0 && t <= 1 {
		return true
	}
	return false
}

func (rs *RegulatorServer) RegulateSetsOfFourPoints(ctx context.Context, req *reg.PointsRequest, res *reg.PointsResponse) error {
	res.Points = req.GetPoints()
	i := 0
	for rs.intersect(res.Points) {
		i++
		if i > 3 {
			return errors.New("Emergency exit")
		}
		if i%3 == 0 {
			res.Points = req.GetPoints()
			res.Points[0], res.Points[2] = res.Points[2], res.Points[0]
			continue
		}
		if i%2 == 0 {
			res.Points = req.GetPoints()
			res.Points[0], res.Points[3] = res.Points[3], res.Points[0]
			continue
		}
		res.Points[1], res.Points[2] = res.Points[2], res.Points[1]
	}
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("points-regulator"),
	)

	service.Init()

	reg.RegisterPointsRegulatorHandler(service.Server(), new(RegulatorServer))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
