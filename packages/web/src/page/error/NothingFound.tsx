import { Button, Container, Group, Text, Title } from '@mantine/core';

const NothingFound = () => {
  return (
    <Container className={['pt-5 pb-5'].join(' ')}>
      <Text className={['text-center'].join(' ')}>404</Text>
      <Title>You have found a secret place.</Title>
      <Text c="dimmed" size="lg" ta="center">
        Unfortunately, this is only a 404 page. You may have mistyped the address, or the page has been moved to another
        URL.
      </Text>
      <Group justify="center">
        <Button variant="subtle" size="md">
          Take me back to home page
        </Button>
      </Group>
    </Container>
  );
};

export { NothingFound };
