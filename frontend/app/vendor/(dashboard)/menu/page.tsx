"use client";

import { useState } from "react";
import { PlusIcon, UtensilsCrossedIcon } from "lucide-react";
import { useListVendorCategories, useListVendorProducts } from "@/lib/api/generated";
import type { Category, Product } from "@/lib/api/generated/models";
import { CategoryDialog } from "@/components/vendor/menu/category-dialog";
import { CategoryTier } from "@/components/vendor/menu/category-tier";
import { DeleteCategoryDialog } from "@/components/vendor/menu/delete-category-dialog";
import { ProductDialog } from "@/components/vendor/menu/product-dialog";
import { Tier } from "@/components/vendor/tier";
import { Button } from "@/components/ui/button";
import { Skeleton } from "@/components/ui/skeleton";

// At most one dialog is open at a time; this says which one and for what.
type OpenDialog =
  | { kind: "category"; category?: Category }
  | { kind: "deleteCategory"; category: Category }
  | { kind: "product"; product?: Product; categoryId?: number }
  | null;

export default function VendorMenuPage() {
  const categories = useListVendorCategories();
  const products = useListVendorProducts();
  const [dialog, setDialog] = useState<OpenDialog>(null);
  const close = () => setDialog(null);

  // A 401 is handled by the layout's redirect to login, so it isn't shown as an error here.
  const failed = [categories, products].some(
    (q) => q.isError && q.error.response?.status !== 401,
  );

  if (failed) {
    return (
      <div className="flex flex-col items-start gap-4">
        <p className="text-base">Can&apos;t load your menu. Check your connection and try again.</p>
        <Button
          className="h-11 px-5 text-base"
          onClick={() => {
            categories.refetch();
            products.refetch();
          }}
        >
          Try again
        </Button>
      </div>
    );
  }

  if (!categories.data || !products.data) return <MenuSkeleton />;

  const allCategories = categories.data;
  const allProducts = products.data;
  const itemsIn = (categoryId: number) => allProducts.filter((p) => p.categoryId === categoryId);
  const hasCategories = allCategories.length > 0;

  return (
    <div className="flex flex-col gap-6">
      <header className="flex flex-wrap items-end justify-between gap-4">
        <div className="flex flex-col gap-1">
          <h1 className="text-2xl font-semibold">Menu</h1>
          <p className="text-base text-muted-foreground">
            {summary(allCategories.length, allProducts)}
          </p>
        </div>
        {hasCategories && (
          <div className="flex w-full gap-2 sm:w-auto">
            <Button
              variant="outline"
              className="h-11 flex-1 gap-2 px-4 text-base sm:flex-none"
              onClick={() => setDialog({ kind: "category" })}
            >
              <PlusIcon aria-hidden className="size-5" />
              Category
            </Button>
            <Button
              className="h-11 flex-1 gap-2 px-4 text-base font-semibold sm:flex-none"
              onClick={() => setDialog({ kind: "product" })}
            >
              <PlusIcon aria-hidden className="size-5" />
              Item
            </Button>
          </div>
        )}
      </header>

      {hasCategories ? (
        <div className="grid items-start gap-6 lg:grid-cols-2">
          {allCategories.map((category) => (
            <CategoryTier
              key={category.id}
              category={category}
              products={itemsIn(category.id)}
              onRename={() => setDialog({ kind: "category", category })}
              onDelete={() => setDialog({ kind: "deleteCategory", category })}
              onAddItem={() => setDialog({ kind: "product", categoryId: category.id })}
              onEditItem={(product) => setDialog({ kind: "product", product })}
            />
          ))}
        </div>
      ) : (
        <EmptyMenu onAddCategory={() => setDialog({ kind: "category" })} />
      )}

      {dialog?.kind === "category" && <CategoryDialog category={dialog.category} onClose={close} />}
      {dialog?.kind === "deleteCategory" && (
        <DeleteCategoryDialog
          category={dialog.category}
          itemCount={itemsIn(dialog.category.id).length}
          onClose={close}
        />
      )}
      {dialog?.kind === "product" && (
        <ProductDialog
          product={dialog.product}
          defaultCategoryId={dialog.categoryId}
          categories={allCategories}
          onClose={close}
        />
      )}
    </div>
  );
}

function summary(categoryCount: number, products: Product[]) {
  const available = products.filter((p) => p.isAvailable).length;
  const categories = `${categoryCount} ${categoryCount === 1 ? "category" : "categories"}`;
  const items = `${products.length} ${products.length === 1 ? "item" : "items"}`;
  return `${categories} · ${items}, ${available} on sale`;
}

function EmptyMenu({ onAddCategory }: { onAddCategory: () => void }) {
  return (
    <Tier title="Menu">
      <div className="flex flex-col items-center gap-3 px-4 py-10 text-center">
        <UtensilsCrossedIcon aria-hidden className="size-10 text-primary" />
        <p className="text-lg font-semibold">Your menu is empty</p>
        <p className="max-w-sm text-base text-muted-foreground">
          Start with a category such as Noodles or Drinks, then add the items you sell in it.
        </p>
        <Button className="mt-2 h-12 gap-2 px-5 text-base font-semibold" onClick={onAddCategory}>
          <PlusIcon aria-hidden className="size-5" />
          Add your first category
        </Button>
      </div>
    </Tier>
  );
}

function MenuSkeleton() {
  return (
    <div className="flex flex-col gap-6" aria-label="Loading menu">
      <div className="flex flex-col gap-2">
        <Skeleton className="h-8 w-24" />
        <Skeleton className="h-5 w-56" />
      </div>
      <div className="grid items-start gap-6 lg:grid-cols-2">
        {[0, 1].map((tier) => (
          <div key={tier} className="overflow-hidden rounded-2xl bg-card ring-1 ring-border">
            <Skeleton className="h-10 rounded-none" />
            <div className="flex flex-col gap-4 p-5">
              <Skeleton className="h-11 w-full" />
              <Skeleton className="h-11 w-5/6" />
              <Skeleton className="h-11 w-4/6" />
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
