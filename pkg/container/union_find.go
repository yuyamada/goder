package container

type UnionFind []int

func NewUnionFind(n int) UnionFind {
	u := make(UnionFind, n)
	for i := range u {
		u[i] = -1
	}
	return u
}

func (u UnionFind) Root(x int) int {
	if u[x] == -1 {
		return x
	}
	r := u.Root(u[x])
	u[x] = r
	return r
}

func (u UnionFind) IsSame(x, y int) bool {
	return u.Root(x) == u.Root(y)
}

func (u UnionFind) Unite(x, y int) {
	u[u.Root(y)] = u.Root(x)
}
