// Prices travel as integer satang (1 baht = 100 satang), the same as the API.
// Conversions work on the digits as text: 1.15 * 100 is 114.99999999999999 in JavaScript.

const BAHT_INPUT = /^(\d{1,7})(?:\.(\d{1,2}))?$/;

const thb = new Intl.NumberFormat("th-TH", { style: "currency", currency: "THB" });

// "65" -> 6500, "65.5" -> 6550, "1,250.50" -> 125050. Anything else -> null.
export function parseBahtToSatang(input: string): number | null {
  const match = BAHT_INPUT.exec(input.replace(/[฿,\s]/g, ""));
  if (!match) return null;

  const [, baht, fraction = ""] = match;
  return Number(baht) * 100 + Number(fraction.padEnd(2, "0"));
}

// 6550 -> "฿65.50", for showing a price.
export function formatSatang(satang: number): string {
  return thb.format(satang / 100);
}

// 6550 -> "65.50", for putting a saved price back into an input.
export function satangToInput(satang: number): string {
  return `${Math.floor(satang / 100)}.${String(satang % 100).padStart(2, "0")}`;
}
