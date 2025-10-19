package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/ifs21014-itdel/log-analyzer/internal/domain"
)

type LogAnalysisRepository interface {
	Create(a *domain.LogAnalysis) error
	GetAll() ([]domain.LogAnalysis, error)
	GetByID(id uint) (*domain.LogAnalysis, error)
	Update(a *domain.LogAnalysis) error
	Delete(id uint) error
}

type logAnalysisRepo struct {
	db *sql.DB
}

func NewLogAnalysisRepo(db *sql.DB) LogAnalysisRepository {
	return &logAnalysisRepo{db: db}
}

func (r *logAnalysisRepo) Create(a *domain.LogAnalysis) error {
	query := `
		INSERT INTO log_analysis 
			(filename, total_requests, unique_ips, error_count, average_response, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`
	return r.db.QueryRow(query,
		a.Filename,
		a.TotalRequests,
		a.UniqueIPs,
		a.ErrorCount,
		a.AverageResponse,
		time.Now(),
		time.Now(),
	).Scan(&a.ID)
}

func (r *logAnalysisRepo) GetAll() ([]domain.LogAnalysis, error) {
	rows, err := r.db.Query(`SELECT id, filename, total_requests, unique_ips, error_count, average_response, created_at, updated_at FROM log_analysis`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []domain.LogAnalysis
	for rows.Next() {
		var a domain.LogAnalysis
		if err := rows.Scan(&a.ID, &a.Filename, &a.TotalRequests, &a.UniqueIPs, &a.ErrorCount, &a.AverageResponse, &a.CreatedAt, &a.UpdatedAt); err != nil {
			return nil, err
		}
		list = append(list, a)
	}
	return list, nil
}

func (r *logAnalysisRepo) GetByID(id uint) (*domain.LogAnalysis, error) {
	var a domain.LogAnalysis
	query := `SELECT id, filename, total_requests, unique_ips, error_count, average_response, created_at, updated_at FROM log_analysis WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(
		&a.ID,
		&a.Filename,
		&a.TotalRequests,
		&a.UniqueIPs,
		&a.ErrorCount,
		&a.AverageResponse,
		&a.CreatedAt,
		&a.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *logAnalysisRepo) Update(a *domain.LogAnalysis) error {
	query := `
		UPDATE log_analysis 
		SET filename=$1, total_requests=$2, unique_ips=$3, error_count=$4, average_response=$5, updated_at=$6
		WHERE id=$7`
	_, err := r.db.Exec(query,
		a.Filename,
		a.TotalRequests,
		a.UniqueIPs,
		a.ErrorCount,
		a.AverageResponse,
		time.Now(),
		a.ID,
	)
	return err
}

func (r *logAnalysisRepo) Delete(id uint) error {
	_, err := r.db.Exec(`DELETE FROM log_analysis WHERE id=$1`, id)
	return err
}
