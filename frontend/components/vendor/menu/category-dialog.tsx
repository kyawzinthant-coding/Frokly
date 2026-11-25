"use client";

import { useState, type FormEvent } from "react";
import { useCreateVendorCategory, useUpdateVendorCategory } from "@/lib/api/generated";
import type { Category } from "@/lib/api/generated/models";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { FormError, FormField, fieldProps } from "./form-field";
import { NETWORK_ERROR, errorStatus, useRefreshMenu } from "./menu-api";

// Mounted only while open, so its state starts fresh each time.
// `category` set = rename, unset = create.
export function CategoryDialog({ category, onClose }: { category?: Category; onClose: () => void }) {
  const refreshMenu = useRefreshMenu();
  const [name, setName] = useState(category?.name ?? "");
  const [nameError, setNameError] = useState<string>();
  const [formError, setFormError] = useState<string>();

  const create = useCreateVendorCategory();
  const update = useUpdateVendorCategory();
  const saving = create.isPending || update.isPending;

  function handleError(error: unknown, trimmed: string) {
    switch (errorStatus(error)) {
      case 409:
        setNameError(`You already have a category called “${trimmed}”.`);
        break;
      case 400:
        setNameError("Use 1 to 100 characters.");
        break;
      case 404:
        setFormError("This category was deleted. Close this and refresh.");
        refreshMenu();
        break;
      default:
        setFormError(NETWORK_ERROR);
    }
  }

  function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    setNameError(undefined);
    setFormError(undefined);

    const trimmed = name.trim();
    if (!trimmed) {
      setNameError("Enter a category name.");
      return;
    }

    const options = {
      onSuccess: async () => {
        await refreshMenu();
        onClose();
      },
      onError: (error: unknown) => handleError(error, trimmed),
    };

    if (category) {
      update.mutate({ categoryId: category.id, data: { name: trimmed } }, options);
    } else {
      create.mutate({ data: { name: trimmed } }, options);
    }
  }

  return (
    <Dialog open onOpenChange={(open) => !open && onClose()}>
      <DialogContent className="gap-5 p-6 sm:max-w-md">
        <DialogHeader>
          <DialogTitle className="text-xl font-semibold">
            {category ? "Rename category" : "Add category"}
          </DialogTitle>
          <DialogDescription className="text-base">
            {category
              ? "The new name shows on your menu straight away."
              : "Group your items, for example Noodles, Rice or Drinks."}
          </DialogDescription>
        </DialogHeader>

        <form onSubmit={handleSubmit} className="flex flex-col gap-5" noValidate>
          <FormField id="category-name" label="Name" error={nameError}>
            <Input
              {...fieldProps("category-name", nameError)}
              className="h-12 px-3 text-base md:text-base"
              autoFocus
              maxLength={100}
              value={name}
              onChange={(e) => setName(e.target.value)}
            />
          </FormField>

          {formError && <FormError message={formError} />}

          <DialogFooter className="-mx-6 -mb-6 px-6 py-4">
            <Button type="button" variant="outline" className="h-11 px-5 text-base" onClick={onClose}>
              Cancel
            </Button>
            <Button type="submit" className="h-11 px-5 text-base font-semibold" disabled={saving}>
              {saving ? "Saving…" : category ? "Save name" : "Add category"}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}
