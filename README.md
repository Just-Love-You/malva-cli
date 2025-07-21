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

## Available Commands

- `mlv cut`  
  Cut a segment from a video file  
  Flags:
    - `--start`, `-s` <HH:MM:SS[.ms]>
    - `--finish`, `-f` <HH:MM:SS[.ms]>
    - `--off-audio` (remove audio track)
