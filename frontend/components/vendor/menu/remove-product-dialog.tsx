"use client";

import { useState } from "react";
import { useDeleteVendorProduct } from "@/lib/api/generated";
import type { Product } from "@/lib/api/generated/models";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "@/components/ui/alert-dialog";
import { FormError } from "./form-field";
import { NETWORK_ERROR, errorStatus, useRefreshMenu } from "./menu-api";

// Removing archives the item on the server; it can't be brought back from the UI,
// so the copy points to the "On sale" switch for a temporary sell-out.
export function RemoveProductDialog({
  product,
  onCancel,
  onRemoved,
}: {
  product: Product;
  onCancel: () => void;
  onRemoved: () => void;
}) {
  const refreshMenu = useRefreshMenu();
  const remove = useDeleteVendorProduct();
  const [error, setError] = useState<string>();

  function handleRemove() {
    setError(undefined);
    remove.mutate(
      { productId: product.id },
      {
        onSuccess: async () => {
          await refreshMenu();
          onRemoved();
        },
        onError: async (err) => {
          // 404: someone else already removed it, which is what we wanted.
          if (errorStatus(err) === 404) {
            await refreshMenu();
            onRemoved();
          } else {
            setError(NETWORK_ERROR);
          }
        },
      },
    );
  }

  return (
    <AlertDialog open onOpenChange={(open) => !open && onCancel()}>
      <AlertDialogContent className="gap-5 p-6 sm:max-w-md">
        <AlertDialogHeader>
          <AlertDialogTitle className="text-xl font-semibold">
            Remove {product.name} from your menu?
          </AlertDialogTitle>
          <AlertDialogDescription className="text-base">
            To sell it again you&apos;ll have to add it again. If it&apos;s only sold out for now,
            turn off On sale instead.
          </AlertDialogDescription>
        </AlertDialogHeader>

        {error && <FormError message={error} />}

        <AlertDialogFooter className="-mx-6 -mb-6 px-6 py-4">
          <AlertDialogCancel className="h-11 px-5 text-base">Keep it</AlertDialogCancel>
          <AlertDialogAction
            variant="destructive"
            className="h-11 px-5 text-base font-semibold"
            disabled={remove.isPending}
            onClick={handleRemove}
          >
            {remove.isPending ? "Removing…" : "Remove item"}
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  );
}
