package entity

type FileResource struct {
	BaseModel
	Md5    string
	Bucket string
	Object string
	Mime   string
	Size   int64
}
