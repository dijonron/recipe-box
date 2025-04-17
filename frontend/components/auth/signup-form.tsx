"use client";

import { signup } from "@/actions/auth/signup";
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
import Link from "next/link";
import { startTransition, useActionState, useEffect } from "react";
import { useForm } from "react-hook-form";
import { toast } from "sonner";
import { SignupInput, signupSchema } from "./types";
import { Loader2 as Spinner } from "lucide-react";

export function SignupForm({ className }: React.ComponentProps<"form">) {
  const [state, formAction, pending] = useActionState(signup, {
    success: false,
    errors: {},
  });

  const form = useForm<SignupInput>({
    resolver: zodResolver(signupSchema),
    defaultValues: {
      name: "",
      email: "",
      password: "",
      confirmPassword: "",
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
          form.setError(key as keyof SignupInput, {
            type: "server",
            message: value?.[0],
          });
        }
      });
    }
    if (touchedOrSubmitted && !state?.success) {
      toast.error("Sign up failed.", {
        description: Object.values(state?.errors || {})
          .flat()
          .join(", "),
      });
    }
  }, [state, form]);

  function onSubmit(data: SignupInput) {
    startTransition(() => {
      formAction(data as any as FormData);
    });
  }

  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className={cn("flex flex-col gap-6", className)}
      >
        <div className="flex flex-col items-center gap-2 text-center">
          <h1 className="text-2xl font-bold">Sign up </h1>
          <p className="text-muted-foreground text-sm">
            Fill in the details below to get started
          </p>
        </div>
        <FormField
          control={form.control}
          name="name"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Name</FormLabel>
              <FormControl>
                <Input placeholder="Hank Kuhnan" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
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
              <FormLabel>Password</FormLabel>
              <FormControl>
                <Input type="password" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="confirmPassword"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Confirm Password</FormLabel>
              <FormControl>
                <Input type="password" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        <Button type="submit" className="w-full" disabled={pending}>
          {pending && <Spinner className="animate-spin" />}
          Sign up
        </Button>
        <div className="text-center text-sm">
          Already have an account?{" "}
          <Link href="/login" className="underline underline-offset-4">
            Log in
          </Link>
        </div>
      </form>
    </Form>
  );
}
