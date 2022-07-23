package mykey

import (
	"SK-builder/internal/conf"

	"github.com/bwmarrin/snowflake"
)

type SnowNode struct {
	Node *snowflake.Node
}

func NewSnowNode(c *conf.Server) *SnowNode {
	node, err := snowflake.NewNode(int64(c.Node))
	if err != nil {
		return nil
	}

	return &SnowNode{
		Node: node,
	}
}

func (s *SnowNode) GetID() ([8]byte, int64) {
	id := s.Node.Generate()
	return id.IntBytes(), id.Int64()
}
