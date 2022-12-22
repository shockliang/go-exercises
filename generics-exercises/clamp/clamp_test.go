package clamp

import "testing"

type Distance int32
type Velocity float64

func TestClampInt32(t *testing.T) {
	var (
		min int8 = -10
		max int8 = 10
	)
	if clamp(-30, min, max) != min {
		t.Errorf("clamp failed for int8")
	}
}

func TestClampUint32(t *testing.T) {
	var (
		min uint32 = 0
		max uint32 = 10
	)
	if clamp(20, min, max) != max {
		t.Errorf("clamp failed for uint32")
	}
}

func TestClampFloat32(t *testing.T) {
	var (
		min float32 = 5.5
		max float32 = 9.9
	)
	if clamp(0, min, max) != min {
		t.Errorf("clamp failed for float32")
	}
}

func TestClampDistance(t *testing.T) {
	var (
		min Distance = 0
		max Distance = 100
	)
	if clamp(Distance(99), min, max) != 99 {
		t.Errorf("clamp failed for Distance")
	}
}

func TestClampVelocity(t *testing.T) {
	var (
		min Velocity = 0.0
		max Velocity = 99.9
	)
	if clamp(Velocity(100), min, max) != max {
		t.Errorf("clamp failed for Velocity")
	}
}
