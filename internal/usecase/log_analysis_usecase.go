package usecase

import (
	"errors"
	"fmt"
	"sync"

	"github.com/ifs21014-itdel/log-analyzer/internal/domain"
	"github.com/ifs21014-itdel/log-analyzer/internal/repository"
)

type LogAnalysisUsecase struct {
	repo repository.LogAnalysisRepository
}

func NewLogAnalysisUsecase(r repository.LogAnalysisRepository) *LogAnalysisUsecase {
	return &LogAnalysisUsecase{repo: r}
}

// CRUD
func (u *LogAnalysisUsecase) Create(a *domain.LogAnalysis) error {
	return u.repo.Create(a)
}

func (u *LogAnalysisUsecase) GetAll() ([]domain.LogAnalysis, error) {
	return u.repo.GetAll()
}

func (u *LogAnalysisUsecase) GetByID(id uint) (*domain.LogAnalysis, error) {
	a, err := u.repo.GetByID(id)
	if a == nil {
		return nil, errors.New("record not found")
	}
	return a, err
}

func (u *LogAnalysisUsecase) Update(a *domain.LogAnalysis) error {
	return u.repo.Update(a)
}

func (u *LogAnalysisUsecase) Delete(id uint) error {
	return u.repo.Delete(id)
}

// 🧠 ProcessLogs — concurrent log analyzer with progress logs
func (u *LogAnalysisUsecase) ProcessLogs(lines []string) (*domain.LogAnalysis, error) {
	totalRequests := 0
	errorCount := 0
	ipSet := make(map[string]struct{})
	var mu sync.Mutex
	var wg sync.WaitGroup

	jobs := make(chan string, len(lines))

	// Producer: kirim semua baris log ke channel
	go func() {
		for _, line := range lines {
			jobs <- line
		}
		close(jobs)
	}()

	// Worker pool
	workerCount := 5
	progressChan := make(chan int, workerCount)

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			processed := 0
			for line := range jobs {
				log := parseLineToLogAnalysis(line)
				mu.Lock()
				totalRequests += log.TotalRequests
				errorCount += log.ErrorCount
				ipSet[log.Filename] = struct{}{}
				mu.Unlock()
				processed++
				if processed%10 == 0 {
					progressChan <- processed
				}
			}
			fmt.Printf("[Worker-%d] Finished processing %d lines\n", workerID, processed)
		}(i + 1)
	}

	// Progress monitor goroutine
	go func() {
		for p := range progressChan {
			fmt.Printf("[Progress] Processed %d lines so far...\n", p)
		}
	}()

	wg.Wait()
	close(progressChan)

	analysis := &domain.LogAnalysis{
		Filename:        "uploaded_file.log",
		TotalRequests:   totalRequests,
		ErrorCount:      errorCount,
		UniqueIPs:       len(ipSet),
		AverageResponse: 0,
	}

	fmt.Println("[Analysis] ✅ Completed log analysis successfully")
	fmt.Printf("[Analysis] Total Requests: %d | Errors: %d | Unique IPs: %d\n",
		totalRequests, errorCount, len(ipSet))

	return analysis, nil
}

// 🧩 ParseAndSaveLog — baca file log & panggil ProcessLogs
func (u *LogAnalysisUsecase) ParseAndSaveLog(path string, userID uint) error {
	lines, err := readFileLines(path)
	if err != nil {
		return err
	}

	fmt.Printf("[Log Parser] Starting log processing for %d lines...\n", len(lines))

	// jalankan concurrent log analysis
	analysis, err := u.ProcessLogs(lines)
	if err != nil {
		return err
	}

	analysis.UserID = userID

	// simpan ke database
	err = u.repo.Create(analysis)
	if err != nil {
		return err
	}

	fmt.Println("[Database] Log analysis result saved successfully!")
	return nil
}
