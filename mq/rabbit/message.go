package rabbit

// Message 消息体
type Message struct {
	ProductID uint
	UserID    uint
}

// NewMessage 消息体构造方法
func NewMessage(userId, productId uint) *Message {
	return &Message{
		ProductID: userId,
		UserID:    productId,
	}
}
