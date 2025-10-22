package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time" // Thêm import này

	"github.com/gin-gonic/gin"
	"github.com/tongthanhdat009/CCNLTHD/internal/models"
	"github.com/tongthanhdat009/CCNLTHD/internal/services"
)

type HangHoaHandler struct {
    service services.HangHoaService
}

func NewHangHoaHandler(service services.HangHoaService) *HangHoaHandler {
    return &HangHoaHandler{service: service}
}

func (h *HangHoaHandler) GetAllHangHoa(c *gin.Context) {
    limitStr := c.DefaultQuery("limit","3") //số bảng ghi cần lấy
    offsetStr := c.DefaultQuery("offset","0") //bắt đầu từ bảng ghi thứ bao nhiêu
    limit, _ := strconv.Atoi(limitStr)
    offset, _ := strconv.Atoi(offsetStr)
    hangHoas, err := h.service.GetAllHangHoa(limit, offset)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch products"})
        return
    }
    // Thêm domain vào đường dẫn ảnh
    for i := range hangHoas {
        if hangHoas[i].AnhDaiDien != "" {
            hangHoas[i].AnhDaiDien = "/AnhHangHoa/" + hangHoas[i].AnhDaiDien
        }
    }
    c.JSON(http.StatusOK, hangHoas)
}

func (h *HangHoaHandler) CreateHangHoa(c *gin.Context) {
    if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form"})
        return
    }

    var hh models.HangHoa
    if payload := c.PostForm("payload"); payload != "" {
        if err := json.Unmarshal([]byte(payload), &hh); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
            return
        }
    }

    if file, err := c.FormFile("anh_dai_dien"); err == nil {
        filename, err := saveHangHoaImage(c, file)
        log.Println(filename)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        // Lưu tên file vào database
        hh.AnhDaiDien = filename
    }

    if err := h.service.CreateHangHoa(&hh); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    // Trả về đường dẫn đầy đủ trong response
    response := hh
    if response.AnhDaiDien != "" {
        response.AnhDaiDien = "/AnhHangHoa/" + response.AnhDaiDien
    }
    
    c.JSON(http.StatusCreated, gin.H{
        "success": true,
        "message": "Tạo hàng hóa thành công",
        "data": response,
    })
}

func (h *HangHoaHandler) UpdateHangHoa(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
        return
    }

    if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form"})
        return
    }

    existing, err := h.service.GetHangHoaByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    var hh models.HangHoa
    if payload := c.Request.FormValue("payload"); payload != "" {
        if err := json.Unmarshal([]byte(payload), &hh); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
            return
        }
    }
    hh.MaHangHoa = id
    hh.AnhDaiDien = existing.AnhDaiDien // Giữ tên file ảnh cũ

    if file, err := c.FormFile("anh_dai_dien"); err == nil {
        filename, err := saveHangHoaImage(c, file)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
    
        // Lưu tên file mới vào database
        hh.AnhDaiDien = filename
    }

    if err := h.service.UpdateHangHoa(&hh); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    // Trả về đường dẫn đầy đủ trong response
    response := hh
    if response.AnhDaiDien != "" {
        response.AnhDaiDien = "/AnhHangHoa/" + response.AnhDaiDien
    }
    
    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "message": "Cập nhật hàng hóa thành công",
        "data": response,
    })
}

func saveHangHoaImage(c *gin.Context, file *multipart.FileHeader) (string, error) {
    ext := strings.ToLower(filepath.Ext(file.Filename))
    if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
        return "", errors.New("only jpg/jpeg/png allowed")
    }
    if file.Size > 5*1024*1024 {
        return "", errors.New("max size 5MB")
    }
    
    // Tạo thư mục nếu không tồn tại
    uploadDir := "./static/AnhHangHoa"
    if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
        if err := os.MkdirAll(uploadDir, 0755); err != nil {
            return "", fmt.Errorf("failed to create upload directory: %v", err)
        }
    }
    
    safe := filepath.Base(file.Filename)
    // Thêm timestamp vào tên file để tránh trùng lặp
    filename := time.Now().UnixNano()
    safe = fmt.Sprintf("%d_%s", filename, safe)
    
    dst := filepath.Join(uploadDir, safe)
    if err := c.SaveUploadedFile(file, dst); err != nil {
        return "", fmt.Errorf("failed to save file: %v", err)
    }
    
    // Chỉ trả về tên file, không kèm đường dẫn
    return safe, nil
}

func (h *HangHoaHandler) SearchHangHoa(c *gin.Context) {
    tenHangHoa := c.Query("ten_hang_hoa")
    tenDanhMuc := c.Query("ten_danh_muc")
    tenHang := c.Query("ten_hang")
    mau := c.Query("mau")
    trangThai := c.Query("trang_thai")
    maKhuyenMai := c.Query("ma_khuyen_mai")

    hangHoas, err := h.service.SearchHangHoa(tenHangHoa, tenDanhMuc, tenHang, mau, trangThai, maKhuyenMai)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    if len(hangHoas) == 0 {
        c.JSON(http.StatusOK, gin.H{"message": "Không có hàng hóa phù hợp"})
        return
    }
    c.JSON(http.StatusOK, hangHoas)
}