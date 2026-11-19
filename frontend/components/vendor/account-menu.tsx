"use client";

import { useRouter } from "next/navigation";
import { useQueryClient } from "@tanstack/react-query";
import { ChevronDownIcon, CircleAlertIcon, LogOutIcon } from "lucide-react";
import { useVendorLogout } from "@/lib/api/generated";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";

export function AccountMenu({ username }: { username: string }) {
  const router = useRouter();
  const queryClient = useQueryClient();

  const logout = useVendorLogout({
    mutation: {
      onSuccess: () => {
        router.replace("/vendor/login");
        queryClient.clear();
      },
    },
  });

  return (
    <DropdownMenu>
      <DropdownMenuTrigger className="flex h-11 items-center gap-2 rounded-xl px-2 text-primary-foreground outline-none hover:bg-primary-foreground/10 focus-visible:ring-3 focus-visible:ring-primary-foreground/60 data-popup-open:bg-primary-foreground/10">
        <span
          aria-hidden
          className="flex size-8 items-center justify-center rounded-full bg-primary-foreground text-sm font-bold text-primary uppercase"
        >
          {username.charAt(0)}
        </span>
        <span className="max-w-32 truncate text-base font-medium">{username}</span>
        <ChevronDownIcon aria-hidden className="size-4" />
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end" className="min-w-56">
        <DropdownMenuGroup>
          <DropdownMenuLabel className="px-2 py-1.5 text-sm">
            Signed in as <span className="font-semibold text-foreground">{username}</span>
          </DropdownMenuLabel>
        </DropdownMenuGroup>
        <DropdownMenuSeparator />
        <DropdownMenuItem
          className="h-11 px-2 text-base"
          closeOnClick={false}
          disabled={logout.isPending}
          onClick={() => logout.mutate()}
        >
          <LogOutIcon aria-hidden />
          {logout.isPending ? "Signing out…" : "Sign out"}
        </DropdownMenuItem>
        {logout.isError && (
          <p role="alert" className="flex items-start gap-2 px-2 py-1.5 text-sm text-destructive">
            <CircleAlertIcon aria-hidden className="mt-0.5 size-4 shrink-0" />
            Couldn&apos;t sign out. Try again.
          </p>
        )}
      </DropdownMenuContent>
    </DropdownMenu>
  );
}
