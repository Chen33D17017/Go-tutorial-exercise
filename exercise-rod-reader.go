package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot(mov byte, char byte, test []byte) byte{
  	mod := (char + mov - test[0]) % (test[1] - test[0] + 1)
  	return mod + test[0]
}

func (coder rot13Reader) Read(b []byte) (int,error){
	n, e := coder.r.Read(b)
	var char byte
	if e == nil {
		for i := 0; i<n ; i++ {
			char = b[i]
			if char > 96 && char < 122{
    			b[i] = rot(13, char, []byte{97, 122})
  				} else if char > 64  && char < 90  {
    			b[i] = rot(13, char, []byte{65, 90})
  				}
			}
			return n, nil
		}
	return 0, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	s.Len()
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
