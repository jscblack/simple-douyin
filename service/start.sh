#!/bin/bash

# Start the user service
go run ./user/. &

# Start the feed service
go run ./feed/. &

# Start the publish service
go run ./publish/. &

# Start the favorite service
go run ./favorite/. &

# Start the relation service
go run ./relation/. &

# Start the comment service
go run ./comment/. &

# Start the message service
go run ./message/. &

# Wait for all background processes to finish
wait
