#!/bin/bash

# This script dynamically resizes a video and adjusts its bitrate, CRF, and audio bitrate based on the selected resolution.

input_file=$1       # First argument: input video file
resolution=$2       # Second argument: the target resolution (e.g., 1080p, 720p, etc.)
output_file=$3      # Third argument: output video file 

case $resolution in
    # 1080p resolution
    "1080p")
        width=1920                # Width of 1080p resolution
        height=1080               # Height of 1080p resolution
        video_bitrate=5000k       # Set the video bitrate to 5000 kbps (5 Mbps) for high quality
        crf=20                    # Set CRF (quality factor) to 20 for high quality encoding
        audio_bitrate=192k        # Set the audio bitrate to 192 kbps for higher audio quality
        ;;

    # 720p resolution
    "720p")
        width=1280                # Width of 720p resolution
        height=720                # Height of 720p resolution
        video_bitrate=2500k       # Set the video bitrate to 2500 kbps (2.5 Mbps) for good quality at medium resolution
        crf=22                    # Set CRF to 22 for balanced quality and compression
        audio_bitrate=128k        # Set the audio bitrate to 128 kbps for standard audio quality
        ;;

    # 480p resolution
    "480p")
        width=854                 # Width of 480p resolution
        height=480                # Height of 480p resolution
        video_bitrate=1000k       # Set the video bitrate to 1000 kbps (1 Mbps) for low-to-medium quality
        crf=24                    # Set CRF to 24 for more compression with acceptable quality
        audio_bitrate=96k         # Set the audio bitrate to 96 kbps for lower audio quality
        ;;

    # 360p resolution
    "360p")
        width=640                 # Width of 360p resolution
        height=360                # Height of 360p resolution
        video_bitrate=800k        # Set the video bitrate to 800 kbps for lower quality (optimized for small file size)
        crf=26                    # Set CRF to 26 for even more compression
        audio_bitrate=64k         # Set the audio bitrate to 64 kbps for basic audio quality
        ;;

    # 240p resolution
    "240p")
        width=426                 # Width of 240p resolution
        height=240                # Height of 240p resolution
        video_bitrate=500k        # Set the video bitrate to 500 kbps for small file sizes and very low quality
        crf=28                    # Set CRF to 28 for high compression and smaller file size
        audio_bitrate=48k         # Set the audio bitrate to 48 kbps for minimal audio quality
        ;;

    # 144p resolution
    "144p")
        width=256                 # Width of 144p resolution
        height=144                # Height of 144p resolution
        video_bitrate=300k        # Set the video bitrate to 300 kbps for minimal quality (suitable for very small screens)
        crf=30                    # Set CRF to 30 for maximum compression and smallest file size
        audio_bitrate=48k         # Keep the audio bitrate low at 48 kbps, as video is very low quality
        ;;

    # If the resolution is invalid or not specified  Exit with an error
    *)
        echo "Invalid resolution. Please choose from 1080p, 720p, 480p, 360p, 240p, or 144p."
        exit 1          
        ;;
esac

# Run ffmpeg with dynamic values for video and audio encoding
# -i "$input_file": Input file
# -vf "scale=$width:$height": Video filter to scale the video to the specified width and height
# -c:v libx264: Encode the video using the H.264 codec (libx264)
# -preset slow: Use the "slow" preset for better compression and quality, at the cost of processing speed
# -crf $crf: Set the CRF (constant rate factor) dynamically based on the resolution (controls video quality and file size)
# -b:v $video_bitrate: Set the video bitrate dynamically based on the resolution
# -c:a aac: Encode the audio using the AAC codec (standard for most devices)
# -b:a $audio_bitrate: Set the audio bitrate dynamically based on the resolution
ffmpeg -y  -i "$input_file" -vf "scale=$width:$height" -c:v libx264 -preset slow -crf $crf -b:v $video_bitrate -c:a aac -b:a $audio_bitrate "$output_file"