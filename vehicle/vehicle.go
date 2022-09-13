package vehicle

import (
	"fmt"

	context "golang.org/x/net/context"
)

var listOfCars []Vehicle

type VehicleRegistration struct {
	Brand string `json:"brand,omitempty"`
	Model string `json:"model,omitempty"`
	Year  int32  `json:"year,omitempty"`
}

type Vehicle struct {
	Id    int32  `json:"id,omitempty"`
	Brand string `json:"brand,omitempty"`
	Model string `json:"model,omitempty"`
	Year  int32  `json:"year,omitempty"`
}

type VehicleId struct {
	VehicleId int32 `json:"vehicleId,omitempty"`
}

type VehicleServer struct {
	UnimplementedVehicleServiceServer
}

func (v *VehicleServer) StoreVehicle(ctx context.Context, vehicleRegistration *VehicleRegistration) (*Vehicle, error) {
	vehicle := Vehicle{
		Id:    int32(len(listOfCars)),
		Brand: vehicleRegistration.GetBrand(),
		Model: vehicleRegistration.GetModel(),
		Year:  vehicleRegistration.GetYear(),
	}
	listOfCars = append(listOfCars, vehicle)
	fmt.Printf("new vehicle added: %v", vehicle.String())
	return &vehicle, nil
}

func (v *VehicleServer) UpdateVehicle(ctx context.Context, vehicle *Vehicle) (*Vehicle, error) {
	return &Vehicle{}, nil
}

func (v *VehicleServer) GetVehicle(ctx context.Context, vehicleId *VehicleId) (*Vehicle, error) {
	vehicle := &listOfCars[vehicleId.GetVehicleId()]
	return vehicle, nil
}

// TODO list vehicles
func (v *VehicleServer) ListVehicles(ctx context.Context, void *Void) (*Vehicle, error) {

	for _, c := range listOfCars {
		fmt.Print(c.String())
	}
	return &Vehicle{}, nil
}

// TODO delete vehicles
func (v *VehicleServer) DeleteVehicle(ctx context.Context, vehicleId *VehicleId) (*Void, error) {
	i := vehicleId.GetVehicleId()
	copy(listOfCars[i:], listOfCars[i+1:])
	listOfCars[len(listOfCars)-1] = Vehicle{}
	listOfCars = listOfCars[:len(listOfCars)-1]
	return &Void{}, nil
}
