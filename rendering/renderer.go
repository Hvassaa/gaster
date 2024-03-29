package rendering

import "github.com/hvassaa/gaster/raycasting"

type Renderer interface {
	Render(rays []raycasting.Ray)
}
