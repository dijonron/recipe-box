import { getCurrentUser } from "@/lib/auth";
import { redirect } from "next/navigation";

export default async function OnboardingLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const user = await getCurrentUser();
  if (!user) {
    redirect("/login");
  }

  return (
    <div className="flex flex-col items-center justify-center w-full h-full">
      <h1 className="text-3xl font-bold mb-6 mt-32">Welcome to Recipe Box!</h1>
      <p className="text-muted-foreground mb-2">
        Before you can start adding recipes, you need to set up or join a recipe
        box.
      </p>
      <p className="text-muted-foreground mb-10">
        Choose an option to proceed.
      </p>

      {children}
    </div>
  );
}
