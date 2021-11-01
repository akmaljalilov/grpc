```
go mod download
protoc --go_out=plugins=grpc:recommendation --go_opt=paths=source_relative recommendations.proto

cd python-server
pip install -r requirements.txt
make proto
```