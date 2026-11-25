"use client";

import { useUpdateVendorProduct } from "@/lib/api/generated";
import type { Product } from "@/lib/api/generated/models";
import { Switch } from "@/components/ui/switch";
import { useRefreshMenu } from "./menu-api";

// Saves on every flip. The API only has a full update (PUT), so the rest of the
// product is sent back unchanged with the new isAvailable.
export function AvailabilitySwitch({ product }: { product: Product }) {
  const refreshMenu = useRefreshMenu();
  // onSettled returns the refetch promise, so the mutation stays pending until the
  // list holds the new value; the switch never flips back for a moment.
  const update = useUpdateVendorProduct({ mutation: { onSettled: () => refreshMenu() } });

  // While saving, show what was asked for instead of the old value.
  const checked = update.isPending ? update.variables.data.isAvailable : product.isAvailable;

  return (
    <div className="flex flex-col items-end gap-1">
      <Switch
        size="lg"
        aria-label={`${product.name} on sale`}
        checked={checked}
        disabled={update.isPending}
        onCheckedChange={(next) =>
          update.mutate({
            productId: product.id,
            data: {
              name: product.name,
              description: product.description,
              price: product.price,
              imageUrl: product.imageUrl,
              categoryId: product.categoryId,
              isAvailable: next,
            },
          })
        }
      />
      {update.isError && (
        <p role="alert" className="text-sm text-destructive">
          Not saved
        </p>
      )}
    </div>
  );
}
