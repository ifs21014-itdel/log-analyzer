package usecase

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ifs21014-itdel/log-analyzer/internal/domain"
)

// readFileLines baca file dan kembalikan array string per line
func readFileLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func parseLineToLogAnalysis(line string) domain.LogAnalysis {
	line = strings.TrimSpace(line)
	if line == "" {
		return domain.LogAnalysis{}
	}

	// buang timestamp di awal, contoh: [2025-10-17 10:00:00]
	if idx := strings.Index(line, "]"); idx != -1 {
		line = strings.TrimSpace(line[idx+1:])
	}

	// setelah buang timestamp: "GET /api/users 200 120ms 192.168.1.1"
	parts := strings.Fields(line)
	if len(parts) < 5 {
		fmt.Println("SKIP, format tidak valid:", line)
		return domain.LogAnalysis{}
	}

	status := parts[2]
	respTimeRaw := parts[3]

	// hapus akhiran 'ms' jika ada
	respTimeStr := strings.TrimSuffix(respTimeRaw, "ms")

	// parse angka response time
	respTime, err := strconv.Atoi(respTimeStr)
	if err != nil {
		fmt.Println("Gagal parse response time:", respTimeStr, "line:", line)
		respTime = 0
	}

	errorCount := 0
	if len(status) > 0 && (status[0] == '4' || status[0] == '5') {
		errorCount = 1
	}

	return domain.LogAnalysis{
		Filename:        line,
		TotalRequests:   1,
		UniqueIPs:       1,
		ErrorCount:      errorCount,
		AverageResponse: float64(respTime),
	}
}
