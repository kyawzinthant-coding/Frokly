import { test } from "node:test";
import assert from "node:assert/strict";
import { formatSatang, parseBahtToSatang, satangToInput } from "./money.ts";

test("parseBahtToSatang reads what staff type", () => {
  assert.equal(parseBahtToSatang("65"), 6500);
  assert.equal(parseBahtToSatang("65.5"), 6550);
  assert.equal(parseBahtToSatang("65.50"), 6550);
  assert.equal(parseBahtToSatang(" ฿1,250.05 "), 125005);
  assert.equal(parseBahtToSatang("0.29"), 29);
});

test("parseBahtToSatang avoids floating point rounding", () => {
  // 1.15 * 100 === 114.99999999999999 in JavaScript.
  assert.equal(parseBahtToSatang("1.15"), 115);
});

test("parseBahtToSatang rejects anything that isn't a price", () => {
  for (const input of ["", "abc", "-5", "1.234", "1.", ".5", "1e3", "12345678"]) {
    assert.equal(parseBahtToSatang(input), null, `"${input}"`);
  }
});

test("formatSatang and satangToInput", () => {
  assert.equal(formatSatang(125050), "฿1,250.50");
  assert.equal(satangToInput(6550), "65.50");
  assert.equal(satangToInput(5), "0.05");
});
