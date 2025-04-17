"use client";

import { login } from "@/actions/auth/login";
import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { cn } from "@/lib/utils";
import { zodResolver } from "@hookform/resolvers/zod";
import { Loader2 as Spinner } from "lucide-react";
import Link from "next/link";
import { startTransition, useActionState, useEffect } from "react";
import { useForm } from "react-hook-form";
import { toast } from "sonner";
import { LoginInput, loginSchema } from "./types";

export function LoginForm({ className }: React.ComponentProps<"form">) {
  const [state, formAction, pending] = useActionState(login, {
    success: false,
    errors: {},
  });

  const form = useForm<LoginInput>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  useEffect(() => {
    const touchedOrSubmitted =
      form.formState.isSubmitted ||
      (form.getFieldState("email").isTouched &&
        form.getFieldState("password").isTouched);

    if (state?.errors) {
      Object.entries(state.errors).forEach(([key, value]) => {
        if (key !== "_form") {
          form.setError(key as keyof LoginInput, {
            type: "server",
            message: value?.[0],
          });
        }
      });
    }
    if (touchedOrSubmitted && !state?.success) {
      toast.error("Login failed.", {
        description: Object.values(state?.errors || {})
          .flat()
          .join(", "),
      });
    }
  }, [state, form]);

  const onSubmit = (data: LoginInput) => {
    startTransition(() => {
      formAction(data as any as FormData);
    });
  };

  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className={cn("flex flex-col gap-6", className)}
      >
        <div className="flex flex-col items-center gap-2 text-center">
          <h1 className="text-2xl font-bold">Login to your account</h1>
          <p className="text-muted-foreground text-sm text-balance">
            Enter your email below to login to your account
          </p>
        </div>
        <FormField
          control={form.control}
          name="email"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Email</FormLabel>
              <FormControl>
                <Input placeholder="me@example.com" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="password"
          render={({ field }) => (
            <FormItem>
              <div className="flex items-center">
                <FormLabel>Password</FormLabel>
                <a
                  href="#"
                  className="ml-auto text-sm underline-offset-4 hover:underline"
                >
                  Forgot your password?
                </a>
              </div>
              <FormControl>
                <Input type="password" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <Button type="submit" className="w-full" disabled={pending}>
          {pending && <Spinner className="animate-spin" />}
          Login
        </Button>
        <div className="text-center text-sm">
          Don&apos;t have an account?{" "}
          <Link href="/signup" className="underline underline-offset-4">
            Sign up
          </Link>
        </div>
      </form>
    </Form>
  );
}
