package main

func Direction(d string) string {
	switch d {
		case "N":
			return "North"
		case "S":
			return "South"
		case "E":
			return "East"
		case "W":
			return "West"
		default:
			return "unknown"
	}
}
