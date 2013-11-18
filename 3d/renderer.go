package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	renderer := NewRenderer(60, 20)

	camera := NewCamera(50, 1, 0.1, 100)
	camera.Pos = Vector{0, 0, -10}
	camera.Rot = LookAt(Vector{}, camera.Pos, Vector{0, 1, 0})

	ca := NewCube(5)
	ca.Pos = Vector{5, 0, 0}

	cb := NewCube(3)
	cb.Pos = Vector{-5, 0, 0}
	scene := []*Mesh{ca, cb}

	start := time.Now()
	for now := range time.Tick(50 * time.Millisecond) {
		delta := now.Sub(start).Seconds()
		ca.Rot = QuaternionFromEuler(math.Pi/8.0, delta*math.Pi, math.Pi/8.0, "XYZ")
		cb.Rot = QuaternionFromEuler(math.Pi/8.0, delta*math.Pi/2.0, math.Pi/8.0, "XYZ")

		renderer.Clear()
		renderer.Render(scene, camera)
	}
}

func Round(v float64) float64 {
	return math.Floor(v + 0.5)
}

func DegToRad(d float64) float64 {
	return d * math.Pi / 180.0
}

type Vector [4]float64

func QuaternionFromEuler(x, y, z float64, order string) Vector {
	s1, c1 := math.Sincos(x / 2.0)
	s2, c2 := math.Sincos(y / 2.0)
	s3, c3 := math.Sincos(z / 2.0)

	switch order {
	case "XYZ":
		return Vector{
			s1*c2*c3 + c1*s2*s3,
			c1*s2*c3 - s1*c2*s3,
			c1*c2*s3 + s1*s2*c3,
			c1*c2*c3 - s1*s2*s3,
		}
	case "YXZ":
		return Vector{
			s1*c2*c3 + c1*s2*s3,
			c1*s2*c3 - s1*c2*s3,
			c1*c2*s3 - s1*s2*c3,
			c1*c2*c3 + s1*s2*s3,
		}
	}

	return Vector{0, 0, 0, 1}
}

func QuaternionFromMatrix(m Matrix) Vector {
	trace := m[0] + m[5] + m[10]

	if trace > 0 {
		s := 0.5 / math.Sqrt(trace+1.0)
		return Vector{
			(m[6] - m[9]) * s,
			(m[8] - m[2]) * s,
			(m[1] - m[4]) * s,
			0.25 / s,
		}
	} else if m[0] > m[5] && m[0] > m[10] {
		s := 2.0 * math.Sqrt(1.0+m[0]-m[5]-m[10])
		return Vector{
			0.25 * s,
			(m[4] + m[1]) / s,
			(m[8] + m[2]) / s,
			(m[6] - m[9]) / s,
		}
	} else if m[5] > m[10] {
		s := 2.0 * math.Sqrt(1.0+m[5]-m[0]-m[10])
		return Vector{
			(m[4] + m[1]) / s,
			0.25 * s,
			(m[9] + m[6]) / s,
			(m[8] - m[2]) / s,
		}
	} else {
		s := 2.0 * math.Sqrt(1.0+m[10]-m[0]-m[5])
		return Vector{
			(m[8] + m[2]) / s,
			(m[9] + m[6]) / s,
			0.25 * s,
			(m[1] - m[4]) / s,
		}
	}
}

func LookAt(target, eye, up Vector) Vector {
	z := eye.Sub(target).Normalize()
	if z.Length() == 0 {
		z[2] = 1
	}

	x := up.Cross(z).Normalize()
	if x.Length() == 0 {
		z[0] += 0.0001
		x = up.Cross(z).Normalize()
	}

	y := z.Cross(x)

	return QuaternionFromMatrix(Matrix{
		x[0], x[1], x[2], 0,
		y[0], y[1], y[2], 0,
		z[0], z[1], z[2], 0,
		0, 0, 0, 1,
	})
}

func (v Vector) Sub(b Vector) Vector {
	return Vector{
		v[0] - b[0],
		v[1] - b[1],
		v[2] - b[2],
		v[3] - b[3],
	}
}

func (v Vector) Normalize() Vector {
	l := v.Length()
	return Vector{
		v[0] / l,
		v[1] / l,
		v[2] / l,
		v[3] / l,
	}
}

func (v Vector) Cross(b Vector) Vector {
	return Vector{
		v[1]*b[2] - v[2]*b[1],
		v[2]*b[0] - v[0]*b[2],
		v[0]*b[1] - v[1]*b[0],
	}
}

func (v Vector) Length() float64 {
	return math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2] + v[3]*v[3])
}

func (v Vector) Mix(b Vector, r float64) Vector {
	return Vector{
		v[0]*(1-r) + b[0]*r,
		v[1]*(1-r) + b[1]*r,
		v[2]*(1-r) + b[2]*r,
		v[3]*(1-r) + b[3]*r,
	}
}

func (v Vector) Transform(m Matrix) Vector {
	return Vector{
		m[0]*v[0] + m[4]*v[1] + m[8]*v[2] + m[12]*v[3],
		m[1]*v[0] + m[5]*v[1] + m[9]*v[2] + m[13]*v[3],
		m[2]*v[0] + m[6]*v[1] + m[10]*v[2] + m[14]*v[3],
		m[3]*v[0] + m[7]*v[1] + m[11]*v[2] + m[15]*v[3],
	}
}

type Matrix [4 * 4]float64

func Identity() Matrix {
	m := Matrix{}
	m[0], m[5], m[10], m[15] = 1, 1, 1, 1
	return m
}

func Frustum(left, right, bottom, top, near, far float64) Matrix {
	m := Matrix{}

	m[0] = 2 * near / (right - left)
	m[8] = (right + left) / (right - left)
	m[5] = 2 * near / (top - bottom)
	m[9] = (top + bottom) / (top - bottom)
	m[10] = -(far + near) / (far - near)
	m[14] = -2 * far * near / (far - near)
	m[11] = -1
	m[15] = 1

	return m
}

func Perspective(fov, aspect, near, far float64) Matrix {
	ymax := near * math.Tan(DegToRad(fov*0.5))
	ymin := -ymax
	xmin := ymin * aspect
	xmax := ymax * aspect

	return Frustum(xmin, xmax, ymin, ymax, near, far)
}

func (m Matrix) Scale(s Vector) Matrix {
	return m.Mul(Matrix{
		s[0], 0, 0, 0,
		0, s[1], 0, 0,
		0, 0, s[2], 0,
		0, 0, 0, 1,
	})
}

func (m Matrix) Rotate(q Vector) Matrix {
	x, y, z, w := q[0], q[1], q[2], q[3]
	x2, y2, z2 := x+x, y+y, z+z
	xx, xy, xz := x*x2, x*y2, x*z2
	yy, yz, zz := y*y2, y*z2, z*z2
	wx, wy, wz := w*x2, w*y2, w*z2

	return m.Mul(Matrix{
		1 - (yy + zz), xy + wz, xz - wy, 0,
		xy - wz, 1 - (xx + zz), yz + wx, 0,
		xz + wy, yz - wx, 1 - (xx + yy), 0,
		0, 0, 0, 1,
	})
}

func (m Matrix) Translate(v Vector) Matrix {
	return m.Mul(Matrix{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		v[0], v[1], v[2], 1,
	})
}

func Compose(pos, rot, scale Vector) Matrix {
	return Identity().Scale(scale).Rotate(rot).Translate(pos)
}

func (m Matrix) Mul(b Matrix) Matrix {
	return Matrix{
		m[0]*b[0] + m[4]*b[1] + m[8]*b[2] + m[12]*b[3],
		m[1]*b[0] + m[5]*b[1] + m[9]*b[2] + m[13]*b[3],
		m[2]*b[0] + m[6]*b[1] + m[10]*b[2] + m[14]*b[3],
		m[3]*b[0] + m[7]*b[1] + m[11]*b[2] + m[15]*b[3],

		m[0]*b[4] + m[4]*b[5] + m[8]*b[6] + m[12]*b[7],
		m[1]*b[4] + m[5]*b[5] + m[9]*b[6] + m[13]*b[7],
		m[2]*b[4] + m[6]*b[5] + m[10]*b[6] + m[14]*b[7],
		m[3]*b[4] + m[7]*b[5] + m[11]*b[6] + m[15]*b[7],

		m[0]*b[8] + m[4]*b[9] + m[8]*b[10] + m[12]*b[11],
		m[1]*b[8] + m[5]*b[9] + m[9]*b[10] + m[13]*b[11],
		m[2]*b[8] + m[6]*b[9] + m[10]*b[10] + m[14]*b[11],
		m[3]*b[8] + m[7]*b[9] + m[11]*b[10] + m[15]*b[11],

		m[0]*b[12] + m[4]*b[13] + m[8]*b[14] + m[12]*b[15],
		m[1]*b[12] + m[5]*b[13] + m[9]*b[14] + m[13]*b[15],
		m[2]*b[12] + m[6]*b[13] + m[10]*b[14] + m[14]*b[15],
		m[3]*b[12] + m[7]*b[13] + m[11]*b[14] + m[15]*b[15],
	}
}

func (m Matrix) SetPosition(p Vector) Matrix {
	m[12], m[13], m[14] = p[0], p[1], p[2]
	return m
}

type Face [3]Vector

type Object struct {
	Pos, Rot, Scale Vector
}

func NewObject() Object {
	return Object{
		Rot:   Vector{0, 0, 0, 1},
		Scale: Vector{1, 1, 1},
	}
}

func (o Object) Matrix() Matrix {
	return Compose(o.Pos, o.Rot, o.Scale)
}

type Camera struct {
	Object
	Projection Matrix
}

func NewCamera(fov, aspect, near, far float64) *Camera {
	return &Camera{
		Object:     NewObject(),
		Projection: Perspective(fov, aspect, near, far),
	}
}

type Mesh struct {
	Object
	Geometry []Face
}

func NewCube(size float64) *Mesh {
	half := size / 2.0
	rtf := Vector{half, half, half, 1}
	ltf := Vector{-half, half, half, 1}
	lbf := Vector{-half, -half, half, 1}
	rbf := Vector{half, -half, half, 1}
	rtb := Vector{half, half, -half, 1}
	ltb := Vector{-half, half, -half, 1}
	lbb := Vector{-half, -half, -half, 1}
	rbb := Vector{half, -half, -half, 1}

	return &Mesh{
		Object: NewObject(),
		Geometry: []Face{
			Face{rtf, ltf, lbf}, Face{lbf, rbf, rtf},
			Face{rtb, rbb, lbb}, Face{lbb, ltb, rtb},
			Face{rtb, ltb, ltf}, Face{ltf, rtf, rtb},
			Face{rbb, rbf, lbf}, Face{lbf, lbb, rbb},
			Face{ltf, ltb, lbb}, Face{lbb, lbf, ltf},
			Face{rtb, rtf, rbf}, Face{rbf, rbb, rtb},
		},
	}
}

func NewPlane(size float64) *Mesh {
	half := size / 2.0
	rtf := Vector{half, half, half, 1}
	ltf := Vector{-half, half, half, 1}
	lbf := Vector{-half, -half, half, 1}
	rbf := Vector{half, -half, half, 1}

	return &Mesh{
		Object: NewObject(),
		Geometry: []Face{
			Face{rtf, ltf, lbf}, Face{lbf, rbf, rtf},
		},
	}
}

type Renderer struct {
	width, height int
	buffer        []uint8
	zbuffer       []float64
	colors        string
}

func NewRenderer(w, h int) *Renderer {
	r := &Renderer{
		width:   w,
		height:  h,
		buffer:  make([]uint8, (w+1)*h),
		zbuffer: make([]float64, w*h),
		colors:  ".,-~:;=!*#$@",
	}
	r.Clear()
	return r
}

func (r *Renderer) set(x, y int, c float64) {
	if x >= r.width || x < 0 || y >= r.height || y < 0 {
		return
	}
	if c > 1 {
		c = 1
	} else if c < 0 {
		c = 0
	}
	pb := x + (y * (r.width + 1))
	pz := x + (y * (r.width))
	pc := int(float64(len(r.colors)-1) * c)
	if r.zbuffer[pz] < c {
		r.buffer[pb] = r.colors[pc]
		r.zbuffer[pz] = c
	}
}

func (r *Renderer) Clear() {
	for p := 0; p < len(r.buffer); p++ {
		if p%(r.width+1) == r.width {
			r.buffer[p] = '\n'
		} else {
			r.buffer[p] = ' '
		}
	}
	r.zbuffer = make([]float64, r.width*r.height)
}

func (r *Renderer) renderFace(f Face, mvp Matrix) {
	a := f[0].Transform(mvp)
	b := f[1].Transform(mvp)
	c := f[2].Transform(mvp)

	distAB := a.Sub(b).Length()
	stepAB := distAB / Round(distAB)
	distBC := b.Sub(c).Length()
	stepBC := distBC / Round(distBC)

	hw := float64(r.width) / 2.0
	hh := float64(r.height) / 2.0
	maxz := float64(r.width)

	for ab := 0.0; ab <= distAB; ab += stepAB {
		for bc := 0.0; bc <= distBC; bc += stepBC {
			v := a.Mix(b, ab/distAB).Mix(c, bc/distBC)
			r.set(int(Round(v[0]+hw)), int(Round(v[1]+hh)), (v[2]+hw)/maxz)
		}
	}
}

func (r Renderer) Render(meshes []*Mesh, camera *Camera) {
	v := camera.Matrix()
	p := camera.Projection
	var mvp Matrix
	for _, mesh := range meshes {
		mvp = v.Mul(p).Mul(mesh.Matrix())
		for _, f := range mesh.Geometry {
			r.renderFace(f, mvp)
		}
	}
	fmt.Print("\x0c", string(r.buffer))
}
