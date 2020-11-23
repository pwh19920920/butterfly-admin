package security

type EncodeService interface {
	// 数据加密
	Encode(text, salt string) string
}
