import { Button, Group, LoadingOverlay, PasswordInput, Stack, TextInput } from '@mantine/core';
import { isEmail, useForm } from '@mantine/form';
import { FormEvent, useCallback } from 'react';
import { useMutation } from '@tanstack/react-query';
import { login } from '../../api/auth.ts';
import { useAuthStore } from '../../store/useAuthStore.ts';
import { useNavigate } from 'react-router';
import { useUserStore } from '../../store/useUserStore.tsx';

const LoginForm = () => {
  const { setToken } = useAuthStore();
  const { setUser } = useUserStore();
  const navigate = useNavigate();
  const form = useForm({
    initialValues: {
      email: '',
      password: '',
    },

    validate: {
      email: (val) => {
        return isEmail(val) ? null : 'Invalid email';
        // return /^\S+@\S+$/.test(val) ? null : 'Invalid email';
      },
      password: (val) => (val.length <= 6 ? 'Password should include at least 6 characters' : null),
    },
  });

  const { mutate, isPending } = useMutation({
    mutationKey: ['loginForm'],
    mutationFn: login,
    onSuccess(data) {
      console.log(data);
      setToken(data.data.token);
      setUser(data.data);
      navigate('/file');
    },
  });

  const submit = useCallback(
    (e: FormEvent) => {
      e.preventDefault();
      if (!form.validate().hasErrors) {
        mutate(form.values);
      }
    },
    [form],
  );
  return (
    <>
      <form onSubmit={submit}>
        <LoadingOverlay visible={isPending} zIndex={1000} overlayProps={{ radius: 'sm', blur: 2 }} />
        <Stack>
          <TextInput
            required
            label="Email"
            placeholder="hello@mantine.dev"
            value={form.values.email}
            onChange={(event) => form.setFieldValue('email', event.currentTarget.value)}
            error={form.errors.email && 'Invalid email'}
            radius="md"
          />

          <PasswordInput
            required
            label="Password"
            placeholder="Your password"
            value={form.values.password}
            onChange={(event) => form.setFieldValue('password', event.currentTarget.value)}
            error={form.errors.password && 'Password should include at least 6 characters'}
            radius="md"
          />
        </Stack>

        <Group justify="end" mt="xl">
          <Button type="submit" radius="xl">
            Login
          </Button>
        </Group>
      </form>
    </>
  );
};

export { LoginForm };
