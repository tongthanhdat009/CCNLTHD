package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"github.com/tongthanhdat009/CCNLTHD/internal/services"
	"fmt"
	"strings"
	"log"
)
type NhaCungCapHandler struct{
	service services.NhaCungCapService
}
func NewNhaCungCapHandler(service services.NhaCungCapService) *NhaCungCapHandler {
	return &NhaCungCapHandler{service: service}
}

func (h *NhaCungCapHandler) GetAllNhaCungCap(c *gin.Context) {
	NhaCungCaps, err := h.service.GetAllNhaCungCap()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch products"})
		return
	}
	c.JSON(http.StatusOK, NhaCungCaps)
}

func (ncc *NhaCungCapHandler) CreateNhaCungCap(c *gin.Context) {
    var nhacungcap models.NhaCungCap
    if err := c.ShouldBindJSON(&nhacungcap); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ", "details": err.Error()})
        return
    }

    if err := ncc.service.CreateNhaCungCap(&nhacungcap); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Nhà cung cấp đã được thêm thành công"})
}

func (ncc *NhaCungCapHandler) GetNhaCungCapByID(c *gin.Context) {
    // Lấy ID từ URL
    idParam := c.Param("id")
    var id int
    if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Mã nhà cung cấp không hợp lệ"})
        return
    }

    // Gọi service để lấy nhà cung cấp theo ID
    nhacungcap, err := ncc.service.GetNhaCungCapByID(id)
    if err != nil {
        if strings.Contains(err.Error(), "không tồn tại") {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lấy thông tin nhà cung cấp"})
        return
    }

    // Trả về kết quả
    c.JSON(http.StatusOK, nhacungcap)
}

func (ncc *NhaCungCapHandler) GetNhaCungCapByName(c *gin.Context) {
    // Lấy tên nhà cung cấp từ URL
    name := c.Param("tenncc")
    // Gọi service để lấy danh sách nhà cung cấp theo tên
    nhacungcaps, err := ncc.service.GetNhaCungCapByName(name)
    if err != nil {
        log.Printf("Error getting nha cung cap by name %s: %v", name, err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lấy thông tin nhà cung cấp"})
        return
    }

    // Trả về kết quả
    c.JSON(http.StatusOK, nhacungcaps)
}

func (ncc *NhaCungCapHandler) UpdateNhaCungCap(c *gin.Context) {
	var nhacungcap models.NhaCungCap
	if err := c.ShouldBindJSON(&nhacungcap); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ", "details": err.Error()})
        return
    }
	if err := ncc.service.UpdateNhaCungCap(&nhacungcap); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nhacungcap)
}

func (ncc *NhaCungCapHandler) DeleteNhaCungCap(c *gin.Context) {
	var id int
	if _, err := fmt.Sscanf(c.Param("id"), "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mã nhà cung cấp không hợp lệ"})
		return
	}
	if err := ncc.service.DeleteNhaCungCap(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Nhà cung cấp đã được xóa thành công"})
}