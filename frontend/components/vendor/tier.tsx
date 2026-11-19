import { cn } from "@/lib/utils";

// One job per tier, like a tiffin carrier: a steel band names the tier, the body holds the work.
export function Tier({
  title,
  className,
  children,
}: {
  title: string;
  className?: string;
  children: React.ReactNode;
}) {
  return (
    <section className={cn("overflow-hidden rounded-2xl bg-card ring-1 ring-border", className)}>
      <h2 className="bg-band px-5 py-2.5 text-sm font-semibold text-band-foreground">{title}</h2>
      <div className="p-5">{children}</div>
    </section>
  );
}
