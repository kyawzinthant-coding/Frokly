import { useQueryClient } from "@tanstack/react-query";
import type { AxiosError } from "axios";
import {
  getListVendorCategoriesQueryKey,
  getListVendorProductsQueryKey,
} from "@/lib/api/generated";
import type { ErrorResponse } from "@/lib/api/generated/models";

// After any menu change, refetch both lists so counts, rows and the summary agree.
export function useRefreshMenu() {
  const queryClient = useQueryClient();
  return () =>
    Promise.all([
      queryClient.invalidateQueries({ queryKey: getListVendorCategoriesQueryKey() }),
      queryClient.invalidateQueries({ queryKey: getListVendorProductsQueryKey() }),
    ]);
}

export function errorStatus(error: unknown): number | undefined {
  return (error as AxiosError | null)?.response?.status;
}

// The API's own message, e.g. "invalid input: price must be greater than 0".
export function errorText(error: unknown): string {
  return (error as AxiosError<ErrorResponse> | null)?.response?.data?.error ?? "";
}

export const NETWORK_ERROR = "Couldn't save. Check your connection and try again.";
