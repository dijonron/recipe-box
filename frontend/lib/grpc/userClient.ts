import * as grpc from "@grpc/grpc-js";
import * as protoLoader from "@grpc/proto-loader";

const packageDefinition = protoLoader.loadSync("../proto/user.proto", {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});

const protoDescriptor = grpc.loadPackageDefinition(packageDefinition) as any;
const UserService = protoDescriptor.user.User as grpc.ServiceClientConstructor;

export const userClient = new UserService(
  "localhost:50051", // TODO: env
  grpc.credentials.createInsecure()
);
