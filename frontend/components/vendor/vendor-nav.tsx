"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { ClipboardListIcon, UtensilsCrossedIcon } from "lucide-react";
import { cn } from "@/lib/utils";

const links = [
  { href: "/vendor", label: "Orders", Icon: ClipboardListIcon },
  { href: "/vendor/menu", label: "Menu", Icon: UtensilsCrossedIcon },
];

// "/vendor" is the parent of every vendor page, so it only matches exactly;
// the others also match their sub-pages.
function isActive(pathname: string, href: string) {
  return href === "/vendor" ? pathname === href : pathname.startsWith(href);
}

// Laptop and tablet: a segmented control inside the jade rail.
export function RailNav() {
  const pathname = usePathname();

  return (
    <nav aria-label="Vendor" className="hidden md:block">
      <ul className="flex gap-1 rounded-full bg-black/15 p-1">
        {links.map(({ href, label, Icon }) => {
          const active = isActive(pathname, href);
          return (
            <li key={href}>
              <Link
                href={href}
                aria-current={active ? "page" : undefined}
                className={cn(
                  "flex h-11 items-center gap-2 rounded-full px-4 text-base outline-none transition-[background-color,color,transform] duration-150",
                  "focus-visible:ring-3 focus-visible:ring-primary-foreground/60 active:scale-[0.97]",
                  active
                    ? "bg-primary-foreground font-semibold text-primary"
                    : "font-medium text-primary-foreground hover:bg-primary-foreground/10",
                )}
              >
                <Icon aria-hidden className="size-[18px]" />
                {label}
              </Link>
            </li>
          );
        })}
      </ul>
    </nav>
  );
}

// Phone: a bottom tab bar, where a thumb reaches.
export function TabBarNav() {
  const pathname = usePathname();

  return (
    <nav
      aria-label="Vendor"
      className="fixed inset-x-0 bottom-0 z-40 border-t border-border bg-card pb-[env(safe-area-inset-bottom)] md:hidden"
    >
      <ul className="grid grid-cols-2">
        {links.map(({ href, label, Icon }) => {
          const active = isActive(pathname, href);
          return (
            <li key={href}>
              <Link
                href={href}
                aria-current={active ? "page" : undefined}
                className="group flex h-16 flex-col items-center justify-center gap-1 outline-none focus-visible:bg-accent"
              >
                <span
                  className={cn(
                    "flex h-8 w-14 items-center justify-center rounded-full transition-colors duration-150 group-active:scale-95",
                    active ? "bg-accent text-primary" : "text-muted-foreground",
                  )}
                >
                  <Icon aria-hidden className="size-5" />
                </span>
                <span
                  className={cn(
                    "text-sm",
                    active ? "font-semibold text-foreground" : "font-medium text-muted-foreground",
                  )}
                >
                  {label}
                </span>
              </Link>
            </li>
          );
        })}
      </ul>
    </nav>
  );
}
