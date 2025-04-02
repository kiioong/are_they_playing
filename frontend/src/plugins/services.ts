import { App } from "vue";
import { GrpcClientService } from "@/services/grpc/grpc-client.service";
import { AuthService } from "@/services/auth.service";
import { SERVICES } from "@/keys";
import { LeagueManagementService } from "@/services/league-management.service";

export default {
  install: (app: App) => {
    const services = {
      grpcClient: GrpcClientService.getInstance(),
      authService: new AuthService(),
      leagueManagementService: new LeagueManagementService(),
    };

    app.provide(SERVICES, services);
  },
};
