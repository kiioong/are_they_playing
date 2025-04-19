import { App } from "vue";
import { GrpcClientService } from "@/services/grpc/grpc-client.service";
import { AuthService } from "@/services/auth.service";
import { SERVICES } from "@/keys";
import { LeagueManagementService } from "@/services/league-management.service";
import { IconService } from "@/services/icon.service";
export default {
  install: (app: App) => {
    const services = {
      grpcClient: GrpcClientService.getInstance(),
      authService: new AuthService(),
      leagueManagementService: new LeagueManagementService(),
      iconService: new IconService(),
    };

    app.provide(SERVICES, services);
  },
};
