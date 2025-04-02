import { GrpcClientService } from "@/services/grpc/grpc-client.service";
import { AuthService } from "@/services/auth.service";
import { LeagueManagementService } from "@/services/league-management.service";

export interface ServicesPlugin {
  grpcClient: GrpcClientService;
  authService: AuthService;
  leagueManagementService: LeagueManagementService;
}
