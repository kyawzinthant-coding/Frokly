import { defineConfig } from "orval";

export default defineConfig({
  forkly: {
    input: "../openapi/openapi.yaml",
    output: {
      mode: "tags-split",
      target: "./lib/api/generated",
      schemas: "./lib/api/generated/models",
      client: "react-query",
      httpClient: "axios",
      override: {
        mutator: {
          path: "./lib/api/axios-instance.ts",
          name: "customInstance",
        },
      },
    },
  },
});
