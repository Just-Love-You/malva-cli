# Malva CLI

Malva (`mlv`) is a command-line tool for cutting MP4 videos, removing audio tracks, overlaying PNG watermarks, and replacing audio in parallel

## Installation

later...

## Usage

### Cut a video segment

Cut from `00:00:10` to `00:00:20` and save as `cut_input.mp4`:

```bash
mlv cut input.mp4 --start 00:00:10 --finish 00:00:20
```

Use short flags:

```bash
mlv cut input.mp4 -s 00:00:05 -f 00:00:15
```

Remove the audio track:

```bash
mlv cut input.mp4 --start 00:01:00 --finish 00:02:30 --off-audio
```

### Supported time formats

You can specify time positions using any of these formats:

- `HH:MM:SS` (e.g., `00:01:05`)
- `HH:MM:SS.ms` (e.g., `00:01:05.500`)
- Shorthand durations as accepted by Go’s `time.ParseDuration`:
  - `14s`
  - `1m9s`
  - `1h0m34s430ms`
  - `500ms`

## Available Commands

- `mlv cut`  
  Cut a segment from a video file.
  **Flags:**
  - `--start`, `-s` <HH:MM:SS[.ms]> — start time  
  - `--finish`, `-f` <HH:MM:SS[.ms]> — finish time  
  - `--off-audio` — remove audio track