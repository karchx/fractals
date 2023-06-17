package main

// https://github.com/ncruces/go-image/blob/v0.1.0/imageutil/srgb.go

var rgb2lin = [...]uint16{
	0x0000, 0x0014, 0x0028, 0x003c, 0x0050, 0x0063, 0x0077, 0x008b,
	0x009f, 0x00b3, 0x00c7, 0x00db, 0x00f1, 0x0108, 0x0120, 0x0139,
	0x0154, 0x016f, 0x018c, 0x01ab, 0x01ca, 0x01eb, 0x020e, 0x0232,
	0x0257, 0x027d, 0x02a5, 0x02ce, 0x02f9, 0x0325, 0x0353, 0x0382,
	0x03b3, 0x03e5, 0x0418, 0x044d, 0x0484, 0x04bc, 0x04f6, 0x0532,
	0x056f, 0x05ad, 0x05ed, 0x062f, 0x0673, 0x06b8, 0x06fe, 0x0747,
	0x0791, 0x07dd, 0x082a, 0x087a, 0x08ca, 0x091d, 0x0972, 0x09c8,
	0x0a20, 0x0a79, 0x0ad5, 0x0b32, 0x0b91, 0x0bf2, 0x0c55, 0x0cba,
	0x0d20, 0x0d88, 0x0df2, 0x0e5e, 0x0ecc, 0x0f3c, 0x0fae, 0x1021,
	0x1097, 0x110e, 0x1188, 0x1203, 0x1280, 0x1300, 0x1381, 0x1404,
	0x1489, 0x1510, 0x159a, 0x1625, 0x16b2, 0x1741, 0x17d3, 0x1866,
	0x18fb, 0x1993, 0x1a2c, 0x1ac8, 0x1b66, 0x1c06, 0x1ca7, 0x1d4c,
	0x1df2, 0x1e9a, 0x1f44, 0x1ff1, 0x20a0, 0x2150, 0x2204, 0x22b9,
	0x2370, 0x242a, 0x24e5, 0x25a3, 0x2664, 0x2726, 0x27eb, 0x28b1,
	0x297b, 0x2a46, 0x2b14, 0x2be3, 0x2cb6, 0x2d8a, 0x2e61, 0x2f3a,
	0x3015, 0x30f2, 0x31d2, 0x32b4, 0x3399, 0x3480, 0x3569, 0x3655,
	0x3742, 0x3833, 0x3925, 0x3a1a, 0x3b12, 0x3c0b, 0x3d07, 0x3e06,
	0x3f07, 0x400a, 0x4110, 0x4218, 0x4323, 0x4430, 0x453f, 0x4651,
	0x4765, 0x487c, 0x4995, 0x4ab1, 0x4bcf, 0x4cf0, 0x4e13, 0x4f39,
	0x5061, 0x518c, 0x52b9, 0x53e9, 0x551b, 0x5650, 0x5787, 0x58c1,
	0x59fe, 0x5b3d, 0x5c7e, 0x5dc2, 0x5f09, 0x6052, 0x619e, 0x62ed,
	0x643e, 0x6591, 0x66e8, 0x6840, 0x699c, 0x6afa, 0x6c5b, 0x6dbe,
	0x6f24, 0x708d, 0x71f8, 0x7366, 0x74d7, 0x764a, 0x77c0, 0x7939,
	0x7ab4, 0x7c32, 0x7db3, 0x7f37, 0x80bd, 0x8246, 0x83d1, 0x855f,
	0x86f0, 0x8884, 0x8a1b, 0x8bb4, 0x8d50, 0x8eef, 0x9090, 0x9235,
	0x93dc, 0x9586, 0x9732, 0x98e2, 0x9a94, 0x9c49, 0x9e01, 0x9fbb,
	0xa179, 0xa339, 0xa4fc, 0xa6c2, 0xa88b, 0xaa56, 0xac25, 0xadf6,
	0xafca, 0xb1a1, 0xb37b, 0xb557, 0xb737, 0xb919, 0xbaff, 0xbce7,
	0xbed2, 0xc0c0, 0xc2b1, 0xc4a5, 0xc69c, 0xc895, 0xca92, 0xcc91,
	0xce94, 0xd099, 0xd2a1, 0xd4ad, 0xd6bb, 0xd8cc, 0xdae0, 0xdcf7,
	0xdf11, 0xe12e, 0xe34e, 0xe571, 0xe797, 0xe9c0, 0xebec, 0xee1b,
	0xf04d, 0xf282, 0xf4ba, 0xf6f5, 0xf933, 0xfb74, 0xfdb8, 0xffff,
}
var lin2rgb = [...]uint16{
	0x0000, 0x0cfc, 0x15f9, 0x1c6b, 0x21ce, 0x2671, 0x2a93, 0x2e53,
	0x31c6, 0x34fb, 0x37fd, 0x3ad3, 0x3d83, 0x4013, 0x4286, 0x44e0,
	0x4722, 0x4950, 0x4b6b, 0x4d75, 0x4f6f, 0x515b, 0x5339, 0x550a,
	0x56d0, 0x588b, 0x5a3c, 0x5be3, 0x5d82, 0x5f17, 0x60a5, 0x622b,
	0x63a9, 0x6521, 0x6692, 0x67fd, 0x6962, 0x6ac1, 0x6c1a, 0x6d6f,
	0x6ebe, 0x7008, 0x714e, 0x7290, 0x73cc, 0x7505, 0x763a, 0x776b,
	0x7898, 0x79c1, 0x7ae7, 0x7c0a, 0x7d29, 0x7e45, 0x7f5e, 0x8074,
	0x8187, 0x8297, 0x83a4, 0x84af, 0x85b7, 0x86bd, 0x87c0, 0x88c0,
	0x89be, 0x8aba, 0x8bb4, 0x8cab, 0x8da1, 0x8e94, 0x8f85, 0x9074,
	0x9161, 0x924d, 0x9336, 0x941e, 0x9503, 0x95e7, 0x96ca, 0x97aa,
	0x9889, 0x9967, 0x9a42, 0x9b1d, 0x9bf5, 0x9ccc, 0x9da2, 0x9e76,
	0x9f49, 0xa01b, 0xa0eb, 0xa1b9, 0xa287, 0xa353, 0xa41e, 0xa4e7,
	0xa5b0, 0xa677, 0xa73d, 0xa802, 0xa8c5, 0xa988, 0xaa49, 0xab09,
	0xabc8, 0xac87, 0xad44, 0xae00, 0xaebb, 0xaf75, 0xb02d, 0xb0e5,
	0xb19d, 0xb253, 0xb308, 0xb3bc, 0xb46f, 0xb522, 0xb5d3, 0xb684,
	0xb734, 0xb7e3, 0xb891, 0xb93e, 0xb9ea, 0xba96, 0xbb41, 0xbbeb,
	0xbc94, 0xbd3d, 0xbde4, 0xbe8b, 0xbf32, 0xbfd7, 0xc07c, 0xc120,
	0xc1c3, 0xc266, 0xc308, 0xc3a9, 0xc44a, 0xc4ea, 0xc589, 0xc628,
	0xc6c6, 0xc763, 0xc800, 0xc89c, 0xc937, 0xc9d2, 0xca6d, 0xcb06,
	0xcb9f, 0xcc38, 0xccd0, 0xcd67, 0xcdfe, 0xce94, 0xcf2a, 0xcfbf,
	0xd053, 0xd0e7, 0xd17b, 0xd20e, 0xd2a0, 0xd332, 0xd3c3, 0xd454,
	0xd4e5, 0xd574, 0xd604, 0xd693, 0xd721, 0xd7af, 0xd83c, 0xd8c9,
	0xd956, 0xd9e2, 0xda6d, 0xdaf8, 0xdb83, 0xdc0d, 0xdc97, 0xdd20,
	0xdda9, 0xde32, 0xdeba, 0xdf41, 0xdfc8, 0xe04f, 0xe0d6, 0xe15b,
	0xe1e1, 0xe266, 0xe2eb, 0xe36f, 0xe3f3, 0xe477, 0xe4fa, 0xe57c,
	0xe5ff, 0xe681, 0xe702, 0xe784, 0xe804, 0xe885, 0xe905, 0xe985,
	0xea04, 0xea83, 0xeb02, 0xeb80, 0xebfe, 0xec7c, 0xecf9, 0xed76,
	0xedf3, 0xee6f, 0xeeeb, 0xef67, 0xefe2, 0xf05d, 0xf0d8, 0xf152,
	0xf1cc, 0xf246, 0xf2bf, 0xf338, 0xf3b1, 0xf429, 0xf4a1, 0xf519,
	0xf591, 0xf608, 0xf67f, 0xf6f6, 0xf76c, 0xf7e2, 0xf858, 0xf8cd,
	0xf942, 0xf9b7, 0xfa2c, 0xfaa0, 0xfb14, 0xfb88, 0xfbfc, 0xfc6f,
	0xfce2, 0xfd54, 0xfdc7, 0xfe39, 0xfeab, 0xff1d, 0xff8e, 0xffff,
}

func RGBToLinear(rgb uint8) uint16 {
	return rgb2lin[rgb]
}

func LinearToRGB(lin uint16) uint8 {
	mul := uint64(lin) * 0xff0100
	div := uint32(mul >> 32)
	l := uint32(lin2rgb[uint8(div)])

	return uint8((uint64(257*l+uint32(uint64(uint32(mul))*257>>32)*
		(uint32(lin2rgb[uint8(div+1)])-l)+0x8100) * 0x1fc05f9) >> 41)
}
