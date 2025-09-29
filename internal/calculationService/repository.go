package calculationService

import "gorm.io/gorm"

// Основные методы CRUD

type CalculationRepository interface {
	CreateCalculation(calc Calculation) error
	GetAllCalculations() ([]Calculation, error)
	UpdateCalculation(calc Calculation) error
	DeleteCalculation(id string) error
	GetCalculationById(id string) (Calculation, error)
}

type calcRepository struct {
	db *gorm.DB
}

func NewCalculationRepository(db *gorm.DB) CalculationRepository {
	return &calcRepository{db: db}
}

func (c *calcRepository) CreateCalculation(calc Calculation) error {
	return c.db.Create(&calc).Error
}

func (c *calcRepository) GetAllCalculations() ([]Calculation, error) {
	var calculations []Calculation

	if err := c.db.Find(&calculations).Error; err != nil {
		return nil, err
	}
	return calculations, nil
}

func (c *calcRepository) GetCalculationById(id string) (Calculation, error) {
	var calculation Calculation

	if err := c.db.First(&calculation, "id = ?", id).Error; err != nil {
		return calculation, err
	}
	return calculation, nil
}

func (c *calcRepository) UpdateCalculation(calc Calculation) error {
	return c.db.Save(&calc).Error
}

func (c *calcRepository) DeleteCalculation(id string) error {
	return c.db.Delete(&Calculation{}, "id = ?", id).Error
}
