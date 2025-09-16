-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Máy chủ: 127.0.0.1
-- Thời gian đã tạo: Th9 16, 2025 lúc 07:38 AM
-- Phiên bản máy phục vụ: 10.4.32-MariaDB
-- Phiên bản PHP: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Cơ sở dữ liệu: `ql_cuahangdungcu`
--

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `bienthe`
--

CREATE TABLE `bienthe` (
  `MaBienThe` int(11) NOT NULL,
  `MaHangHoa` int(11) NOT NULL,
  `Size` varchar(10) NOT NULL,
  `Gia` decimal(10,2) NOT NULL,
  `SoLuongTon` int(11) DEFAULT 0,
  `TrangThai` enum('DangBan','NgungBan') DEFAULT 'DangBan'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Đang đổ dữ liệu cho bảng `bienthe`
--

INSERT INTO `bienthe` (`MaBienThe`, `MaHangHoa`, `Size`, `Gia`, `SoLuongTon`, `TrangThai`) VALUES
(4, 1, '40', 2500000.00, 10, 'DangBan'),
(5, 1, '41', 2500000.00, 5, 'DangBan'),
(6, 1, '42', 2600000.00, 0, 'DangBan'),
(7, 2, '39', 3000000.00, 0, 'DangBan'),
(8, 2, '40', 3000000.00, 0, 'DangBan'),
(9, 3, '41', 2200000.00, 0, 'DangBan'),
(10, 3, '42', 2200000.00, 0, 'DangBan'),
(11, 3, '43', 2300000.00, 0, 'DangBan'),
(12, 4, '38', 800000.00, 0, 'DangBan'),
(13, 4, '39', 800000.00, 0, 'DangBan'),
(14, 4, '40', 850000.00, 0, 'DangBan'),
(15, 5, '38', 750000.00, 0, 'DangBan'),
(16, 5, '39', 750000.00, 0, 'DangBan'),
(17, 5, '40', 800000.00, 0, 'DangBan');

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `chitietchucnang`
--

CREATE TABLE `chitietchucnang` (
  `MaChiTietChucNang` int(11) NOT NULL,
  `MaChucNang` int(11) NOT NULL,
  `TenChiTietChucNang` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Đang đổ dữ liệu cho bảng `chitietchucnang`
--

INSERT INTO `chitietchucnang` (`MaChiTietChucNang`, `MaChucNang`, `TenChiTietChucNang`) VALUES
(1, 1, 'Xem'),
(2, 11, 'Xem'),
(3, 6, 'Xem'),
(4, 6, 'Xử lý'),
(5, 8, 'Xem'),
(6, 8, 'Xử lý'),
(7, 12, 'Xem'),
(8, 12, 'Xử lý'),
(9, 2, 'Xem'),
(10, 2, 'Thêm'),
(11, 2, 'Sửa'),
(12, 2, 'Xóa'),
(13, 3, 'Xem'),
(14, 3, 'Thêm'),
(15, 3, 'Sửa'),
(16, 3, 'Xóa'),
(17, 4, 'Xem'),
(18, 4, 'Thêm'),
(19, 4, 'Sửa'),
(20, 4, 'Xóa'),
(21, 5, 'Xem'),
(22, 5, 'Thêm'),
(23, 5, 'Sửa'),
(24, 5, 'Xóa'),
(25, 7, 'Xem'),
(26, 7, 'Thêm'),
(27, 7, 'Sửa'),
(28, 7, 'Xóa'),
(29, 9, 'Xem'),
(30, 9, 'Thêm'),
(31, 9, 'Sửa'),
(32, 9, 'Xóa'),
(33, 10, 'Xem'),
(34, 10, 'Thêm'),
(35, 10, 'Sửa'),
(36, 10, 'Xóa');

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `chitietdonhang`
--

CREATE TABLE `chitietdonhang` (
  `MaChiTiet` int(11) NOT NULL,
  `MaDonHang` int(11) NOT NULL,
  `MaSanPham` int(11) NOT NULL,
  `GiaBan` decimal(12,2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `chitietphieunhap`
--

CREATE TABLE `chitietphieunhap` (
  `MaChiTiet` int(11) NOT NULL,
  `MaPhieuNhap` int(11) NOT NULL,
  `MaBienthe` int(11) NOT NULL,
  `SoLuong` int(11) NOT NULL,
  `GiaNhap` decimal(12,2) NOT NULL,
  `NgaySanXuat` date DEFAULT NULL,
  `ThoiGianBaoHanh` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Đang đổ dữ liệu cho bảng `chitietphieunhap`
--

INSERT INTO `chitietphieunhap` (`MaChiTiet`, `MaPhieuNhap`, `MaBienthe`, `SoLuong`, `GiaNhap`, `NgaySanXuat`, `ThoiGianBaoHanh`) VALUES
(1, 1, 4, 10, 1200000.00, '2025-08-01', 12),
(2, 1, 5, 5, 1200000.00, '2025-08-01', 12);

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `danhgia`
--

CREATE TABLE `danhgia` (
  `MaDanhGia` int(11) NOT NULL,
  `MaSanPham` int(11) NOT NULL,
  `MaNguoiDung` int(11) NOT NULL,
  `Diem` int(11) NOT NULL,
  `NoiDung` varchar(255) DEFAULT NULL,
  `TrangThai` varchar(50) DEFAULT 'Chưa duyệt',
  `NgayDanhGia` datetime DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Đang đổ dữ liệu cho bảng `danhgia`
--

INSERT INTO `danhgia` (`MaDanhGia`, `MaSanPham`, `MaNguoiDung`, `Diem`, `NoiDung`, `TrangThai`, `NgayDanhGia`) VALUES
(1, 1, 3, 5, 'Giày đẹp, vừa vặn, chất lượng tốt', 'Chưa duyệt', '2025-09-16 11:59:42');

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `danhmuc`
--

CREATE TABLE `danhmuc` (
  `MaDanhMuc` int(11) NOT NULL,
  `TenDanhMuc` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Đang đổ dữ liệu cho bảng `danhmuc`
--

INSERT INTO `danhmuc` (`MaDanhMuc`, `TenDanhMuc`) VALUES
(4, 'Dép'),
(1, 'Giày chạy bộ'),
(3, 'Giày thể thao'),
(2, 'Giày đá bóng'),
(5, 'Sandal');

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `donhang`
--

CREATE TABLE `donhang` (
  `MaDonHang` int(11) NOT NULL,
  `MaNguoiDung` int(11) NOT NULL,
  `NgayTao` datetime DEFAULT current_timestamp(),
  `TrangThai` varchar(50) DEFAULT 'Đang xử lý',
  `TongTien` decimal(12,2) DEFAULT 0.00,
  `TinhThanh` varchar(100) DEFAULT NULL,
  `QuanHuyen` varchar(100) DEFAULT NULL,
  `PhuongXa` varchar(100) DEFAULT NULL,
  `DuongSoNha` varchar(150) DEFAULT NULL,
  `PhuongThucThanhToan` varchar(50) DEFAULT 'Tiền mặt'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `giohang`
--

CREATE TABLE `giohang` (
  `MaNguoiDung` int(11) NOT NULL,
  `MaBienThe` int(11) NOT NULL,
  `SoLuong` int(11) NOT NULL DEFAULT 1
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `hang`
--

CREATE TABLE `hang` (
  `MaHang` int(11) NOT NULL,
  `TenHang` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Đang đổ dữ liệu cho bảng `hang`
--

INSERT INTO `hang` (`MaHang`, `TenHang`) VALUES
(2, 'Adidas'),
(5, 'New Balance'),
(1, 'Nike'),
(3, 'Puma'),
(4, 'Reebok');

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `hanghoa`
--

CREATE TABLE `hanghoa` (
  `MaHangHoa` int(11) NOT NULL,
  `TenHangHoa` varchar(100) NOT NULL,
  `MaHang` int(11) NOT NULL,
  `MaDanhMuc` int(11) NOT NULL,
  `Mau` varchar(50) NOT NULL,
  `MoTa` varchar(255) DEFAULT NULL,
  `TrangThai` enum('DangBan','NgungBan') DEFAULT 'DangBan',
  `MaKhuyenMai` int(11) DEFAULT NULL,
  `AnhDaiDien` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Đang đổ dữ liệu cho bảng `hanghoa`
--

INSERT INTO `hanghoa` (`MaHangHoa`, `TenHangHoa`, `MaHang`, `MaDanhMuc`, `Mau`, `MoTa`, `TrangThai`, `MaKhuyenMai`, `AnhDaiDien`) VALUES
(1, 'Nike Air Zoom', 1, 1, 'Đỏ', 'Giày chạy bộ siêu nhẹ', 'DangBan', 1, 'nike_air_zoom_do.jpg'),
(2, 'Adidas Predator', 2, 2, 'Xanh', 'Giày đá bóng chính hãng', 'DangBan', 2, 'adidas_predator_xanh.jpg'),
(3, 'Puma RS-X', 3, 3, 'Trắng', 'Giày thể thao năng động', 'DangBan', NULL, 'puma_rsx_trang.jpg'),
(4, 'Nike Kawa Slide', 1, 4, 'Xám', 'Dép thoải mái', 'DangBan', NULL, 'nike_kawa_slide_xam.jpg'),
(5, 'Adidas Adilette', 2, 5, 'Đen', 'Sandal nam nữ', 'DangBan', NULL, 'adidas_adilette_den.jpg');

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `khuyenmai`
--

CREATE TABLE `khuyenmai` (
  `MaKhuyenMai` int(11) NOT NULL,
  `TenKhuyenMai` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `MoTa` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `GiaTri` decimal(5,2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Đang đổ dữ liệu cho bảng `khuyenmai`
--

INSERT INTO `khuyenmai` (`MaKhuyenMai`, `TenKhuyenMai`, `MoTa`, `GiaTri`) VALUES
(1, 'Mua 1 tặng 10%', 'Khuyến mãi đặc biệt mùa hè', 10.00),
(2, 'Giảm giá cuối tuần', 'Áp dụng cho giày chạy bộ', 15.00);

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `machucnang`
--

CREATE TABLE `machucnang` (
  `MaChucNang` int(11) NOT NULL,
  `TenChucNang` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Đang đổ dữ liệu cho bảng `machucnang`
--

INSERT INTO `machucnang` (`MaChucNang`, `TenChucNang`) VALUES
(1, 'Trang chủ'),
(2, 'Quản lý khuyến mãi'),
(3, 'Quản lý hãng'),
(4, 'Quản lý nhà cung cấp'),
(5, 'Quản lý phiếu nhập'),
(6, 'Quản lý hàng hóa'),
(7, 'Quản lý danh mục'),
(8, 'Quản lý đơn hàng'),
(9, 'Quản lý người dùng'),
(10, 'Quản lý phân quyền'),
(11, 'Tra cứu sản phẩm'),
(12, 'Quản lý đánh giá');

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `nguoidung`
--

CREATE TABLE `nguoidung` (
  `MaNguoiDung` int(11) NOT NULL,
  `TenDangNhap` varchar(50) NOT NULL,
  `MatKhau` varchar(255) NOT NULL,
  `HoTen` varchar(100) NOT NULL,
  `Email` varchar(100) NOT NULL,
  `SoDienThoai` varchar(15) DEFAULT NULL,
  `TinhThanh` varchar(100) DEFAULT NULL,
  `QuanHuyen` varchar(100) DEFAULT NULL,
  `PhuongXa` varchar(100) DEFAULT NULL,
  `DuongSoNha` varchar(150) DEFAULT NULL,
  `MaQuyen` int(11) DEFAULT NULL,
  `NgayTao` datetime DEFAULT current_timestamp(),
  `NgayCapNhat` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Đang đổ dữ liệu cho bảng `nguoidung`
--

INSERT INTO `nguoidung` (`MaNguoiDung`, `TenDangNhap`, `MatKhau`, `HoTen`, `Email`, `SoDienThoai`, `TinhThanh`, `QuanHuyen`, `PhuongXa`, `DuongSoNha`, `MaQuyen`, `NgayTao`, `NgayCapNhat`) VALUES
(1, 'admin', '5994471abb01112afcc18159f6cc74b4f511b99806da59b3caf5a9c173cacfc5', 'Nguyễn Văn Admin', 'admin@shop.com', '0900000000', 'TP. Hồ Chí Minh', 'Quận 1', 'Phường Bến Nghé', 'Số 1 Nguyễn Huệ', 1, '2025-09-16 09:57:10', '2025-09-16 09:57:10'),
(2, 'nhanvien1', '5994471abb01112afcc18159f6cc74b4f511b99806da59b3caf5a9c173cacfc5', 'Trần Thị Nhân Viên', 'nhanvien1@shop.com', '0911111111', 'TP. Hà Nội', 'Quận Hoàn Kiếm', 'Phường Hàng Trống', 'Số 10 Hàng Bài', 2, '2025-09-16 09:57:10', '2025-09-16 09:57:10'),
(3, 'khach1', '5994471abb01112afcc18159f6cc74b4f511b99806da59b3caf5a9c173cacfc5', 'Lê Văn Khách', 'khach1@shop.com', '0922222222', 'TP. Đà Nẵng', 'Quận Hải Châu', 'Phường Thạch Thang', 'Số 20 Lê Duẩn', 3, '2025-09-16 09:57:10', '2025-09-16 09:57:10');

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `nhacungcap`
--

CREATE TABLE `nhacungcap` (
  `MaNCC` int(11) NOT NULL,
  `TenNCC` varchar(100) NOT NULL,
  `DiaChi` varchar(255) DEFAULT NULL,
  `SoDienThoai` varchar(20) DEFAULT NULL,
  `Email` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Đang đổ dữ liệu cho bảng `nhacungcap`
--

INSERT INTO `nhacungcap` (`MaNCC`, `TenNCC`, `DiaChi`, `SoDienThoai`, `Email`) VALUES
(1, 'Adidas', '123 Duong Le Lai, TP.HCM', '0900000000', NULL);

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `phanquyen`
--

CREATE TABLE `phanquyen` (
  `MaQuyen` int(11) NOT NULL,
  `MaChiTietChucNang` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Đang đổ dữ liệu cho bảng `phanquyen`
--

INSERT INTO `phanquyen` (`MaQuyen`, `MaChiTietChucNang`) VALUES
(1, 1),
(1, 2),
(1, 3),
(1, 4),
(1, 5),
(1, 6),
(1, 7),
(1, 8),
(1, 9),
(1, 10),
(1, 11),
(1, 12),
(1, 13),
(1, 14),
(1, 15),
(1, 16),
(1, 17),
(1, 18),
(1, 19),
(1, 20),
(1, 21),
(1, 22),
(1, 23),
(1, 24),
(1, 25),
(1, 26),
(1, 27),
(1, 28),
(1, 29),
(1, 30),
(1, 31),
(1, 32),
(1, 33),
(1, 34),
(1, 35),
(1, 36);

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `phieunhap`
--

CREATE TABLE `phieunhap` (
  `MaPhieuNhap` int(11) NOT NULL,
  `MaNguoiDung` int(11) NOT NULL,
  `MaNCC` int(11) NOT NULL,
  `NgayNhap` datetime DEFAULT current_timestamp(),
  `TrangThai` varchar(50) DEFAULT 'Chưa duyệt',
  `MoTa` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Đang đổ dữ liệu cho bảng `phieunhap`
--

INSERT INTO `phieunhap` (`MaPhieuNhap`, `MaNguoiDung`, `MaNCC`, `NgayNhap`, `TrangThai`, `MoTa`) VALUES
(1, 2, 1, '2025-09-16 10:00:00', 'Đã duyệt', 'Nhập lô giày Nike đỏ');

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `quyen`
--

CREATE TABLE `quyen` (
  `MaQuyen` int(11) NOT NULL,
  `TenQuyen` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Đang đổ dữ liệu cho bảng `quyen`
--

INSERT INTO `quyen` (`MaQuyen`, `TenQuyen`) VALUES
(1, 'Admin'),
(3, 'Khàch hàng'),
(2, 'Nhân viên');

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `sanpham`
--

CREATE TABLE `sanpham` (
  `MaSanPham` int(11) NOT NULL,
  `MaChiTietPhieuNhap` int(11) NOT NULL,
  `Seri` varchar(50) NOT NULL,
  `TrangThai` varchar(50) DEFAULT 'Chưa bán'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Đang đổ dữ liệu cho bảng `sanpham`
--

INSERT INTO `sanpham` (`MaSanPham`, `MaChiTietPhieuNhap`, `Seri`, `TrangThai`) VALUES
(1, 1, 'NIKE-RED-40-001', 'Chưa bán'),
(2, 1, 'NIKE-RED-40-002', 'Chưa bán'),
(3, 1, 'NIKE-RED-40-003', 'Chưa bán'),
(4, 1, 'NIKE-RED-40-004', 'Chưa bán'),
(5, 1, 'NIKE-RED-40-005', 'Chưa bán'),
(6, 1, 'NIKE-RED-40-006', 'Chưa bán'),
(7, 1, 'NIKE-RED-40-007', 'Chưa bán'),
(8, 1, 'NIKE-RED-40-008', 'Chưa bán'),
(9, 1, 'NIKE-RED-40-009', 'Chưa bán'),
(10, 1, 'NIKE-RED-40-010', 'Chưa bán'),
(11, 2, 'NIKE-RED-41-001', 'Chưa bán'),
(12, 2, 'NIKE-RED-41-002', 'Chưa bán'),
(13, 2, 'NIKE-RED-41-003', 'Chưa bán'),
(14, 2, 'NIKE-RED-41-004', 'Chưa bán'),
(15, 2, 'NIKE-RED-41-005', 'Chưa bán');

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `thanhtoanonline`
--

CREATE TABLE `thanhtoanonline` (
  `MaThanhToan` int(11) NOT NULL,
  `MaDonHang` int(11) NOT NULL,
  `MaGiaoDich` varchar(100) DEFAULT NULL,
  `TrangThai` varchar(50) DEFAULT 'Chưa thanh toán',
  `NgayThanhToan` datetime DEFAULT NULL,
  `TongTien` decimal(12,2) NOT NULL,
  `MoTaGiaoDich` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Chỉ mục cho các bảng đã đổ
--

--
-- Chỉ mục cho bảng `bienthe`
--
ALTER TABLE `bienthe`
  ADD PRIMARY KEY (`MaBienThe`),
  ADD KEY `MaHangHoa` (`MaHangHoa`);

--
-- Chỉ mục cho bảng `chitietchucnang`
--
ALTER TABLE `chitietchucnang`
  ADD PRIMARY KEY (`MaChiTietChucNang`),
  ADD KEY `MaChucNang` (`MaChucNang`);

--
-- Chỉ mục cho bảng `chitietdonhang`
--
ALTER TABLE `chitietdonhang`
  ADD PRIMARY KEY (`MaChiTiet`),
  ADD KEY `MaDonHang` (`MaDonHang`),
  ADD KEY `MaSanPham` (`MaSanPham`);

--
-- Chỉ mục cho bảng `chitietphieunhap`
--
ALTER TABLE `chitietphieunhap`
  ADD PRIMARY KEY (`MaChiTiet`),
  ADD KEY `MaPhieuNhap` (`MaPhieuNhap`),
  ADD KEY `MaBienthe` (`MaBienthe`);

--
-- Chỉ mục cho bảng `danhgia`
--
ALTER TABLE `danhgia`
  ADD PRIMARY KEY (`MaDanhGia`),
  ADD KEY `MaSanPham` (`MaSanPham`),
  ADD KEY `MaNguoiDung` (`MaNguoiDung`);

--
-- Chỉ mục cho bảng `danhmuc`
--
ALTER TABLE `danhmuc`
  ADD PRIMARY KEY (`MaDanhMuc`),
  ADD UNIQUE KEY `TenDanhMuc` (`TenDanhMuc`);

--
-- Chỉ mục cho bảng `donhang`
--
ALTER TABLE `donhang`
  ADD PRIMARY KEY (`MaDonHang`),
  ADD KEY `MaNguoiDung` (`MaNguoiDung`);

--
-- Chỉ mục cho bảng `giohang`
--
ALTER TABLE `giohang`
  ADD PRIMARY KEY (`MaNguoiDung`,`MaBienThe`),
  ADD KEY `MaBienThe` (`MaBienThe`);

--
-- Chỉ mục cho bảng `hang`
--
ALTER TABLE `hang`
  ADD PRIMARY KEY (`MaHang`),
  ADD UNIQUE KEY `TenHang` (`TenHang`);

--
-- Chỉ mục cho bảng `hanghoa`
--
ALTER TABLE `hanghoa`
  ADD PRIMARY KEY (`MaHangHoa`),
  ADD KEY `MaHang` (`MaHang`),
  ADD KEY `MaDanhMuc` (`MaDanhMuc`),
  ADD KEY `MaKhuyenMai` (`MaKhuyenMai`);

--
-- Chỉ mục cho bảng `khuyenmai`
--
ALTER TABLE `khuyenmai`
  ADD PRIMARY KEY (`MaKhuyenMai`);

--
-- Chỉ mục cho bảng `machucnang`
--
ALTER TABLE `machucnang`
  ADD PRIMARY KEY (`MaChucNang`);

--
-- Chỉ mục cho bảng `nguoidung`
--
ALTER TABLE `nguoidung`
  ADD PRIMARY KEY (`MaNguoiDung`),
  ADD UNIQUE KEY `TenDangNhap` (`TenDangNhap`),
  ADD UNIQUE KEY `Email` (`Email`),
  ADD KEY `MaQuyen` (`MaQuyen`);

--
-- Chỉ mục cho bảng `nhacungcap`
--
ALTER TABLE `nhacungcap`
  ADD PRIMARY KEY (`MaNCC`);

--
-- Chỉ mục cho bảng `phanquyen`
--
ALTER TABLE `phanquyen`
  ADD PRIMARY KEY (`MaQuyen`,`MaChiTietChucNang`),
  ADD KEY `MaChiTietChucNang` (`MaChiTietChucNang`);

--
-- Chỉ mục cho bảng `phieunhap`
--
ALTER TABLE `phieunhap`
  ADD PRIMARY KEY (`MaPhieuNhap`),
  ADD KEY `MaNguoiDung` (`MaNguoiDung`),
  ADD KEY `MaNCC` (`MaNCC`);

--
-- Chỉ mục cho bảng `quyen`
--
ALTER TABLE `quyen`
  ADD PRIMARY KEY (`MaQuyen`),
  ADD UNIQUE KEY `TenQuyen` (`TenQuyen`);

--
-- Chỉ mục cho bảng `sanpham`
--
ALTER TABLE `sanpham`
  ADD PRIMARY KEY (`MaSanPham`),
  ADD UNIQUE KEY `Seri` (`Seri`),
  ADD KEY `MaChiTietPhieuNhap` (`MaChiTietPhieuNhap`);

--
-- Chỉ mục cho bảng `thanhtoanonline`
--
ALTER TABLE `thanhtoanonline`
  ADD PRIMARY KEY (`MaThanhToan`),
  ADD KEY `MaDonHang` (`MaDonHang`);

--
-- AUTO_INCREMENT cho các bảng đã đổ
--

--
-- AUTO_INCREMENT cho bảng `bienthe`
--
ALTER TABLE `bienthe`
  MODIFY `MaBienThe` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;

--
-- AUTO_INCREMENT cho bảng `chitietchucnang`
--
ALTER TABLE `chitietchucnang`
  MODIFY `MaChiTietChucNang` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=37;

--
-- AUTO_INCREMENT cho bảng `chitietdonhang`
--
ALTER TABLE `chitietdonhang`
  MODIFY `MaChiTiet` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT cho bảng `chitietphieunhap`
--
ALTER TABLE `chitietphieunhap`
  MODIFY `MaChiTiet` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT cho bảng `danhgia`
--
ALTER TABLE `danhgia`
  MODIFY `MaDanhGia` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT cho bảng `danhmuc`
--
ALTER TABLE `danhmuc`
  MODIFY `MaDanhMuc` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT cho bảng `donhang`
--
ALTER TABLE `donhang`
  MODIFY `MaDonHang` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT cho bảng `hang`
--
ALTER TABLE `hang`
  MODIFY `MaHang` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=19;

--
-- AUTO_INCREMENT cho bảng `hanghoa`
--
ALTER TABLE `hanghoa`
  MODIFY `MaHangHoa` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT cho bảng `khuyenmai`
--
ALTER TABLE `khuyenmai`
  MODIFY `MaKhuyenMai` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT cho bảng `machucnang`
--
ALTER TABLE `machucnang`
  MODIFY `MaChucNang` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT cho bảng `nguoidung`
--
ALTER TABLE `nguoidung`
  MODIFY `MaNguoiDung` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT cho bảng `nhacungcap`
--
ALTER TABLE `nhacungcap`
  MODIFY `MaNCC` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT cho bảng `phieunhap`
--
ALTER TABLE `phieunhap`
  MODIFY `MaPhieuNhap` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT cho bảng `quyen`
--
ALTER TABLE `quyen`
  MODIFY `MaQuyen` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT cho bảng `sanpham`
--
ALTER TABLE `sanpham`
  MODIFY `MaSanPham` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- AUTO_INCREMENT cho bảng `thanhtoanonline`
--
ALTER TABLE `thanhtoanonline`
  MODIFY `MaThanhToan` int(11) NOT NULL AUTO_INCREMENT;

--
-- Các ràng buộc cho các bảng đã đổ
--

--
-- Các ràng buộc cho bảng `bienthe`
--
ALTER TABLE `bienthe`
  ADD CONSTRAINT `bienthe_ibfk_1` FOREIGN KEY (`MaHangHoa`) REFERENCES `hanghoa` (`MaHangHoa`);

--
-- Các ràng buộc cho bảng `chitietchucnang`
--
ALTER TABLE `chitietchucnang`
  ADD CONSTRAINT `chitietchucnang_ibfk_1` FOREIGN KEY (`MaChucNang`) REFERENCES `machucnang` (`MaChucNang`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Các ràng buộc cho bảng `chitietdonhang`
--
ALTER TABLE `chitietdonhang`
  ADD CONSTRAINT `chitietdonhang_ibfk_1` FOREIGN KEY (`MaDonHang`) REFERENCES `donhang` (`MaDonHang`),
  ADD CONSTRAINT `chitietdonhang_ibfk_2` FOREIGN KEY (`MaSanPham`) REFERENCES `sanpham` (`MaSanPham`);

--
-- Các ràng buộc cho bảng `chitietphieunhap`
--
ALTER TABLE `chitietphieunhap`
  ADD CONSTRAINT `chitietphieunhap_ibfk_1` FOREIGN KEY (`MaPhieuNhap`) REFERENCES `phieunhap` (`MaPhieuNhap`),
  ADD CONSTRAINT `chitietphieunhap_ibfk_2` FOREIGN KEY (`MaBienthe`) REFERENCES `bienthe` (`MaBienThe`);

--
-- Các ràng buộc cho bảng `danhgia`
--
ALTER TABLE `danhgia`
  ADD CONSTRAINT `danhgia_ibfk_1` FOREIGN KEY (`MaSanPham`) REFERENCES `sanpham` (`MaSanPham`),
  ADD CONSTRAINT `danhgia_ibfk_2` FOREIGN KEY (`MaNguoiDung`) REFERENCES `nguoidung` (`MaNguoiDung`);

--
-- Các ràng buộc cho bảng `donhang`
--
ALTER TABLE `donhang`
  ADD CONSTRAINT `donhang_ibfk_1` FOREIGN KEY (`MaNguoiDung`) REFERENCES `nguoidung` (`MaNguoiDung`);

--
-- Các ràng buộc cho bảng `giohang`
--
ALTER TABLE `giohang`
  ADD CONSTRAINT `giohang_ibfk_1` FOREIGN KEY (`MaNguoiDung`) REFERENCES `nguoidung` (`MaNguoiDung`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `giohang_ibfk_2` FOREIGN KEY (`MaBienThe`) REFERENCES `bienthe` (`MaBienThe`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Các ràng buộc cho bảng `hanghoa`
--
ALTER TABLE `hanghoa`
  ADD CONSTRAINT `hanghoa_ibfk_1` FOREIGN KEY (`MaHang`) REFERENCES `hang` (`MaHang`),
  ADD CONSTRAINT `hanghoa_ibfk_2` FOREIGN KEY (`MaDanhMuc`) REFERENCES `danhmuc` (`MaDanhMuc`),
  ADD CONSTRAINT `hanghoa_ibfk_3` FOREIGN KEY (`MaKhuyenMai`) REFERENCES `khuyenmai` (`MaKhuyenMai`);

--
-- Các ràng buộc cho bảng `nguoidung`
--
ALTER TABLE `nguoidung`
  ADD CONSTRAINT `nguoidung_ibfk_1` FOREIGN KEY (`MaQuyen`) REFERENCES `quyen` (`MaQuyen`);

--
-- Các ràng buộc cho bảng `phanquyen`
--
ALTER TABLE `phanquyen`
  ADD CONSTRAINT `phanquyen_ibfk_1` FOREIGN KEY (`MaQuyen`) REFERENCES `quyen` (`MaQuyen`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `phanquyen_ibfk_2` FOREIGN KEY (`MaChiTietChucNang`) REFERENCES `chitietchucnang` (`MaChiTietChucNang`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Các ràng buộc cho bảng `phieunhap`
--
ALTER TABLE `phieunhap`
  ADD CONSTRAINT `phieunhap_ibfk_1` FOREIGN KEY (`MaNguoiDung`) REFERENCES `nguoidung` (`MaNguoiDung`),
  ADD CONSTRAINT `phieunhap_ibfk_2` FOREIGN KEY (`MaNCC`) REFERENCES `nhacungcap` (`MaNCC`);

--
-- Các ràng buộc cho bảng `sanpham`
--
ALTER TABLE `sanpham`
  ADD CONSTRAINT `sanpham_ibfk_1` FOREIGN KEY (`MaChiTietPhieuNhap`) REFERENCES `chitietphieunhap` (`MaChiTiet`);

--
-- Các ràng buộc cho bảng `thanhtoanonline`
--
ALTER TABLE `thanhtoanonline`
  ADD CONSTRAINT `thanhtoanonline_ibfk_1` FOREIGN KEY (`MaDonHang`) REFERENCES `donhang` (`MaDonHang`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
