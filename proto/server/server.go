package main

import (
	"context"
	"log"
	"net"
	"project-employee/config"
	"project-employee/dao"
	"project-employee/model"
	pb "project-employee/proto/model"
	"project-employee/service"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type server struct {
	pb.UnimplementedEmployeeServiceServer
}

var employeeService service.EmployeeServiceImpl
var companyService service.CompanyServiceImpl
var positionService service.PositionServiceImpl
var unitService service.UnitServiceImpl

func (*server) GetEmployees(c context.Context, empty *pb.Empty) (data *pb.Employees, e error) {
	var employees []*pb.Employee
	result, err := employeeService.GetEmployeesWithToken(empty.Token)

	if err != nil {
		return data, err
	}
	for _, emp := range result {
		employee := &pb.Employee{
			Id:   emp.Id,
			Nik:  emp.Nik,
			Name: emp.Name,
		}
		employees = append(employees, employee)
	}
	list := pb.Employees{Employee: employees}

	return &list, nil
}

func (*server) GetCompanies(c context.Context, empty *pb.Empty) (data *pb.Companies, e error) {
	var companies []*pb.Company
	result, err := companyService.GetCompanies()

	if err != nil {
		return data, err
	}

	for _, com := range result {
		company := &pb.Company{
			Id:               com.Id,
			Code:             com.Code,
			Name:             com.Name,
			Description:      com.Description,
			CompanyTaxNumber: com.CompanyTaxNumber,
		}
		companies = append(companies, company)
	}
	list := pb.Companies{Company: companies}

	return &list, nil
}

func (*server) GetPositions(c context.Context, empty *pb.Empty) (data *pb.Positions, e error) {
	var listPosition []*pb.Position

	result, err := positionService.GetPositions()

	if err != nil {
		return data, err
	}

	for _, pos := range result {
		position := &pb.Position{
			Id:          pos.Id,
			Code:        pos.Code,
			Name:        pos.Name,
			Description: pos.Description,
			Level:       string(pos.Level),
		}
		listPosition = append(listPosition, position)
	}

	positions := pb.Positions{Position: listPosition}

	return &positions, nil
}

func (*server) GetUnits(c context.Context, empty *pb.Empty) (data *pb.Units, e error) {
	var listUnit []*pb.Unit

	result, err := unitService.GetUnits()

	if err != nil {
		return data, err
	}

	for _, un := range result {
		unit := &pb.Unit{
			Id:   un.Id,
			Code: un.Code,
			Name: un.Name,
		}

		listUnit = append(listUnit, unit)
	}

	units := pb.Units{Unit: listUnit}

	return &units, nil
}

func (*server) CreateEmployee(c context.Context, emp *pb.AddEmployee) (resp *pb.ResponseAddEmployee, e error) {

	employee := model.Employee{
		PersonId:   emp.PersonId,
		CompanyId:  emp.CompanyId,
		UnitId:     emp.UnitId,
		PositionId: emp.PositionId,
		Nik:        emp.Nik,
		HiredDate:  emp.HiredDate,
	}

	err := employeeService.AddEmployee(&employee)
	if err == nil {
		msg := pb.ResponseAddEmployee{
			Msg: "Success add employee",
		}
		return &msg, nil
	}
	msg := pb.ResponseAddEmployee{
		Msg: err.Error(),
	}
	return &msg, err
}

func main() {
	g := initDb()
	dao.SetDao(g)
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal("Failed to listen with err =>", err)
	}

	s := grpc.NewServer()
	pb.RegisterEmployeeServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve with err =>", err)
	}
}

func initDb() *gorm.DB {
	g, err := config.Conn()
	if err == nil {
		config.MigrateSchema(g)
		return g
	}
	panic(err)
}
