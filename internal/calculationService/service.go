package calculationService

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/google/uuid"
)

// бизнес-логика, файл для подсчётов

type CalculationService interface {
	CreateCalculation(expression string) (Calculation, error)
	GetAllCalculations() ([]Calculation, error)
	GetCalculationByID(id string) (Calculation, error)
	UpdateCalculation(id, expression string) (Calculation, error)
	DeleteCalculation(id string) error
}

type calculationService struct {
	repo CalculationRepository
}

func NewCalculationService(r CalculationRepository) CalculationService {
	return &calculationService{repo: r}
}

func (s *calculationService) calculateExpression(expression string) (string, error) {
	expr, err := govaluate.NewEvaluableExpression(expression) // создание самого выражения
	if err != nil {
		return "", err // передали невалидную операцию
	}
	result, err := expr.Evaluate(nil) // считаем
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", result), err
}

func (s *calculationService) CreateCalculation(expression string) (Calculation, error) {
	result, err := s.calculateExpression(expression)
	if err != nil {
		return Calculation{}, err
	}

	calc := Calculation{
		ID:         uuid.NewString(),
		Expression: expression,
		Result:     result,
	}

	if err := s.repo.CreateCalculation(calc); err != nil {
		return Calculation{}, err
	}

	return calc, nil
}

func (s *calculationService) GetAllCalculations() ([]Calculation, error) {
	return s.repo.GetAllCalculations()
}

func (s *calculationService) GetCalculationByID(id string) (Calculation, error) {
	return s.repo.GetCalculationByID(id)
}

func (s *calculationService) UpdateCalculation(id, expression string) (Calculation, error) {
	calc, err := s.GetCalculationByID(id)
	if err != nil {
		return Calculation{}, err
	}
	result, err := s.calculateExpression(expression)
	if err != nil {
		return Calculation{}, err
	}
	calc.Result = result
	calc.Expression = expression

	if err := s.repo.UpdateCalculation(calc); err != nil {
		return Calculation{}, err
	}

	return calc, nil

}

func (s *calculationService) DeleteCalculation(id string) error {
	return s.repo.DeleteCalculation(id)
}
