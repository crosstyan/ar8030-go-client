#!/usr/bin/env bash

gst-launch-1.0 -e videotestsrc pattern=ball ! \
    video/x-raw,width=640,height=480,framerate=30/1 ! \
    x264enc tune=zerolatency bitrate=500 speed-preset=ultrafast ! \
    rtph264pay ! \
    udpsink host=127.0.0.1 port=8001
