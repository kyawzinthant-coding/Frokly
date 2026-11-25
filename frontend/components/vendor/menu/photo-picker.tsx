"use client";

import { useEffect, useMemo, useRef, useState } from "react";
import { ImagePlusIcon, XIcon } from "lucide-react";
import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";

const ACCEPTED_TYPES = ["image/jpeg", "image/png", "image/webp"];
const MAX_BYTES = 5 * 1024 * 1024;

// Returns why a file can't be used, or undefined when it can.
export function checkPhoto(file: File): string | undefined {
  if (!ACCEPTED_TYPES.includes(file.type)) return "Use a JPG, PNG or WebP photo.";
  if (file.size > MAX_BYTES) return "Choose a photo under 5 MB.";
  return undefined;
}

// One tile that is the whole control: click or drop a photo to pick it, then it shows the photo.
// It doesn't upload: the parent decides what to do with the file.
// `existingUrl` is the photo the item already has, shown until a new one is picked or removed.
export function PhotoPicker({
  id,
  file,
  existingUrl,
  describedBy,
  invalid,
  onPick,
  onRemove,
}: {
  id: string;
  file: File | null;
  existingUrl: string | null;
  describedBy?: string;
  invalid?: boolean;
  onPick: (file: File) => void;
  onRemove: () => void;
}) {
  const inputRef = useRef<HTMLInputElement>(null);
  const [dragging, setDragging] = useState(false);
  const preview = useObjectUrl(file) ?? existingUrl;
  const openPicker = () => inputRef.current?.click();

  const input = (
    <input
      ref={inputRef}
      id={id}
      type="file"
      accept={ACCEPTED_TYPES.join(",")}
      aria-describedby={describedBy}
      aria-invalid={invalid || undefined}
      className="sr-only"
      tabIndex={-1}
      onChange={(e) => {
        const picked = e.target.files?.[0];
        // Clear it so picking the same file again still fires onChange.
        e.target.value = "";
        if (picked) onPick(picked);
      }}
    />
  );

  if (!preview) {
    return (
      <>
        <button
          type="button"
          onClick={openPicker}
          onDragOver={(e) => {
            e.preventDefault();
            setDragging(true);
          }}
          onDragLeave={() => setDragging(false)}
          onDrop={(e) => {
            e.preventDefault();
            setDragging(false);
            const dropped = e.dataTransfer.files?.[0];
            if (dropped) onPick(dropped);
          }}
          className={cn(
            "flex w-full items-center gap-4 rounded-xl border-2 border-dashed p-3 text-left outline-none transition-colors",
            "hover:border-primary hover:bg-accent/60 focus-visible:ring-3 focus-visible:ring-ring/50 active:scale-[0.99]",
            dragging ? "border-primary bg-accent" : "border-input bg-card",
            invalid && "border-destructive",
          )}
        >
          <span className="flex size-14 shrink-0 items-center justify-center rounded-lg bg-accent text-primary">
            <ImagePlusIcon aria-hidden className="size-6" />
          </span>
          <span className="flex min-w-0 flex-col">
            <span className="text-base font-semibold">Add a photo</span>
            <span className="text-sm text-muted-foreground">JPG, PNG or WebP · up to 5 MB</span>
          </span>
        </button>
        {input}
      </>
    );
  }

  return (
    <div className="flex w-full items-center gap-4 rounded-xl border border-border bg-card p-3">
      {/* A blob: preview can't go through next/image, so a plain img is used here. */}
      {/* eslint-disable-next-line @next/next/no-img-element */}
      <img src={preview} alt="" className="size-14 shrink-0 rounded-lg object-cover" />
      <span className="flex min-w-0 flex-1 flex-col">
        <span className="truncate text-base font-semibold">{file ? file.name : "Current photo"}</span>
        {file && (
          <span className="text-sm text-muted-foreground tabular-nums">{formatSize(file.size)}</span>
        )}
      </span>
      <Button type="button" variant="ghost" className="h-11 px-3 text-base" onClick={openPicker}>
        Change
      </Button>
      <Button
        type="button"
        variant="ghost"
        aria-label="Remove photo"
        className="size-11 text-muted-foreground hover:text-foreground"
        onClick={onRemove}
      >
        <XIcon aria-hidden className="size-5" />
      </Button>
      {input}
    </div>
  );
}

function formatSize(bytes: number) {
  return bytes < 1024 * 1024
    ? `${Math.max(1, Math.round(bytes / 1024))} KB`
    : `${(bytes / (1024 * 1024)).toFixed(1)} MB`;
}

// A temporary URL for previewing a local file, released when the file changes or the picker closes.
function useObjectUrl(file: File | null) {
  const url = useMemo(() => (file ? URL.createObjectURL(file) : null), [file]);

  useEffect(() => {
    return () => {
      if (url) URL.revokeObjectURL(url);
    };
  }, [url]);

  return url;
}
