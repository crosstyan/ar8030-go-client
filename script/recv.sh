#!/usr/bin/env bash

# Run the GStreamer pipeline with the dynamic filename
gst-launch-1.0 -e udpsrc port=9000 \
    ! 'application/x-rtp, encoding-name=H264, payload=96' \
    ! rtph264depay \
    ! avdec_h264 ! autovideosink
