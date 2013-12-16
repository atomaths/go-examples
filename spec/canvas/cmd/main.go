package main

import (
        "github.com/atomaths/go-examples/spec/canvas/pkg"
)

func main() {
        r := canvas.NewRect(0, 0, 10, 10)
        // Error: r.draw undefined (cannot refer to unexported field or method canvas.(*Rect)."".draw)
        // r.draw()

        c := canvas.Canvas{Name: "test"}
        c.Paint(r)
}
