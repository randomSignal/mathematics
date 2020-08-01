package mathematics

type Point struct {
	X, Y float64
}

type Matrix struct {
	PointList [4]Point
}

func NewMatrix(min Point, max Point) Matrix {
	topLeftPoint := min
	topRightPoint := Point{X: max.X, Y: min.Y}
	lowerRightPoint := max
	lowerLeftPoint := Point{X: min.X, Y: max.Y}

	return Matrix{[4]Point{topLeftPoint, topRightPoint, lowerRightPoint, lowerLeftPoint}}
}

func (m Matrix) Width() float64 {
	return m.MaxX() - m.MinX()
}

func (m Matrix) Height() float64 {
	return m.MaxY() - m.MinY()
}

func (m Matrix) MinX() float64 {
	minX := m.PointList[0].X
	for _, v := range m.PointList {
		if v.X < minX {
			minX = v.X
		}
	}
	return minX
}

func (m Matrix) MinY() float64 {
	minY := m.PointList[0].Y
	for _, v := range m.PointList {
		if v.Y < minY {
			minY = v.Y
		}
	}
	return minY
}

func (m Matrix) MaxX() float64 {
	maxX := m.PointList[0].X
	for _, v := range m.PointList {
		if v.X > maxX {
			maxX = v.X
		}
	}
	return maxX
}

func (m Matrix) MaxY() float64 {
	maxY := m.PointList[0].Y
	for _, v := range m.PointList {
		if v.Y > maxY {
			maxY = v.Y
		}
	}
	return maxY
}

// 左上角点位
func (m Matrix) TopLeftPoint() Point {
	for _, v := range m.PointList {
		if v.X == m.MinX() && v.Y == m.MinY() {
			return v
		}
	}
	return Point{}
}

// 右上角点位
func (m Matrix) TopRightPoint() Point {
	for _, v := range m.PointList {
		if v.X == m.MaxX() && v.Y == m.MinY() {
			return v
		}
	}
	return Point{}
}

// 左下角点位
func (m Matrix) LowerLeftPoint() Point {
	for _, v := range m.PointList {
		if v.X == m.MinX() && v.Y == m.MaxY() {
			return v
		}
	}
	return Point{}
}

// 右下角点位
func (m Matrix) LowerRightPoint() Point {
	for _, v := range m.PointList {
		if v.X == m.MaxX() && v.Y == m.MaxY() {
			return v
		}
	}
	return Point{}
}

// 点是否在矩阵里面
func (m Matrix) PointInMatrix(point Point) bool {
	if point.X >= m.MinX() && point.X <= m.MaxX() &&
		point.Y >= m.MinY() && point.Y <= m.MaxY() {
		return true
	}

	return false
}

// 矩阵面积
func (m Matrix) Area() float64 {
	width := m.MaxX() - m.MinX()
	height := m.MaxY() - m.MinY()

	return width * height
}

// 矩阵是否相交
func (m Matrix) Intersection(args Matrix) bool {
	for _, v := range args.PointList {
		if m.PointInMatrix(v) {
			return true
		}
	}

	return false
}

// 矩阵包含点数量
func (m Matrix) ContainPointNum(args Matrix) []int {
	var res []int

	for i := 0; i < len(args.PointList); i++ {
		if m.PointInMatrix(args.PointList[i]) {
			res = append(res, i)
		}
	}

	return res
}

// 矩阵包含
func (m Matrix) Contain(args Matrix) bool {
	for _, v := range args.PointList {
		if !m.PointInMatrix(v) {
			return false
		}
	}

	return true
}

// 矩阵矫正
func (m Matrix) Correction() Matrix {
	return Matrix{[4]Point{m.TopLeftPoint(), m.TopRightPoint(), m.LowerRightPoint(), m.LowerLeftPoint()}}
}

// 矩阵相交面积
func (m Matrix) IntersectionArea(args Matrix) float64 {
	// 1: 矩阵包含参数矩阵点位数量为0
	//	1.1 没有任何交集
	//	1.2 矩阵被参数矩阵包含
	// 2: 矩阵包含参数矩阵点位数量为1
	// 3: 矩阵包含参数矩阵点位数量为2
	// 4: 矩阵包含参数矩阵点位数量为3, 不存在
	// 5: 矩阵包含参数矩阵点位数量为4
	mCorrection := m.Correction()
	argsCorrection := args.Correction()

	containOffset := mCorrection.ContainPointNum(argsCorrection)
	switch len(containOffset) {
	case 0:
		if argsCorrection.Contain(mCorrection) {
			return mCorrection.Area()
		} else {
			return 0
		}
	case 1:
		mInnerPoint := Point{}
		argsInnerPoint := Point{}
		for _, v := range mCorrection.PointList {
			if argsCorrection.PointInMatrix(v) {
				argsInnerPoint = v
				break
			}
		}
		for _, v := range argsCorrection.PointList {
			if mCorrection.PointInMatrix(v) {
				mInnerPoint = v
				break
			}
		}
		var matrix Matrix
		if mInnerPoint.X <= argsInnerPoint.X && mInnerPoint.Y <= argsInnerPoint.Y {
			matrix = NewMatrix(mInnerPoint, argsInnerPoint)
		} else {
			matrix = NewMatrix(argsInnerPoint, mInnerPoint)
		}
		return matrix.Area()
	case 2:
		// 判断点位为0，1
		// 判断点位为1，2
		// 判断点位为2，3
		// 判断点位为3，0
		width := float64(0)
		height := float64(0)
		if containOffset[0] == 0 && containOffset[1] == 1 {
			width = argsCorrection.PointList[1].X - argsCorrection.PointList[0].X
			height = mCorrection.PointList[3].Y - argsCorrection.PointList[1].Y
		} else if containOffset[0] == 1 && containOffset[1] == 2 {
			height = argsCorrection.PointList[2].Y - argsCorrection.PointList[1].Y
			width = argsCorrection.PointList[1].X - mCorrection.PointList[0].X
		} else if containOffset[0] == 2 && containOffset[1] == 3 {

			width = argsCorrection.PointList[2].X - argsCorrection.PointList[3].X
			height = argsCorrection.PointList[2].Y - mCorrection.PointList[1].Y

		} else {
			height = argsCorrection.PointList[3].Y - argsCorrection.PointList[0].Y
			width = mCorrection.PointList[1].X - argsCorrection.PointList[0].X
		}

		return width * height

	case 4:
		return argsCorrection.Area()
	}

	return 0
}
