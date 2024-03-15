package main

import (
	"fmt"

	"github.com/skip2/go-qrcode"
)

func main() {
	data := "https://netscapeservice.com" // Data to encode into the QR code

	// Generate QR code with the given data
	err := qrcode.WriteFile(data, qrcode.Medium, 256, "qr.png")
	if err != nil {
		fmt.Println("Error generating QR code:", err)
		return
	}

	fmt.Println("QR code generated successfully!")
}
