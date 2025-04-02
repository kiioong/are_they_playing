import { InjectionKey } from "vue";
import { ServicesPlugin } from "@/types";

export const SERVICES: InjectionKey<ServicesPlugin> = Symbol("SERVICES");
