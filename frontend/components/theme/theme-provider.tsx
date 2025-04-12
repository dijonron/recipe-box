"use client";

import * as React from "react";
import { ThemeProvider as NextThemeProvider } from "next-themes";

export function ThemeProvider({
  children,
  ...props
}: React.ComponentProps<typeof NextThemeProvider> & {
  children: React.ReactNode;
}) {
  return <NextThemeProvider {...props}>{children}</NextThemeProvider>;
}
