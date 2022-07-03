package fic

func Fic() func() int {
	var a, b int = 0,1
	return func() int {
		a, b = b, a+b
		return b
	}
}
