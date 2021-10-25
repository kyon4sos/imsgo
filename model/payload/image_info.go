package payload

type ImageInfo struct {
	//宽度
	Width	int
	//高度
	Height	int
	//图片地址，可用于渲染
	Url	string
	//图片大小，单位：Byte
	Size int
	//0	原图
	//1	198p压缩图
	//2	720p压缩图
	Type byte
}

