package main

import (
	"errors"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Status int

const (
	Pause Status = iota
	Playing
	Stopped
)

func (s Status) String() string {
	statuses := [...]string{
		"Pause",
		"Playing",
		"Stopped",
	}

	if s < 0 || int(s) >= len(statuses) {
		return fmt.Sprintf("Status(%d)", s)
	}

	return statuses[s]
}

func main() {

	var age int
	fmt.Printf("Valeur: %-12v | Type: %T\n", age, age)

	var b bool
	fmt.Printf("Valeur: %-12v | Type: %T\n", b, b)

	var s string
	fmt.Printf("Valeur: %-12v | Type: %T\n", s, s)

	var uptr uintptr
	fmt.Printf("Valeur: %-12v | Type: %T\n", uptr, uptr)

	var r rune
	fmt.Printf("Valeur: %-12v | Type: %T\n", r, r)

	var fairWeather Celsius = 25.0
	var convertedFahrenheit Fahrenheit = convCelsiusToFahrenheit(fairWeather)

	fmt.Printf("%.1f°C is equal to %.1f°F\n", fairWeather, convertedFahrenheit)

	mp3PlayerStatuses := []Status{
		Playing,
		Pause,
		Stopped,
	}

	fmt.Println("--- Affichage direct de la liste des status de MP3 Player ---")
	fmt.Println(mp3PlayerStatuses)

	fmt.Println("\n--- Itération et affichage détaillé de la liste des status de MP3 Player ---")
	for i, status := range mp3PlayerStatuses {
		fmt.Printf("Element %d -> Valeur brute: %d | Valeur texte: %v\n", i, status, status)
	}

	valuesToConvertToUint8 := []int{0, 255, 256}
	fmt.Println("--- Test de conversion sécurisée ---")

	for _, v := range valuesToConvertToUint8 {
		result, err := SafeIntToUint8(v)

		if err != nil {
			fmt.Printf("Entrée: %-3d | ❌ Erreur : %v\n", v, err)
		} else {
			fmt.Printf("Entrée: %-3d | ✅ Succès : %v (Type: %T)\n", v, result, result)
		}
	}

}

func convCelsiusToFahrenheit(v Celsius) Fahrenheit {
	return Fahrenheit((v * 9 / 5) + 32)
}

func convFahrenheitToCelsius(v Fahrenheit) Celsius {
	return Celsius((v - 32) * 5 / 9)
}

// ErrOutOfRange is thrown when number exceeds the uint8 limits
var ErrOutOfRange = errors.New("La valeur est en dehors des limites d'un uint8 (0-255)")

func SafeIntToUint8(val int) (uint8, error) {
	if val < 0 || val > 255 {
		return 0, ErrOutOfRange
	}

	return uint8(val), nil
}
