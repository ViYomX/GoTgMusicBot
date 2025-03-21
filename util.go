package main

import (
    "os/exec"
    "strconv"
)

func convertToSle3(file string) (string, error) {
	// ffmpeg -i song.ogg -f s16le -ac 2 -ar 96000 -v quiet sintel_audio.s16le
	cmd := exec.Command("ffmpeg", "-i", file, "-f", "s16le", "-ac", "2", "-ar", "128000", "-v", "quiet", file+".s16le", "-y")
	err := cmd.Run()

	return file + ".s16le", err
}

func Atoi(s string) int {
    i, err := strconv.Atoi(s)
    if err != nil {
        panic("Invalid Integar: " + s)
    }
    return i
}