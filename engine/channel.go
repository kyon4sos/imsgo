package engine



type ByteChannel struct {
	Pipeline
}
func (bc *ByteChannel) ChannelRead() {

}
type ChannelHandler interface {
	ChannelRead(ctx Context)
}