package main

import (
	"os/exec"
)

func convertToSle3(file string) (string, error) {
	// ffmpeg -i song.ogg -f s16le -ac 2 -ar 96000 -v quiet sintel_audio.s16le
	cmd := exec.Command("ffmpeg", "-i", file, "-f", "s16le", "-ac", "2", "-ar", "128000", "-v", "quiet", file+".s16le", "-y")
	err := cmd.Run()

	return file + ".s16le", err
}
