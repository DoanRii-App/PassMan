import React from 'react';
import {
  IconShieldLock,
  IconSquarePlus,
} from '@tabler/icons';
import { ThemeIcon, UnstyledButton, Group, Text } from '@mantine/core';

interface NavLinkProps {
  icon: React.ReactNode;
  color: string;
  label: string;
}

function NavLink({ icon, color, label }: NavLinkProps) {
  return (
    <UnstyledButton
      sx={(theme) => ({
        display: 'block',
        width: '100%',
        padding: theme.spacing.xs,
        borderRadius: theme.radius.sm,
        color:
          theme.colorScheme === 'dark' ? theme.colors.dark[0] : theme.black,

        '&:hover': {
          backgroundColor:
            theme.colorScheme === 'dark'
              ? theme.colors.dark[6]
              : theme.colors.gray[0],
        },
      })}
    >
      <Group>
        <ThemeIcon color={color} variant="light">
          {icon}
        </ThemeIcon>

        <Text size="sm">{label}</Text>
      </Group>
    </UnstyledButton>
  );
}

const data = [
  { icon: <IconShieldLock size={16} />, color: 'teal', label: 'Password Vault' },
  { icon: <IconSquarePlus size={16} />, color: 'grape', label: 'Add Password' },
];

function NavLinks() {
  const links = data.map((link) => <NavLink {...link} key={link.label} />);
  return <div>{links}</div>;
}

export default NavLinks;
