package rtpaac

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var configCases = []struct {
	name string
	enc  []byte
	dec  MPEG4AudioConfig
}{
	{
		name: "test 1",
		enc:  []byte{17, 144},
		dec: MPEG4AudioConfig{
			ObjectType:   2,
			SampleRate:   48000,
			ChannelCount: 2,
		},
	},
}

func TestConfigDecode(t *testing.T) {
	for _, ca := range configCases {
		t.Run(ca.name, func(t *testing.T) {
			var dec MPEG4AudioConfig
			err := dec.Decode(ca.enc)
			require.NoError(t, err)
			require.Equal(t, ca.dec, dec)
		})
	}
}

var cases = []struct {
	name string
	dec  *AUAndTimestamp
	enc  []byte
}{
	{
		"single",
		&AUAndTimestamp{
			Timestamp: 20 * time.Millisecond,
			AU: []byte{
				0x21, 0x1a, 0xd4, 0xf5, 0x9e, 0x20, 0xc5, 0x42,
				0x89, 0x40, 0xa2, 0x9b, 0x3c, 0x94, 0xdd, 0x28,
				0x94, 0x48, 0xd5, 0x8b, 0xb0, 0x2, 0xdb, 0x1b,
				0xeb, 0xe0, 0xfa, 0x9f, 0xea, 0x91, 0xa7, 0x3,
				0xe8, 0x6b, 0xe5, 0x5, 0x95, 0x6, 0x62, 0x88,
				0x13, 0xa, 0x15, 0xa0, 0xeb, 0xef, 0x40, 0x82,
				0xdf, 0x49, 0xf2, 0xe0, 0x26, 0xfc, 0x52, 0x5b,
				0x6c, 0x2a, 0x2d, 0xe8, 0xa5, 0x70, 0xc5, 0xaf,
				0xfc, 0x98, 0x9a, 0x2f, 0x1f, 0xbb, 0xa2, 0xcb,
				0xb8, 0x26, 0xb6, 0x6e, 0x4c, 0x15, 0x6c, 0x21,
				0x3d, 0x35, 0xf6, 0xcf, 0xa4, 0x3b, 0x72, 0x26,
				0xe1, 0x3a, 0x3a, 0x99, 0xd8, 0x2d, 0x6a, 0x22,
				0xcd, 0x97, 0xa, 0xef, 0x52, 0x9c, 0x5f, 0xcd,
				0x5c, 0xd9, 0xd3, 0x12, 0x7e, 0x45, 0x45, 0xb3,
				0x24, 0xef, 0xd3, 0x4f, 0x2f, 0x96, 0xd9, 0x8b,
				0x9c, 0xc2, 0xcd, 0x54, 0xb, 0x6e, 0x19, 0x84,
				0x56, 0xeb, 0x85, 0x52, 0x63, 0x64, 0x28, 0xb2,
				0xf2, 0xcf, 0xb8, 0xa8, 0x71, 0x53, 0x6, 0x82,
				0x88, 0xf2, 0xc4, 0xe1, 0x7d, 0x65, 0x54, 0xe0,
				0x5e, 0xc8, 0x38, 0x75, 0x9d, 0xb0, 0x58, 0x65,
				0x41, 0xa2, 0xcd, 0xdb, 0x1b, 0x9e, 0xac, 0xd1,
				0xbe, 0xc9, 0x22, 0xf5, 0xe9, 0xc6, 0x6f, 0xaf,
				0xf8, 0xb1, 0x4c, 0xcb, 0xa2, 0x56, 0x11, 0xa4,
				0xd7, 0xfd, 0xe5, 0xef, 0x8e, 0xbf, 0xce, 0x4b,
				0xef, 0xe1, 0xd, 0xc0, 0x27, 0x18, 0xe2, 0x64,
				0x63, 0x5, 0x16, 0x6, 0xc, 0x34, 0xe, 0xf3, 0x62,
				0xc2, 0xd6, 0x42, 0x5d, 0x66, 0x81, 0x4, 0x65,
				0x76, 0xaa, 0xe7, 0x39, 0xdd, 0x8e, 0xfe, 0x48,
				0x23, 0x3a, 0x1, 0xc4, 0xd3, 0x65, 0x80, 0x28,
				0x6f, 0x9b, 0xc9, 0xb7, 0x4e, 0x44, 0x4c, 0x98,
				0x6a, 0x5f, 0x3b, 0x97, 0x81, 0x9b, 0xa9, 0xab,
				0xfd, 0xcf, 0x8e, 0x78, 0xbd, 0x4d, 0x70, 0x81,
				0x9b, 0x2d, 0x85, 0x94, 0x74, 0x2a, 0x3a, 0xb4,
				0xff, 0x4a, 0x13, 0x70, 0x76, 0x2c, 0x2f, 0x13,
				0x5b, 0x43, 0xf9, 0x17, 0xee, 0x26, 0x37, 0x1,
				0xbc, 0x9f, 0xb, 0xe, 0x68, 0xcb, 0x87, 0x65,
				0x86, 0xcc, 0x4c, 0x2f, 0x7a, 0x14, 0xd, 0xd1,
				0xb9, 0x57, 0xbd, 0x50, 0xb6, 0x95, 0x44, 0x1a,
				0xd, 0xc0, 0x15, 0xf, 0xd2, 0xc3, 0x72, 0x4d,
				0x6e, 0x4f, 0x8e, 0x6d, 0x64, 0xdc, 0x64, 0x1f,
				0x33, 0x53, 0x4e, 0xd8, 0xa4, 0x74, 0xf3, 0x33,
				0x4, 0x68, 0xd9, 0x92, 0xf3, 0x6e, 0xb7, 0x5b,
				0xe6, 0xf6, 0xc3, 0x55, 0x14, 0x54, 0x87, 0x0,
				0xaf, 0x7,
			},
		},
		[]byte{
			0x80, 0xe0, 0x44, 0xed, 0x88, 0x77, 0x6a, 0x15,
			0x9d, 0xbb, 0x78, 0x12, 0x00, 0x10, 0x0a, 0xd8,
			0x21, 0x1a, 0xd4, 0xf5, 0x9e, 0x20, 0xc5, 0x42,
			0x89, 0x40, 0xa2, 0x9b, 0x3c, 0x94, 0xdd, 0x28,
			0x94, 0x48, 0xd5, 0x8b, 0xb0, 0x02, 0xdb, 0x1b,
			0xeb, 0xe0, 0xfa, 0x9f, 0xea, 0x91, 0xa7, 0x03,
			0xe8, 0x6b, 0xe5, 0x05, 0x95, 0x06, 0x62, 0x88,
			0x13, 0x0a, 0x15, 0xa0, 0xeb, 0xef, 0x40, 0x82,
			0xdf, 0x49, 0xf2, 0xe0, 0x26, 0xfc, 0x52, 0x5b,
			0x6c, 0x2a, 0x2d, 0xe8, 0xa5, 0x70, 0xc5, 0xaf,
			0xfc, 0x98, 0x9a, 0x2f, 0x1f, 0xbb, 0xa2, 0xcb,
			0xb8, 0x26, 0xb6, 0x6e, 0x4c, 0x15, 0x6c, 0x21,
			0x3d, 0x35, 0xf6, 0xcf, 0xa4, 0x3b, 0x72, 0x26,
			0xe1, 0x3a, 0x3a, 0x99, 0xd8, 0x2d, 0x6a, 0x22,
			0xcd, 0x97, 0x0a, 0xef, 0x52, 0x9c, 0x5f, 0xcd,
			0x5c, 0xd9, 0xd3, 0x12, 0x7e, 0x45, 0x45, 0xb3,
			0x24, 0xef, 0xd3, 0x4f, 0x2f, 0x96, 0xd9, 0x8b,
			0x9c, 0xc2, 0xcd, 0x54, 0x0b, 0x6e, 0x19, 0x84,
			0x56, 0xeb, 0x85, 0x52, 0x63, 0x64, 0x28, 0xb2,
			0xf2, 0xcf, 0xb8, 0xa8, 0x71, 0x53, 0x06, 0x82,
			0x88, 0xf2, 0xc4, 0xe1, 0x7d, 0x65, 0x54, 0xe0,
			0x5e, 0xc8, 0x38, 0x75, 0x9d, 0xb0, 0x58, 0x65,
			0x41, 0xa2, 0xcd, 0xdb, 0x1b, 0x9e, 0xac, 0xd1,
			0xbe, 0xc9, 0x22, 0xf5, 0xe9, 0xc6, 0x6f, 0xaf,
			0xf8, 0xb1, 0x4c, 0xcb, 0xa2, 0x56, 0x11, 0xa4,
			0xd7, 0xfd, 0xe5, 0xef, 0x8e, 0xbf, 0xce, 0x4b,
			0xef, 0xe1, 0x0d, 0xc0, 0x27, 0x18, 0xe2, 0x64,
			0x63, 0x05, 0x16, 0x06, 0x0c, 0x34, 0x0e, 0xf3,
			0x62, 0xc2, 0xd6, 0x42, 0x5d, 0x66, 0x81, 0x04,
			0x65, 0x76, 0xaa, 0xe7, 0x39, 0xdd, 0x8e, 0xfe,
			0x48, 0x23, 0x3a, 0x01, 0xc4, 0xd3, 0x65, 0x80,
			0x28, 0x6f, 0x9b, 0xc9, 0xb7, 0x4e, 0x44, 0x4c,
			0x98, 0x6a, 0x5f, 0x3b, 0x97, 0x81, 0x9b, 0xa9,
			0xab, 0xfd, 0xcf, 0x8e, 0x78, 0xbd, 0x4d, 0x70,
			0x81, 0x9b, 0x2d, 0x85, 0x94, 0x74, 0x2a, 0x3a,
			0xb4, 0xff, 0x4a, 0x13, 0x70, 0x76, 0x2c, 0x2f,
			0x13, 0x5b, 0x43, 0xf9, 0x17, 0xee, 0x26, 0x37,
			0x01, 0xbc, 0x9f, 0x0b, 0x0e, 0x68, 0xcb, 0x87,
			0x65, 0x86, 0xcc, 0x4c, 0x2f, 0x7a, 0x14, 0x0d,
			0xd1, 0xb9, 0x57, 0xbd, 0x50, 0xb6, 0x95, 0x44,
			0x1a, 0x0d, 0xc0, 0x15, 0x0f, 0xd2, 0xc3, 0x72,
			0x4d, 0x6e, 0x4f, 0x8e, 0x6d, 0x64, 0xdc, 0x64,
			0x1f, 0x33, 0x53, 0x4e, 0xd8, 0xa4, 0x74, 0xf3,
			0x33, 0x04, 0x68, 0xd9, 0x92, 0xf3, 0x6e, 0xb7,
			0x5b, 0xe6, 0xf6, 0xc3, 0x55, 0x14, 0x54, 0x87,
			0x00, 0xaf, 0x07,
		},
	},
}

func TestEncode(t *testing.T) {
	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			sequenceNumber := uint16(0x44ed)
			ssrc := uint32(0x9dbb7812)
			initialTs := uint32(0x88776655)
			e := NewEncoder(96, 48000, &sequenceNumber, &ssrc, &initialTs)
			enc, err := e.Encode(ca.dec)
			require.NoError(t, err)
			require.Equal(t, ca.enc, enc)
		})
	}
}

func TestDecode(t *testing.T) {
	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			d := NewDecoder(48000)

			// send an initial packet downstream
			// in order to correctly compute the timestamp
			_, err := d.Decode([]byte{
				0x80, 0xe0, 0x44, 0xed, 0x88, 0x77, 0x66, 0x55,
				0x9d, 0xbb, 0x78, 0x12, 0x00, 0x10, 0x00, 0x08, 0x0,
			})
			require.NoError(t, err)

			dec, err := d.Decode(ca.enc)
			require.NoError(t, err)
			require.Equal(t, []*AUAndTimestamp{ca.dec}, dec)
		})
	}
}

func TestDecodeAggregated(t *testing.T) {
	enc := []byte{
		0x80, 0xe1, 0x0e, 0xaf, 0xd1, 0xec, 0xbd, 0x7e,
		0x2d, 0x47, 0x97, 0x58, 0x00, 0x40, 0x09, 0x38,
		0x08, 0x68, 0x09, 0x10, 0x08, 0xb8, 0x21, 0x1a,
		0xd5, 0x1d, 0x1c, 0x86, 0xc6, 0x83, 0xb1, 0x08,
		0xa0, 0x42, 0x08, 0x19, 0xb5, 0x5d, 0x5d, 0x70,
		0xb5, 0x6d, 0x46, 0x70, 0x09, 0xa1, 0x6e, 0x84,
		0x49, 0x8c, 0xc7, 0x83, 0x0d, 0xb3, 0xd7, 0x1e,
		0x5d, 0x4d, 0x93, 0x52, 0x92, 0xd2, 0x4a, 0xdf,
		0x40, 0x39, 0xbe, 0x02, 0x09, 0x73, 0xc5, 0xdc,
		0xa8, 0xd6, 0x12, 0xdb, 0xd8, 0x53, 0x07, 0x42,
		0x35, 0x68, 0x63, 0x81, 0xa6, 0xd2, 0x63, 0x2a,
		0x54, 0x84, 0x0b, 0x24, 0x9d, 0xec, 0xb9, 0x1b,
		0x2b, 0x00, 0xcf, 0xc2, 0x8b, 0x71, 0xe8, 0x90,
		0x50, 0x96, 0x58, 0x69, 0xb3, 0x63, 0x2e, 0xfd,
		0x40, 0xd1, 0x96, 0xca, 0xaf, 0x77, 0x09, 0x04,
		0x60, 0xce, 0x60, 0x57, 0x37, 0x22, 0x4e, 0xae,
		0x39, 0x82, 0x63, 0x5c, 0x20, 0x34, 0xc9, 0x24,
		0x10, 0x31, 0x81, 0xb7, 0xed, 0x2a, 0xfc, 0x2e,
		0xff, 0xcf, 0xee, 0xc5, 0xf5, 0x2f, 0xca, 0x05,
		0x2e, 0xbb, 0xc9, 0x1b, 0x1f, 0xad, 0x7e, 0x1f,
		0xff, 0x47, 0xc1, 0xaf, 0x0c, 0x2c, 0x18, 0x02,
		0xf5, 0xb1, 0xc6, 0x3a, 0xa6, 0x4e, 0x41, 0x0d,
		0x4c, 0x34, 0x71, 0x0d, 0x02, 0x22, 0x03, 0x00,
		0x00, 0xef, 0xb6, 0x5c, 0xba, 0xd2, 0xb4, 0x4a,
		0xb4, 0x01, 0xca, 0x85, 0xf0, 0x95, 0x4f, 0x31,
		0xfc, 0xff, 0xb2, 0xe2, 0x84, 0x21, 0x89, 0xae,
		0x12, 0xe9, 0xbf, 0x80, 0x55, 0x0a, 0x15, 0xba,
		0xd7, 0xae, 0x92, 0xf3, 0x43, 0xef, 0xa2, 0xa0,
		0x75, 0xab, 0xfc, 0x74, 0x75, 0x51, 0x2d, 0x68,
		0xf1, 0x27, 0x34, 0x05, 0xf6, 0xf9, 0x49, 0x28,
		0x47, 0x77, 0x1a, 0x71, 0xbb, 0xff, 0xc4, 0x96,
		0xcd, 0x9a, 0x89, 0x62, 0xb7, 0xaa, 0xa2, 0xed,
		0x0e, 0xe1, 0x67, 0xbd, 0x15, 0xdd, 0x65, 0xef,
		0xe0, 0xa6, 0xbb, 0x75, 0xd5, 0x83, 0x88, 0xc8,
		0x17, 0xa9, 0x0d, 0xf4, 0x4b, 0xc2, 0x32, 0x79,
		0x73, 0xf6, 0x4e, 0x17, 0x66, 0x72, 0xc0, 0xc1,
		0xd0, 0x60, 0x5e, 0xa8, 0x6f, 0xab, 0x34, 0x78,
		0x94, 0x36, 0xde, 0x77, 0xa5, 0x55, 0xa2, 0x7d,
		0xed, 0x84, 0x29, 0x98, 0xe4, 0x14, 0xd3, 0xe6,
		0xb4, 0xf1, 0xe7, 0x5a, 0x4e, 0x21, 0x1a, 0xd5,
		0x05, 0x9d, 0x93, 0x01, 0x63, 0xa0, 0xc4, 0xc0,
		0x31, 0x08, 0x18, 0x02, 0xf7, 0x6e, 0x58, 0x2a,
		0x55, 0xef, 0xac, 0x68, 0x3a, 0xb0, 0x12, 0x03,
		0xc4, 0xe4, 0x25, 0x2c, 0x8b, 0xf3, 0xf1, 0x9a,
		0x4a, 0x40, 0xad, 0x15, 0x92, 0x9b, 0x99, 0xe7,
		0x34, 0xe2, 0x6b, 0x72, 0xba, 0xb4, 0xdc, 0xdc,
		0x01, 0xd6, 0xbc, 0x31, 0x7a, 0x88, 0xaa, 0xbc,
		0xb7, 0x43, 0xdb, 0x5e, 0xa7, 0x4f, 0x8f, 0x9e,
		0xc1, 0x1b, 0x27, 0x9f, 0xa7, 0x8b, 0x15, 0x2f,
		0xc9, 0xfc, 0x2d, 0xdf, 0x55, 0x04, 0xf3, 0x6d,
		0x1a, 0x32, 0xae, 0x15, 0xa5, 0x6a, 0x27, 0xb2,
		0xe7, 0xe9, 0xa5, 0x6f, 0xc0, 0x42, 0x27, 0x70,
		0xe2, 0x30, 0x70, 0x3f, 0x5d, 0x9f, 0x83, 0x1b,
		0xb7, 0x46, 0x33, 0xb9, 0xc7, 0xdc, 0x48, 0xb8,
		0x7c, 0x7d, 0x2b, 0xde, 0xd2, 0xc1, 0x5b, 0xd1,
		0x3b, 0x57, 0x5a, 0xb9, 0xde, 0xd8, 0x5a, 0x50,
		0xd7, 0x43, 0x3c, 0x7c, 0x1e, 0xed, 0x75, 0x52,
		0x97, 0x65, 0x8a, 0x71, 0x72, 0x2d, 0x40, 0x9d,
		0x63, 0x8a, 0x51, 0xe2, 0x60, 0x30, 0x09, 0x61,
		0xf2, 0x6d, 0x48, 0xb4, 0x59, 0x02, 0x96, 0x01,
		0xf2, 0x05, 0xb7, 0x78, 0xab, 0xba, 0xcc, 0xbf,
		0x0f, 0xb6, 0xec, 0x5d, 0x41, 0xe2, 0x0e, 0x29,
		0x16, 0x81, 0x0c, 0xce, 0xe7, 0x89, 0x2c, 0xcd,
		0x83, 0x75, 0x7c, 0xcf, 0xda, 0x02, 0x7e, 0x9a,
		0x1d, 0x62, 0x4f, 0xe0, 0x1f, 0xd2, 0x3a, 0x44,
		0xcb, 0x70, 0xb0, 0xce, 0x1e, 0xae, 0xab, 0xdd,
		0x3a, 0x5f, 0x8b, 0x79, 0x6e, 0x1c, 0x37, 0x02,
		0xc8, 0xed, 0x38, 0x76, 0xeb, 0x16, 0x54, 0x62,
		0xc1, 0xab, 0x2e, 0x33, 0x1a, 0x19, 0x66, 0x0f,
		0x99, 0xef, 0x04, 0x40, 0xcb, 0xfc, 0x0c, 0x57,
		0x85, 0x00, 0x84, 0x39, 0x32, 0x1e, 0xdf, 0x65,
		0x5a, 0x11, 0xa0, 0x55, 0x6b, 0xee, 0x22, 0xd9,
		0x44, 0x02, 0xf0, 0x27, 0x90, 0x29, 0xa5, 0x6a,
		0x38, 0x1c, 0x21, 0x1a, 0xd4, 0xf5, 0x9d, 0x93,
		0x45, 0x61, 0xa0, 0xc8, 0xa2, 0x40, 0x38, 0x00,
		0xdb, 0x00, 0x37, 0xa3, 0x41, 0xa6, 0x80, 0x33,
		0x0e, 0x52, 0xf3, 0x2c, 0x5a, 0xe6, 0xa8, 0x40,
		0x91, 0xcd, 0x91, 0x5d, 0x5a, 0xff, 0x6e, 0x32,
		0xd2, 0x35, 0x2b, 0x8d, 0x43, 0x40, 0x34, 0x06,
		0xda, 0x8c, 0x22, 0x24, 0x96, 0x31, 0xd5, 0x7e,
		0x26, 0xd2, 0x66, 0xea, 0xcb, 0x0a, 0xf3, 0x83,
		0x2d, 0x86, 0xf1, 0x5a, 0x88, 0xcd, 0x56, 0xf8,
		0x58, 0xd2, 0x53, 0x9e, 0xbc, 0xe6, 0x73, 0xed,
		0x97, 0x32, 0x37, 0x13, 0x6e, 0xea, 0x71, 0x49,
		0xa4, 0xdb, 0x09, 0xca, 0xeb, 0x6f, 0xee, 0x4a,
		0xbf, 0x33, 0x42, 0xb9, 0x83, 0xd2, 0x5f, 0xe7,
		0xaf, 0x72, 0x43, 0x5f, 0x43, 0xf7, 0xfe, 0x3f,
		0x39, 0xfc, 0x80, 0x77, 0xc9, 0x82, 0x06, 0x95,
		0xcb, 0x34, 0xa9, 0x0d, 0x44, 0xcf, 0x86, 0x1f,
		0x59, 0x30, 0x37, 0x05, 0x18, 0xe9, 0x4a, 0x0e,
		0x72, 0x5f, 0x6a, 0xc8, 0x0e, 0x72, 0x00, 0x28,
		0x80, 0x20, 0xa6, 0x9a, 0x56, 0xe0, 0x02, 0xaa,
		0xca, 0x29, 0xd1, 0x13, 0x18, 0x4c, 0xc1, 0x48,
		0x11, 0x80, 0xc2, 0x00, 0x0f, 0x7d, 0xca, 0x5c,
		0x42, 0xd0, 0x58, 0xb1, 0x24, 0xf9, 0xfa, 0x3b,
		0x4e, 0x76, 0xa3, 0x4a, 0xb9, 0xde, 0x32, 0x01,
		0xce, 0xe1, 0xdb, 0x64, 0x31, 0xdf, 0x48, 0x02,
		0x14, 0x02, 0x89, 0xac, 0x24, 0x7f, 0x54, 0xba,
		0xd0, 0x42, 0x3e, 0x27, 0xe8, 0x9d, 0xf6, 0xff,
		0xfb, 0x01, 0xd1, 0x97, 0x21, 0xef, 0xfe, 0xaf,
		0x40, 0x65, 0x00, 0xb0, 0x5d, 0x2c, 0x68, 0x47,
		0x4a, 0x78, 0x97, 0xe9, 0xf4, 0x5e, 0x9b, 0xae,
		0x9e, 0x86, 0x15, 0xad, 0x2e, 0xe8, 0x7a, 0xf1,
		0xe8, 0x05, 0x65, 0x54, 0x39, 0x57, 0xd5, 0x0a,
		0x19, 0xfe, 0xd7, 0x5c, 0x76, 0x54, 0x8b, 0x19,
		0x46, 0x46, 0xfa, 0x16, 0x52, 0xc6, 0x00, 0x0e,
		0x20, 0xe3, 0x82, 0x6d, 0xb2, 0x4f, 0x46, 0xaf,
		0x7e, 0x93, 0xce, 0xa1, 0x03, 0xe0, 0xa4, 0x75,
		0xa0, 0x12, 0x35, 0x4a, 0x17, 0x3d, 0xde, 0xde,
		0x27, 0xa0, 0x3d, 0xc0, 0x21, 0x1a, 0xd5, 0x05,
		0xa5, 0x93, 0x06, 0x42, 0x09, 0x00, 0xa4, 0x10,
		0x29, 0x41, 0x6a, 0xac, 0x36, 0xa6, 0x95, 0x62,
		0xb4, 0x1c, 0x74, 0x1e, 0x8e, 0x74, 0xdd, 0x37,
		0x0f, 0xd6, 0xff, 0x55, 0x8d, 0x94, 0xb7, 0x73,
		0x9a, 0x91, 0x91, 0xd1, 0xbc, 0x14, 0x94, 0xc1,
		0x22, 0x10, 0x22, 0x52, 0xa6, 0xa7, 0x42, 0x1d,
		0x01, 0x2d, 0xe2, 0x55, 0x42, 0x49, 0x4a, 0xfe,
		0xb2, 0xd9, 0xa9, 0xb6, 0xfb, 0x78, 0x4e, 0x47,
		0xb1, 0x13, 0x51, 0x46, 0xd5, 0xd2, 0xd9, 0x3d,
		0xca, 0x0c, 0xc5, 0x1c, 0xb2, 0xd1, 0x5b, 0x22,
		0x2a, 0x1a, 0xa2, 0x9b, 0xaf, 0x6f, 0x60, 0x90,
		0x1b, 0x6c, 0x67, 0x97, 0x08, 0x21, 0x3b, 0xa0,
		0x71, 0x40, 0x9c, 0x79, 0x1a, 0x68, 0xf7, 0xdd,
		0x23, 0xf6, 0xb8, 0x67, 0x37, 0xc8, 0x05, 0x01,
		0x6d, 0x9e, 0xc1, 0x68, 0xaf, 0xb4, 0xb8, 0xb0,
		0x2b, 0xe1, 0x86, 0x0e, 0x21, 0x6f, 0xbb, 0x0e,
		0xa8, 0xc6, 0xaf, 0xcb, 0xd1, 0x71, 0x0d, 0x1b,
		0x9f, 0x0e, 0x22, 0x8e, 0x58, 0x7f, 0xe6, 0x68,
		0xa7, 0x16, 0x0c, 0x74, 0x28, 0xa8, 0x0a, 0xa0,
		0x01, 0x9e, 0x00, 0x2c, 0x42, 0xc0, 0xd0, 0x99,
		0xe7, 0x9c, 0xf8, 0x7f, 0x0a, 0x2b, 0x9e, 0xae,
		0x6d, 0x50, 0x04, 0xba, 0xc3, 0x15, 0x9b, 0x40,
		0xa5, 0x28, 0xce, 0x76, 0x93, 0x51, 0x06, 0xab,
		0x01, 0x35, 0xae, 0x95, 0x9d, 0xae, 0x55, 0x82,
		0x18, 0x8d, 0xa8, 0x90, 0x71, 0x71, 0xcb, 0xed,
		0x30, 0x03, 0x1a, 0xe8, 0x6f, 0x1e, 0x0c, 0x0d,
		0x30, 0x8a, 0xb9, 0xfd, 0xc3, 0x60, 0xf8, 0x27,
		0x7b, 0x40, 0x63, 0x05, 0xaf, 0xd5, 0xf4, 0x9c,
		0xd3, 0x55, 0x5b, 0xa1, 0x95, 0x1a, 0x5f, 0x44,
		0xe8, 0x3e, 0x1e, 0xf5, 0x12, 0x97, 0x9f, 0xf0,
		0xc5, 0xe5, 0xff, 0x7e, 0x48, 0x71, 0xf5, 0xf6,
		0xd5, 0xfd, 0xea, 0x0f, 0x7c, 0xf9, 0x9f, 0x95,
		0x6c, 0x50, 0x5e, 0x70, 0xbe, 0x4d, 0xd9, 0xd7,
		0x82, 0x18, 0x8c, 0x2a, 0x53, 0xa4, 0x5d, 0x2e,
		0x89, 0x71, 0xc0,
	}

	d := NewDecoder(48000)

	// send an initial packet downstream
	// in order to correctly compute the timestamp
	_, err := d.Decode([]byte{
		0x80, 0xe0, 0x44, 0xed, 0x88, 0x77, 0x66, 0x55,
		0x9d, 0xbb, 0x78, 0x12, 0x00, 0x10, 0x00, 0x08, 0x0,
	})
	require.NoError(t, err)

	dec, err := d.Decode(enc)
	require.NoError(t, err)

	exp := []*AUAndTimestamp{
		{
			Timestamp: 25675558187500,
			AU: []byte{
				0x21, 0x1a,
				0xd5, 0x1d, 0x1c, 0x86, 0xc6, 0x83, 0xb1, 0x08,
				0xa0, 0x42, 0x08, 0x19, 0xb5, 0x5d, 0x5d, 0x70,
				0xb5, 0x6d, 0x46, 0x70, 0x09, 0xa1, 0x6e, 0x84,
				0x49, 0x8c, 0xc7, 0x83, 0x0d, 0xb3, 0xd7, 0x1e,
				0x5d, 0x4d, 0x93, 0x52, 0x92, 0xd2, 0x4a, 0xdf,
				0x40, 0x39, 0xbe, 0x02, 0x09, 0x73, 0xc5, 0xdc,
				0xa8, 0xd6, 0x12, 0xdb, 0xd8, 0x53, 0x07, 0x42,
				0x35, 0x68, 0x63, 0x81, 0xa6, 0xd2, 0x63, 0x2a,
				0x54, 0x84, 0x0b, 0x24, 0x9d, 0xec, 0xb9, 0x1b,
				0x2b, 0x00, 0xcf, 0xc2, 0x8b, 0x71, 0xe8, 0x90,
				0x50, 0x96, 0x58, 0x69, 0xb3, 0x63, 0x2e, 0xfd,
				0x40, 0xd1, 0x96, 0xca, 0xaf, 0x77, 0x09, 0x04,
				0x60, 0xce, 0x60, 0x57, 0x37, 0x22, 0x4e, 0xae,
				0x39, 0x82, 0x63, 0x5c, 0x20, 0x34, 0xc9, 0x24,
				0x10, 0x31, 0x81, 0xb7, 0xed, 0x2a, 0xfc, 0x2e,
				0xff, 0xcf, 0xee, 0xc5, 0xf5, 0x2f, 0xca, 0x05,
				0x2e, 0xbb, 0xc9, 0x1b, 0x1f, 0xad, 0x7e, 0x1f,
				0xff, 0x47, 0xc1, 0xaf, 0x0c, 0x2c, 0x18, 0x02,
				0xf5, 0xb1, 0xc6, 0x3a, 0xa6, 0x4e, 0x41, 0x0d,
				0x4c, 0x34, 0x71, 0x0d, 0x02, 0x22, 0x03, 0x00,
				0x00, 0xef, 0xb6, 0x5c, 0xba, 0xd2, 0xb4, 0x4a,
				0xb4, 0x01, 0xca, 0x85, 0xf0, 0x95, 0x4f, 0x31,
				0xfc, 0xff, 0xb2, 0xe2, 0x84, 0x21, 0x89, 0xae,
				0x12, 0xe9, 0xbf, 0x80, 0x55, 0x0a, 0x15, 0xba,
				0xd7, 0xae, 0x92, 0xf3, 0x43, 0xef, 0xa2, 0xa0,
				0x75, 0xab, 0xfc, 0x74, 0x75, 0x51, 0x2d, 0x68,
				0xf1, 0x27, 0x34, 0x05, 0xf6, 0xf9, 0x49, 0x28,
				0x47, 0x77, 0x1a, 0x71, 0xbb, 0xff, 0xc4, 0x96,
				0xcd, 0x9a, 0x89, 0x62, 0xb7, 0xaa, 0xa2, 0xed,
				0x0e, 0xe1, 0x67, 0xbd, 0x15, 0xdd, 0x65, 0xef,
				0xe0, 0xa6, 0xbb, 0x75, 0xd5, 0x83, 0x88, 0xc8,
				0x17, 0xa9, 0x0d, 0xf4, 0x4b, 0xc2, 0x32, 0x79,
				0x73, 0xf6, 0x4e, 0x17, 0x66, 0x72, 0xc0, 0xc1,
				0xd0, 0x60, 0x5e, 0xa8, 0x6f, 0xab, 0x34, 0x78,
				0x94, 0x36, 0xde, 0x77, 0xa5, 0x55, 0xa2, 0x7d,
				0xed, 0x84, 0x29, 0x98, 0xe4, 0x14, 0xd3, 0xe6,
				0xb4, 0xf1, 0xe7, 0x5a, 0x4e,
			},
		},
		{
			Timestamp: 25675558187500 + (1000 * time.Second / 48000),
			AU: []byte{
				0x21, 0x1a, 0xd5, 0x05,
				0x9d, 0x93, 0x01, 0x63, 0xa0, 0xc4, 0xc0,
				0x31, 0x08, 0x18, 0x02, 0xf7, 0x6e, 0x58, 0x2a,
				0x55, 0xef, 0xac, 0x68, 0x3a, 0xb0, 0x12, 0x03,
				0xc4, 0xe4, 0x25, 0x2c, 0x8b, 0xf3, 0xf1, 0x9a,
				0x4a, 0x40, 0xad, 0x15, 0x92, 0x9b, 0x99, 0xe7,
				0x34, 0xe2, 0x6b, 0x72, 0xba, 0xb4, 0xdc, 0xdc,
				0x01, 0xd6, 0xbc, 0x31, 0x7a, 0x88, 0xaa, 0xbc,
				0xb7, 0x43, 0xdb, 0x5e, 0xa7, 0x4f, 0x8f, 0x9e,
				0xc1, 0x1b, 0x27, 0x9f, 0xa7, 0x8b, 0x15, 0x2f,
				0xc9, 0xfc, 0x2d, 0xdf, 0x55, 0x04, 0xf3, 0x6d,
				0x1a, 0x32, 0xae, 0x15, 0xa5, 0x6a, 0x27, 0xb2,
				0xe7, 0xe9, 0xa5, 0x6f, 0xc0, 0x42, 0x27, 0x70,
				0xe2, 0x30, 0x70, 0x3f, 0x5d, 0x9f, 0x83, 0x1b,
				0xb7, 0x46, 0x33, 0xb9, 0xc7, 0xdc, 0x48, 0xb8,
				0x7c, 0x7d, 0x2b, 0xde, 0xd2, 0xc1, 0x5b, 0xd1,
				0x3b, 0x57, 0x5a, 0xb9, 0xde, 0xd8, 0x5a, 0x50,
				0xd7, 0x43, 0x3c, 0x7c, 0x1e, 0xed, 0x75, 0x52,
				0x97, 0x65, 0x8a, 0x71, 0x72, 0x2d, 0x40, 0x9d,
				0x63, 0x8a, 0x51, 0xe2, 0x60, 0x30, 0x09, 0x61,
				0xf2, 0x6d, 0x48, 0xb4, 0x59, 0x02, 0x96, 0x01,
				0xf2, 0x05, 0xb7, 0x78, 0xab, 0xba, 0xcc, 0xbf,
				0x0f, 0xb6, 0xec, 0x5d, 0x41, 0xe2, 0x0e, 0x29,
				0x16, 0x81, 0x0c, 0xce, 0xe7, 0x89, 0x2c, 0xcd,
				0x83, 0x75, 0x7c, 0xcf, 0xda, 0x02, 0x7e, 0x9a,
				0x1d, 0x62, 0x4f, 0xe0, 0x1f, 0xd2, 0x3a, 0x44,
				0xcb, 0x70, 0xb0, 0xce, 0x1e, 0xae, 0xab, 0xdd,
				0x3a, 0x5f, 0x8b, 0x79, 0x6e, 0x1c, 0x37, 0x02,
				0xc8, 0xed, 0x38, 0x76, 0xeb, 0x16, 0x54, 0x62,
				0xc1, 0xab, 0x2e, 0x33, 0x1a, 0x19, 0x66, 0x0f,
				0x99, 0xef, 0x04, 0x40, 0xcb, 0xfc, 0x0c, 0x57,
				0x85, 0x00, 0x84, 0x39, 0x32, 0x1e, 0xdf, 0x65,
				0x5a, 0x11, 0xa0, 0x55, 0x6b,
				0xee, 0x22, 0xd9, 0x44, 0x02, 0xf0, 0x27, 0x90,
				0x29, 0xa5, 0x6a, 0x38, 0x1c,
			},
		},
		{
			Timestamp: 25675558187500 + (2000 * time.Second / 48000),
			AU: []byte{
				0x21, 0x1a, 0xd4, 0xf5, 0x9d, 0x93,
				0x45, 0x61, 0xa0, 0xc8, 0xa2, 0x40, 0x38, 0x00,
				0xdb, 0x00, 0x37, 0xa3, 0x41, 0xa6, 0x80, 0x33,
				0x0e, 0x52, 0xf3, 0x2c, 0x5a, 0xe6, 0xa8, 0x40,
				0x91, 0xcd, 0x91, 0x5d, 0x5a, 0xff, 0x6e, 0x32,
				0xd2, 0x35, 0x2b, 0x8d, 0x43, 0x40, 0x34, 0x06,
				0xda, 0x8c, 0x22, 0x24, 0x96, 0x31, 0xd5, 0x7e,
				0x26, 0xd2, 0x66, 0xea, 0xcb, 0x0a, 0xf3, 0x83,
				0x2d, 0x86, 0xf1, 0x5a, 0x88, 0xcd, 0x56, 0xf8,
				0x58, 0xd2, 0x53, 0x9e, 0xbc, 0xe6, 0x73, 0xed,
				0x97, 0x32, 0x37, 0x13, 0x6e, 0xea, 0x71, 0x49,
				0xa4, 0xdb, 0x09, 0xca, 0xeb, 0x6f, 0xee, 0x4a,
				0xbf, 0x33, 0x42, 0xb9, 0x83, 0xd2, 0x5f, 0xe7,
				0xaf, 0x72, 0x43, 0x5f, 0x43, 0xf7, 0xfe, 0x3f,
				0x39, 0xfc, 0x80, 0x77, 0xc9, 0x82, 0x06, 0x95,
				0xcb, 0x34, 0xa9, 0x0d, 0x44, 0xcf, 0x86, 0x1f,
				0x59, 0x30, 0x37, 0x05, 0x18, 0xe9, 0x4a, 0x0e,
				0x72, 0x5f, 0x6a, 0xc8, 0x0e, 0x72, 0x00, 0x28,
				0x80, 0x20, 0xa6, 0x9a, 0x56, 0xe0, 0x02, 0xaa,
				0xca, 0x29, 0xd1, 0x13, 0x18, 0x4c, 0xc1, 0x48,
				0x11, 0x80, 0xc2, 0x00, 0x0f, 0x7d, 0xca, 0x5c,
				0x42, 0xd0, 0x58, 0xb1, 0x24, 0xf9, 0xfa, 0x3b,
				0x4e, 0x76, 0xa3, 0x4a, 0xb9, 0xde, 0x32, 0x01,
				0xce, 0xe1, 0xdb, 0x64, 0x31, 0xdf, 0x48, 0x02,
				0x14, 0x02, 0x89, 0xac, 0x24, 0x7f, 0x54, 0xba,
				0xd0, 0x42, 0x3e, 0x27, 0xe8, 0x9d, 0xf6, 0xff,
				0xfb, 0x01, 0xd1, 0x97, 0x21, 0xef, 0xfe, 0xaf,
				0x40, 0x65, 0x00, 0xb0, 0x5d, 0x2c, 0x68, 0x47,
				0x4a, 0x78, 0x97, 0xe9, 0xf4, 0x5e, 0x9b, 0xae,
				0x9e, 0x86, 0x15, 0xad, 0x2e, 0xe8, 0x7a, 0xf1,
				0xe8, 0x05, 0x65, 0x54, 0x39, 0x57, 0xd5, 0x0a,
				0x19, 0xfe, 0xd7, 0x5c, 0x76, 0x54, 0x8b, 0x19,
				0x46, 0x46, 0xfa, 0x16, 0x52, 0xc6, 0x00, 0x0e,
				0x20, 0xe3, 0x82, 0x6d, 0xb2, 0x4f, 0x46, 0xaf,
				0x7e, 0x93, 0xce, 0xa1, 0x03, 0xe0, 0xa4, 0x75,
				0xa0, 0x12, 0x35, 0x4a, 0x17, 0x3d, 0xde, 0xde,
				0x27, 0xa0, 0x3d, 0xc0,
			},
		},
		{
			Timestamp: 25675558187500 + (3000 * time.Second / 48000),
			AU: []byte{
				0x21, 0x1a, 0xd5, 0x05,
				0xa5, 0x93, 0x06, 0x42, 0x09, 0x00, 0xa4, 0x10,
				0x29, 0x41, 0x6a, 0xac, 0x36, 0xa6, 0x95, 0x62,
				0xb4, 0x1c, 0x74, 0x1e, 0x8e, 0x74, 0xdd, 0x37,
				0x0f, 0xd6, 0xff, 0x55, 0x8d, 0x94, 0xb7, 0x73,
				0x9a, 0x91, 0x91, 0xd1, 0xbc, 0x14, 0x94, 0xc1,
				0x22, 0x10, 0x22, 0x52, 0xa6, 0xa7, 0x42, 0x1d,
				0x01, 0x2d, 0xe2, 0x55, 0x42, 0x49, 0x4a, 0xfe,
				0xb2, 0xd9, 0xa9, 0xb6, 0xfb, 0x78, 0x4e, 0x47,
				0xb1, 0x13, 0x51, 0x46, 0xd5, 0xd2, 0xd9, 0x3d,
				0xca, 0x0c, 0xc5, 0x1c, 0xb2, 0xd1, 0x5b, 0x22,
				0x2a, 0x1a, 0xa2, 0x9b, 0xaf, 0x6f, 0x60, 0x90,
				0x1b, 0x6c, 0x67, 0x97, 0x08, 0x21, 0x3b, 0xa0,
				0x71, 0x40, 0x9c, 0x79, 0x1a, 0x68, 0xf7, 0xdd,
				0x23, 0xf6, 0xb8, 0x67, 0x37, 0xc8, 0x05, 0x01,
				0x6d, 0x9e, 0xc1, 0x68, 0xaf, 0xb4, 0xb8, 0xb0,
				0x2b, 0xe1, 0x86, 0x0e, 0x21, 0x6f, 0xbb, 0x0e,
				0xa8, 0xc6, 0xaf, 0xcb, 0xd1, 0x71, 0x0d, 0x1b,
				0x9f, 0x0e, 0x22, 0x8e, 0x58, 0x7f, 0xe6, 0x68,
				0xa7, 0x16, 0x0c, 0x74, 0x28, 0xa8, 0x0a, 0xa0,
				0x01, 0x9e, 0x00, 0x2c, 0x42, 0xc0, 0xd0, 0x99,
				0xe7, 0x9c, 0xf8, 0x7f, 0x0a, 0x2b, 0x9e, 0xae,
				0x6d, 0x50, 0x04, 0xba, 0xc3, 0x15, 0x9b, 0x40,
				0xa5, 0x28, 0xce, 0x76, 0x93, 0x51, 0x06, 0xab,
				0x01, 0x35, 0xae, 0x95, 0x9d, 0xae, 0x55, 0x82,
				0x18, 0x8d, 0xa8, 0x90, 0x71, 0x71, 0xcb, 0xed,
				0x30, 0x03, 0x1a, 0xe8, 0x6f, 0x1e, 0x0c, 0x0d,
				0x30, 0x8a, 0xb9, 0xfd, 0xc3, 0x60, 0xf8, 0x27,
				0x7b, 0x40, 0x63, 0x05, 0xaf, 0xd5, 0xf4, 0x9c,
				0xd3, 0x55, 0x5b, 0xa1, 0x95, 0x1a, 0x5f, 0x44,
				0xe8, 0x3e, 0x1e, 0xf5, 0x12, 0x97, 0x9f, 0xf0,
				0xc5, 0xe5, 0xff, 0x7e, 0x48, 0x71, 0xf5, 0xf6,
				0xd5, 0xfd, 0xea, 0x0f, 0x7c, 0xf9, 0x9f, 0x95,
				0x6c, 0x50, 0x5e, 0x70, 0xbe, 0x4d, 0xd9, 0xd7,
				0x82, 0x18, 0x8c, 0x2a, 0x53, 0xa4, 0x5d, 0x2e,
				0x89, 0x71, 0xc0,
			},
		},
	}
	require.Equal(t, exp, dec)
}
