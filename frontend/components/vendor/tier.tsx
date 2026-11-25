import { cn } from "@/lib/utils";

// One job per tier, like a tiffin carrier: a steel band names the tier, the body holds the work.
// `action` sits at the right end of the band (a count, a menu).
export function Tier({
  title,
  action,
  className,
  bodyClassName,
  children,
}: {
  title: string;
  action?: React.ReactNode;
  className?: string;
  bodyClassName?: string;
  children: React.ReactNode;
}) {
  return (
    <section className={cn("overflow-hidden rounded-2xl bg-card ring-1 ring-border", className)}>
      <div className="flex min-h-10 items-center gap-3 bg-band px-5 py-1.5 text-band-foreground">
        <h2 className="min-w-0 flex-1 truncate text-sm font-semibold">{title}</h2>
        {action}
      </div>
      <div className={cn("p-5", bodyClassName)}>{children}</div>
    </section>
  );
}
