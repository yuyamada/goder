package math

type Mod struct {
	mod int
}

func NewMod(mod int) Mod {
	return Mod{mod: mod}
}

func (m Mod) Mul(x, y int) int {
	x, y = x%m.mod, y%m.mod
	if y == 0 {
		return 0
	}
	if x*y/y != x {
		panic("overflow")
	}
	return x * y % m.mod
}

func (m Mod) Pow(base, n int) int {
	if n < 0 {
		return m.Inv(m.Pow(base, -n))
	}
	ret := 1
	for n > 0 {
		if n&1 == 1 {
			ret = m.Mul(ret, base)
		}
		base = m.Mul(base, base)
		n >>= 1
	}
	return ret
}

func (m Mod) Inv(x int) int {
	if x <= 0 {
		panic("x must be positive")
	}
	return m.Pow(x, m.mod-2)
}

var (
	modFactorials    = map[int]int{0: 1}
	modInvFactorials = map[int]int{0: 1}
)

func (m Mod) Factorial(x int) int {
	if v, ok := modFactorials[x]; ok {
		return v
	}
	modFactorials[x-1] = m.Factorial(x - 1)
	return m.Mul(x, modFactorials[x-1])
}

func (m Mod) InitInvFactorials(x int) {
	modInvFactorials[x] = m.Inv(m.Factorial(x))
	for i := 1; x-i > 0; i++ {
		modInvFactorials[x-i] = m.Mul(x-i+1, modInvFactorials[x-i+1])
	}
}

func (m Mod) InvFactorial(x int) int {
	if v, ok := modInvFactorials[x]; ok {
		return v
	}
	return m.Inv(m.Factorial(x))
}

func (m Mod) Combination(n, k int) int {
	if n < 0 || k < 0 || n < k {
		return 0
	}
	return m.Mul(
		m.Factorial(n),
		m.Mul(
			m.InvFactorial(k),
			m.InvFactorial(n-k),
		),
	)
}
