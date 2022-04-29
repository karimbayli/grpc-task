package v1

import (
	"context"
	"errors"
	v1 "github.com/karimbayli/mgid-task/pkg/api/v1"
	"github.com/karimbayli/mgid-task/pkg/models"
	"github.com/karimbayli/mgid-task/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sort"
	"time"
)

var ErrEmptyEmployeeList = errors.New("empty employee list")

type service struct {
	// a database dependency imitation
	m map[int64]models.Employee
}

func (s *service) GetImitatedTimeout(ctx context.Context, request *v1.GetImitatedTimeoutRequest) (*v1.Empty, error) {
	for i := 0; i < 19; i++ {
		if ctx.Err() == context.Canceled {
			return nil, status.Errorf(codes.Canceled, "Imitate Timeout Service - canceled")
		}

		time.Sleep(1 * time.Second)
	}
	return &v1.Empty{}, nil
}

func (s *service) GetSortedEmployees(ctx context.Context, request *v1.GetSortedEmployeesRequest) (*v1.GetSortedEmployeesResponse, error) {
	allEmployees := s.GetAllEmployees()
	if len(allEmployees) == 0 {
		return nil, ErrEmptyEmployeeList
	}
	switch request.GetField() {
	case "name":
		sort.Sort(utils.ByName(allEmployees))
	case "yearOfBirth":
		sort.Sort(utils.ByAge(allEmployees))
	case "salary":
		sort.Sort(utils.BySalary(allEmployees))
	}

	var sortedList []*v1.Employee
	for _, v := range allEmployees {
		var employee = &v1.Employee{
			Name:   v.Name,
			Salary: float32(v.Salary),
			Yob:    int32(v.YearOfBirth),
		}
		sortedList = append(sortedList, employee)
	}

	return &v1.GetSortedEmployeesResponse{Employee: sortedList}, nil
}

func (s *service) GetEldestEmployee(ctx context.Context, request *v1.GetEldestEmployeeRequest) (*v1.GetEldestEmployeeResponse, error) {
	req := &v1.GetSortedEmployeesRequest{Field: "yearOfBirth"}
	eldestEmployee, err := s.GetSortedEmployees(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.GetEldestEmployeeResponse{Employee: &v1.Employee{
		Name:   eldestEmployee.Employee[0].Name,
		Salary: eldestEmployee.Employee[0].Salary,
		Yob:    eldestEmployee.Employee[0].Yob,
	}}, nil
}

func (s *service) GetMedianOfSalaries(ctx context.Context, request *v1.GetMedianOfSalariesRequest) (*v1.GetMedianOfSalariesResponse, error) {
	var employeesSalarySlice = s.GetAllEmployees()
	var salaries []float64
	if len(employeesSalarySlice) < 1 {
		return nil, ErrEmptyEmployeeList
	}
	for _, id := range employeesSalarySlice {
		salaries = append(salaries, id.Salary)
		/*			if u, ok := s.m[id]; ok {
					result[id] = u
				}*/
	}
	median, err := utils.CalcMedian(salaries)
	if err != nil {
		return nil, ErrEmptyEmployeeList
	}
	return &v1.GetMedianOfSalariesResponse{MedianOfSalary: float32(median)}, nil
}

func (s *service) GetAverageOfSalaries(ctx context.Context, request *v1.GetAverageOfSalariesRequest) (*v1.GetAverageOfSalariesResponse, error) {
	//TODO implement me
	var employeesSalarySlice = s.GetAllEmployees()
	var salaries float64
	//employees :=
	for _, em := range employeesSalarySlice {
		salaries += em.Salary
	}

	return &v1.GetAverageOfSalariesResponse{
		AverageOfSalary: float32(salaries / float64(len(employeesSalarySlice))),
	}, nil
}

func (s *service) GetAllEmployees() []models.Employee {
	var employees []models.Employee

	for _, v := range s.m {
		employees = append(employees, v)
	}
	return employees
}
func (s *service) GetHighestPaid(ctx context.Context, request *v1.GetHighestPaidRequest) (*v1.GetHighestPaidResponse, error) {
	req := &v1.GetSortedEmployeesRequest{Field: "salary"}
	highestPaidEmployee, err := s.GetSortedEmployees(ctx, req)
	if err != nil {
		return nil, err
	}
	//return highestPaidEmployee[0], nil

	return &v1.GetHighestPaidResponse{Employee: &v1.Employee{
		Name:   highestPaidEmployee.Employee[0].Name,
		Salary: highestPaidEmployee.Employee[0].Salary,
		Yob:    highestPaidEmployee.Employee[0].Yob,
	}}, nil
}

// NewService instantiates a new Service.
func NewService( /* a database connection would be injected here */ ) v1.EmployeeServiceServer {
	return &service{
		m: map[int64]models.Employee{
			1: {ID: 1, Name: "Volodymyr", YearOfBirth: 1998, Salary: 2678},
			2: {ID: 2, Name: "Mykhailo", YearOfBirth: 1995, Salary: 3001},
			3: {ID: 3, Name: "Oleksiy", YearOfBirth: 2002, Salary: 3987},
			4: {ID: 4, Name: "Test", YearOfBirth: 1983, Salary: 2000},
			5: {ID: 5, Name: "Adam", YearOfBirth: 1979, Salary: 2030},
		},
	}
}
