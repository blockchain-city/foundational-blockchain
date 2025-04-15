package api

import (
	"net/http"

	"github.com/blockchain-city/foundational-blockchain/internal/blockchain"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Router     *gin.Engine
	Blockchain *blockchain.Blockchain
}

func NewServer(bc *blockchain.Blockchain) *Server {
	r := gin.Default()
	server := &Server{Router: r, Blockchain: bc}
	server.setupRoutes()
	return server
}

func (s *Server) setupRoutes() {
	s.Router.GET("/chain", s.getChain)
	s.Router.POST("/mine", s.mineBlock)
}

func (s *Server) Run(addr string) {
	s.Router.Run(addr)
}

func (s *Server) getChain(c *gin.Context) {
	c.JSON(http.StatusOK, s.Blockchain.GetAllBlocks())
}

type MineRequest struct {
	Data string `json:"data" binding:"required"`
}

func (s *Server) mineBlock(c *gin.Context) {
	var req MineRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	newBlock := s.Blockchain.AddBlock(req.Data)
	if err := s.Blockchain.SaveToDisk(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save blockchain"})
		return
	}
	c.JSON(http.StatusCreated, newBlock)
}
