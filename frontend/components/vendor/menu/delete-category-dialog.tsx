"use client";

import { useState } from "react";
import { useDeleteVendorCategory } from "@/lib/api/generated";
import type { Category } from "@/lib/api/generated/models";
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

// The API refuses to delete a category that still has items (409), so when this page
// already knows it has items it explains that instead of offering a button that would fail.
export function DeleteCategoryDialog({
  category,
  itemCount,
  onClose,
}: {
  category: Category;
  itemCount: number;
  onClose: () => void;
}) {
  const refreshMenu = useRefreshMenu();
  const remove = useDeleteVendorCategory();
  const [error, setError] = useState<string>();
  const blocked = itemCount > 0;

  function handleDelete() {
    setError(undefined);
    remove.mutate(
      { categoryId: category.id },
      {
        onSuccess: async () => {
          await refreshMenu();
          onClose();
        },
        onError: async (err) => {
          const status = errorStatus(err);
          if (status === 409) {
            setError(`${category.name} has items again. Move or remove them first.`);
          } else if (status === 404) {
            await refreshMenu();
            onClose();
          } else {
            setError(NETWORK_ERROR);
          }
        },
      },
    );
  }

  return (
    <AlertDialog open onOpenChange={(open) => !open && onClose()}>
      <AlertDialogContent className="gap-5 p-6 sm:max-w-md">
        <AlertDialogHeader>
          <AlertDialogTitle className="text-xl font-semibold">
            {blocked ? `${category.name} still has items` : `Delete ${category.name}?`}
          </AlertDialogTitle>
          <AlertDialogDescription className="text-base">
            {blocked
              ? `Move its ${itemCount} ${itemCount === 1 ? "item" : "items"} to another category or remove them from your menu, then delete it.`
              : "It disappears from your menu. This can't be undone."}
          </AlertDialogDescription>
        </AlertDialogHeader>

        {error && <FormError message={error} />}

        <AlertDialogFooter className="-mx-6 -mb-6 px-6 py-4">
          <AlertDialogCancel className="h-11 px-5 text-base">
            {blocked ? "OK" : "Cancel"}
          </AlertDialogCancel>
          {!blocked && (
            <AlertDialogAction
              variant="destructive"
              className="h-11 px-5 text-base font-semibold"
              disabled={remove.isPending}
              onClick={handleDelete}
            >
              {remove.isPending ? "Deleting…" : "Delete category"}
            </AlertDialogAction>
          )}
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  );
}
