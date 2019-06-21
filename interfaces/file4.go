package main

import "fmt"

type config struct {
	addressNop i1        //This can take a pointer struct or a struct
	addressP   i1
	network    string
	raw1 	rawConfig   //This can take only a struct
	raw2	*rawConfig   //This can take only a pointer struct
}

type rawConfig struct{
	key1 string
}

func main() {

	structNop1 := S1{
		"s1-nop",
	}
	fmt.Println(structNop1)

	structP1 := S1{
		"s1-p",
	}
	fmt.Println(structP1)

	structNop2 := &S1{
		"s1-nop",
	}
	fmt.Println(structNop2)

	structP2 := &S1{
		"s1-p",
	}
	fmt.Println(structP2)

	raw1 := rawConfig{
		key1: "key1",
	}

	raw2 := &rawConfig{
		key1: "key2",
	}

	structConfig1 := config{
		addressNop: structNop1,
		addressP:   structP1,
		network: "tcp",
		raw1: raw1,  
		raw2: raw2,
	}

	structConfig2 := config{}  //This pointers default to nil and others to zero values

	fmt.Println(structConfig1)
	fmt.Println(structConfig2)


}
