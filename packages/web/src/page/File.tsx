import {ActionIcon, Box, Card, Center, Divider, Flex, Grid, Group, Text} from "@mantine/core";
import {IconFolder, IconLayout2} from "@tabler/icons-react";

const folders = [
  {name: 'Documents', icon: <IconFolder size={30}/>, files: 12},
  {name: 'Photos', icon: <IconFolder size={30}/>, files: 43},
  {name: 'Music', icon: <IconFolder size={30}/>, files: 10},
  {name: 'Videos', icon: <IconFolder size={30}/>, files: 28},
  {name: 'Downloads', icon: <IconFolder size={30}/>, files: 5},
  {name: 'Projects', icon: <IconFolder size={30}/>, files: 22},
];

const File = () => {
  return <>
    <Box>
      <Flex justify={'space-between'}>
        <Group>
          面包屑导航
        </Group>
        <Group>
          <ActionIcon>
            <IconLayout2/>
          </ActionIcon>
        </Group>
      </Flex>
      <Divider/>
      <Grid gutter="md" style={{padding: '20px'}}>
        {folders.map((folder, index) => (
          <Grid.Col span={2} key={index}>
            <Card color={'blue'} component={'a'}
                  href="#"
                  shadow="sm" padding="md" radius="md"
                  style={{
                    display: 'flex', flexDirection: 'column', alignItems: 'center',
                    cursor: 'pointer',
                    backgroundColor: '#f0f0f0',  // 设置背景色
                  }}
            >
              <Center style={{marginBottom: '10px'}}>
                {folder.icon}
              </Center>
              <Text size="lg" style={{marginBottom: '8px'}}>
                {folder.name}
              </Text>
              <Text size="sm" color="dimmed">
                {folder.files} files
              </Text>
            </Card>
          </Grid.Col>
        ))}
      </Grid>
    </Box>
  </>
}

export {File}