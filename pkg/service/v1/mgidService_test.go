package v1

import (
	"context"
	v1 "mgidAssignment/pkg/api/v1"
	"reflect"
	"testing"
)

func TestService_GetEldestEmployee(t *testing.T) {
	ctx := context.Background()
	s := NewService()
	type args struct {
		ctx context.Context
		req *v1.GetEldestEmployeeRequest
	}
	tests := []struct {
		name    string
		s       v1.EmployeeServiceServer
		args    args
		want    *v1.GetEldestEmployeeResponse
		wantErr bool
	}{
		{
			name: "OK",
			s:    s,
			args: args{
				ctx: ctx,
				req: &v1.GetEldestEmployeeRequest{},
			},
			want: &v1.GetEldestEmployeeResponse{Employee: &v1.Employee{
				Name:   "Adam",
				Salary: 2030,
				Yob:    1979,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetEldestEmployee(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("EmployeeServer.GetEldestEmployee() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EmployeeServer.GetEldestEmployee() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetHighestPaid(t *testing.T) {
	ctx := context.Background()
	s := NewService()
	type args struct {
		ctx context.Context
		req *v1.GetHighestPaidRequest
	}
	tests := []struct {
		name    string
		s       v1.EmployeeServiceServer
		args    args
		want    *v1.GetHighestPaidResponse
		wantErr bool
	}{
		{
			name: "OK",
			s:    s,
			args: args{
				ctx: ctx,
				req: &v1.GetHighestPaidRequest{},
			},
			want: &v1.GetHighestPaidResponse{Employee: &v1.Employee{
				Name:   "Oleksiy",
				Salary: 3987,
				Yob:    2002},
			}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetHighestPaid(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("EmployeeServer.GetHighestPaid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EmployeeServer.GetHighestPaid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetAverageOfSalaries(t *testing.T) {
	ctx := context.Background()
	s := NewService()
	type args struct {
		ctx context.Context
		req *v1.GetAverageOfSalariesRequest
	}
	tests := []struct {
		name    string
		s       v1.EmployeeServiceServer
		args    args
		want    *v1.GetAverageOfSalariesResponse
		wantErr bool
	}{
		{
			name: "OK",
			s:    s,
			args: args{
				ctx: ctx,
				req: &v1.GetAverageOfSalariesRequest{},
			},
			want: &v1.GetAverageOfSalariesResponse{AverageOfSalary: 2739.2},
		}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetAverageOfSalaries(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("EmployeeServer.GetAverageOfSalaries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EmployeeServer.GetAverageOfSalaries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetMedianOfSalaries(t *testing.T) {
	ctx := context.Background()
	s := NewService()
	type args struct {
		ctx context.Context
		req *v1.GetMedianOfSalariesRequest
	}
	tests := []struct {
		name    string
		s       v1.EmployeeServiceServer
		args    args
		want    *v1.GetMedianOfSalariesResponse
		wantErr bool
	}{
		{
			name: "OK",
			s:    s,
			args: args{
				ctx: ctx,
				req: &v1.GetMedianOfSalariesRequest{},
			},
			want: &v1.GetMedianOfSalariesResponse{MedianOfSalary: 2678},
		}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetMedianOfSalaries(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("EmployeeServer.GetMedianOfSalaries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EmployeeServer.GetMedianOfSalaries() = %v, want %v", got, tt.want)
			}
		})
	}
}
