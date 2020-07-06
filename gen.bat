cd protos
protoc --go_out=plugins=grpc:../pbfiles ProductModel.proto
protoc --go_out=plugins=grpc:../pbfiles OrderModel.proto
protoc --go_out=plugins=grpc:../pbfiles UserModel.proto
protoc --go_out=plugins=grpc:../pbfiles Product.proto
protoc --go_out=plugins=grpc:../pbfiles Order.proto
protoc --go_out=plugins=grpc:../pbfiles User.proto
protoc  --grpc-gateway_out=logtostderr=true:../pbfiles Product.proto
protoc  --grpc-gateway_out=logtostderr=true:../pbfiles Order.proto
cd ../