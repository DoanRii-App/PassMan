import {
  ColorScheme,
  Navbar,
  Text,
  Group,
  Box,
  UnstyledButton,
  useMantineTheme,
  Stack,
  ScrollArea,
} from "@mantine/core";
import { IconLogout } from "@tabler/icons";
import NavLinks from "../components/NavLinks";

function NaviBar() {
  return (
    <Stack justify="space-between" sx={(theme) => ({ height: "100%" })}>
      <Navbar.Section grow component={ScrollArea}>
        <NavLinks />
      </Navbar.Section>

      <Navbar.Section>
        <Box>
          <UnstyledButton
            sx={(theme) => ({
              display: "block",
              width: "100%",
              padding: theme.spacing.xs,
              borderRadius: theme.radius.sm,
              color:
                theme.colorScheme === "dark"
                  ? theme.colors.dark[0]
                  : theme.black,

              "&:hover": {
                backgroundColor:
                  theme.colorScheme === "dark"
                    ? theme.colors.dark[6]
                    : theme.colors.gray[0],
              },
            })}
          >
            <Group>
              <IconLogout size={25} />
              <Text weight={500}>Logout</Text>
            </Group>
          </UnstyledButton>
        </Box>
      </Navbar.Section>
    </Stack>
  );
}

export default NaviBar;
