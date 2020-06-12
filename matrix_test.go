package mathematics_test

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/randomSignal/mathematics"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMatrix_IntersectionArea(t *testing.T) {
	var res float64

	type Sample struct {
		MMatrix    mathematics.Matrix
		ArgsMatrix mathematics.Matrix
		Result     float64
	}

	// 1: 矩阵包含参数矩阵点位数量为0
	//	1.1 没有任何交集
	//	1.2 矩阵被参数矩阵包含
	// 2: 矩阵包含参数矩阵点位数量为1
	// 3: 矩阵包含参数矩阵点位数量为2
	// 4: 矩阵包含参数矩阵点位数量为3, 不存在
	// 5: 矩阵包含参数矩阵点位数量为4
	sampleList := []Sample{
		{
			MMatrix:    mathematics.NewMatrix(mathematics.Point{10, 10}, mathematics.Point{20, 20}),
			ArgsMatrix: mathematics.NewMatrix(mathematics.Point{30, 30}, mathematics.Point{40, 40}),
			Result:     0,
		},
		{
			MMatrix:    mathematics.NewMatrix(mathematics.Point{10, 10}, mathematics.Point{20, 20}),
			ArgsMatrix: mathematics.NewMatrix(mathematics.Point{0, 0}, mathematics.Point{40, 40}),
			Result:     100,
		},
		{
			MMatrix:    mathematics.NewMatrix(mathematics.Point{10, 10}, mathematics.Point{20, 20}),
			ArgsMatrix: mathematics.NewMatrix(mathematics.Point{15, 15}, mathematics.Point{18, 40}),
			Result:     15,
		},

		{
			MMatrix:    mathematics.NewMatrix(mathematics.Point{10, 10}, mathematics.Point{20, 20}),
			ArgsMatrix: mathematics.NewMatrix(mathematics.Point{0, 12}, mathematics.Point{18, 18}),
			Result:     48,
		},

		{
			MMatrix:    mathematics.NewMatrix(mathematics.Point{10, 10}, mathematics.Point{20, 20}),
			ArgsMatrix: mathematics.NewMatrix(mathematics.Point{12, 0}, mathematics.Point{18, 18}),
			Result:     48,
		},

		{
			MMatrix:    mathematics.NewMatrix(mathematics.Point{10, 10}, mathematics.Point{20, 20}),
			ArgsMatrix: mathematics.NewMatrix(mathematics.Point{12, 12}, mathematics.Point{22, 18}),
			Result:     48,
		},
	}
	for i := 0; i < len(sampleList); i++ {
		res = sampleList[i].MMatrix.IntersectionArea(sampleList[i].ArgsMatrix)
		fmt.Println(sampleList[i].MMatrix, sampleList[i].ArgsMatrix)
		require.Equal(t, sampleList[i].Result, res)
		spew.Dump(res)
	}
}
