package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/golang/geo/r3"
)

// id,edsm_id,name,x,y,z,population,is_populated,government_id,government,allegiance_id,allegiance,security_id,security,primary_economy_id,primary_economy,power,power_state,power_state_id,needs_permit,updated_at,minor_factions_updated_at,simbad_ref,controlling_minor_faction_id,controlling_minor_faction,reserve_type_id,reserve_type,ed_system_address
type System struct {
	Name string  `csv:"name"`
	X    float64 `csv:"x"`
	Y    float64 `csv:"y"`
	Z    float64 `csv:"z"`
}

func (s *System) Distance(s2 *System) float64 {
	return r3.Vector{X: s.X, Y: s.Y, Z: s.Z}.
		Distance(r3.Vector{X: s2.X, Y: s2.Y, Z: s2.Z})
}

func Main() error {
	f, err := os.Open("systems.csv")
	if err != nil {
		return fmt.Errorf("open: %w", err)
	}
	defer f.Close()

	log.Print("Decoding systems")

	var systems []*System
	if err := gocsv.Unmarshal(f, &systems); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	log.Print("Decoded systems")

	var anyNa, gabraceni *System
	for _, s := range systems {
		switch s.Name {
		case "Any Na":
			anyNa = s
		case "Gabraceni":
			gabraceni = s
		}
	}

	switch {
	case anyNa == nil:
		return errors.New("any na missing")
	case gabraceni == nil:
		return errors.New("gabraceni missing")
	}

	log.Print(anyNa, gabraceni)

	for _, s := range systems {
		d1 := anyNa.Distance(s)
		if d1 < 274 || d1 > 276 {
			continue
		}
		d2 := gabraceni.Distance(s)
		if d2 < 322 || d2 > 324 {
			continue
		}
		log.Printf("candidate: %s (Any Na: %f ly; Gabraceni: %f ly)\n", s.Name, d1, d2)
	}
	return nil
}

func main() {
	if err := Main(); err != nil {
		log.Fatal(err)
	}
}
