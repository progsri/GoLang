package main

type s2 struct{
	k1 string
}

func (k *s2) m1() string {
	return k.k1
}