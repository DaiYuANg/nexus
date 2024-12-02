import {Center, Container, Divider, Paper, PaperProps, Text} from "@mantine/core";
import {IconCloudDataConnection} from "@tabler/icons-react";
import {LoginForm} from "./LoginForm.tsx";

const Login = (props: PaperProps) => {

  return (
    <>
      <Container style={{height: '100vh'}}>
        <Center style={{height: '80%'}}>
          <Paper radius="md" w={'40%'} p="xl" withBorder {...props}>
            <Text size="lg" fw={500}>
              <IconCloudDataConnection/>Welcome to Nexus
            </Text>
            <Divider label="Or continue with email" labelPosition="center" my="lg"/>
            <LoginForm/>
          </Paper>
        </Center>
      </Container>

    </>
  );
}

export {Login}