"use client";

import { useState, type FormEvent } from "react";
import { useRouter } from "next/navigation";
import { useQueryClient } from "@tanstack/react-query";
import { CircleAlertIcon } from "lucide-react";
import { getVendorMeQueryKey, useVendorLogin } from "@/lib/api/generated";
import { Wordmark } from "@/components/brand/wordmark";
import { FloorPlan } from "@/components/vendor/floor-plan";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

export default function VendorLoginPage() {
  const router = useRouter();
  const queryClient = useQueryClient();
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const login = useVendorLogin({
    mutation: {
      onSuccess: (vendor) => {
        // Login already returns the profile, so seed the /me cache and skip a second request.
        queryClient.setQueryData(getVendorMeQueryKey(), vendor);
        router.replace("/vendor");
      },
    },
  });

  function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    login.mutate({ data: { username, password } });
  }

  const errorMessage = login.isError
    ? login.error.response?.status === 401
      ? "That username and password don't match. Check both and try again."
      : "Can't reach Forkly right now. Check your connection and try again."
    : null;

  return (
    <main className="relative flex flex-1 items-center justify-center overflow-hidden p-4 py-10">
      <FloorPlan />

      <section className="relative w-full max-w-sm rounded-2xl border-3 border-primary bg-card p-6 pt-7">
        <p className="absolute -top-3.5 left-5 rounded-full bg-primary px-3 py-1 text-sm font-semibold text-primary-foreground">
          Your stall
        </p>

        <div className="mb-6 flex items-baseline justify-between gap-4">
          <h1 className="text-2xl font-semibold">Sign in</h1>
          <Wordmark className="text-lg" />
        </div>

        <form onSubmit={handleSubmit} className="flex flex-col gap-5">
          <div className="flex flex-col gap-2">
            <Label htmlFor="username" className="text-base">
              Username
            </Label>
            <Input
              id="username"
              className="h-12 px-3 text-base md:text-base"
              autoComplete="username"
              autoCapitalize="none"
              autoFocus
              required
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              aria-invalid={login.isError || undefined}
            />
          </div>
          <div className="flex flex-col gap-2">
            <Label htmlFor="password" className="text-base">
              Password
            </Label>
            <Input
              id="password"
              type="password"
              className="h-12 px-3 text-base md:text-base"
              autoComplete="current-password"
              required
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              aria-invalid={login.isError || undefined}
            />
          </div>

          {errorMessage && (
            <p role="alert" className="flex items-start gap-2 text-sm text-destructive">
              <CircleAlertIcon aria-hidden className="mt-0.5 size-4 shrink-0" />
              {errorMessage}
            </p>
          )}

          <Button
            type="submit"
            className="mt-1 h-12 w-full text-base font-semibold"
            disabled={login.isPending}
          >
            {login.isPending ? "Signing in…" : "Sign in"}
          </Button>
        </form>
      </section>
    </main>
  );
}
