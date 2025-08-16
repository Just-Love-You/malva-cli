package helpers

import (
	"fmt"
	"strconv"
	"time"
)

func ParseFFmpegTime(hours, minutes, seconds string) float64 {
	hh, _ := strconv.Atoi(hours)
	mm, _ := strconv.Atoi(minutes)
	ss, _ := strconv.ParseFloat(seconds, 64)
	return float64(hh)*3600 + float64(mm)*60 + ss
}

// FormatDurationFFMPEG converts a time.Duration to "HH:MM:SS.ms" format for ffmpeg
func FormatDurationFFMPEG(d time.Duration) string {
	totalMillis := d.Milliseconds()
	h := totalMillis / 3600000
	m := (totalMillis % 3600000) / 60000
	s := (totalMillis % 60000) / 1000
	ms := totalMillis % 1000
	return fmt.Sprintf("%02d:%02d:%02d.%03d", h, m, s, ms)
}
