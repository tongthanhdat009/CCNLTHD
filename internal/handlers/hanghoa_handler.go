package handlers

import (
    "net/http"
    "os"
    "path/filepath"
    "strconv"
    "strings"
    "errors"

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
    hangHoas, err := h.service.GetAllHangHoa()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch products"})
        return
    }
    // Thêm domain vào đường dẫn ảnh (nếu cần)
    for i := range hangHoas {
        if hangHoas[i].AnhDaiDien != "" {
            hangHoas[i].AnhDaiDien = "http://localhost:8080/AnhHangHoa/" + hangHoas[i].AnhDaiDien
        }
    }
    c.JSON(http.StatusOK, hangHoas)
}

func (h *HangHoaHandler) saveImage(c *gin.Context, key string) (string, error) {
    file, err := c.FormFile(key)
    if err != nil {
        return "", nil // Không có file, không lỗi
    }

    // Kiểm tra loại file (chỉ cho phép jpg, png)
    ext := strings.ToLower(filepath.Ext(file.Filename))
    if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
        return "", errors.New("chỉ cho phép file ảnh JPG hoặc PNG")
    }

    // Kiểm tra kích thước (tối đa 5MB)
    if file.Size > 5*1024*1024 {
        return "", errors.New("kích thước ảnh tối đa 5MB")
    }

    // Tạo tên file an toàn (tránh path traversal)
    safeName := filepath.Base(file.Filename)
    dst := "./static/AnhHangHoa/" + safeName

    // Lưu file
    if err := c.SaveUploadedFile(file, dst); err != nil {
        return "", err
    }

    // Trả về đường dẫn tương đối
    return "/AnhHangHoa/" + safeName, nil
}

func (h *HangHoaHandler) CreateHangHoa(c *gin.Context) {
    var hh models.HangHoa

    // Lấy và lưu ảnh (nếu có)
    anhDaiDien, err := h.saveImage(c, "anh_dai_dien")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    hh.AnhDaiDien = anhDaiDien

    // Bind các trường khác từ form (hoặc JSON nếu client gửi JSON)
    // Nếu client gửi JSON + file, dùng c.ShouldBindJSON cho hh, nhưng ảnh riêng
    // Hoặc bind từ form: hh.TenHangHoa = c.PostForm("ten_hang_hoa"), v.v.
    // Ví dụ đơn giản: giả sử client gửi JSON, nhưng ảnh qua form
    if err := c.ShouldBindJSON(&hh); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
        return
    }

    if err := h.service.CreateHangHoa(&hh); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Thêm hàng hóa thành công"})
}

func (h *HangHoaHandler) UpdateHangHoa(c *gin.Context) {
    var hh models.HangHoa

    // Lấy id từ URL
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
        return
    }

    // Lấy và lưu ảnh mới (nếu có), xóa ảnh cũ nếu cần
    anhDaiDien, err := h.saveImage(c, "anh_dai_dien")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if anhDaiDien != "" {
        // Xóa ảnh cũ (nếu có)
        if hh.AnhDaiDien != "" {
            oldPath := strings.Replace(hh.AnhDaiDien, "/AnhHangHoa/", "./static/AnhHangHoa/", 1)
            os.Remove(oldPath)
        }
        hh.AnhDaiDien = anhDaiDien
    }

    // Bind các trường khác
    if err := c.ShouldBindJSON(&hh); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
        return
    }

    // Gán id
    hh.MaHangHoa = id

    if err := h.service.UpdateHangHoa(&hh); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Cập nhật hàng hóa thành công"})
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