import * as grpc from "@grpc/grpc-js";
import * as protoLoader from "@grpc/proto-loader";

const packageDefinition = protoLoader.loadSync("../proto/auth.proto", {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});

const protoDescriptor = grpc.loadPackageDefinition(packageDefinition) as any;
const AuthService = protoDescriptor.auth.Auth as grpc.ServiceClientConstructor;

export const authClient = new AuthService(
  "localhost:50052", // TODO: env
  grpc.credentials.createInsecure()
);
