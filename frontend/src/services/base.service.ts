import { GrpcClientService } from "./grpc/grpc-client.service";
import { RpcError } from "@protobuf-ts/runtime-rpc";

export class BaseService {
  protected grpcClient: GrpcClientService;

  constructor() {
    this.grpcClient = GrpcClientService.getInstance();
  }

  public setToken(token: string) {
    this.grpcClient.setToken(token);
  }

  protected async handleGrpcError(error: unknown): Promise<never> {
    if (error instanceof RpcError) {
      // Handle specific gRPC error codes
      switch (error.code) {
        case "UNAUTHENTICATED":
          // Check the error message for more specific authentication errors
          if (error.message.includes("token has expired")) {
            throw new Error("Your session has expired. Please login again.");
          } else if (error.message.includes("invalid password")) {
            throw new Error("Invalid username or password.");
          } else if (error.message.includes("invalid token")) {
            throw new Error("Invalid authentication token. Please login again.");
          } else {
            throw new Error("Authentication failed. Please login again.");
          }
        case "PERMISSION_DENIED":
          throw new Error("You don't have permission to perform this action.");
        case "NOT_FOUND":
          if (error.message.includes("user not found")) {
            throw new Error("User not found. Please check your credentials.");
          }
          throw new Error("The requested resource was not found.");
        case "INVALID_ARGUMENT":
          throw new Error("Invalid request: " + error.message);
        case "INTERNAL":
          throw new Error("An internal server error occurred. Please try again later.");
        default:
          throw new Error(error.message || "An unexpected error occurred.");
      }
    }
    
    // Handle non-gRPC errors
    if (error instanceof Error) {
      throw error;
    }
    
    throw new Error("An unknown error occurred.");
  }
}
