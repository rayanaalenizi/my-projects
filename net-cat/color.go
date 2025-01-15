package main

var colorCodes = []string{
"\033[95m", // Light Magenta (Pink-like)
"\033[96m", // Light Cyan (Aqua-like)
"\033[91m", // Light Red (Bright Pinkish Red)
"\033[93m", // Light Yellow (Soft Yellow)
"\033[94m", // Light Blue (Sky Blue)
"\033[97m", // White (Bright and Neutral)
}

func assignColor(name string) string {
	index := len(clientColors) % len(colorCodes)
	return colorCodes[index]
}
