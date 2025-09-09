package models

import (
	"database/sql"
)

type Product struct {
	ID        int          `json:"id" db:"id"`
	Name      string       `json:"name" db:"name"`
	Brand     Brand        `json:"brand" db:"brand"`
	Category  Category     `json:"category_id" db:"categories"`
	Battery   Battery      `json:"battery" db:"battery_specs"`
	Platform  Platform     `json:"platform" db:"platform_specs"`
	Network   Network      `json:"network" db:"network_specs"`
	Display   Display      `json:"display" db:"display_specs"`
	Launch    Launch       `json:"launch" db:"launch_specs"`
	Body      Body         `json:"body" db:"body_specs"`
	Memory    Memory       `json:"memory" db:"memory_specs"`
	Sound     Sound        `json:"sound" db:"sound_specs"`
	Comms     Comms        `json:"comms" db:"comms_specs"`
	Features  Features     `json:"features" db:"features_specs"`
	Colors    []string     `json:"colors" db:"product_colors"`
	Models    []string     `json:"models" db:"product_models"`
	Cameras   Cameras      `json:"cameras" db:"cameras"`
	CreatedAt sql.NullTime `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at,omitempty" db:"deleted_at"`
}

type Battery struct {
	Type     string   `json:"type" db:"type"`
	Charging []string `json:"charging" db:"charging"`
}

type Platform struct {
	OS      string `json:"os" db:"os"`
	Chipset string `json:"chipset" db:"chipset"`
	CPU     string `json:"cpu" db:"cpu"`
	GPU     string `json:"gpu" db:"gpu"`
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
	Type       string `json:"type" db:"type"`
	Size       string `json:"size" db:"size"`
	Resolution string `json:"resolution" db:"resolution"`
	Protection string `json:"protection" db:"protection"`
}

type Launch struct {
	Announced sql.NullTime `json:"announced" db:"announced"`
	Released  sql.NullTime `json:"released" db:"released"`
	Status    string       `json:"status" db:"status"`
}

type Body struct {
	Dimensions string `json:"dimensions" db:"dimensions"`
	Weight     string `json:"weight" db:"weight"`
	Build      string `json:"build" db:"build"`
	SIM        string `json:"sim" db:"sim"`
}

type Memory struct {
	CardSlot string `json:"cardSlot" db:"card_slot"`
	Internal string `json:"internal" db:"internal"`
}

type Sound struct {
	Loudspeaker string `json:"loudspeaker" db:"loudspeaker"`
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
	Sensors string `json:"sensors" db:"sensors"`
}

type Cameras struct {
	MainCamera   Camera `json:"mainCamera" db:"main_camera"`
	SelfieCamera Camera `json:"selfieCamera" db:"selfie_camera"`
}

type Camera struct {
	Type        string   `json:"type" db:"type"`
	CameraSpecs []string `json:"cameraSpecs" db:"camera_specs"`
	Features    []string `json:"features" db:"features"`
	Video       []string `json:"video" db:"video"`
}
