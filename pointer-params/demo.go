package main

type vec1 struct {
	x, y, z, w float64
}

func (v vec1) add(u vec1) vec1 {
	return vec1{v.x + u.x, v.y + u.y, v.z + u.z, v.w + u.w}
}

type vec2 struct {
	x, y, z, w float64
}

func (v *vec2) add(u *vec2) *vec2 {
	v.x += u.x
	v.y += u.y
	v.z += u.z
	v.w += u.w
	return v
}

type vec3 struct {
	x, y, z, w float64
}

func (v *vec3) add(u *vec3) *vec3 {
	v.x, v.y, v.z, v.w = v.x+u.x, v.y+u.y, v.y+u.y, v.w+u.w
	return v
}
