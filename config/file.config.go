package config

type ContentType struct {
	JPEG string
	PNG  string
}

type _File struct {
	MaxSize int64
	ContentType
}

var File = _File{
	MaxSize: 10 * 1000 * 1000, // 10 MB
	ContentType: ContentType{
		JPEG: "image/jpeg",
		PNG:  "image/png",
	},
}
