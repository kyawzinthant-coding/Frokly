"use client";

import { useState, type FormEvent } from "react";
import { useCreateVendorProduct, useUpdateVendorProduct } from "@/lib/api/generated";
import type { Category, Product, ProductRequest } from "@/lib/api/generated/models";
import { parseBahtToSatang, satangToInput } from "@/lib/money";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Switch } from "@/components/ui/switch";
import { Textarea } from "@/components/ui/textarea";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { FormError, FormField, fieldProps } from "./form-field";
import { NETWORK_ERROR, errorStatus, errorText, useRefreshMenu } from "./menu-api";
import { PhotoPicker, checkPhoto } from "./photo-picker";
import { RemoveProductDialog } from "./remove-product-dialog";

type Field = "name" | "category" | "price" | "description" | "photo";
type Errors = Partial<Record<Field | "form", string>>;

const inputClass = "h-12 px-3 text-base md:text-base";

// Mounted only while open. `product` set = edit, unset = create (optionally in `defaultCategoryId`).
export function ProductDialog({
  product,
  defaultCategoryId,
  categories,
  onClose,
}: {
  product?: Product;
  defaultCategoryId?: number;
  categories: Category[];
  onClose: () => void;
}) {
  const refreshMenu = useRefreshMenu();
  const [name, setName] = useState(product?.name ?? "");
  const [categoryId, setCategoryId] = useState<number | null>(
    product?.categoryId ?? defaultCategoryId ?? null,
  );
  const [price, setPrice] = useState(product ? satangToInput(product.price) : "");
  const [description, setDescription] = useState(product?.description ?? "");
  // Photo upload isn't built yet: a picked file is previewed but not sent.
  // The item keeps the photo it already has unless staff remove it.
  const [photo, setPhoto] = useState<File | null>(null);
  const [keepExistingPhoto, setKeepExistingPhoto] = useState(true);
  const [isAvailable, setIsAvailable] = useState(product?.isAvailable ?? true);
  const [errors, setErrors] = useState<Errors>({});
  const [confirmRemove, setConfirmRemove] = useState(false);

  const create = useCreateVendorProduct();
  const update = useUpdateVendorProduct();
  const saving = create.isPending || update.isPending;

  // Checks what the browser can check, so most mistakes never reach the API.
  function validate(): { data?: ProductRequest; errors: Errors } {
    const found: Errors = {};
    const trimmedName = name.trim();
    const satang = parseBahtToSatang(price);

    if (!trimmedName) found.name = "Enter the item's name.";
    if (categoryId === null) found.category = "Choose a category.";
    if (satang === null) found.price = "Enter a price like 65 or 65.50.";
    else if (satang === 0) found.price = "The price must be more than ฿0.";

    if (Object.keys(found).length > 0 || categoryId === null || satang === null) {
      return { errors: found };
    }
    return {
      errors: found,
      data: {
        name: trimmedName,
        categoryId,
        price: satang,
        description: description.trim(),
        imageUrl: keepExistingPhoto ? (product?.imageUrl ?? null) : null,
        isAvailable,
      },
    };
  }

  // Puts the API's answer next to the field it is about.
  function handleError(error: unknown, sentName: string) {
    const status = errorStatus(error);
    const text = errorText(error);

    if (status === 409) {
      setErrors({ name: `You already have an item called “${sentName}”.` });
    } else if (status === 404 && text.includes("category")) {
      setErrors({ category: "That category was deleted. Choose another." });
      refreshMenu();
    } else if (status === 404) {
      setErrors({ form: "This item was removed from your menu. Close this and refresh." });
      refreshMenu();
    } else if (status === 400) {
      setErrors(mapBadRequest(text));
    } else {
      setErrors({ form: NETWORK_ERROR });
    }
  }

  function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    const { data, errors: found } = validate();
    setErrors(found);
    if (!data) return;

    const options = {
      onSuccess: async () => {
        await refreshMenu();
        onClose();
      },
      onError: (error: unknown) => handleError(error, data.name),
    };

    if (product) update.mutate({ productId: product.id, data }, options);
    else create.mutate({ data }, options);
  }

  return (
    <>
      <Dialog open={!confirmRemove} onOpenChange={(open) => !open && onClose()}>
        <DialogContent className="max-h-[calc(100dvh-2rem)] gap-5 overflow-y-auto p-6 sm:max-w-lg">
          <DialogHeader>
            <DialogTitle className="text-xl font-semibold">
              {product ? `Edit ${product.name}` : "Add item"}
            </DialogTitle>
            <DialogDescription className="text-base">
              {product ? "Changes show on your menu as soon as you save." : "Customers see this on your menu."}
            </DialogDescription>
          </DialogHeader>

          <form onSubmit={handleSubmit} className="flex flex-col gap-5" noValidate>
            <FormField id="product-name" label="Name" error={errors.name}>
              <Input
                {...fieldProps("product-name", errors.name)}
                className={inputClass}
                autoFocus={!product}
                maxLength={100}
                value={name}
                onChange={(e) => setName(e.target.value)}
              />
            </FormField>

            <div className="grid gap-5 sm:grid-cols-2">
              <FormField id="product-category" label="Category" error={errors.category}>
                <Select
                  items={categories.map((c) => ({ value: c.id, label: c.name }))}
                  value={categoryId}
                  onValueChange={(value) => setCategoryId(value)}
                >
                  <SelectTrigger
                    {...fieldProps("product-category", errors.category)}
                    className="h-12 w-full rounded-lg px-3 text-base data-[size=default]:h-12"
                  >
                    <SelectValue placeholder="Choose" />
                  </SelectTrigger>
                  <SelectContent>
                    {categories.map((c) => (
                      <SelectItem key={c.id} value={c.id} className="h-11 px-3 text-base">
                        {c.name}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
              </FormField>

              <FormField id="product-price" label="Price" error={errors.price}>
                <div className="relative">
                  <span
                    aria-hidden
                    className="pointer-events-none absolute inset-y-0 left-3 flex items-center text-base text-muted-foreground"
                  >
                    ฿
                  </span>
                  <Input
                    {...fieldProps("product-price", errors.price)}
                    className={`${inputClass} pl-7 tabular-nums`}
                    inputMode="decimal"
                    placeholder="65.00"
                    value={price}
                    onChange={(e) => setPrice(e.target.value)}
                  />
                </div>
              </FormField>
            </div>

            <FormField
              id="product-description"
              label="Description (optional)"
              hint="One line about what's in it."
              error={errors.description}
            >
              <Textarea
                {...fieldProps("product-description", errors.description, true)}
                className="min-h-20 px-3 py-2.5 text-base md:text-base"
                maxLength={500}
                value={description}
                onChange={(e) => setDescription(e.target.value)}
              />
            </FormField>

            <FormField
              id="product-photo"
              label="Photo (optional)"
              hint="Photo saving is coming soon."
              error={errors.photo}
            >
              <PhotoPicker
                id="product-photo"
                file={photo}
                existingUrl={keepExistingPhoto ? (product?.imageUrl ?? null) : null}
                describedBy="product-photo-note"
                invalid={Boolean(errors.photo)}
                onPick={(file) => {
                  const problem = checkPhoto(file);
                  setErrors((prev) => ({ ...prev, photo: problem }));
                  if (!problem) setPhoto(file);
                }}
                onRemove={() => {
                  setPhoto(null);
                  setKeepExistingPhoto(false);
                  setErrors((prev) => ({ ...prev, photo: undefined }));
                }}
              />
            </FormField>

            <label className="flex items-center justify-between gap-4 rounded-xl bg-muted px-4 py-3">
              <span className="flex flex-col">
                <span className="text-base font-semibold">On sale</span>
                <span className="text-sm text-muted-foreground">
                  Turn off when it&apos;s sold out. Customers won&apos;t see it.
                </span>
              </span>
              <Switch size="lg" checked={isAvailable} onCheckedChange={setIsAvailable} />
            </label>

            {errors.form && <FormError message={errors.form} />}

            <DialogFooter className="-mx-6 -mb-6 px-6 py-4 sm:items-center">
              {product && (
                <Button
                  type="button"
                  variant="ghost"
                  className="h-11 px-3 text-base text-destructive hover:bg-destructive/10 hover:text-destructive sm:mr-auto"
                  onClick={() => setConfirmRemove(true)}
                >
                  Remove from menu
                </Button>
              )}
              <Button type="button" variant="outline" className="h-11 px-5 text-base" onClick={onClose}>
                Cancel
              </Button>
              <Button type="submit" className="h-11 px-5 text-base font-semibold" disabled={saving}>
                {saving ? "Saving…" : product ? "Save changes" : "Add item"}
              </Button>
            </DialogFooter>
          </form>
        </DialogContent>
      </Dialog>

      {product && confirmRemove && (
        <RemoveProductDialog
          product={product}
          onCancel={() => setConfirmRemove(false)}
          onRemoved={onClose}
        />
      )}
    </>
  );
}

// The API answers 400 with e.g. "invalid input: price must be greater than 0".
function mapBadRequest(text: string): Errors {
  const message = text.replace(/^invalid input:\s*/, "");
  const sentence = message.charAt(0).toUpperCase() + message.slice(1) + ".";

  if (text.includes("price")) return { price: sentence };
  if (text.includes("categoryId")) return { category: "Choose a category." };
  if (text.includes("description")) return { description: sentence };
  if (text.includes("name")) return { name: sentence };
  return { form: NETWORK_ERROR };
}
