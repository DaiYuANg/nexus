import {forwardRef, useState} from "react";
import {IconMoonStars, IconSun} from "@tabler/icons-react";
import {rem, Switch, Tooltip, useMantineColorScheme, useMantineTheme} from "@mantine/core";

const ThemeSwitcher = forwardRef((_props, _ref) => {
  const theme = useMantineTheme();
  const [checked, setChecked] = useState(false);
  const {setColorScheme} = useMantineColorScheme();
  const sunIcon = (
    <IconSun
      style={{width: rem(16), height: rem(16)}}
      stroke={2.5}
      color={theme.colors.yellow[4]}
    />
  );

  const moonIcon = (
    <IconMoonStars
      style={{width: rem(16), height: rem(16)}}
      stroke={2.5}
      color={theme.colors.blue[6]}
    />
  );

  return <>
    <Tooltip label="Switch tooltip" refProp="rootRef">
      <Switch
        checked={checked}
        onChange={(event) => {
          setChecked(event.currentTarget.checked);
          if (checked) {
            setColorScheme('light')
          } else {
            setColorScheme('dark')
          }
        }}
        size="md" color="dark.4" onLabel={sunIcon} offLabel={moonIcon}/>
    </Tooltip>

  </>;
})

export {ThemeSwitcher}