import type { Metadata } from "next";
import './main.css'
import ToastLayout from "@/components/ToastLayout";
import "react-toastify/dist/ReactToastify.css";

export const metadata: Metadata = {
  title: "SuperFin - Your Financial App",
  description: "SuperFin - Your Financial App",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>
        {children}
        <ToastLayout />
      </body>
    </html>
  );
}
