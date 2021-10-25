package payload

type ImagePayload struct {
	//图片唯一标识
	UUID string `json:"uuid" gorm:"primaryKey"`
	//图片格式类型，JPG/JPEG = 1，GIF = 2，PNG = 3，BMP = 4，其他 = 255
	ImageFormat int8
	//图片信息
	ImageInfoArray []interface{}
}

func (i *ImagePayload) PayloadHandler()  {

}