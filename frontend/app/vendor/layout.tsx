export default function VendorLayout({ children }: { children: React.ReactNode }) {
  return (
    <div data-role="vendor" className="flex flex-1 flex-col">
      {children}
    </div>
  );
}
