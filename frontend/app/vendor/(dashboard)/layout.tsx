"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation";
import { useVendorMe } from "@/lib/api/generated";
import { Wordmark } from "@/components/brand/wordmark";
import { AccountMenu } from "@/components/vendor/account-menu";
import { RailNav, TabBarNav } from "@/components/vendor/vendor-nav";
import { Skeleton } from "@/components/ui/skeleton";

// Auth guard + shell for every vendor page except login. The API is on another origin, so the
// Next.js server can't see the cookie — asking /me from the browser is the only check.
// Children always render (pages show their own loading state) so Next.js can prerender the segment.
export default function VendorDashboardLayout({ children }: { children: React.ReactNode }) {
  const router = useRouter();
  // No retries: a 401 won't fix itself, and retrying delays the redirect.
  const me = useVendorMe({ query: { retry: false } });
  const isUnauthorized = me.error?.response?.status === 401;

  useEffect(() => {
    if (isUnauthorized) router.replace("/vendor/login");
  }, [isUnauthorized, router]);

  return (
    <div className="flex flex-1 flex-col">
      <header className="bg-primary text-primary-foreground">
        <div className="mx-auto flex h-16 w-full max-w-6xl items-center gap-4 px-4 sm:px-6">
          <Wordmark />
          <span aria-hidden className="h-6 w-px bg-primary-foreground/30" />
          {me.data ? (
            <p className="min-w-0 truncate text-lg font-semibold">{me.data.name}</p>
          ) : (
            <Skeleton className="h-5 w-32 bg-primary-foreground/20" />
          )}
          <div className="ml-auto flex items-center gap-3">
            <RailNav />
            {me.data && <AccountMenu username={me.data.username} />}
          </div>
        </div>
      </header>
      {/* Bottom padding on phones keeps content clear of the fixed tab bar. */}
      <main className="mx-auto w-full max-w-6xl flex-1 px-4 pt-6 pb-28 sm:px-6 sm:pt-8 md:pb-8">
        {children}
      </main>
      <TabBarNav />
    </div>
  );
}
