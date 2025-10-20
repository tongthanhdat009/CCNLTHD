package handlers

import (
    "net/http"
    "github.com/tongthanhdat009/CCNLTHD/internal/services"
    "github.com/tongthanhdat009/CCNLTHD/internal/models"
    "github.com/gin-gonic/gin"
    "fmt"
    "log"
)

type HangHandler struct {
    service services.HangService
}

func NewHangHandler(service services.HangService) *HangHandler {
    return &HangHandler{service: service}
}

func (h *HangHandler) GetAllHang(c *gin.Context) {
    Hangs, err := h.service.GetAllHang()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch products"})
        return
    }
    c.JSON(http.StatusOK, Hangs)
}

func (h *HangHandler) DeleteHang(c *gin.Context) {
    idParam := c.Param("id")
    var id int
    if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Mã hãng không tồn tại hoặc không hợp lệ"})
        return
    }

    if err := h.service.DeleteHang(id); err != nil {
        log.Printf("Error deleting hang with id %d: %v", id, err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Hãng đã được xóa thành công"})
}

func (h *HangHandler) CreateHang(c *gin.Context) {
    var hang models.Hang
    if err := c.ShouldBindJSON(&hang); err != nil {
        log.Printf("Error binding JSON: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
        return
    }

    if err := h.service.CreateHang(&hang); err != nil {
        log.Printf("Error creating hang: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, hang)
}

func (h *HangHandler) UpdateHang(c *gin.Context) {
    var hang models.Hang
    if err := c.ShouldBindJSON(&hang); err != nil {
        log.Printf("Error binding JSON: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
        return
    }

    if err := h.service.UpdateHang(&hang); err != nil {
        log.Printf("Error updating hang: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, hang)
}

func (h *HangHandler) GetHangByID(c *gin.Context) {
    idParam := c.Param("id")
    var id int
    if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Mã hãng không hợp lệ"})
        return
    }

    hang, err := h.service.GetHangByID(id)
    if err != nil {
        log.Printf("Error getting hang by ID %d: %v", id, err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, hang)
}

func (h *HangHandler) GetHangByName(c *gin.Context) {
    name := c.Param("tenhang")
    if name == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Tên hãng không được để trống"})
        return
    }

    hangs, err := h.service.GetHangByName(name)
    if err != nil {
        log.Printf("Error getting hang by name %s: %v", name, err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if len(hangs) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy hãng nào với tên đã cho"})
        return
    }
    c.JSON(http.StatusOK, hangs)
}