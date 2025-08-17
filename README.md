# Malva CLI

Fast Go CLI wrapping ffmpeg for video processing: cut segments, remove audio, overlay PNG watermarks, resize with aspect ratio preservation, and convert to animated GIFs with custom fps/scale.

## Features

- cut video segments (`cut`)
- change video (remove audio, watermark, resize, replace audio) (`change`)
- convert to GIF with adjustable fps and width (`convert`)
- smart duration parsing: `HH:MM:SS`, `HH:MM:SS.ms`, `14s`, `1m9s`, `1h0m34s430ms`, `500ms`
- MP4 faststart support for better preview (quick look)

## Installation

### via Homebrew (recommended)

Tap the formula and install:
```bash
brew tap Just-Love-You/homebrew-malva-cli  
brew install mlv

# or single command call

brew install just-love-you/malva-cli/mlv
```

To upgrade later:
```bash
brew update && brew upgrade mlv
```

To uninstall:
```bash
brew uninstall mlv
```

### from source (manual)

Clone/build yourself:
```bash
git clone https://github.com/Just-Love-You/malva-cli.git  
cd malva-cli  
go build -o mlv  
sudo mv mlv /usr/local/bin/
```

### Verification

Run:
```bash
mlv --help
```

## Usage

### Cut

Cut a segment:
```bash
mlv cut input.mp4 –-start 00:00:10 –-finish 00:00:20 -o new_segment_name.mp4
```

### Change

Remove audio:

```bash
mlv change input.mp4 –-remove-audio -o noaudio.mp4
```

Resize by height / width:
```bash
mlv change input.mp4 -–resize-height 720 -o resized.mp4
# or
mlv change input.mp4 –-resize-width 480 -o resized_w.mp4
```

Add watermark:
```bash
mlv change input.mp4 -–watermark logo.png -o watermarked.mp4
```

Combine resize + watermark + remove audio:
```bash
mlv change input.mp4 -–resize-height 720 –-watermark logo.png –-remove-audio -o combo.mp4
```

### Convert

Convert to GIF with defaults:
```bash
mlv convert input.mp4 -–gif
```

Custom fps and scale:
```bash
mlv convert input.mp4 -–gif -–fps 20 -–scale 320
```

Combined example (cut → change → GIF)
```bash
mlv cut test_video.mp4 -s 5s -f 20s -o short.mp4 
&& mlv change short.mp4 –-remove-audio –-resize-width 480 –-watermark logo.png -o final.mp4 
&& mlv convert final.mp4 –-gif –-fps 10 –-scale 320
```

## Available Commands Summary
- mlv cut: –-start / -s, –-finish / -f, –-off-audio, -–output / -o 
- mlv change: -–remove-audio, -–watermark, -–replace-audio, –-resize-height, –-resize-width, –-output / -o
- mlv convert: -–gif, –-fps, -–scale

### Supported time formats

Any of:
-	HH:MM:SS (e.g., 00:01:05)
-	HH:MM:SS.ms (e.g., 00:01:05.500)
-	Go durations: 14s, 1m9s, 1h0m34s430ms, 500ms

## Contributing

Open issues or submit pull requests. Feedback and improvements welcome
