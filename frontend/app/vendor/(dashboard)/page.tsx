"use client";

import { ClipboardListIcon } from "lucide-react";
import { useVendorMe } from "@/lib/api/generated";
import { Tier } from "@/components/vendor/tier";
import { Button } from "@/components/ui/button";
import { Skeleton } from "@/components/ui/skeleton";

export default function VendorDashboardPage() {
  // Same query as the layout, so this reads from the shared cache — no second request.
  const me = useVendorMe({ query: { retry: false } });

  if (me.isError && me.error.response?.status !== 401) {
    return (
      <div className="flex flex-col items-start gap-4">
        <p className="text-base">Can&apos;t load your stall. Check your connection and try again.</p>
        <Button className="h-11 px-5 text-base" onClick={() => me.refetch()}>
          Try again
        </Button>
      </div>
    );
  }

  const vendor = me.data;

  return (
    <div className="grid items-start gap-6 lg:grid-cols-3">
      <Tier title="Orders" className="lg:col-span-2">
        <div className="flex flex-col items-center gap-3 px-4 py-10 text-center">
          <ClipboardListIcon aria-hidden className="size-10 text-primary" />
          <p className="text-lg font-semibold">No orders yet</p>
          <p className="max-w-sm text-base text-muted-foreground">
            When a table orders from your stall, it appears here. You&apos;ll have 10 minutes to
            accept or reject it.
          </p>
        </div>
      </Tier>

      <Tier title="Signed in">
        {vendor ? (
          <dl className="flex flex-col gap-4">
            <div>
              <dt className="text-sm text-muted-foreground">Stall</dt>
              <dd className="text-lg font-semibold">{vendor.name}</dd>
            </div>
            <div>
              <dt className="text-sm text-muted-foreground">Staff username</dt>
              <dd className="text-lg font-semibold">{vendor.username}</dd>
            </div>
          </dl>
        ) : (
          <div className="flex flex-col gap-4" aria-label="Loading account">
            <Skeleton className="h-12 w-3/4" />
            <Skeleton className="h-12 w-1/2" />
          </div>
        )}
      </Tier>
    </div>
  );
}
