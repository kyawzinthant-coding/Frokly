"use client";

import { useGetStatus } from "@/lib/api/generated";

export default function Home() {
  const { data, isPending, isError, error } = useGetStatus();

  return (
    <div className="flex flex-col flex-1 items-center justify-center gap-4 bg-zinc-50 font-sans dark:bg-black">
      <h1 className="text-2xl font-semibold text-black dark:text-zinc-50">
        Forkly
      </h1>

      {isPending && <p>Checking backend status…</p>}

      {isError && (
        <p className="text-red-600">
          Backend unhealthy: {error.response?.data.status ?? error.message}{" "}
          (database: {error.response?.data.database ?? "unknown"})
        </p>
      )}

      {data && (
        <pre className="rounded bg-black/[.06] px-4 py-3 font-mono text-sm dark:bg-white/[.08]">
          {JSON.stringify(data, null, 2)}
        </pre>
      )}
    </div>
  );
}
