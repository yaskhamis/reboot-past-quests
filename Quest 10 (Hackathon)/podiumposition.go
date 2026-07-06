package piscine

func PodiumPosition(podium [][]string) [][]string {
	n := len(podium)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if len(podium[j]) > 0 {
				if podium[j][0] == "1st Place" {
					podium[0], podium[j] = podium[j], podium[0]
				} else if podium[j][0] == "2nd Place" {
					podium[1], podium[j] = podium[j], podium[1]
				} else if podium[j][0] == "3rd Place" {
					podium[2], podium[j] = podium[j], podium[2]
				} else if podium[j][0] == "4th Place" {
					podium[3], podium[j] = podium[j], podium[3]
				}
			}
		}
	}
	return podium
}
