# protoc command | generate code to "." (this folder) | generate proto files from this folder (proto/*.proto) * representing all
protoc --go_out=src/ proto/*.proto