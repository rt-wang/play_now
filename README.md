# Now Playing Feed

This project aims to provide a real-time feed of what your friends are currently listening to on Spotify or Apple Music.

## Overview

A Go-based REST and WebSocket API delivers live updates using goroutines for concurrent stream handling. Redis provides Pub/Sub channels for instant event delivery and stores recent listening history in Redis Streams. A lightweight message broker such as NATS JetStream buffers events when users are offline and replays them when they reconnect.

