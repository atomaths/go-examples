package canvas

import (
        "fmt"
)

type Shaper interface {
        draw()
}

type Rect struct {
        x, y    int
        cx, cy  int
}

func NewRect(x, y, cx, cy int) *Rect {
        return &Rect{x, y, cx, cy}
}

func (r *Rect) draw() {
        fmt.Println("Rect:draw()", r.x, r.y, r.cx, r.cy)
}

type Canvas struct {
        Name string
}

func (c *Canvas) Paint(s Shaper) {
        fmt.Println(c.Name)
        s.(Shaper).draw()
}
