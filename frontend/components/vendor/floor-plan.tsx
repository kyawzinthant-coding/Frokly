// Decorative food court plan behind the login form: numbered stalls around the edge, tables in the
// middle. One drawing per screen shape (landscape, tablet portrait, phone), scaled to fit, never cropped.

type Stall = { x: number; y: number; w: number; h: number; n: string };
type Table = { cx: number; cy: number };

const STALL_RADIUS = 10;
const TABLE_RADIUS = 18;

function Plan({ viewBox, stalls, tables, className }: {
  viewBox: string;
  stalls: Stall[];
  tables: Table[];
  className: string;
}) {
  return (
    <svg
      viewBox={viewBox}
      preserveAspectRatio="xMidYMid meet"
      aria-hidden
      className={`absolute inset-0 size-full ${className}`}
    >
      {stalls.map((s) => (
        <g key={s.n}>
          <rect
            x={s.x}
            y={s.y}
            width={s.w}
            height={s.h}
            rx={STALL_RADIUS}
            className="fill-muted stroke-border"
            strokeWidth={2}
          />
          <text x={s.x + 14} y={s.y + 26} className="fill-muted-foreground text-base font-bold">
            {s.n}
          </text>
        </g>
      ))}
      {tables.map((t) => (
        <circle
          key={`${t.cx}-${t.cy}`}
          cx={t.cx}
          cy={t.cy}
          r={TABLE_RADIUS}
          className="fill-none stroke-border"
          strokeWidth={2}
        />
      ))}
    </svg>
  );
}

const LANDSCAPE_STALLS: Stall[] = [
  { x: 80, y: 60, w: 150, h: 90, n: "01" },
  { x: 250, y: 60, w: 150, h: 90, n: "02" },
  { x: 420, y: 60, w: 150, h: 90, n: "03" },
  { x: 630, y: 60, w: 150, h: 90, n: "04" },
  { x: 800, y: 60, w: 150, h: 90, n: "05" },
  { x: 970, y: 60, w: 150, h: 90, n: "06" },
  { x: 1030, y: 190, w: 90, h: 140, n: "07" },
  { x: 1030, y: 370, w: 90, h: 140, n: "08" },
  { x: 970, y: 550, w: 150, h: 90, n: "09" },
  { x: 800, y: 550, w: 150, h: 90, n: "10" },
  { x: 630, y: 550, w: 150, h: 90, n: "11" },
  { x: 420, y: 550, w: 150, h: 90, n: "12" },
  { x: 250, y: 550, w: 150, h: 90, n: "13" },
  { x: 80, y: 550, w: 150, h: 90, n: "14" },
  { x: 80, y: 370, w: 90, h: 140, n: "15" },
  { x: 80, y: 190, w: 90, h: 140, n: "16" },
];

const LANDSCAPE_TABLES: Table[] = [
  { cx: 240, cy: 230 }, { cx: 320, cy: 300 }, { cx: 240, cy: 380 }, { cx: 320, cy: 460 },
  { cx: 880, cy: 230 }, { cx: 960, cy: 300 }, { cx: 880, cy: 380 }, { cx: 960, cy: 460 },
];

const TABLET_STALLS: Stall[] = [
  { x: 30, y: 50, w: 120, h: 80, n: "01" },
  { x: 170, y: 50, w: 120, h: 80, n: "02" },
  { x: 310, y: 50, w: 120, h: 80, n: "03" },
  { x: 450, y: 50, w: 120, h: 80, n: "04" },
  { x: 490, y: 230, w: 80, h: 130, n: "05" },
  { x: 490, y: 500, w: 80, h: 130, n: "06" },
  { x: 450, y: 730, w: 120, h: 80, n: "07" },
  { x: 310, y: 730, w: 120, h: 80, n: "08" },
  { x: 170, y: 730, w: 120, h: 80, n: "09" },
  { x: 30, y: 730, w: 120, h: 80, n: "10" },
  { x: 30, y: 500, w: 80, h: 130, n: "11" },
  { x: 30, y: 230, w: 80, h: 130, n: "12" },
];

const TABLET_TABLES: Table[] = [
  { cx: 160, cy: 190 }, { cx: 300, cy: 190 }, { cx: 440, cy: 190 },
  { cx: 160, cy: 670 }, { cx: 300, cy: 670 }, { cx: 440, cy: 670 },
];

const PHONE_STALLS: Stall[] = [
  { x: 15, y: 40, w: 100, h: 72, n: "01" },
  { x: 130, y: 40, w: 100, h: 72, n: "02" },
  { x: 245, y: 40, w: 100, h: 72, n: "03" },
  { x: 245, y: 668, w: 100, h: 72, n: "04" },
  { x: 130, y: 668, w: 100, h: 72, n: "05" },
  { x: 15, y: 668, w: 100, h: 72, n: "06" },
];

const PHONE_TABLES: Table[] = [
  { cx: 65, cy: 160 }, { cx: 180, cy: 160 }, { cx: 295, cy: 160 },
  { cx: 65, cy: 620 }, { cx: 180, cy: 620 }, { cx: 295, cy: 620 },
];

export function FloorPlan() {
  return (
    <>
      <Plan
        viewBox="0 0 1200 700"
        stalls={LANDSCAPE_STALLS}
        tables={LANDSCAPE_TABLES}
        className="hidden landscape:block"
      />
      <Plan
        viewBox="0 0 600 860"
        stalls={TABLET_STALLS}
        tables={TABLET_TABLES}
        className="hidden md:portrait:block"
      />
      <Plan
        viewBox="0 0 360 780"
        stalls={PHONE_STALLS}
        tables={PHONE_TABLES}
        className="md:hidden landscape:hidden"
      />
    </>
  );
}
