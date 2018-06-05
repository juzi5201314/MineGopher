package data

type Rotation struct {
	Yaw, HeadYaw, Pitch float64
}

func (rotation *Rotation) GetYaw() float64 {
	return rotation.Yaw
}

func (rotation *Rotation) GetHeadYaw() float64 {
	return rotation.HeadYaw
}

func (rotation *Rotation) GetPitch() float64 {
	return rotation.Pitch
}
