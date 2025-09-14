package models

import (
	"database/sql"
)

type Product struct {
	ID         int          `json:"id" db:"id"`
	Name       string       `json:"name" db:"name"`
	Brand      string       `json:"brand_name" db:"brand_name"`
	BrandID    int          `json:"brand_id" db:"brand_id"`
	ImageUrl   *string      `json:"image_url" db:"image_url"`
	CategoryId int          `json:"category_id" db:"category_id"`
	Category   string       `json:"category_name" db:"category_name"`
	Released   string       `json:"released,omitempty" db:"released"`
	Announced  string       `json:"announced,omitempty" db:"announced"`
	Status     string       `json:"status,omitempty" db:"status"`
	DeletedAt  sql.NullTime `json:"-" db:"deleted_at"`
}

type ProductDetail struct {
	Product     Product     `json:"product"`
	PhoneDetail PhoneDetail `json:"phone_detail"`
	Battery     Battery     `json:"battery"`
	Display     Display     `json:"display"`
	Memory      []Memory    `json:"memory"`
	Sound       Sound       `json:"sound" `
	Sensors     Sensors     `json:"features"`
	Colors      []string    `json:"colors_tr"`
	Models      []string    `json:"models"`
	Cameras     Cameras     `json:"cameras"`
}
type PhoneDetail struct {
	CurrentOS     string `json:"current_os" db:"current_os"`
	UpgradableOS  string `json:"upgradable_to" db:"upgradable_to"`
	Chipset       string `json:"chipset" db:"chipset"`
	CPU           string `json:"cpu" db:"cpu"`
	GPU           string `json:"gpu" db:"gpu"`
	Dimensions    string `json:"dimensions" db:"dimensions"`
	Weight        string `json:"weight" db:"weight"`
	Build         string `json:"build" db:"build"`
	SimInfo       string `json:"sim_info" db:"sim_info"`
	NetTechnology string `json:"network_technology" db:"network_technology"`
	NetSpeed      string `json:"network_speed" db:"network_speed"`
	G5            string `json:"5g" db:"g5"`
	G4            string `json:"4g" db:"g4"`
	G3            string `json:"3g" db:"g3"`
	G2            string `json:"2g" db:"g2"`
	GPS           string `json:"gps" db:"gps"`
	NFC           string `json:"nfc" db:"nfc"`
	Radio         string `json:"radio" db:"radio"`
	Wlan          string `json:"wlan" db:"wlan"`
	Bluetooth     string `json:"bluetooth" db:"bluetooth"`
	USB           string `json:"usb" db:"usb"`
	CardSlot      string `json:"card_slot" db:"card_slot"`
}

type Battery struct {
	Technology      string            `json:"technology" db:"technology"`
	Capacity        string            `json:"capacity" db:"capacity"`
	ChargingDetails []ChargingDetails `json:"charging" db:"charging"`
}
type ChargingDetails struct {
	Type        string `json:"type" db:"type"`
	Description string `json:"description" db:"description"`
	Power       string `json:"power,omitempty" db:"power"`
}
type Display struct {
	Type              string   `json:"type" db:"type"`
	Size              string   `json:"size" db:"size"`
	Resolution        string   `json:"resolution" db:"resolution"`
	AspectRatio       string   `json:"aspect_ratio" db:"aspect_ratio"`
	HDR               []string `json:"hdr" db:"hdr"`
	RefreshRate       string   `json:"refreshRate" db:"refresh_rate"`
	PPI               []string `json:"ppi" db:"ppi"`
	BrightnessTypical string   `json:"brightness_typical" db:"brightness_typical"`
	BrightnessHbm     string   `json:"brightness_hbm" db:"brightness_hbm"`
	Protection        string   `json:"protection" db:"protection"`
}

type Memory struct {
	Storage string `json:"storage" db:"storage"`
	RAM     string `json:"ram" db:"ram"`
}

type Sound struct {
	Loudspeaker string `json:"loudspeaker" db:"loudspeaker"`
	Features    string `json:"features" db:"features"`
}

type Sensors struct {
	Sensors []string `json:"sensors" db:"features"`
}

type Cameras struct {
	MainCamera   Camera `json:"mainCamera"`
	SelfieCamera Camera `json:"selfieCamera"`
}

type Camera struct {
	Lenses   []Lens          `json:"lenses"`
	Features []CameraFeature `json:"features"`
	Video    []CameraVideo   `json:"video"`
}
type CameraVideo struct {
	Spec string `json:"video" db:"video_spec"`
}
type CameraFeature struct {
	Spec string `json:"spec" db:"feature"`
}
type Lens struct {
	Type          string `json:"type,omitempty" db:"type"`                 // Örn: "Derinlik" — bazı lenslerde yok
	Megapixels    string `json:"megapixels,omitempty" db:"megapixels"`     // Örn: "48 MP"
	Aperture      string `json:"aperture,omitempty" db:"aperture"`         // Örn: "f/1.8"
	FocalLength   string `json:"focal_length,omitempty" db:"focal_length"` // Örn: "24mm"
	SensorSize    string `json:"sensor_size,omitempty" db:"sensor_size"`   // Örn: "1/1.28\""
	PixelSize     string `json:"pixel_size,omitempty" db:"pixel_size"`     // Örn: "1.22µm"
	Zoom          string `json:"zoom,omitempty" db:"zoom"`                 // Örn: "5x optical zoom"
	OtherFeatures string `json:"other_features,omitempty" db:"other_features"`
}
