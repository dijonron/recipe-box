import { GalleryVerticalEnd } from "lucide-react";

export default async function AuthLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en" suppressHydrationWarning>
      <body>
        <div className="grid min-h-svh lg:grid-cols-2">
          <div className="flex flex-col gap-4 p-6 md:p-10">
            <div className="flex justify-center gap-2 md:justify-start">
              <div className="bg-primary text-primary-foreground flex size-6 items-center justify-center rounded-md">
                <GalleryVerticalEnd className="size-4" />
              </div>
              Recipe Box
            </div>
            <div className="flex flex-1 items-center justify-center">
              <div className="w-full max-w-xs">
                <main>{children}</main>
              </div>
            </div>
          </div>
          <div className="bg-muted relative hidden lg:block">
            <img
              src="/recipe_box.png"
              alt="Image"
              className="absolute inset-0 h-full w-full object-cover dark:brightness-[0.2] dark:grayscale"
            />
          </div>
        </div>{" "}
      </body>
    </html>
  );
}
