package calculationService

import (
	"fmt"

	"github.com/Knetic/govaluate"
	"github.com/google/uuid"
)

type CalculationService interface {
	CreateCalculation(expression string) (Calculation, error)
	GetAllCalculations() ([]Calculation, error)
	GetCalculationById(id string) (Calculation, error)
	UpdateCalculation(id string, expression string) (Calculation, error)
	DeleteCalculationById(id string) error
}

type calcService struct {
	repo CalculationRepository
}

func NewCalculationService(c CalculationRepository) *calcService {
	return &calcService{repo: c}
}

func (s *calcService) calculateExpression(expression string) (string, error) {
	expr, err := govaluate.NewEvaluableExpression(expression) // Создание выражения
	if err != nil {
		return "", err // При подаче невалидного выражения будет ошибка
	}
	result, err := expr.Evaluate(nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", result), err
}

func (s *calcService) CreateCalculation(expression string) (Calculation, error) {
	result, err := s.calculateExpression(expression)
	if err != nil {
		return Calculation{}, err
	}

	calculation := Calculation{
		ID:         uuid.NewString(),
		Expression: expression,
		Result:     result,
	}

	if err := s.repo.CreateCalculation(calculation); err != nil {
		return Calculation{}, err
	}

	return calculation, nil
}

func (s *calcService) GetAllCalculations() ([]Calculation, error) {
	return s.repo.GetAllCalculations()
}

func (s *calcService) GetCalculationById(id string) (Calculation, error) {
	return s.repo.GetCalculationById(id)
}

func (s *calcService) UpdateCalculation(id string, expression string) (Calculation, error) {
	calculation, err := s.repo.GetCalculationById(id)
	if err != nil {
		return Calculation{}, err
	}

	result, err := s.calculateExpression(expression)
	if err != nil {
		return Calculation{}, err
	}

	calculation.Expression = expression
	calculation.Result = result

	if err := s.repo.UpdateCalculation(calculation); err != nil {
		return Calculation{}, err
	}

	return calculation, nil
}

func (s *calcService) DeleteCalculationById(id string) error {
	return s.repo.DeleteCalculation(id)
}
