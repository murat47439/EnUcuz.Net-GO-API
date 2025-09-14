package models

import (
	"database/sql"
)

type Product struct {
	ID         int          `json:"id" db:"id"`
	Name       string       `json:"name" db:"name"`
	Brand      int          `json:"brand_id" db:"brand_id"`
	ImageUrl   *string      `json:"image_url" db:"image_url"`
	CategoryId int          `json:"category_id" db:"category_id"`
	CreatedAt  sql.NullTime `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt  sql.NullTime `json:"-" db:"deleted_at"`
}

type ProductDetail struct {
	ID        int          `json:"id" db:"id"`
	Name      string       `json:"name" db:"name"`
	Brand     *Brand       `json:"brand" db:"brand"`
	ImageUrl  *string      `json:"image_url" db:"image_url"`
	Category  *Category    `json:"category_id" db:"categories"`
	Battery   Battery      `json:"battery" db:"battery_specs"`
	Platform  Platform     `json:"platform" db:"platform_specs"`
	Network   Network      `json:"network" db:"network_specs"`
	Display   Display      `json:"display" db:"display_specs"`
	Launch    Launch       `json:"launch" db:"launch_specs"`
	Body      Body         `json:"body" db:"body_specs"`
	Memory    Memory       `json:"memory" db:"memory_specs"`
	Sound     Sound        `json:"sound" db:"sound_specs"`
	Comms     Comms        `json:"comms" db:"comms_specs"`
	Features  Features     `json:"features" db:"feature_specs"`
	Colors    []string     `json:"colors_tr" db:"colors"`
	Models    []string     `json:"models" db:"models"`
	Cameras   Cameras      `json:"cameras" db:"cameras"`
	CreatedAt sql.NullTime `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt sql.NullTime `json:"-" db:"deleted_at"`
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
type Platform struct {
	CurrentOS    string `json:"current_os" db:"current_os"`
	UpgradableOS string `json:"upgradable_to,omitempty" db:"upgradable_to"`
	Chipset      string `json:"chipset" db:"chipset"`
	CPU          string `json:"cpu" db:"cpu"`
	GPU          string `json:"gpu" db:"gpu"`
}

type Network struct {
	Technology string `json:"technology" db:"technology"`
	Speed      string `json:"speed" db:"speed"`
	G2         string `json:"2g" db:"g2"`
	G3         string `json:"3g" db:"g3"`
	G4         string `json:"4g" db:"g4"`
	G5         string `json:"5g" db:"g5"`
}

type Display struct {
	PanelType        string     `json:"panelType" db:"type"`
	SizeInches       string     `json:"size_inches" db:"size"`
	OtherFeatures    []string   `json:"other_features" db:"other_features"`
	HDR              []string   `json:"hdrSupport" db:"hdr"`
	RefreshRate      string     `json:"refreshRate" db:"refresh_rate"`
	Brightness       Brightness `json:"brightness"`
	ResolutionPixels string     `json:"resolution_pixels" db:"resolution"`
	Protection       string     `json:"protection" db:"protection"`
	AspectRatio      string     `json:"aspect_ratio" db:"aspect_ratio"`
}

type Brightness struct {
	Typical *string `json:"typical" db:"brightness_typical"`
	Hbm     *string `json:"hbm" db:"brightness_hbm"`
}

type Launch struct {
	Announced string `json:"announced" db:"announced"`
	Released  string `json:"released" db:"released"`
	Status    string `json:"status" db:"status"`
}

type Body struct {
	Dimensions string `json:"dimensions" db:"dimensions"`
	Weight     string `json:"weight" db:"weight"`
	Build      string `json:"build" db:"build"`
	SIM        string `json:"sim" db:"sim"`
}

type Memory struct {
	CardSlot        string          `json:"cardSlot" db:"card_slot"`
	InternalOptions []MemoryVariant `json:"internal_options"`
}

type MemoryVariant struct {
	Storage string `json:"storage"`
	RAM     string `json:"ram"`
}

type Sound struct {
	Loudspeaker string `json:"has_loudspeaker" db:"loudspeaker"`
	Features    string `json:"loudspeaker_features" db:"features"`
}

type Comms struct {
	WLAN        string `json:"wlan" db:"wlan"`
	Bluetooth   string `json:"bluetooth" db:"bluetooth"`
	Positioning string `json:"positioning" db:"positioning"`
	NFC         string `json:"nfc" db:"nfc"`
	Radio       string `json:"radio" db:"radio"`
	USB         string `json:"usb" db:"usb"`
}

type Features struct {
	Sensors []string `json:"sensors_tr" db:"sensors"`
}

type Cameras struct {
	MainCamera   Camera `json:"mainCamera" db:"main_camera"`
	SelfieCamera Camera `json:"selfieCamera" db:"selfie_camera"`
}

type Camera struct {
	Lenses   []Lens          `json:"lenses"`
	Features []CameraFeature `json:"features"`
	Video    []CameraFeature `json:"video"`
}

type CameraFeature struct {
	ID   int    `json:"id"`
	Spec string `json:"spec"`
}
type Lens struct {
	ID            int      `json:"id" db:"id"`
	Type          string   `json:"type,omitempty" db:"type"`                 // Örn: "Derinlik" — bazı lenslerde yok
	Megapixels    string   `json:"megapixels,omitempty" db:"megapixels"`     // Örn: "48 MP"
	Aperture      string   `json:"aperture,omitempty" db:"aperture"`         // Örn: "f/1.8"
	FocalLength   string   `json:"focal_length,omitempty" db:"focal_length"` // Örn: "24mm"
	SensorSize    string   `json:"sensor_size,omitempty" db:"sensor_size"`   // Örn: "1/1.28\""
	PixelSize     string   `json:"pixel_size,omitempty" db:"pixel_size"`     // Örn: "1.22µm"
	Zoom          string   `json:"zoom,omitempty" db:"zoom"`                 // Örn: "5x optical zoom"
	OtherFeatures []string `json:"other_features,omitempty" db:"other_features"`
}
