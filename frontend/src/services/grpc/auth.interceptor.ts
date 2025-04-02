import {
  MethodInfo,
  NextServerStreamingFn,
  RpcInterceptor,
  RpcOptions,
  ServerStreamingCall,
  UnaryCall,
} from "@protobuf-ts/runtime-rpc";
import {
  NextClientStreamingFn,
  NextDuplexStreamingFn,
  NextUnaryFn,
} from "@protobuf-ts/runtime-rpc/build/types/rpc-interceptor";
import type { ClientStreamingCall } from "@protobuf-ts/runtime-rpc/build/types/client-streaming-call";
import type { DuplexStreamingCall } from "@protobuf-ts/runtime-rpc/build/types/duplex-streaming-call";

export class AuthInterceptor implements RpcInterceptor {
  private readonly tokenProvider: () => string;

  constructor(tokenProvider: () => string) {
    this.tokenProvider = tokenProvider;
  }

  public interceptUnary(
    next: NextUnaryFn,
    method: MethodInfo,
    input: object,
    options: RpcOptions,
  ): UnaryCall {
    const token = this.tokenProvider();

    if (!options.meta) {
      options.meta = {};
    }

    if (token) {
      options.meta["Authorization"] = "Bearer " + token;
    }

    return next(method, input, options);
  }

  public interceptServerStreaming(
    next: NextServerStreamingFn,
    method: MethodInfo,
    input: object,
    options: RpcOptions,
  ): ServerStreamingCall {
    const token = this.tokenProvider();

    if (!options.meta) {
      options.meta = {};
    }

    if (token) {
      options.meta["Authorization"] = "Bearer " + token;
    }

    return next(method, input, options);
  }

  public interceptClientStreaming?(
    next: NextClientStreamingFn,
    method: MethodInfo,
    options: RpcOptions,
  ): ClientStreamingCall {
    return next(method, options);
  }

  public interceptDuplex(
    next: NextDuplexStreamingFn,
    method: MethodInfo,
    options: RpcOptions,
  ): DuplexStreamingCall {
    return next(method, options);
  }
}
