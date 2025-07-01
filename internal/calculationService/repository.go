package calculationService

import "gorm.io/gorm"

// CRUD
// file for work with database only!

type CalculationRepository interface { // public (с большой буквы)
	CreateCalculation(calculation Calculation) error
	GetAllCalculations() ([]Calculation, error)
	GetCalculationByID(id string) (Calculation, error)
	UpdateCalculation(calculation Calculation) error
	DeleteCalculation(id string) error
}

type calculationRepository struct {
	db *gorm.DB
}

func NewCalculationRepository(db *gorm.DB) CalculationRepository {
	return &calculationRepository{db: db}
}

func (r *calculationRepository) CreateCalculation(calculation Calculation) error {
	return r.db.Create(&calculation).Error
}

func (r *calculationRepository) GetAllCalculations() ([]Calculation, error) {
	var calculations []Calculation

	return calculations, r.db.Find(&calculations).Error
}

func (r *calculationRepository) GetCalculationByID(id string) (Calculation, error) {
	var calculation Calculation

	return calculation, r.db.First(&calculation, "id = ?", id).Error
}

func (r *calculationRepository) UpdateCalculation(calculation Calculation) error {
	return r.db.Save(&calculation).Error
}

func (r *calculationRepository) DeleteCalculation(id string) error {
	return r.db.Delete(&Calculation{}, "id = ?", id).Error
}
