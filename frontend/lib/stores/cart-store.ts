import { create } from "zustand";

export type CartItem = {
  productId: string;
  vendorId: string;
  name: string;
  unitPriceCents: number;
  quantity: number;
};

type CartState = {
  items: CartItem[];
  addItem: (item: CartItem) => void;
  removeItem: (productId: string) => void;
  clear: () => void;
};

// Skeleton only. The actual cart rules — merging a duplicate product instead
// of adding a second line, grouping by vendorId for the split-checkout step,
// computing per-vendor and combined totals — belong to Milestone 4
// (ordering/checkout). Implement addItem/removeItem when you get there
// instead of treating this file as finished.
export const useCartStore = create<CartState>((set) => ({
  items: [],
  addItem: () => {
    throw new Error("not implemented");
  },
  removeItem: () => {
    throw new Error("not implemented");
  },
  clear: () => set({ items: [] }),
}));
