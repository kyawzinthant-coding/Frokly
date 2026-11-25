"use client";

import { EllipsisIcon, PencilIcon, PlusIcon, Trash2Icon } from "lucide-react";
import type { Category, Product } from "@/lib/api/generated/models";
import { formatSatang } from "@/lib/money";
import { cn } from "@/lib/utils";
import { Tier } from "@/components/vendor/tier";
import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { AvailabilitySwitch } from "./availability-switch";

export function CategoryTier({
  category,
  products,
  onRename,
  onDelete,
  onAddItem,
  onEditItem,
}: {
  category: Category;
  products: Product[];
  onRename: () => void;
  onDelete: () => void;
  onAddItem: () => void;
  onEditItem: (product: Product) => void;
}) {
  const offCount = products.filter((p) => !p.isAvailable).length;

  return (
    <Tier
      title={category.name}
      bodyClassName="p-0"
      action={
        <div className="-mr-2 flex shrink-0 items-center gap-1">
          <span className="text-sm tabular-nums">
            {products.length} {products.length === 1 ? "item" : "items"}
            {offCount > 0 && ` · ${offCount} off`}
          </span>
          <DropdownMenu>
            <DropdownMenuTrigger
              aria-label={`${category.name} options`}
              className="relative flex size-9 items-center justify-center rounded-lg outline-none after:absolute after:-inset-1 hover:bg-black/5 focus-visible:ring-3 focus-visible:ring-ring/50 data-popup-open:bg-black/5"
            >
              <EllipsisIcon aria-hidden className="size-5" />
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end" className="min-w-44">
              <DropdownMenuItem className="h-11 px-2 text-base" onClick={onRename}>
                <PencilIcon aria-hidden />
                Rename
              </DropdownMenuItem>
              <DropdownMenuItem variant="destructive" className="h-11 px-2 text-base" onClick={onDelete}>
                <Trash2Icon aria-hidden />
                Delete
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        </div>
      }
    >
      {products.length === 0 ? (
        <p className="px-5 pt-5 pb-2 text-base text-muted-foreground">
          No items in {category.name} yet.
        </p>
      ) : (
        <ul className="divide-y divide-border">
          {products.map((product) => (
            <li key={product.id} className="flex min-h-18 items-center gap-3 py-3 pr-3 pl-5">
              <div className="min-w-0 flex-1">
                <p
                  className={cn(
                    "truncate text-base font-semibold",
                    !product.isAvailable && "text-muted-foreground",
                  )}
                >
                  {product.name}
                </p>
                <p className="truncate text-sm text-muted-foreground tabular-nums">
                  {formatSatang(product.price)}
                  {product.description && ` · ${product.description}`}
                </p>
              </div>
              <AvailabilitySwitch product={product} />
              <Button
                variant="ghost"
                className="size-11 shrink-0 text-muted-foreground hover:text-foreground"
                aria-label={`Edit ${product.name}`}
                onClick={() => onEditItem(product)}
              >
                <PencilIcon aria-hidden className="size-5" />
              </Button>
            </li>
          ))}
        </ul>
      )}
      <div className="border-t border-border p-2">
        <Button
          variant="ghost"
          className="h-11 w-full justify-start gap-2 px-3 text-base font-medium text-primary hover:bg-accent hover:text-primary"
          onClick={onAddItem}
        >
          <PlusIcon aria-hidden className="size-5" />
          Add item to {category.name}
        </Button>
      </div>
    </Tier>
  );
}
