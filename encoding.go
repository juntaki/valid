package valid

const encoding = "23456789CFGHJMPQRVWXcfghjmpqrvwx"

func encode40bytes(dst, id []byte) {
	for i := 0; i < len(id)/5; i++ {
		dst[7+8*i] = encoding[id[4+5*i]&0x1F]
		dst[6+8*i] = encoding[id[4+5*i]>>5|(id[3+5*i]<<3)&0x1F]
		dst[5+8*i] = encoding[(id[3+5*i]>>2)&0x1F]
		dst[4+8*i] = encoding[id[3+5*i]>>7|(id[2+5*i]<<1)&0x1F]
		dst[3+8*i] = encoding[(id[2+5*i]>>4)&0x1F|(id[1+5*i]<<4)&0x1F]
		dst[2+8*i] = encoding[(id[1+5*i]>>1)&0x1F]
		dst[1+8*i] = encoding[(id[1+5*i]>>6)&0x1F|(id[0+5*i]<<2)&0x1F]
		dst[0+8*i] = encoding[id[0+5*i]>>3]
	}
}

func Encode40bytes2(out, in []byte) {
	for i := 0; i < len(in)/5; i++ {
		dst := out[i*8:]
		id := in[i*5:]
		dst[7] = encoding[id[4]&0x1F]
		dst[6] = encoding[id[4]>>5|(id[3]<<3)&0x1F]
		dst[5] = encoding[(id[3]>>2)&0x1F]
		dst[4] = encoding[id[3]>>7|(id[2]<<1)&0x1F]
		dst[3] = encoding[(id[2]>>4)&0x1F|(id[1]<<4)&0x1F]
		dst[2] = encoding[(id[1]>>1)&0x1F]
		dst[1] = encoding[(id[1]>>6)&0x1F|(id[0]<<2)&0x1F]
		dst[0] = encoding[id[0]>>3]
	}
}
