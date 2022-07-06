package config

type ContentType struct {
	JPEG string
	PNG  string
}

var File = struct {
	MaxSize int64
	ContentType
}{
	MaxSize: 10 * 1000 * 1000, // 10 MB
	ContentType: ContentType{
		JPEG: "image/jpeg",
		PNG:  "image/png",
	},
}
