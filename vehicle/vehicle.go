package vehicle

import (
	"fmt"

	context "golang.org/x/net/context"
)

type VehicleServer struct {
	UnimplementedVehicleServiceServer
}

func (v *VehicleServer) StoreVehicle(ctx context.Context, vehicleRegistration *VehicleRegistration) (*Vehicle, error) {
	a := vehicleRegistration.String()
	fmt.Print(a)
	return &Vehicle{}, nil
}

func (v *VehicleServer) UpdateVehicle(ctx context.Context, vehicle *Vehicle) (*Vehicle, error) {
	return &Vehicle{}, nil
}

func (v *VehicleServer) GetVehicle(ctx context.Context, vehicleId *VehicleId) (*Vehicle, error) {
	return &Vehicle{}, nil
}

func (v *VehicleServer) ListVehicles(ctx context.Context, void *Void) (*Vehicle, error) {
	return &Vehicle{}, nil
}

func (v *VehicleServer) DeleteVehicle(ctx context.Context, vehicle *VehicleId) (*Void, error) {
	return &Void{}, nil
}
