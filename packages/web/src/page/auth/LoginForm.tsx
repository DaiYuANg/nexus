import {Button, Group, PasswordInput, Stack, TextInput} from "@mantine/core";
import {useForm} from "@mantine/form";
import {FormEvent, useCallback} from "react";
import {useMutation} from "@tanstack/react-query";
import {login} from "../../api/auth.ts";
import {useAuthStore} from "../../store/useAuthStore.ts";
import {useNavigate} from "react-router";

const LoginForm = () => {
  const {setToken} = useAuthStore()
  const navigate = useNavigate();
  const form = useForm({
    initialValues: {
      email: '',
      password: '',
    },

    validate: {
      email: (val) => (/^\S+@\S+$/.test(val) ? null : 'Invalid email'),
      password: (val) => (val.length <= 6 ? 'Password should include at least 6 characters' : null),
    },
  });

  const {mutate} = useMutation({
    mutationKey: ['loginForm'],
    mutationFn: login,
    onSuccess(data) {
      setToken(data.data)
      navigate('/file');
    },
  })

  const submit = useCallback((e: FormEvent) => {
    e.preventDefault()
    console.log(form.values)
    mutate(form.values)

  }, [form])
  return <>
    <form onSubmit={submit}>
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
}

export {LoginForm}