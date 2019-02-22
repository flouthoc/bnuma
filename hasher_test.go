package bnumamaps

import ("testing")

func BenchmarkFnvHash(b *testing.B){

	for n:=0; n<b.N; n++{
		FnvHash("hello")
	}
}

func BenchmarkXxHash(b *testing.B){

	for n:=0; n<b.N; n++{
		XxHash("hello")
	}
}