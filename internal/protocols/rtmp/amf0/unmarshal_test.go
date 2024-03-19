package amf0

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var cases = []struct {
	name string
	enc  []byte
	dec  []interface{}
}{
	{
		"on metadata",
		[]byte{
			0x02, 0x00, 0x0d, 0x40, 0x73, 0x65, 0x74, 0x44,
			0x61, 0x74, 0x61, 0x46, 0x72, 0x61, 0x6d, 0x65,
			0x02, 0x00, 0x0a, 0x6f, 0x6e, 0x4d, 0x65, 0x74,
			0x61, 0x44, 0x61, 0x74, 0x61, 0x08, 0x00, 0x00,
			0x00, 0x0d, 0x00, 0x08, 0x64, 0x75, 0x72, 0x61,
			0x74, 0x69, 0x6f, 0x6e, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x77,
			0x69, 0x64, 0x74, 0x68, 0x00, 0x40, 0x94, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x06, 0x68,
			0x65, 0x69, 0x67, 0x68, 0x74, 0x00, 0x40, 0x86,
			0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0d,
			0x76, 0x69, 0x64, 0x65, 0x6f, 0x64, 0x61, 0x74,
			0x61, 0x72, 0x61, 0x74, 0x65, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x09,
			0x66, 0x72, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x74,
			0x65, 0x00, 0x40, 0x4d, 0xf8, 0x53, 0xe2, 0x55,
			0x6b, 0x28, 0x00, 0x0c, 0x76, 0x69, 0x64, 0x65,
			0x6f, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x69, 0x64,
			0x00, 0x40, 0x1c, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x6f,
			0x64, 0x61, 0x74, 0x61, 0x72, 0x61, 0x74, 0x65,
			0x00, 0x40, 0x57, 0x58, 0x90, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x0f, 0x61, 0x75, 0x64, 0x69, 0x6f,
			0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x61,
			0x74, 0x65, 0x00, 0x40, 0xe7, 0x70, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x0f, 0x61, 0x75, 0x64,
			0x69, 0x6f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65,
			0x73, 0x69, 0x7a, 0x65, 0x00, 0x40, 0x30, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x06, 0x73,
			0x74, 0x65, 0x72, 0x65, 0x6f, 0x01, 0x01, 0x00,
			0x0c, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x63, 0x6f,
			0x64, 0x65, 0x63, 0x69, 0x64, 0x00, 0x40, 0x24,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x07,
			0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x02,
			0x00, 0x0d, 0x4c, 0x61, 0x76, 0x66, 0x35, 0x36,
			0x2e, 0x33, 0x36, 0x2e, 0x31, 0x30, 0x30, 0x00,
			0x08, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x69, 0x7a,
			0x65, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x09,
		},
		[]interface{}{
			"@setDataFrame",
			"onMetaData",
			ECMAArray{
				{
					Key:   "duration",
					Value: float64(0),
				},
				{
					Key:   "width",
					Value: float64(1280),
				},
				{
					Key:   "height",
					Value: float64(720),
				},
				{
					Key:   "videodatarate",
					Value: float64(0),
				},
				{
					Key:   "framerate",
					Value: float64(59.94005994005994),
				},
				{
					Key:   "videocodecid",
					Value: float64(7),
				},
				{
					Key:   "audiodatarate",
					Value: float64(93.3837890625),
				},
				{
					Key:   "audiosamplerate",
					Value: float64(48000),
				},
				{
					Key:   "audiosamplesize",
					Value: float64(16),
				},
				{
					Key:   "stereo",
					Value: true,
				},
				{
					Key:   "audiocodecid",
					Value: float64(10),
				},
				{
					Key:   "encoder",
					Value: "Lavf56.36.100",
				},
				{
					Key:   "filesize",
					Value: float64(0),
				},
			},
		},
	},
	{
		"connect",
		[]byte{
			0x02, 0x00, 0x07, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
			0x63, 0x74, 0x00, 0x3f, 0xf0, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x03, 0x00, 0x03, 0x61, 0x70,
			0x70, 0x02, 0x00, 0x02, 0x61, 0x70, 0x00, 0x04,
			0x74, 0x79, 0x70, 0x65, 0x02, 0x00, 0x0a, 0x6e,
			0x6f, 0x6e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
			0x65, 0x00, 0x08, 0x66, 0x6c, 0x61, 0x73, 0x68,
			0x56, 0x65, 0x72, 0x02, 0x00, 0x24, 0x46, 0x4d,
			0x4c, 0x45, 0x2f, 0x33, 0x2e, 0x30, 0x20, 0x28,
			0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62,
			0x6c, 0x65, 0x3b, 0x20, 0x4c, 0x61, 0x76, 0x66,
			0x35, 0x36, 0x2e, 0x31, 0x35, 0x2e, 0x31, 0x30,
			0x32, 0x29, 0x00, 0x05, 0x74, 0x63, 0x55, 0x72,
			0x6c, 0x02, 0x00, 0x1c, 0x72, 0x74, 0x6d, 0x70,
			0x3a, 0x2f, 0x2f, 0x31, 0x39, 0x32, 0x2e, 0x31,
			0x36, 0x38, 0x2e, 0x31, 0x2e, 0x32, 0x33, 0x33,
			0x3a, 0x31, 0x39, 0x33, 0x35, 0x2f, 0x61, 0x70,
			0x00, 0x00, 0x09,
		},
		[]interface{}{
			"connect",
			float64(1),
			Object{
				{Key: "app", Value: "ap"},
				{Key: "type", Value: "nonprivate"},
				{Key: "flashVer", Value: "FMLE/3.0 (compatible; Lavf56.15.102)"},
				{Key: "tcUrl", Value: "rtmp://192.168.1.233:1935/ap"},
			},
		},
	},
	{
		"srs",
		[]byte{
			0x02, 0x00, 0x07, 0x5f, 0x72, 0x65, 0x73, 0x75,
			0x6c, 0x74, 0x00, 0x3f, 0xf0, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x03, 0x00, 0x06, 0x66, 0x6d,
			0x73, 0x56, 0x65, 0x72, 0x02, 0x00, 0x0d, 0x46,
			0x4d, 0x53, 0x2f, 0x33, 0x2c, 0x35, 0x2c, 0x33,
			0x2c, 0x38, 0x38, 0x38, 0x00, 0x0c, 0x63, 0x61,
			0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69,
			0x65, 0x73, 0x00, 0x40, 0x5f, 0xc0, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x04, 0x6d, 0x6f, 0x64,
			0x65, 0x00, 0x3f, 0xf0, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x09, 0x03, 0x00, 0x05,
			0x6c, 0x65, 0x76, 0x65, 0x6c, 0x02, 0x00, 0x06,
			0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x00, 0x04,
			0x63, 0x6f, 0x64, 0x65, 0x02, 0x00, 0x1d, 0x4e,
			0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
			0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e,
			0x6e, 0x65, 0x63, 0x74, 0x2e, 0x53, 0x75, 0x63,
			0x63, 0x65, 0x73, 0x73, 0x00, 0x0b, 0x64, 0x65,
			0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
			0x6e, 0x02, 0x00, 0x14, 0x43, 0x6f, 0x6e, 0x6e,
			0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x73,
			0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64,
			0x00, 0x0e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
			0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x04, 0x64, 0x61, 0x74, 0x61, 0x08,
			0x00, 0x00, 0x00, 0x0f, 0x00, 0x07, 0x76, 0x65,
			0x72, 0x73, 0x69, 0x6f, 0x6e, 0x02, 0x00, 0x09,
			0x33, 0x2c, 0x35, 0x2c, 0x33, 0x2c, 0x38, 0x38,
			0x38, 0x00, 0x07, 0x73, 0x72, 0x73, 0x5f, 0x73,
			0x69, 0x67, 0x02, 0x00, 0x03, 0x53, 0x52, 0x53,
			0x00, 0x0a, 0x73, 0x72, 0x73, 0x5f, 0x73, 0x65,
			0x72, 0x76, 0x65, 0x72, 0x02, 0x00, 0x34, 0x53,
			0x52, 0x53, 0x20, 0x31, 0x2e, 0x30, 0x2e, 0x31,
			0x30, 0x20, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75,
			0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x69,
			0x6e, 0x6c, 0x69, 0x6e, 0x76, 0x69, 0x70, 0x2f,
			0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x72,
			0x74, 0x6d, 0x70, 0x2d, 0x73, 0x65, 0x72, 0x76,
			0x65, 0x72, 0x29, 0x00, 0x0b, 0x73, 0x72, 0x73,
			0x5f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
			0x02, 0x00, 0x15, 0x54, 0x68, 0x65, 0x20, 0x4d,
			0x49, 0x54, 0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e,
			0x73, 0x65, 0x20, 0x28, 0x4d, 0x49, 0x54, 0x29,
			0x00, 0x08, 0x73, 0x72, 0x73, 0x5f, 0x72, 0x6f,
			0x6c, 0x65, 0x02, 0x00, 0x12, 0x6f, 0x72, 0x69,
			0x67, 0x69, 0x6e, 0x2f, 0x65, 0x64, 0x67, 0x65,
			0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x00,
			0x07, 0x73, 0x72, 0x73, 0x5f, 0x75, 0x72, 0x6c,
			0x02, 0x00, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x73,
			0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
			0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x69,
			0x6e, 0x6c, 0x69, 0x6e, 0x76, 0x69, 0x70, 0x2f,
			0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x72,
			0x74, 0x6d, 0x70, 0x2d, 0x73, 0x65, 0x72, 0x76,
			0x65, 0x72, 0x00, 0x0b, 0x73, 0x72, 0x73, 0x5f,
			0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x02,
			0x00, 0x06, 0x31, 0x2e, 0x30, 0x2e, 0x31, 0x30,
			0x00, 0x08, 0x73, 0x72, 0x73, 0x5f, 0x73, 0x69,
			0x74, 0x65, 0x02, 0x00, 0x1c, 0x68, 0x74, 0x74,
			0x70, 0x3a, 0x2f, 0x2f, 0x62, 0x6c, 0x6f, 0x67,
			0x2e, 0x63, 0x73, 0x64, 0x6e, 0x2e, 0x6e, 0x65,
			0x74, 0x2f, 0x77, 0x69, 0x6e, 0x5f, 0x6c, 0x69,
			0x6e, 0x00, 0x09, 0x73, 0x72, 0x73, 0x5f, 0x65,
			0x6d, 0x61, 0x69, 0x6c, 0x02, 0x00, 0x12, 0x77,
			0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x40, 0x76, 0x69,
			0x70, 0x2e, 0x31, 0x32, 0x36, 0x2e, 0x63, 0x6f,
			0x6d, 0x00, 0x0d, 0x73, 0x72, 0x73, 0x5f, 0x63,
			0x6f, 0x70, 0x79, 0x72, 0x69, 0x67, 0x68, 0x74,
			0x02, 0x00, 0x1e, 0x43, 0x6f, 0x70, 0x79, 0x72,
			0x69, 0x67, 0x68, 0x74, 0x20, 0x28, 0x63, 0x29,
			0x20, 0x32, 0x30, 0x31, 0x33, 0x2d, 0x32, 0x30,
			0x31, 0x34, 0x20, 0x77, 0x69, 0x6e, 0x6c, 0x69,
			0x6e, 0x00, 0x0b, 0x73, 0x72, 0x73, 0x5f, 0x70,
			0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x02, 0x00,
			0x06, 0x77, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x00,
			0x0b, 0x73, 0x72, 0x73, 0x5f, 0x61, 0x75, 0x74,
			0x68, 0x6f, 0x72, 0x73, 0x02, 0x00, 0x0b, 0x77,
			0x65, 0x6e, 0x6a, 0x69, 0x65, 0x2e, 0x7a, 0x68,
			0x61, 0x6f, 0x00, 0x0d, 0x73, 0x72, 0x73, 0x5f,
			0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69,
			0x70, 0x02, 0x00, 0x0b, 0x31, 0x37, 0x32, 0x2e,
			0x31, 0x37, 0x2e, 0x30, 0x2e, 0x31, 0x30, 0x00,
			0x07, 0x73, 0x72, 0x73, 0x5f, 0x70, 0x69, 0x64,
			0x00, 0x3f, 0xf0, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x06, 0x73, 0x72, 0x73, 0x5f, 0x69,
			0x64, 0x00, 0x40, 0x5a, 0x40, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x09, 0x00, 0x00, 0x09,
		},
		[]interface{}{
			"_result",
			float64(1),
			Object{
				{Key: "fmsVer", Value: "FMS/3,5,3,888"},
				{Key: "capabilities", Value: float64(127)},
				{Key: "mode", Value: float64(1)},
			},
			Object{
				{Key: "level", Value: "status"},
				{Key: "code", Value: "NetConnection.Connect.Success"},
				{Key: "description", Value: "Connection succeeded"},
				{Key: "objectEncoding", Value: float64(0)},
				{
					Key: "data",
					Value: ECMAArray{
						{Key: "version", Value: "3,5,3,888"},
						{Key: "srs_sig", Value: "SRS"},
						{Key: "srs_server", Value: "SRS 1.0.10 (github.com/winlinvip/simple-rtmp-server)"},
						{Key: "srs_license", Value: "The MIT License (MIT)"},
						{Key: "srs_role", Value: "origin/edge server"},
						{Key: "srs_url", Value: "https://github.com/winlinvip/simple-rtmp-server"},
						{Key: "srs_version", Value: "1.0.10"},
						{Key: "srs_site", Value: "http://blog.csdn.net/win_lin"},
						{Key: "srs_email", Value: "winlin@vip.126.com"},
						{Key: "srs_copyright", Value: "Copyright (c) 2013-2014 winlin"},
						{Key: "srs_primary", Value: "winlin"},
						{Key: "srs_authors", Value: "wenjie.zhao"},
						{Key: "srs_server_ip", Value: "172.17.0.10"},
						{Key: "srs_pid", Value: float64(1)},
						{Key: "srs_id", Value: float64(105)},
					},
				},
			},
		},
	},
	{
		"play",
		[]byte{
			0x02, 0x00, 0x04, 0x70, 0x6c, 0x61, 0x79, 0x00,
			0x40, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x05, 0x02, 0x00, 0x01, 0x31, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		},
		[]interface{}{
			"play",
			float64(3),
			nil,
			"1",
			float64(0),
		},
	},
}

func TestUnmarshal(t *testing.T) {
	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			dec, err := Unmarshal(ca.enc)
			require.NoError(t, err)
			require.Equal(t, ca.dec, dec)
		})
	}
}

func FuzzUnmarshal(f *testing.F) {
	for _, ca := range cases {
		f.Add(ca.enc)
	}

	f.Fuzz(func(_ *testing.T, b []byte) {
		Unmarshal(b) //nolint:errcheck
	})
}
