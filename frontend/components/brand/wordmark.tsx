import { cn } from "@/lib/utils";

export function Wordmark({ className }: { className?: string }) {
  return (
    <span className={cn("inline-flex items-baseline gap-1 text-xl font-bold tracking-tight", className)}>
      forkly
      <span aria-hidden className="size-1.5 rounded-full bg-rim" />
    </span>
  );
}
