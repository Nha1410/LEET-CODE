package main

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	// Khởi tạo mảng chuỗi cho từng dòng
	rows := make([]string, numRows)
	curRow := 0
	goingDown := false

	for _, c := range s {
		rows[curRow] += string(c)
		// Đổi chiều khi tới dòng đầu hoặc dòng cuối
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}
		// Di chuyển dòng
		if goingDown {
			curRow++
		} else {
			curRow--
		}
	}

	// Nối tất cả dòng lại
	result := ""
	for _, row := range rows {
		result += row
	}

	return result
}

// PAYPALISHIRING -> PAHNAPLSIIGYIR

// tu duy cho bài toán
// 1. Chia chuỗi thành các dòng theo quy luật zigzag
// 2. Duyệt từng ký tự trong chuỗi, xác định dòng hiện tại và thêm ký tự vào dòng đó
// 3. Khi đến đầu hoặc cuối dòng, đổi chiều di chuyển
// 4. Nối tất cả các dòng lại thành chuỗi kết quả
// 5. Trả về chuỗi kết quả
// 6. Thời gian O(n), không gian O(n)
